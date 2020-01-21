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
	"encoding/json"
	"errors"
	"github.com/kallydev/yogurt/common/context"
	"github.com/kallydev/yogurt/common/restful"
	"github.com/kallydev/yogurt/service/gateway/database/table"
	"log"
	"net"
	"net/http"
	"time"
)

type Server struct {
	Equalizer *Equalizer
	Limit     *Limit

	*http.ServeMux
}

func NewServer() *Server {
	return &Server{
		Equalizer: NewEqualizer(),
		Limit:     NewLimit(),
		ServeMux:  http.NewServeMux(),
	}
}

func (s *Server) EqualizerHandleFunc(rw http.ResponseWriter, r *http.Request) {
	go func(r http.Request) {
		if host, _, err := net.SplitHostPort(r.RemoteAddr); err != nil {
			log.Println(err)
			return
		} else if _, err := table.InsertLog(context.WithTimeoutNoCancel(time.Second*3),
			r.Host, r.Method, r.URL.Path, r.URL.Query().Encode(), r.UserAgent(), host,
		); err != nil {
			log.Println(err)
		}
	}(*r)
	if s.Limit.Handle(r) {
		res := restful.RespondJSON(restful.Error, errors.New("requests are too frequent"), nil)
		if data, err := json.MarshalIndent(res, "", restful.Ident); err != nil {
			log.Println(err)
		} else if _, err := rw.Write(data); err != nil {
			log.Println(err)
		}
	} else {
		s.Equalizer.Handle(rw, r)
	}
}
