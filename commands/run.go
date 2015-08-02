package commands

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/fntlnz/raspberry-bot/actions/sysinfo"
	"github.com/fntlnz/raspberry-bot/configuration"
	"github.com/fntlnz/raspberry-bot/sources"
	"github.com/fntlnz/raspberry-bot/sources/telegram"
)

func cmdRun(c *cli.Context) {
	conf := configuration.ParseFile(c.Args()[0])
	availableCommands := map[string]func() (string, error){
		"ip":     sysinfo.IPAddress,
		"status": sysinfo.SystemStatus,
	}

	// I'm working on making this source indipendent
	for _, source := range conf.TelegramSources {
		telegram := telegram.NewSource(source.Token, source.AllowedUsers)
		log.Printf("Using source %s", telegram.SourceName())
		go telegram.WaitUpdates()
		go telegram.WaitFeedback()

		for update := range telegram.Updates() {
			command, ok := availableCommands[update.Text]
			if !ok {
				log.Printf("Unknown command %s", update.Text)
				continue
			}
			msg, err := command()
			if err != nil {
				log.Printf("%s", err.Error())
				continue
			}
			telegram.Feedback() <- &sources.Message{
				Sender: update.Sender,
				Text:   msg,
			}
		}
	}
}
