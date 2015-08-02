package commands

import "github.com/codegangsta/cli"

var Commands = []cli.Command{
	{
		Name:   "start",
		Usage:  "Start the bot",
		Action: cmdRun,
	},
}
