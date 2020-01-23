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
	"github.com/kallydev/yogurt/service/api"
)

type NodeShadowsocksR struct {
	tableName struct{} `pg:"public.nodes_shadowsocksr"`

	Host             string
	Port             string
	Password         string
	Method           string
	Protocol         string
	ProtocolParam    string
	Obfuscation      string
	ObfuscationParam string
	Plugin           string
	PluginParam      string

	database.Table
}

func QueryNodeShadowsocksRByIDs(ids []string) ([]NodeShadowsocksR, error) {
	var nssrs []NodeShadowsocksR
	err := api.DB.Model(&nssrs).Where("id in (?)", pg.In(ids)).Order("host ASC").Select()
	return nssrs, err
}
