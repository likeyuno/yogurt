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
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

type Server struct {
	*echo.Echo
}

func NewServer() *Server {
	return &Server{
		Echo: func() *echo.Echo {
			e := echo.New()
			e.HideBanner = true
			e.HidePort = true
			e.HTTPErrorHandler = httpErrorHandlerFunc
			return e
		}(),
	}
}

func httpErrorHandlerFunc(err error, ctx echo.Context) {
	var errs = err
	if err, ok := err.(*echo.HTTPError); ok {
		errs = errors.New(strings.ToLower(err.Message.(string)))
	}
	if err := ctx.JSONPretty(http.StatusOK, restful.RespondJSON(restful.Error, errs, nil), restful.Ident); err != nil {
		log.Printf("http error handler error: %s", err)
	}
}
