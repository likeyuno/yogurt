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

type Node struct {
	tableName struct{} `pg:"public.nodes_new"`

	Name        string
	Description string
	Tags        []string `pg:",array"`
	Location    string
	Type        string
	FastOpen    bool

	database.Table
}

func QueryNodeByID(id string) (*Node, error) {
	node := Node{}
	err := api.DB.Model(&node).Where("id = ?", id).Order("name ASC").Select()
	return &node, err
}

// SELECT * FROM public.nodes WHERE array_position($1::uuid[], id) NOTNULL AND deleted_at IS NULL ORDER BY name
func QueryNodeByIDs(ids []string) ([]Node, error) {
	var nodes []Node
	err := api.DB.Model(&nodes).Where("id in (?)", pg.In(ids)).Select()
	return nodes, err
}
