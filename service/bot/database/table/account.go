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
)

type Account struct {
	tableName struct{} `pg:"public.accounts"`

	Username string
	Nickname string
	Email    string
	QQ       string
	Telegram string
	Money    int
	Password string

	database.Table
}

func QueryAccountByUsername(db *pg.DB, username string) (*Account, error) {
	var account Account
	err := db.Model(&account).Where("username = ?", username).Select()
	return &account, err
}

func QueryAccountByTelegram(db *pg.DB, telegram string) (*Account, error) {
	var account Account
	err := db.Model(&account).Where("telegram = ?", telegram).Select()
	return &account, err
}

func UpdateAccountTelegramByUsername(db *pg.DB, username, telegram string) (*Account, error) {
	var account Account
	_, err := db.Model(&account).Set("telegram = ?", telegram).Where("username = ?", username).Returning("*").Update()
	return &account, err
}

func InsertAccount(db *pg.DB, username, email, qq string) (*Account, error) {
	var account = Account{
		Username: username,
		QQ:       qq,
		Email:    email,
	}
	_, err := db.Model(&account).Returning("*").Insert()
	return &account, err
}
