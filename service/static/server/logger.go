/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package server

import (
	"github.com/kallydev/yogurt/service/static/database/table"
	"github.com/labstack/echo/v4"
	"log"
)

func Logger() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			go func(ctx echo.Context) {
				if _, err := table.InsertLog(
					ctx.Request().Host, ctx.Request().URL.Path,
					ctx.QueryParams().Encode(), ctx.Request().UserAgent(),
					ctx.RealIP(),
				); err != nil {
					log.Println(err)
				}
			}(ctx)
			return handlerFunc(ctx)
		}
	}
}
