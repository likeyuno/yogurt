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
	"database/sql"
	"github.com/kallydev/yogurt/common/database"
	"github.com/kallydev/yogurt/service/gateway"
	"github.com/lib/pq"
	"golang.org/x/net/context"
)

type Service struct {
	Name    string
	Host    string
	Servers pq.StringArray

	database.Table
}

const (
	QueryAllServicesSQL = `SELECT * FROM gateway.services`
)

func QueryAllServices(ctx context.Context) ([]Service, error) {
	var services []Service
	if rows, err := gateway.DB.QueryContext(ctx, QueryAllServicesSQL); err != nil {
		return nil, err
	} else if err := scanServices(rows, &services); err != nil {
		return nil, err
	}
	return services, nil
}

func scanService(row *sql.Row, service *Service) error {
	if err := row.Scan(
		&service.ID, &service.Name, &service.Host, &service.Servers,
		&service.CreatedAt, &service.UpdatedAt, &service.DeletedAt,
	); err != nil {
		return err
	}
	return nil
}

func scanServices(rows *sql.Rows, services *[]Service) error {
	for rows.Next() {
		service := Service{}
		if err := rows.Scan(
			&service.ID, &service.Name, &service.Host, &service.Servers,
			&service.CreatedAt, &service.UpdatedAt, &service.DeletedAt,
		); err != nil {
			return err
		}
		*services = append(*services, service)
	}
	return nil
}
