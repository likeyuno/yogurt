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
	"strconv"
	"time"
)

func Subscription(b *tb.BotAPI, update tb.Update) {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
		userID    = update.Message.From.ID
	)
	if update.Message.Chat.Type != "private" {
		if err := subscriptionChatIsGroup(b, update); err != nil {
			log.Println(err)
		}
		return
	}
	if acc, err := table.QueryAccountByTelegram(bot.DB, strconv.Itoa(userID)); err != nil {
		if err == pg.ErrNoRows {
			if err := subscriptionNotBind(b, update); err != nil {
				log.Println(err)
			}
			return
		}
		log.Println(err)
		if err := bindError(b, update); err != nil {
			log.Println(err)
		}
	} else if subs, err := table.QuerySubscriptionsByUsername(bot.DB, acc.Username); err != nil {
		if err == pg.ErrNoRows {
			if err := subscriptionEmpty(b, update); err != nil {
				log.Println(err)
			}
			return
		}
		log.Println(err)
		if err := bindError(b, update); err != nil {
			log.Println(err)
		}
	} else {
		result := ""
		for i, sub := range subs {
			result += fmt.Sprintf("订阅 Key %s\n套餐类型 %s\n到期时间 %s", sub.Key, sub.Package, sub.ExpireAt.Format("2006-01-02"))
			if i != len(subs)-1 {
				result += "\n\n"
			}
		}
		msg := tb.NewMessage(chatID, message.SubscriptionListMessage(result))
		msg.ParseMode = "Markdown"
		msg.ReplyToMessageID = messageID
		if _, err := b.Send(msg); err != nil {
			log.Println(err)
		}
	}
}

func subscriptionChatIsGroup(b *tb.BotAPI, update tb.Update) error {
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

func subscriptionNotBind(b *tb.BotAPI, update tb.Update) error {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
	)
	msg := tb.NewMessage(chatID, message.AccountNotBindMessage())
	msg.ReplyToMessageID = messageID
	msg.ParseMode = "Markdown"
	if _, err := b.Send(msg); err != nil {
		return err
	}
	return nil
}

func subscriptionEmpty(b *tb.BotAPI, update tb.Update) error {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
	)
	msg := tb.NewMessage(chatID, message.SubscriptionEmptyMessage())
	msg.ReplyToMessageID = messageID
	msg.ParseMode = "Markdown"
	if _, err := b.Send(msg); err != nil {
		return err
	}
	return nil
}
