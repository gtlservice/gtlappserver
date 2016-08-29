package server

import "github.com/gtlservice/gtlappserver/api"
import "github.com/gtlservice/gtlappserver/ctrl"
import "github.com/gtlservice/gtlappserver/etc"
import "github.com/gtlservice/gutils/logger"
import "github.com/gtlservice/gutils/system"
import "github.com/gtlservice/gzkwrapper"

import (
	"flag"
)

type AppServer struct {
	AppServerHandler
	Configuration *etc.Configuration
	ApiServer     *api.ApiServer
	Controller    *ctrl.Controller
	Worker        *gzkwrapper.Worker
}

func NewAppServer() (*AppServer, error) {

	path, err := system.GetExecDir()
	if err != nil {
		return nil, err
	}

	key, err := system.MakeKeyFile(path + "/appserver.key")
	if err != nil {
		return nil, err
	}

	var etcfile string
	flag.StringVar(&etcfile, "f", "etc/config.yaml", "server etc file.")
	flag.Parse()
	configuration, err := etc.NewConfiguration(etcfile)
	if err != nil {
		return nil, err
	}

	appserver := &AppServer{}
	largs := configuration.GetLogger()
	logger.OPEN(largs)
	zkargs := configuration.GetZkWrapper()
	worker, err := gzkwrapper.NewWorker(key, zkargs, appserver)
	if err != nil {
		return nil, err
	}

	sargs := configuration.GetService()
	controller := ctrl.NewController(worker, sargs)
	apiargs := configuration.GetApiServer()
	apiserver := api.NewApiServer(controller, apiargs)
	appserver.Configuration = configuration
	appserver.Controller = controller
	appserver.ApiServer = apiserver
	appserver.Worker = worker
	return appserver, nil
}

func (appserver *AppServer) Startup() error {

	if err := appserver.Controller.Initialize(); err != nil {
		logger.ERROR("[#server#] appserver controller initialize error, %s", err)
		return err
	}
	go func() { //启动apiserver
		appserver.ApiServer.Startup()
	}()
	logger.INFO("[#server#] appserver started.")
	logger.INFO("[#server#] zkwrapper %s, key:%s", appserver.Worker.Data.Location, appserver.Worker.Key)
	logger.INFO("[#server#] appserver configuration...\r\n%v", appserver.Configuration)
	return nil
}

func (appserver *AppServer) Stop() error {

	if err := appserver.Controller.UnInitialize(); err != nil {
		logger.ERROR("[#server#] appserver controller uninitialize error, %s", err)
		return err
	}
	logger.INFO("[#server#] appserver closed.")
	logger.CLOSE()
	return nil
}
