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
	"errors"
	"fmt"
)

type Vmess struct {
	Name            string
	Host            string
	Port            string
	UUID            string
	Security        string
	AlertID         string
	TLS             bool
	TLSSecurity     bool
	TLSHost         string
	Obfuscation     string
	ObfuscationHost string
	ObfuscationPath string
}

func (v Vmess) Build(c string) ([]byte, error) {
	switch c {
	case V2RayNGClient:
		return NewV2RayNGVmess(v).Build()
	case NetchClient:
		return NewNetchVmess(v).Build()
	case QuantumultXClient:
		return NewQuantumultXVmess(v).Build()
	default:
		return nil, errors.New(fmt.Sprintf("client %s not support", c))
	}
}

type VmessArray []Vmess

func (va VmessArray) Build(_client string) ([]byte, error) {
	result := ""
	for i, vmess := range va {
		if data, err := vmess.Build(_client); err != nil {
			return nil, err
		} else {
			result += string(data)
		}
		if i < len(va)-1 {
			result += "\n"
		}
	}
	switch _client {
	case V2RayNGClient:
		fallthrough
	case NetchClient:
		result = base64.RawURLEncoding.EncodeToString([]byte(result))
	case QuantumultXClient:
	default:
		return nil, errors.New(fmt.Sprintf("client %s not support", _client))
	}
	return []byte(result), nil
}
