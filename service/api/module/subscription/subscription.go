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
	"github.com/kallydev/yogurt/service/api/database/table"
)

type ISubscription interface {
	Build(_type string) ([]byte, error)
}

type INode interface {
	Build(_type string) ([]byte, error)
}

func BuildVmess(group, client, uuid string, nodes []table.Node) ([]byte, error) {
	va := VmessArray{}
	for _, node := range nodes {
		va = append(va, Vmess{
			Name: func() string {
				return fmt.Sprintf(fmt.Sprintf(
					"[%s | %s] %s %s",
					node.Tags[0], node.Tags[1], node.Location, node.Name,
				))
			}(),
			Host:            node.NodeV2Ray.Host,
			Port:            node.NodeV2Ray.Port,
			UUID:            uuid,
			Security:        node.NodeV2Ray.Security,
			AlertID:         node.NodeV2Ray.AlertID,
			TLS:             node.NodeV2Ray.TLS,
			TLSSecurity:     node.NodeV2Ray.TLSVerification,
			TLSHost:         node.NodeV2Ray.TLSHost,
			Obfuscation:     node.NodeV2Ray.Obfuscation,
			ObfuscationHost: node.NodeV2Ray.ObfuscationHost,
			ObfuscationPath: node.NodeV2Ray.ObfuscationPath,
		})
	}
	return va.Build(group, client)
}

func BuildShadowsocksR(_package string, nodes []table.Node) ([]byte, error) {
	sa := ShadowsocksRArray{}
	for _, node := range nodes {
		sa = append(sa, ShadowsocksR{
			Host:             node.NodeShadowsocksR.Host,
			Port:             node.NodeShadowsocksR.Port,
			Method:           node.NodeShadowsocksR.Method,
			Password:         node.NodeShadowsocksR.Password,
			Obfuscation:      node.NodeShadowsocksR.Obfuscation,
			ObfuscationParam: node.NodeShadowsocksR.ObfuscationParam,
			Protocol:         node.NodeShadowsocksR.Protocol,
			ProtocolParam:    node.NodeShadowsocksR.ProtocolParam,
			Remarks: func() string {
				return fmt.Sprintf(fmt.Sprintf(
					"[%s | %s] %s %s",
					node.Tags[0], node.Tags[1], node.Location, node.Name,
				))
			}(),
			Group: _package,
		})
	}
	return sa.Build()
}
