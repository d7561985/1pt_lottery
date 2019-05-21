package routes

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/websocket"
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
	}))

	router.Any("/iris-ws.js", websocket.ClientHandler())

	api := router.Party("/api")
	{
		api.Post("/", start)
		api.Get("/ws", W.Handler())

	}
}
