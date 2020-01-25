package message

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/image/webp"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type TGBot struct {
	*tgbotapi.BotAPI
}

func NewTGBot(token string) (*TGBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	return &TGBot{
		BotAPI: bot,
	}, err
}

func (bot *TGBot) Run() error {
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}
	for update := range updates {
		if QQBotInstance == nil || update.Message == nil || update.Message.Chat.ID != TGGroupNumber {
			continue
		}
		name := strings.TrimSpace(update.Message.From.FirstName + " " + update.Message.From.LastName)
		_, err := QQBotInstance.SendMessage(QQGroupNumber,
			"group",
			strings.TrimRight(fmt.Sprintf("[%s]:\n%s", name, update.Message.Text), "\n"),
		)
		if err != nil {
			log.Println(err)
		}
		if update.Message.Sticker != nil {
			if u, err := bot.GetFileDirectURL(update.Message.Sticker.FileID); err != nil {
				log.Println(err)
			} else if data, err := downloadFile(u); err != nil {
				log.Println(err)
			} else {
				QQBotInstance.NewMessage(QQGroupNumber, "group").ImageLocal(data).Send()
			}
		}
	}
	return nil
}

func CreatePNGFile(name string) (*os.File, error) {
	return os.Create(name + ".png")
}

func downloadFile(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	buff := bytes.NewBuffer(data)
	s := sha1.New()
	s.Write(data)
	name := hex.EncodeToString(s.Sum(nil))
	file, err := CreatePNGFile(name)
	defer file.Close()
	if m, err := webp.Decode(buff); err != nil {
		log.Fatal(err)
	} else if err := png.Encode(file, m); err != nil {
		log.Fatal(err)
	}
	return ImagePath + name + ".png", nil
}
