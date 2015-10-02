package commands

import "github.com/fntlnz/raspberry-bot/Godeps/_workspace/src/github.com/codegangsta/cli"

var Commands = []cli.Command{
	{
		Name:   "start",
		Usage:  "Start the bot",
		Action: cmdStart,
		Flags: []cli.Flag{
			cli.StringFlag{
				EnvVar: "BOT_CONFIGURATION",
				Name:   "configuration, c",
				Usage:  "Configuration file path",
				Value:  "configuration.json",
			},
		},
	},
}
