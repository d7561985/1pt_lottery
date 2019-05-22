package routes

import (
	"github.com/d7561985/1pt_lottery"
	"github.com/d7561985/1pt_lottery/persistence"
	"github.com/kataras/iris"
	"github.com/rs/zerolog/log"
)

func competitorsList(ctx iris.Context) {
	res := persistence.Competitors{}
	if err := res.All(); err != nil {
		log.Error().Err(err).Msg("competitorsList")
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	_, _ = ctx.JSON(res)
}

func dropDatabase(ctx iris.Context) {
	if err := persistence.Clean(); err != nil {
		log.Error().Err(err).Msg("dropDatabase")
		ctx.StatusCode(iris.StatusInternalServerError)
	}
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON("OK")
}

func lotteryBegin(ctx iris.Context) {
	if err := W.Emit(lottery.WsEventBegin, ""); err != nil {
		ctx.StatusCode(iris.StatusConflict)
		log.Error().Err(err).Msg("lotteryBegin")
		return
	}
	ctx.StatusCode(iris.StatusOK)
}

func lotteryStop(ctx iris.Context) {
	if err := W.Emit(lottery.WsEventStop, ""); err != nil {
		ctx.StatusCode(iris.StatusConflict)
		log.Error().Err(err).Msg("lotteryStop")
		return
	}
	ctx.StatusCode(iris.StatusOK)
}
