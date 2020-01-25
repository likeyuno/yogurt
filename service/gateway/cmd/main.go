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
	"github.com/kallydev/yogurt/service/gateway"
	"github.com/kallydev/yogurt/service/gateway/server"
	"log"
	"net/http"
)

func main() {
	s := server.NewServer()
	s.HandleFunc("/", s.EqualizerHandleFunc)
	log.Fatalln(http.ListenAndServe(
		gateway.Conf.HTTPS.Addr(),  s,
	))
}
