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
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type V2Ray struct {
	Name            string
	Host            string
	Port            string
	UUID            string
	Security        string
	AlertID         string
	TLS             bool
	TLSSecurity     bool
	TLSHost         string
	Obfuscation     string `json:"type"`
	ObfuscationHost string `json:"host"`
	ObfuscationPath string `json:"path"`
}

func (v V2Ray) Build(client string) ([]byte, error) {
	switch client {
	case "v2rayng":
		return v.buildForV2rayNG()
	case "quantumultx":
		return v.buildForQuantumultX()
	default:
		return nil, errors.New("not support")
	}
}

type V2RayNG struct {
	Version         string `json:"v"`
	Name            string `json:"ps"`
	Host            string `json:"add"`
	Port            string `json:"port"`
	UUID            string `json:"id"`
	Security        string `json:"type"`
	AlterID         string `json:"aid"`
	TLS             string `json:"tls"`
	Obfuscation     string `json:"net"`
	ObfuscationHost string `json:"host"`
	ObfuscationPath string `json:"path"`
}

func (v V2Ray) buildForV2rayNG() ([]byte, error) {
	var (
		tls string
		net = v.Obfuscation
	)
	if v.TLS {
		tls = "tls"
	}
	if net == "" {
		net = "tcp"
	}
	if body, err := json.Marshal(&V2RayNG{
		Version:         "2",
		Name:            v.Name,
		Host:            v.Host,
		Port:            v.Port,
		UUID:            v.UUID,
		AlterID:         v.AlertID,
		Security:        v.Security,
		TLS:             tls,
		Obfuscation:     net,
		ObfuscationHost: v.ObfuscationHost,
		ObfuscationPath: v.ObfuscationPath,
	}); err != nil {
		return nil, err
	} else {
		result := "vmess://" + base64.RawStdEncoding.EncodeToString(body)
		return []byte(result), nil
	}
}

type VmessForQuantumultX struct {
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

/*
 * vmess=ws-tls-b.example.com:443, method=chacha20-ietf-poly1305,
 * password= 23ad6b10-8d1a-40f7-8ad0-e3e35cd32291, obfs-host=ws-tls-b.example.com,
 * obfs=wss, obfs-uri=/ws, fast-open=false, udp-relay=false, tag=Sample-I
 */

func (vfqx VmessForQuantumultX) Build() []byte {
	var result string
	result += fmt.Sprintf("vmess=%s, ", vfqx.Addr)
	result += fmt.Sprintf("method=%s, ", vfqx.Security)
	result += fmt.Sprintf("password=%s, ", vfqx.UUID)
	if vfqx.ObfuscationHost != "" {
		result += fmt.Sprintf("obfs-host=%s, ", vfqx.ObfuscationHost)
	}
	if vfqx.Obfuscation != "" {
		result += fmt.Sprintf("obfs=%s, ", vfqx.Obfuscation)
	}
	if vfqx.ObfuscationPath != "" {
		result += fmt.Sprintf("obfs-uri=%s, ", vfqx.ObfuscationPath)
	}
	if vfqx.FastOpen != "" {
		result += fmt.Sprintf("fast-open=%s, ", vfqx.FastOpen)
	}
	if vfqx.UDPRelay != "" {
		result += fmt.Sprintf("udp-relay=%s, ", vfqx.UDPRelay)
	}
	result += fmt.Sprintf("tag=%s", vfqx.Name)
	return []byte(result)
}

func (v V2Ray) buildForQuantumultX() ([]byte, error) {
	return VmessForQuantumultX{
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
	}.Build(), nil
}

type V2Rays []V2Ray

func (vs V2Rays) Build(client string) ([]byte, error) {
	result := ""
	for i, v := range vs {
		data, _ := v.Build(client)
		result += string(data)
		if i < len(vs)-1 {
			result += "\n"
		}
	}
	switch client {
	case "v2rayng":
		result = base64.RawURLEncoding.EncodeToString([]byte(result))
	case "quantumultx":
	default:
		return nil, errors.New("not support")
	}
	return []byte(result), nil
}
