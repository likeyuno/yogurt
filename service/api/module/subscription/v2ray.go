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
)

type V2Ray struct {
	Host     string
	Port     string
	UUID     string
	Security string
}

func (v V2Ray) Build(_type string) ([]byte, error) {
	switch _type {
	case "v2rayNG":
		return v.buildForV2rayNG()
	case "Kitsunebi":
		return v.buildForKitsunebi()
	}
	return nil, errors.New("not support")
}

func (v V2Ray) buildForV2rayNG() ([]byte, error) {
	if body, err := json.Marshal(&v); err != nil {
		return nil, err
	} else {
		result := "vmess://" + base64.RawURLEncoding.EncodeToString(body)
		return []byte(result), nil
	}
}

func (v V2Ray) buildForKitsunebi() ([]byte, error) {
	body := fmt.Sprintf("%s:%s@%s:%s", v.Security, v.UUID, v.Host, v.Port)
	result := "vmess://" + base64.RawURLEncoding.EncodeToString([]byte(body))
	return []byte(result), nil
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
