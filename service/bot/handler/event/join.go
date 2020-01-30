/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package event

import (
	tb "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kallydev/yogurt/service/bot/core"
	"github.com/kallydev/yogurt/service/bot/message"
	"log"
	"time"
)

func Join(bot *tb.BotAPI, update tb.Update) {
	var (
		nickname  = update.Message.From.FirstName + update.Message.From.LastName
		groupName = update.Message.Chat.Title
		userID    = update.Message.From.ID
		chatID    = update.Message.Chat.ID
	)
	as := tb.NewAnimationShare(update.Message.Chat.ID, core.WelcomeAnimationID)
	as.Caption = message.WelcomeMessage(nickname, groupName, userID, chatID)
	as.ParseMode = "Markdown"
	as.ReplyMarkup = message.JoinKeyboard
	msg, err := bot.Send(as)
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		time.Sleep(time.Second * 3)
		if _, err := bot.DeleteMessage(tb.NewDeleteMessage(chatID, msg.MessageID)); err != nil {
			log.Println(err)
			return
		}
	}()
}
