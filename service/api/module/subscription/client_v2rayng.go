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

const (
	V2RayNGClient = "v2rayng"
)

type V2RayNGVmess struct {
	Version         string `json:"v"`
	Name            string `json:"ps"`
	Host            string `json:"add"`
	Port            string `json:"port"`
	UUID            string `json:"id"`
	AlterID         string `json:"aid"`
	TLS             string `json:"tls"`
	Protocol        string `json:"type"`
	Obfuscation     string `json:"net"`
	ObfuscationHost string `json:"host"`
	ObfuscationPath string `json:"path"`
}

func NewV2RayNGVmess(v Vmess) *V2RayNGVmess {
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
	return &V2RayNGVmess{
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

func (v V2RayNGVmess) Build() ([]byte, error) {
	if data, err := json.Marshal(&v); err != nil {
		return nil, err
	} else {
		return []byte("vmess://" + base64.RawStdEncoding.EncodeToString(data)), nil
	}
}
