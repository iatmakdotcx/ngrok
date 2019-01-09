package client

import (
	"os"
	"path"
	"path/filepath"

	"github.com/kardianos/service"
)

var svcConfig = &service.Config{
	Name:        "ngrok",
	DisplayName: "ngrok",
	Description: "ngrok",
	Arguments:   []string{"-log=ngrok.log", "server"},
}

var ggConfig *Configuration

type serviceprogram struct{}

func (p *serviceprogram) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *serviceprogram) run() {
	// 代码写在这儿
	NewController().Run(ggConfig)
}

func (p *serviceprogram) Stop(s service.Service) error {
	return nil
}

func init() {
	p, e := os.Executable()
	if e != nil {
		svcConfig.Arguments[0] = "-log=log/ngrok.log"
	} else {
		svcConfig.Arguments[0] = "-log=" + path.Join(filepath.Dir(p), "log/ngrok.log")
	}
}
