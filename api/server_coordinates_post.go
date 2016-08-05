package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"encoding/json"
	"net/http"
	"sabey.co/spoofgo/coordinates"
)

func (self *Server) SetCoordinates(w http.ResponseWriter, r *http.Request) {
	// SETS COORDINATES/ACCELERATING IF KEY EXISTS
	// returns API
	/*
		Example:
		clear && curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST -d '{"key":"vancouver","accelerating":false}' http://localhost:8844/coordinates/set; echo
	*/
	if r.Method != "POST" {
		// failed
		api := self.GetAPI()
		api.Error = "POST only!"
		self.APIRespond(api, w, r, 400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	post := &APICoordinates{}
	err := decoder.Decode(&post)
	if err != nil {
		// failed
		api := self.GetAPI()
		api.Error = "failed to decode"
		self.APIRespond(api, w, r, 400)
		return
	}
	// success
	coordinates.Set(post.Key, post.Accelerating)
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) AddCoordinates(w http.ResponseWriter, r *http.Request) {
	// ADD COORDINATES
	/*
		Example:
		clear && curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST -d '{"key":"vancouver","latitude":49.282220,"longitude":-123.121165}' http://localhost:8844/coordinates/add; echo
	*/
	if r.Method != "POST" {
		// failed
		api := &APICoordinates{
			Error: "POST only!",
		}
		self.APICoordinatesRespond(api, w, r, 400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	post := &APICoordinates{}
	err := decoder.Decode(&post)
	if err != nil {
		// failed
		api := &APICoordinates{
			Error: "POST only!",
		}
		api.Error = "failed to decode"
		self.APICoordinatesRespond(api, w, r, 400)
		return
	}
	// success
	coordinates.Add(post.Key, post.Latitude, post.Longitude)
	self.APICoordinatesRespond(nil, w, r, 200)
}
func (self *Server) DeleteCoordinates(w http.ResponseWriter, r *http.Request) {
	// ADD COORDINATES
	/*
		Example:
		clear && curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST -d '{"key":"vancouver"}' http://localhost:8844/coordinates/delete; echo
	*/
	if r.Method != "POST" {
		// failed
		api := &APICoordinates{
			Error: "POST only!",
		}
		self.APICoordinatesRespond(api, w, r, 400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	post := &APICoordinates{}
	err := decoder.Decode(&post)
	if err != nil {
		// failed
		api := &APICoordinates{
			Error: "POST only!",
		}
		api.Error = "failed to decode"
		self.APICoordinatesRespond(api, w, r, 400)
		return
	}
	// success
	coordinates.Delete(post.Key)
	self.APICoordinatesRespond(nil, w, r, 200)
}
