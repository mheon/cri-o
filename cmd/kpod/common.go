package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/kubernetes-incubator/cri-o/server"
	"github.com/urfave/cli"
)

const (
	kpodVersion         = "0.0.1"
	defaultHostname     = "kpod-launch"
	defaultCgroupParent = "/kpod-launch"
)

var commonFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "config",
		Usage: "Specify configuration `file`",
		Value: server.CRIOConfigPath,
	},
	cli.BoolFlag{
		Name:  "debug, D",
		Usage: "Enable debug logging",
	},
}

// Does not parse config flag to obtain configuration file path, as that is done elsewhere
// (Usually when creating a runtime)
func parseCommonFlags(c *cli.Context) error {
	if c.GlobalBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	return nil
}

// Runtime for executing kpod commands
type kpodRuntime struct {
	crioConfig *server.Config
	crioServer *server.Server
}

// Get a new kpod runtime from a given configuration file
func getNewRuntime(configPath string) (*kpodRuntime, error) {
	// TODO make crio's mergeConfig more generic and use it here
	config := new(server.Config)

	if err := config.FromFile(configPath); err != nil {
		return nil, fmt.Errorf("error retrieving kpod config: %v", err)
	}

	server, err := server.New(config)
	if err != nil {
		return nil, fmt.Errorf("error creating CRI-O server: %v", err)
	}

	runtime := new(kpodRuntime)
	runtime.crioConfig = config
	runtime.crioServer = server

	return runtime, nil
}
