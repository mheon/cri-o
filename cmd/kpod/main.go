package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	cli.HelpFlag = cli.BoolFlag{
		Name:  "help",
		Usage: "Show help",
	}
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version",
		Usage: "Print version information and exit",
	}

	app := cli.NewApp()
	app.Name = "kpod"
	app.Usage = "manage pods and images"
	app.Version = kpodVersion

	app.Commands = []cli.Command{
		runCommand,
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
