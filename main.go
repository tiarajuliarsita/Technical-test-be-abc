package main

import (
	"fmt"
	"technical-test-be-abc/configs"
	"technical-test-be-abc/configs/db"
	"technical-test-be-abc/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg:=configs.LoadConfig()
	db := db.InitDB()
	routes.Routes(e, db)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	routes.Routes(e, db)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.APP.SERVERPORT)))
}
