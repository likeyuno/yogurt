/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package static

import (
	_ "github.com/lib/pq"
	"log"
	"xorm.io/xorm"
)

var (
	Conf *Config
	DB   *xorm.Engine
)

const confPath = "service/static/config/config.yaml"

func init() {
	var err error
	if Conf, err = ParseConfigFile(confPath); err != nil {
		log.Fatalln(err)
	} else if DB, err = DialPostgres(
		Conf.Postgres.Schema, Conf.Postgres.Username, Conf.Postgres.Password,
		Conf.Postgres.Host, Conf.Postgres.Port, Conf.Postgres.Database, nil,
	); err != nil {
		log.Fatalln(err)
	}
}
