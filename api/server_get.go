package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"net/http"
	"sabey.co/spoofgo/movement"
)

func (self *Server) Get(w http.ResponseWriter, r *http.Request) {
	movement.Calculate()
	self.APIRespond(nil, w, r, 200)
}
