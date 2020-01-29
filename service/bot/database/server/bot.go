/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package server

import (
	"gopkg.in/tucnak/telebot.v2"
	"time"
)

type Bot struct {
	*telebot.Bot
}

func NewBot(token string) (*Bot, error) {
	b, err := telebot.NewBot(telebot.Settings{
		Token: token,
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	})
	return &Bot{
		Bot: b,
	}, err
}
