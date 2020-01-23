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
	"fmt"
)

type ShadowsocksR struct {
	Host             string
	Port             string
	Method           string
	Password         string
	Obfuscation      string
	ObfuscationParam string
	Protocol         string
	ProtocolParam    string
	Remarks          string
	Group            string
}

func (ssr ShadowsocksR) Build() ([]byte, error) {
	return []byte(ssr.build()), nil
}

func (ssr ShadowsocksR) build() string {
	return fmt.Sprintf("ssr://%s", base64.RawURLEncoding.EncodeToString(
		[]byte(fmt.Sprintf(
			"%s:%s:%s:%s:%s:%s/?protoparam=%s&obfsparam=%s&remarks=%s&group=%s",
			ssr.Host, ssr.Port, ssr.Protocol, ssr.Method, ssr.Obfuscation,
			base64.RawURLEncoding.EncodeToString([]byte(ssr.Password)),
			base64.RawURLEncoding.EncodeToString([]byte(ssr.ProtocolParam)),
			base64.RawURLEncoding.EncodeToString([]byte(ssr.ObfuscationParam)),
			base64.RawURLEncoding.EncodeToString([]byte(ssr.Remarks)),
			base64.RawURLEncoding.EncodeToString([]byte(ssr.Group))),
		)),
	)
}

type ShadowsocksRs []ShadowsocksR

func (ssrs ShadowsocksRs) Build() ([]byte, error) {
	result := ""
	for i, ssr := range ssrs {
		data, _ := ssr.Build()
		result += string(data)
		if i < len(ssrs)-1 {
			result += "\n"
		}
	}
	result = base64.RawStdEncoding.EncodeToString([]byte(result))
	return []byte(result), nil
}
