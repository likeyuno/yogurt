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
	"github.com/kallydev/yogurt/service/static"
	"github.com/kallydev/yogurt/service/static/server"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	s := server.NewServer()
	s.Use(server.Logger())
	s.Use(server.CROSS())
	s.Use(middleware.Recover())
	s.Static("/", "service/static/public")
	s.Logger.Fatal(s.Start(static.Conf.HTTPS.Addr()))
}
