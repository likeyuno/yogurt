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
)

const NetchClient = "netch"

type Netch struct{}

type NetchVmess struct {
	Version         string `json:"v"`
	Name            string `json:"ps"`
	Host            string `json:"add"`
	Port            string `json:"port"`
	UUID            string `json:"id"`
	AlterID         string `json:"aid"`
	TLS             string `json:"tls"`
	Protocol        string `json:"type"` // none
	Obfuscation     string `json:"net"`  // ws
	ObfuscationHost string `json:"host"`
	ObfuscationPath string `json:"path"`
}

func NewNetchVmess(v Vmess) *NetchVmess {
	net := v.Obfuscation
	tls := ""
	if v.TLS {
		tls = "tls"
	}
	if net == "websocket" {
		net = "ws"
	}
	return &NetchVmess{
		Version:         "2",
		Name:            v.Name,
		Host:            v.Host,
		Port:            v.Port,
		UUID:            v.UUID,
		AlterID:         v.AlertID,
		TLS:             tls,
		Protocol:        "none",
		Obfuscation:     net,
		ObfuscationHost: v.ObfuscationHost,
		ObfuscationPath: v.ObfuscationPath,
	}
}

func (n *NetchVmess) Build() ([]byte, error) {
	if data, err := json.Marshal(&n); err != nil {
		return nil, err
	} else {
		return []byte("vmess://" + base64.RawStdEncoding.EncodeToString(data)), nil
	}
}
