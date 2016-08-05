package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"net/http"
)

const (
	CONTENT_JSON = iota + 1
	CONTENT_HTML
)

func (self *Server) Headers(
	w http.ResponseWriter,
	content int,
	code int,
) {
	if content == CONTENT_JSON {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	}
	w.Header().Set("Server", self.version.GetUserAgent())
	// no cache
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.WriteHeader(code)
}
