package strFormatters

import (
	"fmt"

	"github.com/deeper-x/hagano-machine/pkgs/configuration"
)

// HTTPPort ... todo description
var HTTPPort = fmt.Sprintf(":%d", configuration.Port)
