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
	"github.com/lib/pq"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type Equalizer struct {
	sync.Map
}

type Service struct {
	Host    string
	Servers []string
}

func NewEqualizer() *Equalizer {
	e := &Equalizer{}
	go e.autoUpdate()
	return e
}

func (e *Equalizer) autoUpdate() {
	for {
		if services, err := table.QueryAllServices(context.WithTimeoutNoCancel(time.Second * 3)); err != nil {
			log.Println(err)
		} else {
			for _, service := range services {
				e.Update(service.Host, service.Servers)
			}
		}
		time.Sleep(time.Second)
	}
}

func (e *Equalizer) Update(host string, services []string) {
	e.Store(host, services)
}

func (e *Equalizer) Handle(rw http.ResponseWriter, r *http.Request) {
	if v, ok := e.Load(r.Host); ok {
		servers := v.(pq.StringArray)
		rand.Seed(time.Now().Unix())
		if remote, err := url.Parse(servers[rand.Intn(len(servers))]); err != nil {
			log.Println(err)
			return
		} else {
			proxy := httputil.NewSingleHostReverseProxy(remote)
			r.Header.Add("Client-IP", r.RemoteAddr)
			proxy.ServeHTTP(rw, r)
		}
	} else {
		res := restful.RespondJSON(restful.Error, errors.New("service not found"), nil)
		if data, err := json.MarshalIndent(res, "", restful.Ident); err != nil {
			log.Println(err)
		} else if _, err := rw.Write(data); err != nil {
			log.Println(err)
		}
	}
}
