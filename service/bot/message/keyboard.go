/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package message

import tb "github.com/go-telegram-bot-api/telegram-bot-api"

var JoinKeyboard = tb.NewInlineKeyboardMarkup(
	tb.NewInlineKeyboardRow(
		tb.NewInlineKeyboardButtonURL("查看群規", "https://t.me/yogurtCloud/5"),
	),
)
