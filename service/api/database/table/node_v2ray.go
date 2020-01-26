package table

import (
	"github.com/kallydev/yogurt/common/database"
	"github.com/kallydev/yogurt/service/api"
)

type NodeV2Ray struct {
	tableName struct{} `pg:"public.nodes_v2ray"`

	NodeID          string
	Host            string
	Port            string
	Security        string
	AlertID         string
	TLS             bool
	TLSHost         string
	TLSVerification bool
	Obfuscation     string
	ObfuscationHost string
	ObfuscationPath string

	database.Table
}

func QueryNodeV2RayByIDs(ids []string) ([]Node, error) {
	var ns []Node
	err := api.DB.Model(&ns).
		Relation("NodeV2Ray").
		Order("name ASC").
		Select()
	return ns, err
}
