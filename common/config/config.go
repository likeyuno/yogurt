/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net"
)

type Config struct {
	HTTPS    HTTPS    `yaml:"https"`
	Postgres Postgres `yaml:"postgres"`
}

type HTTPS struct {
	Host    string   `yaml:"host"`
	Port    string   `yaml:"port"`
	Public  string   `yaml:"public"`
	Private string   `yaml:"private"`
	Cross   []string `yaml:"cross"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Schema   string `yaml:"schema"`
}

func ParseConfig(data []byte) (*Config, error) {
	var conf Config
	return &conf, yaml.Unmarshal(data, &conf)
}

func ParseConfigFile(path string) (*Config, error) {
	var conf Config
	if data, err := ioutil.ReadFile(path); err != nil {
		return &conf, err
	} else {
		return ParseConfig(data)
	}
}

func (h HTTPS) Addr() string {
	return net.JoinHostPort(h.Host, h.Port)
}
