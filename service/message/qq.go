package message

import (
	"fmt"
	"github.com/catsworld/qq-bot-api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

type QQBot struct {
	*qqbotapi.BotAPI
}

func NewQQBot() (*QQBot, error) {
	bot, err := qqbotapi.NewBotAPI("", "ws://localhost:6700", "")
	return &QQBot{
		BotAPI: bot,
	}, err
}

func (bot *QQBot) Run() error {
	bot.Debug = true
	u := qqbotapi.NewUpdate(0)
	u.PreloadUserInfo = true
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}
	for update := range updates {
		if TGBotInstance == nil {
			continue
		}
		if update.Message == nil {
			continue
		}
		if update.GroupID != QQGroupNumber {
			log.Println(update)
			continue
		}
		name := update.Message.From.NickName
		message := strings.Split(update.Message.Text, "[CQ")[0]
		_, err := TGBotInstance.Send(
			tgbotapi.NewMessage(TGGroupNumber, fmt.Sprintf("[%s]:\n%s", name, message)),
		)
		if err != nil {
			log.Println(err)
			continue
		}
		tempURL := strings.Split(update.Message.Message.CQString(), "url=")
		if len(tempURL) > 1 {
			URL := tempURL[1][:len(tempURL[1])-1]
			_, err := TGBotInstance.Send(tgbotapi.NewPhotoShare(TGGroupNumber, URL))
			if err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}
