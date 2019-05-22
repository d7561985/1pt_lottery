package routes

import (
	"github.com/d7561985/1pt_lottery"
	"github.com/d7561985/1pt_lottery/dto"
	"github.com/iris-contrib/go.uuid"
	"github.com/kataras/iris"
	"github.com/rs/zerolog/log"
)

// POST /api/start
func start(ctx iris.Context) {
	start := &dto.StartRequest{}
	if err := ctx.ReadJSON(start); err != nil {
		log.Error().Err(err).Msg("read")
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	id, err := uuid.NewV4()
	if err != nil {
		log.Error().Err(err).Msg("gen uuid")
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	// 2h leave cookie
	ctx.SetCookieKV(lottery.Cookie, id.String())
	ctx.JSON(&dto.StartResponse{UUID: id.String()})
}
