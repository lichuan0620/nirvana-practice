package main

import (
	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/log"

	"github.com/lichuan0620/nirvana-practice/cmd/server/options"
	"github.com/lichuan0620/nirvana-practice/pkg/apis"
)

func main() {
	// parse command line arguments
	cmd := options.ParseCommands()

	// initialize Server config
	config := nirvana.NewDefaultConfig().Configure(nirvana.Port(cmd.HTTPPort))

	// install APIs
	apis.Install(config)

	// create the server and server
	server := nirvana.NewServer(config)
	if err := server.Serve(); err != nil {
		log.Errorf("server failed with error: %s", err.Error())
	}
}
