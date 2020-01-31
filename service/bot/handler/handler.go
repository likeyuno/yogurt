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
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kallydev/yogurt/service/bot/handler/command"
	"github.com/kallydev/yogurt/service/bot/handler/event"
)

func HandleUpdate(bot *tg.BotAPI, update tg.Update) {
	if update.Message == nil {
		return
	}
	if update.Message.NewChatMembers != nil {
		handleJoin(bot, update)
		return
	}
	if update.Message.LeftChatMember != nil {
		handleLeft(bot, update)
		return
	}
	if update.Message.IsCommand() {
		handleCommand(bot, update)
		return
	}
	handleMessage(bot, update)
}

func handleJoin(bot *tg.BotAPI, update tg.Update) {
	event.Join(bot, update)
}

func handleLeft(bot *tg.BotAPI, update tg.Update) {

}

func handleCommand(bot *tg.BotAPI, update tg.Update) {
	switch update.Message.Command() {
	case "start":
		command.Start(bot, update)
	case "bind":
		command.Bind(bot, update)
	case "subscription":
		command.Subscription(bot, update)
	case "link":
		command.Link(bot, update)
	case "info":
		command.Info(bot, update)
	case "create":
		command.Create(bot, update)
	default:
		//if update.Message.Chat.Type == "private" {
		//	b.Send(message.DefaultMessage(update.Message.Chat.ID, update.Message.MessageID))
		//}
	}
}

func handleMessage(bot *tg.BotAPI, update tg.Update) {

}
