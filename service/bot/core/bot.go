/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package core

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	*tg.BotAPI
}

func NewBot(token string) (*Bot, error) {
	b, err := tg.NewBotAPI(token)
	return &Bot{
		BotAPI: b,
	}, err
}

func (b *Bot) Run() (tg.UpdatesChannel, error) {
	u := tg.NewUpdate(0)
	u.Timeout = 60
	return b.GetUpdatesChan(u)
}
