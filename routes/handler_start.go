package routes

import (
	"github.com/d7561985/1pt_lottery"
	"github.com/d7561985/1pt_lottery/dto"
	"github.com/d7561985/1pt_lottery/persistence"
	img "github.com/d7561985/1pt_lottery/pkg/image"
	"github.com/gookit/validate"
	"github.com/iris-contrib/go.uuid"
	"github.com/kataras/iris"
	"github.com/rs/zerolog/log"
)

// POST /api/start
func start(ctx iris.Context) {
	start := &dto.UserRequest{}
	if err := ctx.ReadJSON(start); err != nil {
		log.Error().Err(err).Msg("read")
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	v := validate.Struct(start)
	if !v.Validate() {
		log.Info().Str("vld", v.Errors.String()).Msg("bad validation")
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(&dto.WSEvent{Event: lottery.EventError, Data: v.Errors.String()})
		return
	}

	av, err := img.ReadImage(start.Avatar)
	if err != nil {
		log.Error().Err(err).Msg("decode avatar")
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(&dto.WSEvent{Event: lottery.EventError, Data: "error of decoding avatar"})
		return
	}
	avb, err := img.JPEGwithBase64(av, 80)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		log.Error().Err(err).Msg("decode image to base64")
		return
	}

	id, err := uuid.NewV4()
	if err != nil {
		log.Error().Err(err).Msg("gen uuid")
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	c := persistence.Competitor{Name: start.Name, UUID: id.String(), Avatar: avb}
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
