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
)

type V2Ray struct {
	Name     string
	Host     string
	Port     string
	UUID     string
	Security string
}

func (v V2Ray) Build(_type string) ([]byte, error) {
	switch _type {
	case "v2rayng":
		return v.buildForV2rayNG()
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
	AID             string `json:"aid"`
	Network         string `json:"net"`
	Obfuscation     string `json:"type"`
	ObfuscationHost string `json:"host"`
	ObfuscationPath string `json:"path"`
	// TLS             string `json:"tls"`
}

func (v V2Ray) buildForV2rayNG() ([]byte, error) {
	if body, err := json.Marshal(&V2RayNG{
		Version:     "1",
		Name:        v.Name,
		Host:        v.Host,
		Port:        v.Port,
		UUID:        v.UUID,
		AID:         "0",
		Network:     "tcp",
		Obfuscation: "none",
	}); err != nil {
		return nil, err
	} else {
		result := "vmess://" + base64.RawURLEncoding.EncodeToString(body)
		return []byte(result), nil
	}
}

type V2Rays []V2Ray

func (vs V2Rays) Build(_type string) ([]byte, error) {
	result := ""
	for i, v := range vs {
		data, _ := v.Build(_type)
		result += string(data)
		if i < len(vs)-1 {
			result += "\n"
		}
	}
	result = base64.RawStdEncoding.EncodeToString([]byte(result))
	return []byte(result), nil
}
