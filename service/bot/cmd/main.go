/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	bot2 "github.com/kallydev/yogurt/service/bot"
	"github.com/kallydev/yogurt/service/bot/database/handler"
	"log"
	"time"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(bot2.Token)
	if err != nil {
		log.Println(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Println(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.LeftChatMember != nil {
			continue
		}
		if update.Message.NewChatMembers != nil {
			handler.JoinHandler(bot, update)
			continue
		}
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				handler.StartHandler(bot, update)
			case "bind":
				handler.BindHandler(bot, update)
			case "subscription":
				handler.GetSubscriptionHandler(bot, update)
			case "link":
				handler.GetSubscriptionLinkHandler(bot, update)
			default:
				if update.Message.Chat.Type == "private" {
					bot.Send(handler.DefaultMessage(update.Message.Chat.ID, update.Message.MessageID))
				} else {
					go func(bot *tgbotapi.BotAPI, chatID int64, messageID int) {
						time.Sleep(time.Second * 5)
						if _, err := bot.DeleteMessage(tgbotapi.NewDeleteMessage(chatID, messageID)); err != nil {
							log.Println(err)
						}
					}(bot, update.Message.Chat.ID, update.Message.MessageID)
				}
			}
			continue
		}
		handler.MessageHandler(bot, update)
	}
}
