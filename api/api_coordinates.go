package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sabey.co/spoofgo/coordinates"
)

type APICoordinates struct {
	Key          string                             `json:"key,omitempty"`
	Latitude     float64                            `json:"latitude,omitempty"`
	Longitude    float64                            `json:"longitude,omitempty"`
	Accelerating bool                               `json:"accelerating,omitempty"`
	Coordinates  map[string]*coordinates.Coordinate `json:"coordinates,omitempty"`
	Error        string                             `json:"error,omitempty"`
}

func (self *Server) GetAPICoordinates() *APICoordinates {
	o := &APICoordinates{
		Coordinates: coordinates.List(),
	}
	return o
}
func (self *Server) APICoordinatesRespond(
	api *APICoordinates,
	w http.ResponseWriter,
	r *http.Request,
	code int,
) {
	if code == 0 {
		code = 200
	}
	self.Headers(w, CONTENT_JSON, code)
	if api == nil {
		api = self.GetAPICoordinates()
	}
	bs, _ := json.MarshalIndent(api, "", "\t")
	fmt.Fprintf(w, "%s", bs)
}
