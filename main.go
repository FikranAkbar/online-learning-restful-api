package main

import (
	"online-learning-restful-api/di"
)

func main() {
	server := di.InitializedEchoServer()
	server.Logger.Fatal(server.Start("localhost:8000"))
}
