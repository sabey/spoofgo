package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIPlugins struct {
	Using   string   `json:"using,omitempty"`
	Plugins []string `json:"plugins,omitempty"`
}

func (self *Server) GetAPIPlugins() *APIPlugins {
	o := &APIPlugins{
		Plugins: self.plugins,
	}
	if self.plugin != nil {
		// we can't include our plugins directory because it includes this library
		// we have to get a copy from our plugin object
		o.Using = self.plugin.GetUsing()
	}
	return o
}
func (self *Server) APIPluginsRespond(
	api *APIPlugins,
	w http.ResponseWriter,
	r *http.Request,
) {
	self.Headers(w, CONTENT_JSON, 200)
	if api == nil {
		api = self.GetAPIPlugins()
	}
	bs, _ := json.MarshalIndent(api, "", "\t")
	fmt.Fprintf(w, "%s", bs)
}
