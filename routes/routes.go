package routes

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/basicauth"
	"github.com/kataras/iris/websocket"
	"time"
)

func RegisterRoutes(in *router.APIBuilder) {
	in.Any("/iris-ws.js", websocket.ClientHandler())

	api := in.Party("/api", cors.AllowAll()).AllowMethods(iris.MethodOptions)
	{
		api.Post("/", start)
		api.Get("/ws", W.Handler())

		authentication := basicauth.New(basicauth.Config{
			Users:   map[string]string{"root": "1root1"},
			Realm:   "Authorization Required", // defaults to "Authorization Required"
			Expires: time.Duration(30) * time.Minute,
		})

		admin := api.Party("/admin", authentication)
		{
			admin.Get("/competitors", competitorsList)
			admin.Delete("/database", dropDatabase)
			actions := admin.Party("/actions")
			{
				actions.Get("/begin", lotteryBegin)
				actions.Get("/stop", lotteryStop)
			}
		}
	}
}
