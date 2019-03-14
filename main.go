package main

import (
	"github.com/deeper-x/hagano-machine/pkgs/httpServer"
	"github.com/deeper-x/hagano-machine/pkgs/servicesRegister"
	"github.com/deeper-x/hagano-machine/pkgs/strFormatters"
)

func main() {
	servicesRegister.UrlsRegister()

	httpServer.Run(strFormatters.HTTPPort)
}
