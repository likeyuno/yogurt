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
		ObfuscationPath: v.Obfuscation,
	}); err != nil {
		return nil, err
	} else {
		result := "vmess://" + base64.RawStdEncoding.EncodeToString(body)
		return []byte(result), nil
	}
}

func (v V2Ray) buildForQuantumultX() ([]byte, error) {
	format := "vmess = %s, method=%s, password=%s, tag=%s"
	return []byte(fmt.Sprintf(format, net.JoinHostPort(v.Host, v.Port), v.Security, v.UUID, v.Name)), nil
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
