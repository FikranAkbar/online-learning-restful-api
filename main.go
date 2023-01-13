package main

import (
	"github.com/labstack/echo/v4/middleware"
	"online-learning-restful-api/di"
)

func main() {
	e := di.InitializedEchoServer()
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start("localhost:8000"))
}
