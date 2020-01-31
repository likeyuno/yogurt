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
	"github.com/go-pg/pg/v9"
	"github.com/kallydev/yogurt/common/database"
)

type Package struct {
	tableName   struct{} `pg:"public.packages"`
	ID          string
	Name        string
	Description string
	Nodes       []string `pg:",array"`
	Money       int
	Day         int
	Traffic     int
	Device      int

	database.Table
}

func QueryPackageByName(db *pg.DB, name string) (*Package, error) {
	var _package = Package{}
	err := db.Model(&_package).Where("name = ?", name).Select()
	return &_package, err
}
