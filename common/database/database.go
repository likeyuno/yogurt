/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package database

import (
	"database/sql"
	"github.com/kallydev/yogurt/common/context"
	_ "github.com/lib/pq"
	"net"
	"net/url"
	"strconv"
	"time"
)

func DialPostgres(username, password, host string, port int, database string, options map[string]string) (*sql.DB, error) {
	if db, err := sql.Open(createURL("postgres", username, password, host, port, database, options)); err != nil {
		return nil, err
	} else if err := db.PingContext(context.WithTimeoutNoCancel(time.Second * 3)); err != nil {
		return nil, err
	} else {
		return db, nil
	}
}

func createURL(driver, username, password, host string, port int, database string, options map[string]string) (string, string) {
	u := url.URL{
		Scheme:   driver,
		User:     url.UserPassword(username, password),
		Host:     net.JoinHostPort(host, strconv.Itoa(port)),
		Path:     database,
		RawQuery: createValues(options),
	}
	return driver, u.String()
}

func createValues(options map[string]string) string {
	values := url.Values{}
	for k, v := range options {
		values.Add(k, v)
	}
	return values.Encode()
}
