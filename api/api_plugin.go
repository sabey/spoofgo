package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sabey.co/spoofgo/duration"
	"time"
)

type APIPlugin struct {
	Using       string    `json:"using,omitempty"`
	Every       string    `json:"every,omitempty"`
	SinceLast   string    `json:"since-last,omitempty"`
	LastSuccess time.Time `json:"last-success,omitempty"`
	Offline     bool      `json:"offline"`
	Running     bool      `json:"running"`
}

func (self *Server) GetAPIPlugin() *APIPlugin {
	o := &APIPlugin{}
	if self.plugin != nil {
		o.Using = self.plugin.GetUsing()
		// plugin won't be garaunteed to exist
		if every := self.plugin.GetEvery(); every > 0 {
			o.Every = duration.Round(every, time.Second).String()
		}
		info := self.plugin.GetInfo()
		if info.SinceLast > 0 {
			o.SinceLast = duration.Round(info.SinceLast, time.Second).String()
		}
		if !info.LastSuccess.IsZero() {
			o.LastSuccess = info.LastSuccess
		}
		o.Running = info.Running
		o.Offline = info.Offline
	}
	return o
}
func (self *Server) APIPluginRespond(
	api *APIPlugin,
	w http.ResponseWriter,
	r *http.Request,
) {
	self.Headers(w, CONTENT_JSON, 200)
	if api == nil {
		api = self.GetAPIPlugin()
	}
	bs, _ := json.MarshalIndent(api, "", "\t")
	fmt.Fprintf(w, "%s", bs)
}
