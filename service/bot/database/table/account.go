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
	"github.com/kallydev/yogurt/service/bot"
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

func QueryAccountByUsername(username string) (*Account, error) {
	var account Account
	err := bot.DB.Model(&account).Where("username = ?", username).Select()
	return &account, err
}

func QueryAccountByTelegram(telegram string) (*Account, error) {
	var account Account
	err := bot.DB.Model(&account).Where("telegram = ?", telegram).Select()
	return &account, err
}

func UpdateAccountTelegramByUsername(username, telegram string) (*Account, error) {
	var account Account
	_, err := bot.DB.Model(&account).Set("telegram = ?", telegram).Where("username = ?", username).Returning("*").Update()
	return &account, err
}
