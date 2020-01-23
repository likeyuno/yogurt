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

func QueryAllAccounts() ([]Account, error) {
	var accounts []Account
	err := api.DB.Model(&accounts).Select()
	return accounts, err
}
