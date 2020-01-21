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
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type Limit struct {
	ipMap sync.Map
}

func NewLimit() *Limit {
	l := new(Limit)
	go l.auto()
	return l
}

func (l *Limit) Run() {

}

func (l *Limit) Handle(r *http.Request) bool {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println(err)
		return true
	}
	return l.speedLimit(host)
}

func (l *Limit) auto() {
	for {
		time.Sleep(time.Second * 1)
		l.ipMap.Range(func(key, value interface{}) bool {
			if value.(int) > 0 {
				l.ipMap.Store(key, value.(int)-1)
			}
			return true
		})
	}
}

func (l *Limit) speedLimit(ip string) bool {
	if count := l.getCount(ip); count >= 5 {
		if count < 30 {
			l.addCount(ip)
		}
		return true
	}
	l.addCount(ip)
	return false
}

func (l *Limit) getCount(ip string) int {
	count, _ := l.ipMap.LoadOrStore(ip, 0)
	return count.(int)
}

func (l *Limit) addCount(ip string) {
	l.ipMap.Store(ip, l.getCount(ip)+1)
}
