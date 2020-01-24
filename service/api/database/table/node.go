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
	"github.com/go-pg/pg/v9/orm"
	"github.com/kallydev/yogurt/common/database"
	"github.com/kallydev/yogurt/service/api"
)

type Node struct {
	tableName struct{} `pg:"public.nodes"`

	Name        string
	Description string
	Tags        []string `pg:",array"`
	Location    string
	Type        string
	FastOpen    bool

	NodeShadowsocksR *NodeShadowsocksR
	NodeV2Ray        *NodeV2Ray `pg:"nodes_v2ray"`

	database.Table
}

func QueryNodeByID(id string) (*Node, error) {
	node := Node{}
	err := api.DB.Model(&node).Where("id = ?", id).Order("name ASC").Select()
	return &node, err
}

func QueryNodeByIDsAndType(_type string, ids []string) ([]Node, error) {
	var ns []Node
	q := api.DB.Model(&ns)
	switch _type {
	case "shadowsocksr":
		q = q.Relation("NodeShadowsocksR", func(query2 *orm.Query) (query *orm.Query, err error) {
			return query2.Where("node_shadowsocksr.node_id in (?)", pg.In(ids)), nil
		})
	case "vmess":
		q = q.Relation("NodeV2Ray", func(query2 *orm.Query) (query *orm.Query, err error) {
			return query2.Where("nodes_v2ray.node_id in (?)", pg.In(ids)), nil
		})
	}
	err := q.Where("type = ?", _type).
		Order("name ASC").
		Select()
	return ns, err
}
