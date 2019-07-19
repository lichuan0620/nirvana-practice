package main

import (
	"fmt"
	"os"

	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/log"
	"github.com/spf13/pflag"

	"github.com/lichuan0620/nirvana-practice/pkg/apis"
	"github.com/lichuan0620/nirvana-practice/pkg/info"
)

var (
	httpPort uint16
	version  bool
)

func init() {
	pflag.Uint16VarP(&httpPort, "port", "p", 8081, "the HTTP port used by the server")
	pflag.BoolVarP(&version, "version", "v", false, "show version info")
	pflag.Parse()
}

func main() {
	if version {
		fmt.Printf("practice-cache, %s\n", info.Info())
		os.Exit(0)
	}

	// initialize Server config
	config := nirvana.NewDefaultConfig().Configure(nirvana.Port(httpPort))

	// install APIs
	apis.CacheInstall(config)

	// create the server and server
	server := nirvana.NewServer(config)
	if err := server.Serve(); err != nil {
		log.Errorf("server failed with error: %s", err.Error())
	}
}
