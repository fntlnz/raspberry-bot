package main

import (
	"github.com/codegangsta/cli"
	"github.com/fntlnz/raspberry-bot/commands"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "raspberry-bot"
	app.Author = "Lorenzo Fontana"
	app.Email = "fontanalorenz@gmail.com"
	app.Commands = commands.Commands
	app.Run(os.Args)
}
