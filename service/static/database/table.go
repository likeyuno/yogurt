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
	"time"
)

const Schema = "static"

const (
	Log = "logs"
)

type Table struct {
	ID        string       `xorm:"<- notnull pk UUID 'id'"`
	CreatedAt time.Time    `xorm:"notnull created 'created_at'"`
	UpdatedAt time.Time    `xorm:"notnull updated 'updated_at'"`
	DeletedAt sql.NullTime `xorm:"deleted 'deleted_at'"`
}

func (Table) TableName() string {
	return "logs"
}
