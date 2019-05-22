package routes

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/basicauth"
	"github.com/kataras/iris/websocket"
	"time"
)

func RegisterRoutes(router *router.APIBuilder) {
	// cors
	router.Use(cors.New(cors.Options{
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		// should contain all supported
		AllowedMethods:     []string{"GET", "DELETE", "POST", "PUT"},
		OptionsPassthrough: true,
	}), func(context iris.Context) {
		// hack for OPTIONS, no need handle options method.
		if context.Request().Method != "OPTIONS" {
			context.Next()
			return
		}
		context.StatusCode(iris.StatusNoContent)
	})

	router.Any("/iris-ws.js", websocket.ClientHandler())

	api := router.Party("/api").AllowMethods(iris.MethodOptions)
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
		}
	}
}
