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

type NodeShadowsocksR struct {
	tableName struct{} `pg:"public.nodes_shadowsocksr"`

	NodeID           string
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

func QueryNodeShadowsocksRByIDs(ids []string) ([]Node, error) {
	var nssrs []Node
	err := api.DB.Model(&nssrs).
		Relation("NodeShadowsocksR").
		Order("name ASC").
		Select()
	return nssrs, err
}
