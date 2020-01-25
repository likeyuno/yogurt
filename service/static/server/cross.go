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
	"errors"
	"github.com/kallydev/yogurt/common/restful"
	"github.com/kallydev/yogurt/service/static"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CROSS() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			for _, cross := range static.Conf.HTTPS.Cross {
				if ctx.Request().Referer() == cross {
					return handlerFunc(ctx)
				}
			}
			return ctx.JSONPretty(http.StatusOK, restful.RespondJSON(restful.Error, errors.New("no permission to access"), nil), restful.Ident)
		}
	}
}
