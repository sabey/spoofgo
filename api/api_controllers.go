package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIControllers struct {
	Using       string   `json:"using,omitempty"`
	Controllers []string `json:"controllers,omitempty"`
}

func (self *Server) GetAPIControllers() *APIControllers {
	o := &APIControllers{
		Controllers: self.controllers,
	}
	if self.controller != nil {
		// we can't include our controllers directory because it includes this library
		// we have to get a copy from our controller object
		o.Using = self.controller.GetUsing()
	}
	return o
}
func (self *Server) APIControllersRespond(
	api *APIControllers,
	w http.ResponseWriter,
	r *http.Request,
) {
	self.Headers(w, CONTENT_JSON, 200)
	if api == nil {
		api = self.GetAPIControllers()
	}
	bs, _ := json.MarshalIndent(api, "", "\t")
	fmt.Fprintf(w, "%s", bs)
}
