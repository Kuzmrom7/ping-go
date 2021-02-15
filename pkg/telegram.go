package pkg

import (
	"log"

	"github.com/Kuzmrom7/ping-go/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendPhoto(imageBuf []byte, botConfig *config.Bot) error {
	bot, err := tgbotapi.NewBotAPI(botConfig.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	file := tgbotapi.FileBytes{
		Name:  "scr.jpg",
		Bytes: imageBuf,
	}

	msg := tgbotapi.NewPhotoUpload(botConfig.Chat, file)
	_, err = bot.Send(msg)

	return err
}
