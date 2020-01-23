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
	"fmt"
	"github.com/kallydev/yogurt/service/api/database/table"
	"github.com/kallydev/yogurt/service/api/module/subscription"
)

func GetSubscription(key, protocol, client string) ([]byte, error) {
	sub, err := table.QuerySubscriptionByKey(key)
	if err != nil {
		return nil, err
	}
	pack, err := table.QueryPackageByName(sub.Package)
	if err != nil {
		return nil, err
	}
	nodes, err := table.QueryNodeByIDs(pack.Nodes)
	if err != nil {
		return nil, err
	}
	var nodeIDs []string
	for _, node := range nodes {
		nodeIDs = append(nodeIDs, node.ID)
	}
	ssrs, err := table.QueryNodeShadowsocksRByIDs(nodeIDs)
	if err != nil {
		return nil, err
	}
	newSSRs := subscription.ShadowsocksRs{}
	for i, ssr := range ssrs {
		newSSRs = append(newSSRs, subscription.ShadowsocksR{
			Host:             ssr.Host,
			Port:             ssr.Port,
			Method:           ssr.Method,
			Password:         ssr.Password,
			Obfuscation:      ssr.Obfuscation,
			ObfuscationParam: ssr.ObfuscationParam,
			Protocol:         ssr.Protocol,
			ProtocolParam:    ssr.ProtocolParam,
			Remarks:          fmt.Sprintf("%s %s", nodes[i].Location, nodes[i].Name),
			Group:            pack.Name,
		})
	}
	return newSSRs.Build()
}
