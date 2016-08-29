package api

import "github.com/gtlservice/gtlappserver/ctrl"
import "github.com/labstack/echo"
import mw "github.com/labstack/echo/middleware"
import "github.com/rs/cors"

import (
	"strings"
)

type ApiServerArgs struct {
	Bind string
	Cors map[string]string
}

type ApiServer struct {
	bind       string
	controller *ctrl.Controller
	echo       *echo.Echo
	cors       *cors.Cors
}

func NewApiServer(controller *ctrl.Controller, args *ApiServerArgs) *ApiServer {

	e := echo.New()
	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(args.Cors["origin"], ","),
		AllowedMethods:   strings.Split(args.Cors["methods"], ","),
		AllowCredentials: true,
	})

	e.Use(c.Handler)
	e.Use(mw.Recover())
	apiserver := &ApiServer{
		bind:       args.Bind,
		controller: controller,
		echo:       e,
		cors:       c,
	}
	apiserver.InitRouter()
	return apiserver
}

func (s *ApiServer) Startup() {

	if s.echo != nil {
		s.echo.Run(s.bind)
	}
}
