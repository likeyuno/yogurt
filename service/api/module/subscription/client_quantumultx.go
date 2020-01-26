/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package subscription

import (
	"fmt"
	"net"
)

const QuantumultXClient = "quantumultx"

type QuantumultXVmess struct {
	Name            string // tag=Sample-H
	Addr            string // addr=ws-c.example.com:80
	UUID            string // password=23ad6b10-8d1a-40f7-8ad0-e3e35cd32291
	Security        string // method=chacha20-ietf-poly1305
	Obfuscation     string // obfs=ws
	ObfuscationHost string // obfs-host=ws-c.example.com
	ObfuscationPath string // obfs-uri=/ws
	FastOpen        string // fast-open=false
	UDPRelay        string // udp-relay=false
}

func NewQuantumultXVmess(v Vmess) *QuantumultXVmess {
	return &QuantumultXVmess{
		Name:     v.Name,
		Addr:     net.JoinHostPort(v.Host, v.Port),
		UUID:     v.UUID,
		Security: v.Security,
		Obfuscation: func() string {
			var obf string
			if v.Obfuscation == "websocket" {
				obf = "ws"
			}
			if v.TLS {
				obf += "s"
			}
			return obf
		}(),
		ObfuscationHost: v.ObfuscationHost,
		ObfuscationPath: v.ObfuscationPath,
	}
}

func (c QuantumultXVmess) Build() ([]byte, error) {
	var result string
	result += fmt.Sprintf("vmess=%s, ", c.Addr)
	result += fmt.Sprintf("method=%s, ", c.Security)
	result += fmt.Sprintf("password=%s, ", c.UUID)
	if c.ObfuscationHost != "" {
		result += fmt.Sprintf("obfs-host=%s, ", c.ObfuscationHost)
	}
	if c.Obfuscation != "" {
		result += fmt.Sprintf("obfs=%s, ", c.Obfuscation)
	}
	if c.ObfuscationPath != "" {
		result += fmt.Sprintf("obfs-uri=%s, ", c.ObfuscationPath)
	}
	//if vfqx.FastOpen != "" {
	//	result += fmt.Sprintf("fast-open=%s, ", vfqx.FastOpen)
	//}
	//if vfqx.UDPRelay != "" {
	//	result += fmt.Sprintf("udp-relay=%s, ", vfqx.UDPRelay)
	//}
	result += fmt.Sprintf("tag=%s", c.Name)
	return []byte(result), nil
}
