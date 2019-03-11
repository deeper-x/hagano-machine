package main

import (
	"github.com/deeper-x/hagano-machine/pkgs/httpRegister"
	"github.com/deeper-x/hagano-machine/pkgs/httpServer"
	"github.com/deeper-x/hagano-machine/pkgs/strFormatters"
)

func main() {
	httpRegister.ServiceCaller()

	httpServer.Run(strFormatters.HTTPPort)
}
