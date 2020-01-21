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
	"database/sql"
	"github.com/kallydev/yogurt/common/database"
	"github.com/kallydev/yogurt/service/gateway"
	"golang.org/x/net/context"
)

type Log struct {
	Host      string
	Method    string
	Path      string
	Params    string
	UserAgent string
	IP        string

	database.Table
}

const (
	InsertLogSQL = `INSERT INTO gateway.logs_new(
			host, method, path, params, user_agent, ip
		) VALUES (
			$1, $2, $3, $4, $5, $6
		) RETURNING *`
)

func InsertLog(ctx context.Context, host, method, path, params, userAgent, ip string) (*Log, error) {
	var log Log
	if err := scanLog(gateway.DB.QueryRowContext(ctx, InsertLogSQL,
		host, method, path, params, userAgent, ip,
	), &log); err != nil {
		return nil, err
	}
	return &log, nil
}

func scanLog(row *sql.Row, log *Log) error {
	if err := row.Scan(
		&log.ID, &log.Method, &log.Host, &log.Path, &log.Params, &log.UserAgent, &log.IP,
		&log.CreatedAt, &log.UpdatedAt, &log.DeletedAt,
	); err != nil {
		return err
	}
	return nil
}

func scanLogs(rows *sql.Rows, logs *[]Log) error {
	for rows.Next() {
		log := Log{}
		if err := rows.Scan(
			&log.ID, &log.Method, &log.Host, &log.Path, &log.Params, &log.UserAgent, &log.IP,
			&log.CreatedAt, &log.UpdatedAt, &log.DeletedAt,
		); err != nil {
			return err
		}
		*logs = append(*logs, log)
	}
	return nil
}
