package telegram

import (
	"log"

	"github.com/Syfaro/telegram-bot-api"
	"github.com/fntlnz/raspberry-bot/sources"
)

var updates = make(chan *sources.Message)
var feedback = make(chan *sources.Message)

type Source struct {
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
		Bot:          bot,
		AllowedUsers: allowedUsers,
	}
}

func (s *Source) SourceName() string {
	return "telegram"
}

func (s *Source) Updates() <-chan *sources.Message {
	return updates
}

func (s *Source) Feedback() chan<- *sources.Message {
	return feedback
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
	for feed := range feedback {
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
	updates <- &sources.Message{
		Sender: update.Message.From,
		Text:   update.Message.Text,
	}
}
