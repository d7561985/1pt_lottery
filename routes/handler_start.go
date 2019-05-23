package routes

import (
	"github.com/d7561985/1pt_lottery"
	"github.com/d7561985/1pt_lottery/dto"
	"github.com/d7561985/1pt_lottery/persistence"
	"github.com/gookit/validate"
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

	v := validate.Struct(start)
	if !v.Validate() {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(v.Errors.String())
		return
	}

	id, err := uuid.NewV4()
	if err != nil {
		log.Error().Err(err).Msg("gen uuid")
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	c := persistence.Competitor{Name: start.User, UUID: id.String()}
	if err := c.Create(); err != nil {
		log.Error().Err(err).Msg("create user")
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	// 2h leave cookie
	ctx.SetCookieKV(lottery.Cookie, id.String())
	_, _ = ctx.JSON(&dto.StartResponse{UUID: id.String()})
	log.Info().Str("uuid", id.String()).Msg("start")
}
