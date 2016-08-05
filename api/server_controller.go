package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"bytes"
	"fmt"
	"net/http"
)

func (self *Server) Controller(
	w http.ResponseWriter,
	r *http.Request,
) {
	if self.controller == nil {
		// failed
		self.Headers(w, CONTENT_HTML, 400)
		fmt.Fprintln(w, `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta http-equiv="Cache-Control" content="no-cache, no-store, must-revalidate" />
<meta http-equiv="Pragma" content="no-cache" />
<meta http-equiv="Expires" content="0" />
<title>Spoof.go | Controller Not Selected!</title>
</head>
<body>
<h1>Controller Not Selected!</h1>
</body>
</html>`)
		return
	}
	self.Headers(w, CONTENT_HTML, 200)
	// this should be replaced with html templates
	buffer := &bytes.Buffer{}
	// call controller
	self.controller.Controller(buffer, r)
	// copy the buffer to w
	buffer.WriteTo(w)
	// that's it!!!
}
