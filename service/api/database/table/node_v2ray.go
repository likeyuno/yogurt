package table

import "github.com/kallydev/yogurt/common/database"

type NodeV2Ray struct {
	tableName struct{} `pg:"public.nodes_v2ray"`

	Host     string
	Port     string
	UUID     string
	Security string

	database.Table
}
