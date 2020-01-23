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

type Subscription struct {
	tableName struct{} `pg:"public.subscriptions"`

	Package  string
	Account  string
	Key      string
	Status   string
	ExpireAt string

	database.Table
}

func QuerySubscriptionByKey(key string) (*Subscription, error) {
	sub := Subscription{}
	err := api.DB.Model(&sub).Where("key = ?", key).Select()
	return &sub, err
}
