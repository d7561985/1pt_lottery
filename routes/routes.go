package routes

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/core/router"
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

	api := router.Party("/api")
	{
		api.Post("/", start)
		api.Get("/ws", W.Handler())
	}
}
