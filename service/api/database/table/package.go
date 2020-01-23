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
	"github.com/kallydev/yogurt/service/api"
)

type Package struct {
	tableName struct{} `pg:"public.packages_new"`

	Name        string
	Description string
	Nodes       []string `pg:",array"`
	Money       int
	Day         int
	Traffic     int
	Device      int

	database.Table
}

func QueryPackageByID(id string) (*Package, error) {
	pack := Package{}
	err := api.DB.Model(&pack).Where("id = ?", id).Select()
	return &pack, err
}

func QueryPackageByName(name string) (*Package, error) {
	pack := Package{}
	err := api.DB.Model(&pack).Where("name = ?", name).Select()
	return &pack, err
}
