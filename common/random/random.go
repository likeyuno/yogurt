/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package random

import (
	"math/rand"
)

const strings = "123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func String(l int) string {
	result := ""
	for i := 0; i < l; i++ {
		result += string(strings[rand.Int63()%int64(len(strings))])
	}
	return result
}
