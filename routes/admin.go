package routes

import (
	"github.com/d7561985/1pt_lottery"
	"github.com/d7561985/1pt_lottery/dto"
	"github.com/d7561985/1pt_lottery/persistence"
	"github.com/kataras/iris"
	"github.com/rs/zerolog/log"
	"math/rand"
)

var (
	Start = false
)

func competitorsList(ctx iris.Context) {
	_, _ = ctx.JSON(new(persistence.Competitors).FillByStorage())
}

func lotteryBegin(ctx iris.Context) {
	if Start {
		ctx.StatusCode(iris.StatusNotAcceptable)
		_, _ = ctx.JSON("already started")
		return
	}
	if err := W.BroadCast(&dto.WSEvent{Event: lottery.WsEventBegin}); err != nil {
		ctx.StatusCode(iris.StatusConflict)
		log.Error().Err(err).Msg("lotteryBegin")
		return
	}
	ctx.StatusCode(iris.StatusOK)
	Start = true
}

func lotteryStop(ctx iris.Context) {
	if !Start {
		ctx.StatusCode(iris.StatusNotAcceptable)
		_, _ = ctx.JSON("not started yet")
		return
	}
	num, _ := persistence.S.GetOnline()
	if err := W.BroadCast(&dto.WSEvent{Event: lottery.WsEventStop, Data: &dto.WSNameResponse{
		UserRequest: dice(),
		Competitors: num,
	}}); err != nil {
		ctx.StatusCode(iris.StatusConflict)
		log.Error().Err(err).Msg("lotteryStop")
		return
	}

	ctx.StatusCode(iris.StatusOK)
	Start = false
}

func dice() dto.UserRequest {
	total, list := persistence.S.GetOnline()
	winner := rand.Int31n(int32(total))
	return list[winner]
}
