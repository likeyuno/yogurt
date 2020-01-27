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
	"time"
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
	if protocol == Vmess && client == "" {
		client = V2RayNG
	}
	sub, err := table.QuerySubscriptionByKey(key)
	if err != nil {
		return nil, err
	}
	if sub.ExpireAt.Before(time.Now()) {
		return nil, errors.New("subscription expired")
	} else if sub.Status != "normal" && sub.Status != "free" {
		return nil, errors.New("subscription blocked")
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
		return subscription.BuildShadowsocksR(pack.Name, nodes)
	case Vmess:
		return subscription.BuildVmess(pack.Name, client, sub.UUID, nodes)
	default:
		return nil, errors.New("not support")
	}
}

type SubscriptionInfoData struct {
	Package  string         `json:"package"`
	Username string         `json:"username"`
	Status   string         `json:"status"`
	ExpireAt string         `json:"expire_at"`
	Nodes    []NodeInfoData `json:"nodes"`
}

type NodeInfoData struct {
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
}

func GetSubscriptionInfo(key string) (*SubscriptionInfoData, error) {
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
	var nds []NodeInfoData
	for _, node := range nodes {
		nds = append(nds, NodeInfoData{
			Name: func() string {
				return fmt.Sprintf(fmt.Sprintf(
					"[%s | %s] %s %s",
					node.Tags[0], node.Tags[1], node.Location, node.Name,
				))
			}(),
			Protocol: node.Type,
		})
	}
	return &SubscriptionInfoData{
		Package:  sub.Package,
		Username: sub.Account,
		Status:   sub.Status,
		ExpireAt: sub.ExpireAt.Format("2006-01-02"),
		Nodes:    nds,
	}, nil
}

type GetSubscriptionsRankData struct {
	Subscriptions []struct {
		Name    string `json:"name"`
		Expired string `json:"expired"`
	} `json:"subscriptions"`
}

func GetSubscriptionsRank() (*GetSubscriptionsRankData, error) {
	data := GetSubscriptionsRankData{}
	subs, err := table.QuerySubscriptionsOrder()
	if err != nil {
		return nil, err
	}
	for _, sub := range subs {
		name := func() string {
			l := len(sub.Account) - 3
			left := l / 2
			return sub.Account[:left] + "***" + sub.Account[l+3-left:]
		}()
		data.Subscriptions = append(data.Subscriptions, struct {
			Name    string `json:"name"`
			Expired string `json:"expired"`
		}{
			Name: name, Expired: sub.ExpireAt.Format("2006-01-02"),
		})
	}
	return &data, err
}
