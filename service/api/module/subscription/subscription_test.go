/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package subscription_test

import (
	"fmt"
	"github.com/kallydev/yogurt/service/api/module/subscription"
	"testing"
)

func TestShadowsocksR(t *testing.T) {
	ssrs := createShadowsocksRs()
	if result, err := ssrs.Build(); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(string(result))
	}
}

func createShadowsocksRs() subscription.ShadowsocksRs {
	ssrs := subscription.ShadowsocksRs{}
	for i := 0; i < 10; i++ {
		ssrs = append(ssrs, createShadowsocksR())
	}
	return ssrs
}

func createShadowsocksR() subscription.ShadowsocksR {
	return subscription.ShadowsocksR{}
}

func TestV2Ray(t *testing.T) {
	vs := createV2RAys()
	if result, err := vs.Build(); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(string(result))
	}
}

func createV2RAy() subscription.V2Ray {
	return subscription.V2Ray{}
}

func createV2RAys() subscription.V2Rays {
	vs := subscription.V2Rays{}
	for i := 0; i < 10; i++ {
		vs = append(vs, createV2RAy())
	}
	return vs
}
