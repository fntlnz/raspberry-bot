package configuration

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	TelegramSources []Telegram `json:"telegram_sources"`
}

func ParseFile(path string) Configuration {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatalf("An error occurred opening configuration file: %s", err.Error())
	}
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatalf("An error occurred decoding configuration: %s", err.Error())
	}
	return configuration
}
