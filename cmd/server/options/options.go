package options

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/caicloud/go-common/command"

	"github.com/lichuan0620/nirvana-practice/pkg/info"
)

type CmdLineOpts struct {
	version bool

	HTTPPort uint16
}

// ParseCommands parse command line arguments
// See https://github.com/caicloud/go-common/blob/master/command/README.md
func ParseCommands() *CmdLineOpts {
	ret := &CmdLineOpts{
		HTTPPort: 8080,
	}
	cmd := command.NewCommand("alerting-admin")
	cmd.Add(&ret.version, "version", "v", "show version")
	cmd.Add(&ret.HTTPPort, "httpPort", "p", "port on which the HTTP server would be exposed")
	cmd.Execute()
	if ret.version {
		ShowVersionInfo()
		os.Exit(0)
	}
	b, _ := json.Marshal(ret)
	fmt.Println(string(b))
	return ret
}

func ShowVersionInfo() {
	fmt.Printf("nirvana-practice, %s\n", info.Info())
}
