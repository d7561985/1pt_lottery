package routes

import (
	"encoding/json"
	"github.com/d7561985/1pt_lottery"
	"github.com/d7561985/1pt_lottery/dto"
	"github.com/d7561985/1pt_lottery/persistence"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/errors"
	"github.com/kataras/iris/websocket"
	"github.com/rs/zerolog/log"
)

var (
	W        *WsController = nil
	errEmpty               = errors.New("no connections")
)

type WsController struct {
	ws *websocket.Server
}

func init() {
	srv := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// EvtMessagePrefix: []byte("123"),
	})

	W = &WsController{srv}
	srv.OnConnection(W.handleConnection)
}

func (w *WsController) Handler() iris.Handler {
	return w.ws.Handler()
}

func (w *WsController) BroadCast(msg interface{}) error {
	con := w.ws.GetConnections()
	for _, c := range con {
		return w.Emit(c, websocket.All, msg)
	}
	return errEmpty
}

func (w *WsController) Emit(c websocket.Connection, tp string, msg interface{}) error {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return c.To(tp).EmitMessage(bytes)
}

func (w *WsController) handleConnection(c websocket.Connection) {
	uid := c.Context().URLParam("uuid")
	participant := new(persistence.Competitor)
	if participant.Find(uid) != nil {
		log.Info().Str("uuid", uid).Msg("disconnect as not exist")

		if err := c.Disconnect(); err != nil {
			log.Error().Err(err).Msg("ws connect")
		}
		return
	}

	// register as online with this value we can send private message
	persistence.Online.Store(uid, *participant)

	total, list := persistence.S.Online()
	// broadcast to all
	if err := w.Emit(c, websocket.Broadcast, &dto.WSEvent{
		Event: lottery.WsEventEnter,
		Data: &dto.WSNameResponse{
			Name:        participant.Name,
			Competitors: total,
		},
	}); err != nil {
		log.Error().Err(err).Msg("fail send broadcast on connect")
	}

	// private message
	if err := w.Emit(c, c.ID(), &dto.WSEvent{
		Event: lottery.WsList,
		Data:  &dto.WSListResponse{Me: participant.Name, List: list},
	}); err != nil {
		log.Error().Err(err).Str("event", lottery.WsList).Msg("whisper fail")
	}

	log.Info().Str("name", participant.Name).Str("uuid", participant.UUID).
		Int("total", total+1).Msg("join")

	// hook at disconnect
	c.OnDisconnect(func() {
		persistence.Online.Delete(uid)
		total, _ := persistence.S.Online()

		if err := w.Emit(c, websocket.Broadcast, &dto.WSEvent{
			Event: lottery.WsEventLeave,
			Data: &dto.WSNameResponse{
				Name:        participant.Name,
				Competitors: total,
			},
		}); err != nil {
			log.Error().Err(err).Msg("fail send broadcast on connect")
		}
	})
}
