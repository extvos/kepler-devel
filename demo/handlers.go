package main

import (
	"github.com/extvos/kepler/contrib/restlet"
	"github.com/extvos/kepler/service"
	"github.com/extvos/kepler/servlet"
	log "github.com/sirupsen/logrus"
)

func hello(ctx servlet.RequestContext) error {
	log.Debugln("hello:> ", ctx.Ctx().Method(), ctx.Ctx().Path())
	if nil != ctx.Session() {
		log.Debugln("hello:session:> ", ctx.Session())
	}
	return ctx.Ctx().SendString("Hello")
}

func middleware(ctx servlet.RequestContext) error {
	log.Debugln("middleware:> ", ctx.Ctx().Method(), ctx.Ctx().Path())
	return ctx.Next()
}

func init() {
	service.Use("/hello", middleware)
	service.Get("/", hello)
	service.Get("/hello", hello)
	service.All("/student", restlet.NewHandler[Student]())
}
