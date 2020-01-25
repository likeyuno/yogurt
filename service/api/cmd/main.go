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
	"github.com/kallydev/yogurt/common/restful"
	"github.com/kallydev/yogurt/service/api"
	"github.com/kallydev/yogurt/service/api/handler"
	"github.com/labstack/echo/v4"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
	"strings"
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
					protocol = strings.ToLower(ctx.QueryParam("protocol"))
					client   = strings.ToLower(ctx.QueryParam("client"))
				)
				if result, err := handler.GetSubscription(key, protocol, client); err != nil {
					return err
				} else {
					return ctx.String(http.StatusOK, string(result))
				}
			})
			subGroup.GET("/:key/info", func(ctx echo.Context) error {
				var key = ctx.Param("key")
				if result, err := handler.GetSubscriptionInfo(key); err != nil {
					return err
				} else {
					return ctx.JSONPretty(http.StatusOK, restful.RespondJSON(restful.OK, nil, result), restful.Ident)
				}
			})
			subGroup.GET("/:key/qrcode", func(ctx echo.Context) error {
				if png, err := qrcode.Encode(
					`vmess = hk-b2.yogurtcloud.com:80, method=chacha20-ietf-poly1305, password=41d7b431-6d72-4a93-876e-70b983765e9b, tag=🇭🇰 [测试 | 视频] 香港 HK-B2a
vmess = hk-b2.yogurtcloud.com:443, method=chacha20-ietf-poly1305, password=41d7b431-6d72-4a93-876e-70b983765e9b, tag=🇭🇰 [测试 | 视频] 香港 HK-B2b
vmess = hk-b2.yogurtcloud.com:8443, method=chacha20-ietf-poly1305, password=41d7b431-6d72-4a93-876e-70b983765e9b, tag=🇭🇰 [测试 | 视频] 香港 HK-B2c`, qrcode.Medium, 256); err != nil {
					return err
				} else {
					ctx.Response().Header().Set("content-type", "image/png")
					return ctx.String(http.StatusOK, string(png))
				}
			})
		}
	}
	log.Fatalln(s.Start(api.Conf.HTTPS.Addr()))
}
