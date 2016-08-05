package plugins

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"sabey.co/spoofgo/api"
	"sabey.co/spoofgo/plugins/plugin"
	"sabey.co/spoofgo/plugins/plugin/standard"
)

var (
	plugins = map[string]func(
		*plugin.Plugin,
		*api.API,
	) bool{ // success
		// default is a reserved keyword in golang so we're calling this package standard
		// this is the standard plugin
		"standard-http": standard.HTTP,
		"standard-log":  standard.Log,
	}
)

func GetPlugin(
	key string,
) (
	func(*plugin.Plugin, *api.API) bool, // func
	bool, // exists
) {
	f, ok := plugins[key]
	return f, ok
}
func GetPlugins() []string {
	list := make([]string, 0, len(plugins))
	for plug, _ := range plugins {
		list = append(list, plug)
	}
	return list
}
