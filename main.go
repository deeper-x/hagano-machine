package main

import (
	"fmt"

	"github.com/deeper-x/hagano-machine/pkgs/httpRegister"
	"github.com/deeper-x/hagano-machine/pkgs/httpServer"
	"github.com/deeper-x/hagano-machine/pkgs/strFormatters"
)

func main() {
	fmt.Println("Running server...")

	httpRegister.ServiceCaller()

	httpServer.Run(strFormatters.HTTPPort)
}
