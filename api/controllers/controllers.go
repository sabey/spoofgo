package controllers

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"bytes"
	"net/http"
	"sabey.co/spoofgo/api"
	"sabey.co/spoofgo/api/controllers/controller"
	"sabey.co/spoofgo/api/controllers/controller/standard"
)

var (
	controllers = map[string]func(
		*controller.Controller,
		*api.API,
		*bytes.Buffer,
		*http.Request,
	){ // success
		// default is a reserved keyword in golang so we're calling this package standard
		// this is the standard plugin
		"standard": standard.Controller,
	}
)

func GetController(
	key string,
) (
	func(*controller.Controller, *api.API, *bytes.Buffer, *http.Request), // func
	bool, // exists
) {
	f, ok := controllers[key]
	return f, ok
}
func GetControllers() []string {
	list := make([]string, 0, len(controllers))
	for control, _ := range controllers {
		list = append(list, control)
	}
	return list
}
