/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package static

import (
	"github.com/kallydev/yogurt/service/static/database/table"
	"github.com/labstack/echo/v4"
	"log"
)

func Logger() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			go func(ctx echo.Context) {
				l := table.Log{
					Host:      ctx.Request().Host,
					Path:      ctx.Request().URL.Path,
					Params:    ctx.QueryParams().Encode(),
					UserAgent: ctx.Request().UserAgent(),
					IP:        ctx.RealIP(),
				}
				if res, err := DB.Insert(&l); err != nil {
					log.Println(err)
				} else {
					log.Println(l, res)
				}
			}(ctx)
			return handlerFunc(ctx)
		}
	}
}
