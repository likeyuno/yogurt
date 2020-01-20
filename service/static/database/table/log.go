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
	"github.com/kallydev/yogurt/service/static/database"
)

type Log struct {
	Host      string `xorm:"'host'"`
	Path      string `xorm:"'path'"`
	Params    string `xorm:"'params'"`
	UserAgent string `xorm:"'user_agent'"`
	IP        string `xorm:"'ip'"`

	database.Table `xorm:"extends"`
}
