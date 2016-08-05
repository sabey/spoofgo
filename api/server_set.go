package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"encoding/json"
	"net/http"
	"sabey.co/spoofgo/movement"
)

func (self *Server) Set(w http.ResponseWriter, r *http.Request) {
	// SETS EVERYTHING EXCEPT COORDINATES!
	/*
		Example:
		clear && curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST -d '{"accelerating":true,"angle":33,"mode":1,"modifier":13}' http://localhost:8844/set; echo
	*/
	if r.Method != "POST" {
		// failed
		api := self.GetAPI()
		api.Error = "POST only!"
		self.APIRespond(api, w, r, 400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	post := &API{}
	err := decoder.Decode(&post)
	if err != nil {
		// failed
		api := self.GetAPI()
		api.Error = "failed to decode"
		self.APIRespond(api, w, r, 400)
		return
	}
	// success
	if post.Accelerating {
		movement.Accelerate()
	} else {
		movement.Decelerate()
	}
	movement.SetAngle(post.Angle)
	movement.SetModifier(post.Mode, post.Modifier)
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) SetAngle(w http.ResponseWriter, r *http.Request) {
	// SETS ONLY ANGLE
	/*
		Example:
		clear && curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST -d '{"angle":359}' http://localhost:8844/angle/set; echo
	*/
	if r.Method != "POST" {
		// failed
		api := self.GetAPI()
		api.Error = "POST only!"
		self.APIRespond(api, w, r, 400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	post := &API{}
	err := decoder.Decode(&post)
	if err != nil {
		// failed
		api := self.GetAPI()
		api.Error = "failed to decode"
		self.APIRespond(api, w, r, 400)
		return
	}
	// success
	movement.SetAngle(post.Angle)
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) SetModifier(w http.ResponseWriter, r *http.Request) {
	// SETS MODE AND MODIFIER
	/*
		Example:
		clear && curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST -d '{"mode":1,"modifier":13}' http://localhost:8844/modifier/set; echo
	*/
	if r.Method != "POST" {
		// failed
		api := self.GetAPI()
		api.Error = "POST only!"
		self.APIRespond(api, w, r, 400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	post := &API{}
	err := decoder.Decode(&post)
	if err != nil {
		// failed
		api := self.GetAPI()
		api.Error = "failed to decode"
		self.APIRespond(api, w, r, 400)
		return
	}
	// success
	movement.SetModifier(post.Mode, post.Modifier)
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) SetCoords(w http.ResponseWriter, r *http.Request) {
	// SETS LAT/LONG/ACCELERATING
	/*
		Example:
		clear && curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST -d '{"latitude":49.282523,"longitude":-123.035301,"accelerating":false}' http://localhost:8844/coords/set; echo
	*/
	if r.Method != "POST" {
		// failed
		api := self.GetAPI()
		api.Error = "POST only!"
		self.APIRespond(api, w, r, 400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	post := &API{}
	err := decoder.Decode(&post)
	if err != nil {
		// failed
		api := self.GetAPI()
		api.Error = "failed to decode"
		self.APIRespond(api, w, r, 400)
		return
	}
	// success
	movement.SetLatLong(post.Latitude, post.Longitude, post.Accelerating)
	self.APIRespond(nil, w, r, 200)
}
