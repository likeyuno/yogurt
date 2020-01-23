/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package subscription

type ISubscription interface {
	Build(_type string) ([]byte, error)
}

type INode interface {
	Build(_type string) ([]byte, error)
}
