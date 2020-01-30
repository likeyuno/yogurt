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
	"github.com/go-pg/pg/v9"
	tb "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kallydev/yogurt/service/bot"
	"github.com/kallydev/yogurt/service/bot/database/table"
	"github.com/kallydev/yogurt/service/bot/message"
	"log"
	"strconv"
	"time"
)

func Bind(b *tb.BotAPI, update tb.Update) {
	if update.Message.Chat.Type != "private" {
		if err := bindChatIsGroup(b, update); err != nil {
			log.Println(err)
		}
		return
	}
	if key := update.Message.CommandArguments(); key == "" {
		if err := bindKeyIsEmpty(b, update); err != nil {
			log.Println(err)
		}
	} else if sub, err := table.QuerySubscriptionByKey(bot.DB, key); err != nil {
		if err == pg.ErrNoRows {
			if err := bindSubscriptionNotFound(b, update); err != nil {
				log.Println(err)
			}
			return
		}
		log.Println(err)
		if err := bindError(b, update); err != nil {
			log.Println(err)
		}
	} else if acc, err := table.QueryAccountByUsername(bot.DB, sub.Account); err != nil {
		if err == pg.ErrNoRows {
			if err := bindSubscriptionNotFound(b, update); err != nil {
				log.Println(err)
			}
			return
		}
		log.Println(err)
		if err := bindError(b, update); err != nil {
			log.Println(err)
		}
	}else if acc.Telegram != "" {
		if err := bindAlready(b, update); err != nil {
			log.Println(err)
		}
	} else if _, err := table.UpdateAccountTelegramByUsername(bot.DB, acc.Username, strconv.Itoa(update.Message.From.ID)); err != nil {
		log.Println(err)
		if err := bindError(b, update); err != nil {
			log.Println(err)
		}
	} else {
		if err := bindSuccess(b, update,acc.Username); err != nil {
			log.Println(err)
		}
	}
}

func bindChatIsGroup(b *tb.BotAPI, update tb.Update) error {
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

func bindKeyIsEmpty(b *tb.BotAPI, update tb.Update) error {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
	)
	msg := tb.NewMessage(chatID, message.BindKeyEmptyMessage())
	msg.ReplyToMessageID = messageID
	msg.ParseMode = "Markdown"
	if _, err := b.Send(msg); err != nil {
		return err
	}
	return nil
}

func bindSubscriptionNotFound(b *tb.BotAPI, update tb.Update) error {
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

func bindError(b *tb.BotAPI, update tb.Update) error {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
	)
	msg := tb.NewMessage(chatID, message.ErrorMessage())
	msg.ReplyToMessageID = messageID
	msg.ParseMode = "Markdown"
	if _, err := b.Send(msg); err != nil {
		return err
	}
	return nil
}

func bindAlready(b *tb.BotAPI, update tb.Update) error {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
	)
	msg := tb.NewMessage(chatID, message.AccountAlreadyBindMessage())
	msg.ReplyToMessageID = messageID
	msg.ParseMode = "Markdown"
	if _, err := b.Send(msg); err != nil {
		return err
	}
	return nil
}

func bindSuccess(b *tb.BotAPI, update tb.Update, username string) error {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
	)
	msg := tb.NewMessage(chatID, message.AccountBindSuccessMessage(username))
	msg.ReplyToMessageID = messageID
	msg.ParseMode = "Markdown"
	if _, err := b.Send(msg); err != nil {
		return err
	}
	return nil
}
