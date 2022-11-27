package main

import (
	"recipe-tracker-ms/handler"
)

func main() {
	c := handler.NewHandler()
	c.AddRoute("/hello", handler.HelloWorld)
	c.Start()
}
