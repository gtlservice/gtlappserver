package main

import "github.com/gtlservice/gtlappserver/base"
import "github.com/gtlservice/gtlappserver/server"
import "github.com/gtlservice/gutils/system"

import (
	"os"
)

func main() {

	appserver, err := server.NewAppServer()
	if err != nil {
		panic(err)
		os.Exit(base.EXITCODE_INITFAILED)
	}

	defer func() {
		appserver.Stop()
		os.Exit(base.EXITCODE_EXITED)
	}()

	if err := appserver.Startup(); err != nil {
		panic(err)
		os.Exit(base.EXITCODE_STARTFAILED)
	}
	system.InitSignal(nil)
}
