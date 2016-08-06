package controller

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"bytes"
	"log"
	"net/http"
	p "sabey.co/spoofgo/plugins/plugin"
	"sabey.co/spoofgo/version"
)

type Controller struct {
	// safe
	state       *State
	version     *version.Version
	using       string
	controllers []string // list of controllers, since we can't include controllers package
	// we have to use func() to get around include loop issues
	controller func(*bytes.Buffer, *http.Request)
}
type State struct {
	// unsafe
	Plugin *p.Plugin
	// KeyValue can be set with the flag -controller-kv
	KeyValue map[string]interface{}
}

func Create(
	version *version.Version,
	state *State,
	using string,
	controllers []string, // list of controllers, since we can't include controllers package
	controller func(*bytes.Buffer, *http.Request),
) *Controller {
	if version == nil {
		log.Fatalln("./spoofgo/api/controllers/controller.Create(): Version was nil")
		return nil
	}
	if state == nil {
		log.Fatalln("./spoofgo/api/controllers/controller.Create(): State was nil")
		return nil
	}
	return &Controller{
		state:       state,
		version:     version,
		using:       using,
		controllers: controllers, // list of controllers, since we can't include controllers package
		controller:  controller,
	}
}
func (self *Controller) Controller(
	w *bytes.Buffer,
	r *http.Request,
) {
	// we can only safely get the information about the current request and the buffer to write to
	// to make the controller work, the controller function must find get the api object itself
	self.controller(w, r)
}
func (self *Controller) GetVersion() *version.Version {
	// dereference version
	return self.version.Clone()
}
func (self *Controller) GetUsing() string {
	return self.using
}
func (self *Controller) GetControllers() []string {
	controllers := make([]string, len(self.controllers))
	copy(controllers, self.controllers)
	return controllers
}
func (self *Controller) GetState() *State {
	return self.state
}
