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
	"github.com/go-pg/pg/v9"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kallydev/yogurt/service/bot/database/table"
	"log"
	"regexp"
	"strconv"
	"time"
)

func StartHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
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
	bot.Send(StartMessage(chatID, messageID))
}

func BindHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
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
		bot.Send(AccountBindEmptyMessage(chatID, messageID))
		return
	}
	if sub, err := table.QuerySubscriptionByKey(key); err != nil {
		if err == pg.ErrNoRows {
			bot.Send(SubscriptionNotFoundMessage(chatID, messageID))
			return
		}
		log.Println(err)
		bot.Send(ErrorMessage(chatID, messageID))
	} else if acc, err := table.QueryAccountByUsername(sub.Account); err != nil {
		if err == pg.ErrNoRows {
			bot.Send(AccountNotFoundMessage(chatID, messageID))
			return
		}
		log.Println(err)
		bot.Send(ErrorMessage(chatID, messageID))
	} else if acc.Telegram != "" {
		// Not Empty
		bot.Send(AccountAlreadyBindMessage(chatID, messageID))
	} else if _, err := table.UpdateAccountTelegramByUsername(acc.Username, strconv.Itoa(userID)); err != nil {
		log.Println(err)
		bot.Send(ErrorMessage(chatID, messageID))
	} else {
		bot.Send(AccountBindSuccessMessage(chatID, messageID, acc.Username))
	}
}

func JoinHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var (
		chatID    = update.Message.Chat.ID
		userID    = update.Message.From.ID
		nickname  = update.Message.From.LastName + update.Message.From.FirstName
		groupName = update.Message.Chat.Title
	)
	if ok, err := regexp.MatchString("{InArabic}{0, 1}", nickname); ok || err != nil {
		message, err := bot.Send(ArabicMessage(chatID, 0, userID, nickname))
		if err != nil {
			log.Println(err)
		}
		go func(bot *tgbotapi.BotAPI, chatID int64, messageID int) {
			time.Sleep(time.Second * 10)
			if _, err := bot.DeleteMessage(tgbotapi.NewDeleteMessage(chatID, messageID)); err != nil {
				log.Println(err)
			}
		}(bot, chatID, message.MessageID)
		return
	}
	message, err := bot.Send(WelcomeMessage(chatID, nickname, groupName, userID))
	if err != nil {
		log.Println(err)
	}
	go func(bot *tgbotapi.BotAPI, chatID int64, messageID int) {
		time.Sleep(time.Second * 10)
		if _, err := bot.DeleteMessage(tgbotapi.NewDeleteMessage(chatID, messageID)); err != nil {
			log.Println(err)
		}
	}(bot, chatID, message.MessageID)
}

func MessageHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var (
		chatID = update.Message.Chat.ID
		// username  = update.Message.From.UserName
		nickname  = update.Message.From.LastName + update.Message.From.FirstName
		text      = update.Message.Text
		messageID = update.Message.MessageID
		userID    = update.Message.From.ID
		// groupName = update.Message.Chat.Title
	)
	if update.Message.Chat.Type == "private" {
		bot.Send(DefaultMessage(update.Message.Chat.ID, update.Message.MessageID))
		return
	}
	if ok, err := regexp.MatchString("\\p{Arabic}+", nickname+text); ok || err != nil {
		_, err := bot.Send(ArabicMessage(chatID, messageID, userID, nickname))
		if err != nil {
			log.Println(err)
		}
		if _, err := bot.KickChatMember(tgbotapi.KickChatMemberConfig{
			ChatMemberConfig: tgbotapi.ChatMemberConfig{
				ChatID: chatID,
				UserID: userID,
			},
		}); err != nil {
			log.Println(err)
		}
		return
	}
}
