package commands

import "github.com/codegangsta/cli"

var Commands = []cli.Command{
	{
		Name:   "example",
		Usage:  "boh",
		Action: cmdRun,
	},
}
