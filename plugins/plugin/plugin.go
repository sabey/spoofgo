package plugin

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"log"
	"net/http"
	clog "sabey.co/spoofgo/log"
	"sabey.co/spoofgo/version"
	"sync"
	"time"
)

type Plugin struct {
	// safe
	state   *State
	version *version.Version
	using   string
	plugins []string // list of plugins, since we can't include plugins package
	// we have to use func() bool to get around include loop issues
	call  func() bool
	every time.Duration
	// unsafe
	last_success time.Time
	running      bool
	mu           sync.RWMutex
}
type State struct {
	// unsafe
	Addr       string
	HTTPClient *http.Client
	// KeyValue can be set with the flag -plugin-kv
	KeyValue map[string]interface{}
}
type Info struct {
	SinceLast   time.Duration
	LastSuccess time.Time
	Offline     bool
	Running     bool
}

func Create(
	version *version.Version,
	state *State,
	using string,
	plugins []string, // list of plugins, since we can't include plugins package
	call func() bool,
	every time.Duration,
) *Plugin {
	if version == nil {
		log.Fatalln("./spoofgo/plugins/plugin.Create(): Version was nil")
		return nil
	}
	if state == nil {
		log.Fatalln("./spoofgo/plugins/plugin.Create(): State was nil")
		return nil
	}
	return &Plugin{
		state:   state,
		version: version,
		using:   using,
		plugins: plugins, // list of plugins, since we can't include plugins package
		call:    call,
		every:   every,
	}
}
func (self *Plugin) Start() {
	self.mu.Lock()
	defer self.mu.Unlock()
	if !self.running {
		self.running = true
		go func() {
			for _ = range time.NewTicker(self.every).C {
				if self.call() {
					// online
					self.mu.Lock()
					self.last_success = time.Now()
					self.mu.Unlock()
				} // offline
			}
			clog.Log("stopped calling plugins for unknown reason")
			log.Fatalln("stopped calling plugins for unknown reason")
		}()
	}
}
func (self *Plugin) GetVersion() *version.Version {
	// dereference version
	return self.version.Clone()
}
func (self *Plugin) GetUsing() string {
	return self.using
}
func (self *Plugin) GetPlugins() []string {
	plugins := make([]string, len(self.plugins))
	copy(plugins, self.plugins)
	return plugins
}
func (self *Plugin) GetState() *State {
	return self.state
}
func (self *Plugin) GetCall() func() bool {
	return self.call
}
func (self *Plugin) GetEvery() time.Duration {
	return self.every
}
func (self *Plugin) GetLastSuccess() time.Time {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.last_success
}
func (self *Plugin) GetSinceLast() time.Duration {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return time.Since(self.last_success)
}
func (self *Plugin) IsRunning() bool {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.running
}
func (self *Plugin) IsOffline() bool {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return time.Since(self.last_success) >= (self.every * 5)
}
func (self *Plugin) GetInfo() *Info {
	self.mu.RLock()
	defer self.mu.RUnlock()
	o := &Info{
		LastSuccess: self.last_success,
		Running:     self.running,
	}
	if self.last_success.IsZero() {
		o.Offline = true
	} else {
		o.SinceLast = time.Since(self.last_success)
		o.Offline = (o.SinceLast >= (self.every * 5))
	}
	return o
}
