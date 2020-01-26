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
	"crypto/tls"
	"flag"
	"github.com/kallydev/yogurt/service/reverse"
	"log"
	"net"
)

var (
	localAddr, remoteAddr string
	publicPem, privateKey string
)

func main() {
	flag.StringVar(&localAddr, "l", "0.0.0.0:443", "local addr")
	flag.StringVar(&remoteAddr, "r", "127.0.0.1:10000", "remote addr")
	flag.StringVar(&publicPem, "p", "public.pem", "public pem")
	flag.StringVar(&privateKey, "k", "private.key", "private key")
	flag.Parse()

	cert, err := tls.LoadX509KeyPair(publicPem, privateKey)
	if err != nil {
		log.Println(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	tcpListener, err := tls.Listen("tcp", localAddr, config)
	if err != nil {
		log.Println(err)
		return
	}
	defer tcpListener.Close()
	for {
		tcpConn, err := tcpListener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go tcpHandler(tcpConn)
	}
}

func tcpHandler(tcpConn net.Conn) {
	dstConn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		log.Println(err)
		return
	}
	go pipe(tcpConn, dstConn)
	pipe(dstConn, tcpConn)
}

func pipe(src, dst net.Conn) {
	defer dst.Close()
	buf := reverse.BufInstance.Get()
	defer reverse.BufInstance.Put(buf)
	for {
		n, err := src.Read(buf)
		if n > 0 {
			if _, err := dst.Write(buf[0:n]); err != nil {
				log.Println(err)
				break
			}
		}
		if err != nil {
			break
		}
	}
	return
}
