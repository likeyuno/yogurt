/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package handler

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kallydev/yogurt/service/bot/database/table"
	"log"
	"strconv"
	"time"
)

func GetSubscriptionHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
		userID    = update.Message.From.ID
		nickname  = update.Message.From.LastName + update.Message.From.FirstName
	)
	if update.Message.Chat.Type != "private" {
		bot.DeleteMessage(tgbotapi.DeleteMessageConfig{ChatID: chatID, MessageID: messageID})
		message, _ := bot.Send(ChatNotPrivateMessage(chatID, 0, nickname, userID))
		go func(bot *tgbotapi.BotAPI, chatID int64, messageID int) {
			time.Sleep(time.Second * 10)
			if _, err := bot.DeleteMessage(tgbotapi.NewDeleteMessage(chatID, messageID)); err != nil {
				log.Println(err)
			}
		}(bot, chatID, message.MessageID)
		return
	}

	// Get subscription
	if acc, err := table.QueryAccountByTelegram(strconv.Itoa(userID)); err != nil {
		if err == pg.ErrNoRows {
			bot.Send(AccountNotBindMessage(chatID, messageID))
			return
		}
		log.Println(err)
		bot.Send(ErrorMessage(chatID, messageID))
	} else if subs, err := table.QuerySubscriptionsByUsername(acc.Username); err != nil {
		if err == pg.ErrNoRows {
			bot.Send(SubscriptionNotHaveMessage(chatID, messageID))
			return
		}
		log.Println(err)
		bot.Send(ErrorMessage(chatID, messageID))
	} else {
		resutl := ""
		for i, sub := range subs {
			resutl += fmt.Sprintf("订阅 Key %s\n套餐类型 %s\n到期时间 %s", sub.Key, sub.Package,sub.ExpireAt.Format("2006-01-02"))
			if i != len(subs) - 1 {
				resutl += "\n\n"
			}
		}
		bot.Send(SubscriptionListMessage(chatID, messageID, resutl))
	}

}
