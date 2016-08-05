package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"net/http"
)

func (self *Server) Version(w http.ResponseWriter, r *http.Request) {
	self.APIVersionRespond(nil, w, r)
}
func (self *Server) VersionStart(w http.ResponseWriter, r *http.Request) {
	self.version.Start()
	self.APIVersionRespond(nil, w, r)
}
func (self *Server) VersionOnce(w http.ResponseWriter, r *http.Request) {
	self.version.StartOnce()
	self.APIVersionRespond(nil, w, r)
}
func (self *Server) VersionStop(w http.ResponseWriter, r *http.Request) {
	self.version.Stop()
	self.APIVersionRespond(nil, w, r)
}
