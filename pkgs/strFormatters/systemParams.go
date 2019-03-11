package strFormatters

import (
	"fmt"

	"github.com/deeper-x/hagano-machine/pkgs/configuration"
)

// HTTPPort ... todo description
var HTTPPort = fmt.Sprintf(":%d", configuration.Port)

// ConnStr ... to do description
var ConnStr = fmt.Sprintf("postgres://%s:%s@%s/%s", configuration.User, configuration.Password, configuration.Host, configuration.Db)
