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
	"github.com/kallydev/yogurt/service/gateway"
	"github.com/lib/pq"
)

type Service struct {
	tableName struct{} `pg:"gateway.services"`

	Name    string
	Host    string
	Servers pq.StringArray

	database.Table
}

func QueryAllServices() ([]Service, error) {
	var services []Service
	if err := gateway.DB.Model(&services).Select(); err != nil {
		return nil, err
	}
	return services, nil
}
