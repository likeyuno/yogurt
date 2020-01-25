/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package table

import (
	"github.com/kallydev/yogurt/common/database"
	"github.com/kallydev/yogurt/service/gateway"
)

type Log struct {
	tableName struct{} `pg:"gateway.logs"`

	Host      string
	Method    string
	Path      string
	Params    string
	UserAgent string
	IP        string

	database.Table
}

func InsertLog(host, method, path, params, userAgent, ip string) (*Log, error) {
	var log = Log{
		Host:      host,
		Method:    method,
		Path:      path,
		Params:    params,
		UserAgent: userAgent,
		IP:        ip,
	}
	if _, err := gateway.DB.Model(&log).Returning("*").Insert(); err != nil {
		return nil, err
	}
	return &log, nil
}
