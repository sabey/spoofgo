package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sabey.co/spoofgo/movement"
)

type API struct {
	Angle        int     `json:"angle"`
	Mode         int     `json:"mode"`
	ModeString   string  `json:"mode-string,omitempty"`
	Modifier     int     `json:"modifier"`
	Speed        float64 `json:"speed"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Accelerating bool    `json:"accelerating"`
	Error        string  `json:"error,omitempty"`
}

func (self *Server) GetAPI() *API {
	movement.Calculate()
	x := &API{
		Angle:        movement.GetAngle(),
		Speed:        movement.GetSpeed(),
		Accelerating: movement.IsAccelerating(),
	}
	x.Mode, x.Modifier = movement.GetModeModifier()
	x.Latitude, x.Longitude = movement.GetLatLong()
	x.ModeString = movement.PrintMode(x.Mode, true)
	return x
}
func (self *Server) APIRespond(
	api *API,
	w http.ResponseWriter,
	r *http.Request,
	code int,
) {
	if r.URL != nil && r.URL.Query().Get("controller") == "index" {
		// client is using the controller
		if r.URL.Query().Get("reload") == "no" {
			w.Header().Set("Location", fmt.Sprintf("/controller?reload=no&%s", no_cache_query()))
		} else {
			w.Header().Set("Location", fmt.Sprintf("/controller?%s", no_cache_query()))
		}
		self.Headers(w, CONTENT_JSON, 301)
	} else {
		if code == 0 {
			code = 200
		}
		self.Headers(w, CONTENT_JSON, code)
	}
	if api == nil {
		api = self.GetAPI()
	}
	bs, _ := json.MarshalIndent(api, "", "\t")
	fmt.Fprintf(w, "%s", bs)
}
