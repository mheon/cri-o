package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: "launch a container",
	Action: func(context *cli.Context) error {
		return fmt.Errorf("this functionality is not yet implemented")
	},
}
