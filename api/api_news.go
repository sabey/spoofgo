package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sabey.co/spoofgo/news"
	"time"
)

type APINews struct {
	News        []*news.News `json:"news,omitempty"`
	LastSuccess time.Time    `json:"last-success,omitempty"`
}

func (self *Server) GetAPINews() *APINews {
	o := &APINews{
		News: self.news.GetNews(),
	}
	o.LastSuccess = self.news.GetLastSuccess()
	return o
}
func (self *Server) APINewsRespond(
	api *APINews,
	w http.ResponseWriter,
	r *http.Request,
) {
	self.Headers(w, CONTENT_JSON, 200)
	if api == nil {
		api = self.GetAPINews()
	}
	bs, _ := json.MarshalIndent(api, "", "\t")
	fmt.Fprintf(w, "%s", bs)
}
