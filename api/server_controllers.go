package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"net/http"
)

func (self *Server) Controllers(w http.ResponseWriter, r *http.Request) {
	self.APIControllersRespond(nil, w, r)
}
