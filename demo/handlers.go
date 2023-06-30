package main

import (
	"github.com/extvos/kepler/contrib/restlet"
	"github.com/extvos/kepler/service"
	"github.com/extvos/kepler/servlet"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func hello(ctx servlet.RequestContext) error {
	log.Debugln("hello:> ", ctx.Ctx().Method(), ctx.Ctx().Path())
	if nil != ctx.Session() {
		log.Debugln("hello:session:> ", ctx.Session())
	}
	p := ctx.Ctx().Params("param")
	return ctx.Ctx().SendString("Hello," + p)
}

func middleware(ctx servlet.RequestContext) error {
	log.Debugln("middleware:> ", ctx.Ctx().Method(), ctx.Ctx().Path())
	return ctx.Next()
}

func soapHandler(ctx servlet.RequestContext) error {
	log.Debugln("soapHandler:> ", ctx.Ctx().Method(), ctx.Ctx().Path(), string(ctx.Ctx().Request().Header.ContentType()))
	log.Debugln("soapHandler:> ", string(ctx.Ctx().Request().Body()))
	s := "<soapenv:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:iptv=\"iptv\">\n\n\t<soapenv:Header/>\n\t<soapenv:Body>\n\t   <iptv:ExecCmdResponse soapenv:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">\n\t         <Result xsi:type=\"xsd:int\">0</Result>\n\t         <ErrorDescription xsi:type=\"soapenc:string\" xmlns:soapenc=\"http://schemas.xmlsoap.org/soap/encoding/\">?</ErrorDescription>\n\t</iptv:ExecCmdResponse>\n\t</soapenv:Body>\n\n</soapenv:Envelope>"
	return ctx.Ctx().Status(fiber.StatusOK).SendString(s)
}

func init() {
	service.Use("/hello", middleware)
	service.Get("/", hello)
	service.Get("/hello/:param?", hello)
	service.All("/student/:id?", restlet.NewHandler[Student]())
	service.All("/soap", soapHandler)
}
