/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package command

import (
	tb "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kallydev/yogurt/service/bot/message"
	"log"
	"time"
)

func Start(b *tb.BotAPI, update tb.Update) {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
		userID    = update.Message.From.ID
		nickname  = update.Message.From.LastName + update.Message.From.FirstName
	)
	if update.Message.Chat.Type != "private" {
		if _, err := b.DeleteMessage(tb.NewDeleteMessage(chatID, messageID)); err != nil {
			log.Println(err)
			return
		}
		if msg, err := b.Send(tb.NewMessage(chatID, message.ChatNotPrivateMessage(nickname, userID))); err != nil {
			log.Println(err)
			return
		} else {
			go func() {
				time.Sleep(time.Second * 20)
				if _, err := b.DeleteMessage(tb.NewDeleteMessage(chatID, msg.MessageID)); err != nil {
					log.Println(err)
				}
			}()
		}
	} else {
		msg := tb.NewMessage(chatID, message.StartMessage())
		msg.ReplyToMessageID = messageID
		msg.ParseMode = "Markdown"
		if _, err := b.Send(msg); err != nil {
			log.Println(err)
		}
	}
}
