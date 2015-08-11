package commands

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/fntlnz/raspberry-bot/actions/sysinfo"
	"github.com/fntlnz/raspberry-bot/configuration"
	"github.com/fntlnz/raspberry-bot/sources"
	"github.com/fntlnz/raspberry-bot/sources/telegram"
)

func cmdStart(c *cli.Context) {
	conf := configuration.ParseFile(c.String("configuration"))
	availableCommands := map[string]func() (string, error){
		"ip":     sysinfo.IPAddress,
		"status": sysinfo.SystemStatus,
	}

	// I'm working on making this source indipendent
	for _, source := range conf.TelegramSources {
		telegram := telegram.NewSource(source.Token, source.AllowedUsers)
		log.Printf("Initializing source %s", telegram.Name())
		go telegram.WaitUpdates()
		go telegram.WaitFeedback()
	}

	for update := range sources.Updates() {
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
		sources.Feedback() <- &sources.Message{
			SourceName: update.SourceName,
			Sender:     update.Sender,
			Text:       msg,
		}
	}
}
