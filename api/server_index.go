package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
	"net/http"
	"sabey.co/spoofgo/duration"
	"time"
)

func (self *Server) Index(w http.ResponseWriter, r *http.Request) {
	self.Headers(w, CONTENT_HTML, 200)
	fmt.Fprintln(w, `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta http-equiv="Cache-Control" content="no-cache, no-store, must-revalidate" />
<meta http-equiv="Pragma" content="no-cache" />
<meta http-equiv="Expires" content="0" />
<title>Spoof.go</title>
</head>
<body>`)
	fmt.Fprintf(w, "Version: %s %s<br />\n", self.version.GetUserAgent(), self.version.PrintVersionState())
	fmt.Fprintf(w, "Website: <a href=\"http://spoofgo.com/?spoofgo=%s\">http://spoofgo.com</a> - for the latest news, guides and downloads<br />\n", self.version.GetVersion())
	fmt.Fprintf(w, "Email: <a href=\"mailto:spoofgo@gmail.com?subject=%s\">spoofgo@gmail.com</a> - feel free to send feeback<br />\n", self.version.GetUserAgent())
	fmt.Fprintf(w, "Donations: <a href=\"http://spoofgo.com/donate?spoofgo=%s\">http://spoofgo.com/donate</a><br />\n", self.version.GetVersion())
	fmt.Fprintln(w, "we are an open source project, contributions are appreciated!<br /><hr />")
	if self.plugin == nil {
		fmt.Fprintln(w, "<h3>Plugin Not Selected!</h3>")
	} else {
		fmt.Fprintf(w, "<h3>Running Plugin \"%s\" every %s!</h3>\n", self.plugin.GetUsing(), self.plugin.GetEvery().String())
		info := self.plugin.GetInfo()
		if info.LastSuccess.IsZero() {
			fmt.Fprintln(w, "<h1 style=\"color: red;\">Plugin Not Online Yet!</h1>")
		} else {
			if info.Offline {
				fmt.Fprintf(w, "<h1 style=\"color: red;\">Plugin Offline For %s</h1>\n", duration.Round(info.SinceLast, time.Second).String())
			}
		}
	}
	fmt.Fprintln(w, "<hr />")
	if self.controller == nil {
		fmt.Fprintln(w, "<h1 style=\"color: red;\">Controller Not Selected!</a></h1>")
	} else {
		fmt.Fprintf(w, "<h1><a href=\"/controller?%s\">Controller</a></h1>\n", no_cache_query())
	}
	fmt.Fprintln(w, "<hr />")
	fmt.Fprintln(w, "API:<br />\n<ul>")
	for _, h := range self.GetHandlers() {
		if h.Pattern != "/" &&
			h.Pattern != "/controller" {
			fmt.Fprintf(w, "<li><a href=\"%s?%s\">%s</a></li>\n", h.Pattern, no_cache_query(), h.Pattern)
		}
	}
	fmt.Fprintln(w, `</ul>
</body>
</html>`)
}
