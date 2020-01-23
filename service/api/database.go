/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package api

import (
	"github.com/go-pg/pg/v9"
	"net"
	"strconv"
)

func DialPostgres(username, password, host string, port int, database string, options map[string]string) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     net.JoinHostPort(host, strconv.Itoa(port)),
		User:     username,
		Password: password,
		Database: database,
	})
}
