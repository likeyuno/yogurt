/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package handler

import (
	"errors"
	"fmt"
	"github.com/kallydev/yogurt/service/api/database/table"
	"github.com/kallydev/yogurt/service/api/module/subscription"
)

const (
	ShadowsocksR = "shadowsocksr"
	Vmess        = "vmess"

	V2RayNG = "v2rayng"
)

func GetSubscription(key, protocol, client string) ([]byte, error) {
	if protocol == "" {
		protocol = ShadowsocksR
	}
	if protocol == Vmess && client != V2RayNG {
		client = V2RayNG
	}
	sub, err := table.QuerySubscriptionByKey(key)
	if err != nil {
		return nil, err
	}
	pack, err := table.QueryPackageByName(sub.Package)
	if err != nil {
		return nil, err
	}
	nodes, err := table.QueryNodeByIDsAndType(protocol, pack.Nodes)
	if err != nil {
		return nil, err
	}
	var nodeIDs []string
	for _, node := range nodes {
		nodeIDs = append(nodeIDs, node.ID)
	}
	switch protocol {
	case ShadowsocksR:
		return buildShadowsocksr(pack.Name, nodes)
	case Vmess:
		return buildVmess(client, sub.UUID, nodes)
	default:
		return nil, errors.New("not support")
	}
}

func buildShadowsocksr(packageName string, nodes []table.Node) ([]byte, error) {
	ssrs := subscription.ShadowsocksRs{}
	for _, node := range nodes {
		ssrs = append(ssrs, subscription.ShadowsocksR{
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
			Group: packageName,
		})
	}
	return ssrs.Build()
}

func buildVmess(_type, uuid string, nodes []table.Node) ([]byte, error) {
	vs := subscription.V2Rays{}
	for _, node := range nodes {
		vs = append(vs, subscription.V2Ray{
			Name: func() string {
				return fmt.Sprintf(fmt.Sprintf(
					"[%s | %s] %s %s",
					node.Tags[0], node.Tags[1], node.Location, node.Name,
				))
			}(),
			Host:     node.NodeV2Ray.Host,
			Port:     node.NodeV2Ray.Port,
			UUID:     uuid,
			Security: "auto",
		})
	}
	return vs.Build(_type)
}
