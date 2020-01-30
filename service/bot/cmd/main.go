/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://githubot.com/kallydev/yogurt/blob/master/LICENSE
 */

package main

import (
	"github.com/kallydev/yogurt/service/bot"
	"github.com/kallydev/yogurt/service/bot/core"
	"github.com/kallydev/yogurt/service/bot/handler"
	"log"
)

func main() {
	b, err := core.NewBot(bot.Token)
	if err != nil {
		log.Fatalln(err)
	}
	b.Debug = true
	if updates, err := b.Run(); err != nil {
		log.Fatalln(err)
	} else {
		for update := range updates {
			handler.HandleUpdate(b.BotAPI, update)
		}
	}
}
