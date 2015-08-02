package telegram

import (
	"fmt"
	"log"

	"github.com/Syfaro/telegram-bot-api"
	"github.com/fntlnz/raspberry-bot/sources"
	"github.com/fntlnz/raspberry-bot/utils"
)

type Source struct {
	name         string
	Bot          *tgbotapi.BotAPI
	AllowedUsers []int
}

func NewSource(token string, allowedUsers []int) sources.Source {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return &Source{
		name:         fmt.Sprintf("telegram_%s", utils.RandomString(5)),
		Bot:          bot,
		AllowedUsers: allowedUsers,
	}
}

func (s *Source) Type() string {
	return "telegram"
}

func (s *Source) Name() string {
	return s.name
}

func (s *Source) WaitUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	err := s.Bot.UpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range s.Bot.Updates {
		s.handleUpdate(update)
	}
}

func (s *Source) WaitFeedback() {
	for feed := range sources.Feedback() {
		log.Printf("F: %s, S: %s", feed.SourceName, s.Name())
		if feed.SourceName != s.Name() {
			continue
		}
		msg := tgbotapi.NewMessage(feed.Sender.(tgbotapi.User).ID, feed.Text)
		s.Bot.SendMessage(msg)
	}
}

func (s *Source) handleUpdate(update tgbotapi.Update) {
	for _, u := range s.AllowedUsers {
		if update.Message.From.ID != u {
			log.Printf("Ignoring update message from unknown user: [%d]", update.Message.From.ID)
			return
		}
	}
	log.Printf("[%d] %s", update.Message.From.ID, update.Message.Text)
	sources.Updates() <- &sources.Message{
		SourceName: s.Name(),
		Sender:     update.Message.From,
		Text:       update.Message.Text,
	}
}
