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
	"github.com/kallydev/yogurt/common/random"
	"time"
)

type Subscription struct {
	tableName struct{} `pg:"public.subscriptions"`

	Package  string
	Account  string
	Key      string
	UUID     string
	Status   string
	ExpireAt time.Time

	database.Table
}

func QuerySubscriptionByKey(db *pg.DB, key string) (*Subscription, error) {
	sub := Subscription{}
	err := db.Model(&sub).Where("key = ?", key).Select()
	return &sub, err
}

func QuerySubscriptionsByUsername(db *pg.DB, username string) ([]Subscription, error) {
	var subs []Subscription
	err := db.Model(&subs).Where("account = ?", username).Select()
	return subs, err
}

func InsertSubscription(db *pg.DB, account, _package string, day time.Duration) (*Subscription, error) {
	var subscription = Subscription{
		Package:  _package,
		Account:  account,
		Key:      random.String(8),
		ExpireAt: time.Now().Add(day),
	}
	_, err := db.Model(&subscription).Returning("*").Insert()
	return &subscription, err
}
