package routes

import (
	"github.com/d7561985/1pt_lottery"
	"github.com/icrowley/fake"
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

func (w *WsController) BroadCast(msg []byte) error {
	con := w.ws.GetConnections()
	for _, c := range con {
		return c.To(websocket.All).EmitMessage(msg)
	}
	return errEmpty
}

func (w *WsController) Emit(event, msg string) error {
	con := w.ws.GetConnections()
	for _, c := range con {
		return c.To(websocket.All).Emit(event, msg)
	}
	return errEmpty
}

func (w *WsController) handleConnection(c websocket.Connection) {
	// register messages
	c.On(lottery.WsEventEnter, w.enter(c))

	if err := c.To(websocket.All).Emit(lottery.WsEventEnter, fake.FullName()); err != nil {
		log.Error().Err(err).Msg("fail send broadcast on connect")
	}
}

// client send event: WsEventEnter - enter
// return name which was already create by POST form and and all users
func (w *WsController) enter(c websocket.Connection) websocket.MessageFunc {
	return func(msg string) {
		log.Info().Str("ip", c.Context().RemoteAddr()).Str("id", c.ID()).Str("msg", msg).Msg(lottery.WsEventEnter)

		// Write message back to the client message owner:
		_ = c.Emit(lottery.WsEventEnter, fake.FullName())
		_ = c.To(websocket.Broadcast).Emit("join", msg)
	}
}
