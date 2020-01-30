/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package bot

import (
	"github.com/go-pg/pg/v9"
	"github.com/kallydev/yogurt/common/config"
	"github.com/kallydev/yogurt/common/database"
	"github.com/kallydev/yogurt/service/bot/core"
	_ "github.com/lib/pq"
	"log"
)

var (
	Conf    *config.Config
	DB      *pg.DB
	TeleBot *core.Bot
	Token   string
)

const confPath = "config/config_service-bot.yaml"

func init() {
	var err error
	if Conf, err = config.ParseConfigFile(confPath); err != nil {
		log.Fatalln(err)
	} else {
		DB = database.DialPostgres(
			Conf.Postgres.Username, Conf.Postgres.Password,
			Conf.Postgres.Host, Conf.Postgres.Port, Conf.Postgres.Database, nil,
		)
	}
	Token = Conf.Telegram.Token
}
