package command

import (
	"github.com/RobyFerro/go-web/app/config"
	"github.com/RobyFerro/go-web/app/http"
	"github.com/RobyFerro/go-web/app/kernel"
	"github.com/RobyFerro/go-web/exception"
	daemon "github.com/sevlyar/go-daemon"
	"log"
)

type ServerDaemon struct {
	Signature   string
	Description string
}

func (c *ServerDaemon) Register() {
	c.Signature = "server:daemon"
	c.Description = "Run Go-Web server as a daemon"
}

// Run Go-Web as a daemon
func (c *ServerDaemon) Run(kernel *kernel.HttpKernel, args string, console map[string]interface{}) {
	err := kernel.Container.Invoke(func(conf config.Conf) {

		// Simple way to check is a string contains only digits
		cntxt := &daemon.Context{
			PidFileName: "storage/log/go-web.pid",
			PidFilePerm: 0754,
			LogFileName: "storage/log/go-webd.log",
			LogFilePerm: 0754,
			Umask:       027,
		}

		daemonBg, daemonError := cntxt.Reborn()

		if daemonError != nil {
			log.Fatal("Troubles to start daemon ", daemonError, daemonError.Error())
		}

		if daemonBg != nil {
			return
		}

		defer func() {
			_ = cntxt.Release()
		}()

		http.StartServer(kernel.Container)

	})

	if err != nil {
		exception.ProcessError(err)
	}
}