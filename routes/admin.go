package routes

import (
	"github.com/d7561985/1pt_lottery/persistence"
	"github.com/kataras/iris"
)

func competitorsList(ctx iris.Context) {
	res := persistence.Competitors{}
	if err := res.All(); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	_, _ = ctx.JSON(res)
}

func dropDatabase(ctx iris.Context) {
	if err := persistence.Clean(); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
	}
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON("OK")
}
