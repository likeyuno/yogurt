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
	tb "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kallydev/yogurt/service/bot"
	"github.com/kallydev/yogurt/service/bot/database/table"
	"github.com/kallydev/yogurt/service/bot/message"
	"log"
	"strings"
	"time"
)

func Create(b *tb.BotAPI, update tb.Update) {
	if update.Message.Chat.Type != "private" {
		if err := createChatIsGroup(b, update); err != nil {
			log.Println(err)
		}
		return
	}
	if update.Message.From.ID != bot.Conf.Telegram.Master {
		if err := createNotMaster(b, update); err != nil {
			log.Println(err)
		}
		return
	}
	param := update.Message.CommandArguments()
	if param == "" || len(strings.Split(param, " ")) < 3 {
		if err := createParamError(b, update); err != nil {
			log.Println(err)
		}
		return
	}
	params := strings.Split(param, " ")
	switch params[0] {
	case "account":
		var (
			username = params[1]
			email    = params[2]
			qq       = params[3]
		)
		if _, err := table.InsertAccount(bot.DB, username, email, qq); err != nil {
			log.Println(err)
			if err := createError(b, update); err != nil {
				log.Println(err)
			}
		} else {
			if err := createAccountSuccess(b, update); err != nil {
				log.Println(err)
			}
		}
	case "subscription":
		var (
			account  = params[1]
			_package = params[2]
		)
		if pack, err := table.QueryPackageByName(bot.DB, _package); err != nil {
			log.Println(err)
			if err := createPackageNotFound(b, update); err != nil {
				log.Println(err)
			}
		} else if d, err := time.ParseDuration(fmt.Sprintf("%dh", pack.Day*24)); err != nil {
			log.Println(err)
			if err := createError(b, update); err != nil {
				log.Println(err)
			}
		} else if sub, err := table.InsertSubscription(bot.DB, account, _package, d); err != nil {
			log.Println(err)
			if err := createError(b, update); err != nil {
				log.Println(err)
			}
		} else {
			if err := createSubscriptionSuccess(b, update, sub.Key, sub.ExpireAt.Format("2006-01-02")); err != nil {
				log.Println(err)
			}
		}
	default:
		if err := createParamError(b, update); err != nil {
			log.Println(err)
		}
		return
	}
}

func createChatIsGroup(b *tb.BotAPI, update tb.Update) error {
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

func createParamError(b *tb.BotAPI, update tb.Update) error {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
	)
	msg := tb.NewMessage(chatID, message.CreateParamErrorMessage())
	msg.ReplyToMessageID = messageID
	msg.ParseMode = "Markdown"
	if _, err := b.Send(msg); err != nil {
		return err
	}
	return nil
}

func createError(b *tb.BotAPI, update tb.Update) error {
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

func createAccountSuccess(b *tb.BotAPI, update tb.Update) error {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
	)
	msg := tb.NewMessage(chatID, message.CreateAccountSuccessMessage())
	msg.ReplyToMessageID = messageID
	msg.ParseMode = "Markdown"
	if _, err := b.Send(msg); err != nil {
		return err
	}
	return nil
}

func createSubscriptionSuccess(b *tb.BotAPI, update tb.Update, key, expiredAt string) error {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
	)
	msg := tb.NewMessage(chatID, message.CreateSubscriptionSuccessMessage(key, expiredAt))
	msg.ReplyToMessageID = messageID
	msg.ParseMode = "Markdown"
	if _, err := b.Send(msg); err != nil {
		return err
	}
	return nil
}

func createNotMaster(b *tb.BotAPI, update tb.Update) error {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
	)
	msg := tb.NewMessage(chatID, message.NotIsMasterMessage())
	msg.ReplyToMessageID = messageID
	msg.ParseMode = "Markdown"
	if _, err := b.Send(msg); err != nil {
		return err
	}
	return nil
}

func createPackageNotFound(b *tb.BotAPI, update tb.Update) error {
	var (
		chatID    = update.Message.Chat.ID
		messageID = update.Message.MessageID
	)
	msg := tb.NewMessage(chatID, message.PackageNotFoundMessage())
	msg.ReplyToMessageID = messageID
	msg.ParseMode = "Markdown"
	if _, err := b.Send(msg); err != nil {
		return err
	}
	return nil
}
