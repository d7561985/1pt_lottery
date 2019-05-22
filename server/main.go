package main

import (
	"fmt"
	"github.com/d7561985/1pt_lottery/persistence"
	"github.com/d7561985/1pt_lottery/routes"
	"github.com/d7561985/heroku_boilerplate/pkg/config"
	"github.com/kataras/iris"
)

func main() {
	if err := persistence.InitDB(); err != nil {
		panic(err)
	}

	app := iris.Default()

	// expose root `asset` dir  for distributing static files
	app.StaticWeb("/", "./static")

	// register routes
	routes.RegisterRoutes(app.APIBuilder)

	_ = app.Run(
		iris.Addr(fmt.Sprintf(":%s", config.V.Port)),
		iris.WithOptimizations,
		iris.WithPostMaxMemory(1<<20), // 1Mb post limit
	)
}
