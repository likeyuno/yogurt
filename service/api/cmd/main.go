/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package main

import (
	"github.com/kallydev/yogurt/service/api"
	"github.com/kallydev/yogurt/service/api/handler"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	s := api.NewServer()
	v1Group := s.Group("/v1")
	{
		subGroup := v1Group.Group("/subscriptions")
		{
			subGroup.GET("/:key", func(ctx echo.Context) error {
				var (
					key      = ctx.Param("key")
					protocol = ctx.QueryParam("protocol")
					client   = ctx.QueryParam("client")
				)
				result, err := handler.GetSubscription(key, protocol, client)
				if err != nil {
					return err
				}
				return ctx.String(http.StatusOK, string(result))
			})
		}
	}
	log.Fatalln(s.Start(api.Conf.HTTPS.Addr()))
}
