package main

import (
	"fmt"
	"os"

	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/log"
	"github.com/caicloud/nirvana/rest"
	"github.com/spf13/pflag"

	"github.com/lichuan0620/nirvana-practice/client"
	"github.com/lichuan0620/nirvana-practice/pkg/apis"
	"github.com/lichuan0620/nirvana-practice/pkg/info"
)

func main() {
	var (
		httpPort uint16
		version  bool
	)

	pflag.Uint16VarP(&httpPort, "port", "p", 8080, "the HTTP port used by the server")
	pflag.BoolVarP(&version, "version", "v", false, "show version info")
	pflag.Parse()

	if version {
		fmt.Printf("practice-server, %s\n", info.Info())
		os.Exit(0)
	}

	cli := client.MustNewClient(&rest.Config{
		Scheme: "http",
		Host:   "localhost:8081",
	})

	// initialize Server config
	config := nirvana.NewDefaultConfig().Configure(nirvana.Port(httpPort))

	// install APIs
	apis.Install(config, cli)

	// create the server and server
	server := nirvana.NewServer(config)
	if err := server.Serve(); err != nil {
		log.Errorf("server failed with error: %s", err.Error())
	}
}
