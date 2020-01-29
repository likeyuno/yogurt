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
	"time"
)

func GetSubscriptionLinkHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
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
	key := update.Message.CommandArguments()
	if key == "" {
		bot.Send(GetSubscriptionLinEmptyMessage(chatID, messageID))
		return
	}
	if sub, err := table.QuerySubscriptionByKey(key); err != nil {
		if err == pg.ErrNoRows {
			bot.Send(SubscriptionNotFoundMessage(chatID, messageID))
			return
		}
		log.Println(err)
		bot.Send(ErrorMessage(chatID, messageID))
	} else {
		var url = "https://api.yogurtcloud.com/v1/subscriptions/" + sub.Key
		result := fmt.Sprintf("- SSR 通用訂閱鏈接\n%s%s\n\n", url, "")
		result += fmt.Sprintf("- Vmess V2RayNG / Shadorocket 訂閱鏈接\n%s%s\n\n", url, "?protocol=vmess")
		result += fmt.Sprintf("- Vmess Quantumult 訂閱鏈接\n%s%s\n\n", url, "?protocol=vmess&client=quantumult")
		result += fmt.Sprintf("- Vmess QuantumultX 訂閱鏈接\n%s%s\n\n", url, "?protocol=vmess&client=quantumultx")
		result += fmt.Sprintf("- Vmess Netch 訂閱鏈接\n%s%s", url, "?protocol=vmess&client=netch")
		bot.Send(GetSubscriptionLinkMessage(chatID, messageID, result))
	}
}
