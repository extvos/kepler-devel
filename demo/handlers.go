package main

import (
	"github.com/extvos/kepler/service"
	"github.com/extvos/kepler/servlet"
)

func hello(ctx servlet.RequestContext) error {
	return ctx.Ctx().SendString("Hello")
}

func init() {
	service.Get("/", hello)
	service.Get("/hello", hello)
}
