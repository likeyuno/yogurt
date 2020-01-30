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
	"fmt"
	"github.com/go-pg/pg/v9"
	tb "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kallydev/yogurt/service/bot"
	"github.com/kallydev/yogurt/service/bot/database/table"
	"github.com/kallydev/yogurt/service/bot/message"
	"log"
	"time"
)

func Link(b *tb.BotAPI, update tb.Update) {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
	)
	if update.Message.Chat.Type != "private" {
		if err := linkChatIsGroup(b, update); err != nil {
			log.Println(err)
		}
		return
	}
	if key := update.Message.CommandArguments(); key == "" {
		if err := linkKeyIsEmpty(b, update); err != nil {
			log.Println(err)
		}
	} else if sub, err := table.QuerySubscriptionByKey(bot.DB, key); err != nil {
		if err == pg.ErrNoRows {
			if err := linkNotFound(b, update); err != nil {
				log.Println(err)
			}
			return
		}
		log.Println(err)
		if err := bindError(b, update); err != nil {
			log.Println(err)
		}
	} else {
		var url = "https://api.yogurtcloud.com/v1/subscriptions/" + sub.Key
		result := fmt.Sprintf("- SSR 通用訂閱鏈接\n%s%s\n\n", url, "")
		result += fmt.Sprintf("- Vmess V2RayNG / Shadorocket 訂閱鏈接\n%s%s\n\n", url, "?protocol=vmess")
		result += fmt.Sprintf("- Vmess Quantumult 訂閱鏈接\n%s%s\n\n", url, "?protocol=vmess&client=quantumult")
		result += fmt.Sprintf("- Vmess QuantumultX 訂閱鏈接\n%s%s\n\n", url, "?protocol=vmess&client=quantumultx")
		result += fmt.Sprintf("- Vmess Netch 訂閱鏈接\n%s%s", url, "?protocol=vmess&client=netch")
		msg := tb.NewMessage(chatID, message.GetSubscriptionLinkMessage(result))
		msg.ReplyToMessageID = messageID
		msg.ParseMode = "Markdown"
		if _, err := b.Send(msg); err != nil {
			log.Println(err)
		}
	}
}

func linkChatIsGroup(b *tb.BotAPI, update tb.Update) error {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
		nickname  = update.Message.From.FirstName + update.Message.From.LastName
		userID    = update.Message.From.ID
	)
	if _, err := b.DeleteMessage(tb.NewDeleteMessage(chatID, messageID)); err != nil {
		return err
	}
	if msg, err := b.Send(tb.NewMessage(chatID, message.ChatNotPrivateMessage(nickname, userID))); err != nil {
		return err
	} else {
		go func() {
			time.Sleep(time.Second * 10)
			if _, err := b.DeleteMessage(tb.NewDeleteMessage(chatID, msg.MessageID)); err != nil {
				log.Println(err)
			}
		}()
	}
	return nil
}

func linkKeyIsEmpty(b *tb.BotAPI, update tb.Update) error {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
	)
	msg := tb.NewMessage(chatID, message.LinkKeyEmptyMessage())
	msg.ReplyToMessageID = messageID
	msg.ParseMode = "Markdown"
	if _, err := b.Send(msg); err != nil {
		return err
	}
	return nil
}

func linkNotFound(b *tb.BotAPI, update tb.Update) error {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
	)
	msg := tb.NewMessage(chatID, message.SubscriptionNotFoundMessage())
	msg.ReplyToMessageID = messageID
	msg.ParseMode = "Markdown"
	if _, err := b.Send(msg); err != nil {
		return err
	}
	return nil
}
