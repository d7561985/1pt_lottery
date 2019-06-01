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
	"time"
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
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
		EnableCompression: true,
		// EvtMessagePrefix: []byte("123"),
	})

	W = &WsController{srv}
	srv.OnConnection(W.handleConnection)

	go func() {
		c := time.Tick(time.Second * 5)
		for t := range c {
			if err := W.BroadCast(&dto.WSEvent{Event: lottery.WSEventServerTime, Data: t}); err != nil {
				log.Error().Err(err).Msg("time")
			}
		}
	}()
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
	persistence.S.Online.Store(uid, *participant)

	total, list := persistence.S.GetOnline()
	// broadcast to all
	if err := w.Emit(c, websocket.Broadcast, &dto.WSEvent{
		Event: lottery.WsEventEnter,
		Data: &dto.WSNameResponse{
			UserRequest: dto.UserRequest{
				Name:   participant.Name,
				Avatar: participant.Avatar,
			},
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

	// listener of any data
	c.OnMessage(func(b []byte) {
		// no escapes required => use interface unmarshal
		var i interface{}
		if err := json.Unmarshal(b, &i); err != nil {
			log.Error().Err(err).Msg("ws data read")
			return
		}

		log.Info().Interface("data", i).Msg("take ws data")
		if err := w.Emit(c, websocket.All, i); err != nil {
			log.Error().Err(err).Msg("ws data send")
		}
	})

	// hook at disconnect
	c.OnDisconnect(func() {
		persistence.S.Online.Delete(uid)
		total, _ := persistence.S.GetOnline()

		if err := w.Emit(c, websocket.Broadcast, &dto.WSEvent{
			Event: lottery.WsEventLeave,
			Data: &dto.WSNameResponse{
				UserRequest: dto.UserRequest{
					Name: participant.Name,
				},
				Competitors: total,
			},
		}); err != nil {
			log.Error().Err(err).Msg("fail send broadcast on connect")
		}
	})
}
