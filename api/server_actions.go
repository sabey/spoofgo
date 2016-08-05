package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"net/http"
	"sabey.co/spoofgo/movement"
)

func (self *Server) Accelerate(w http.ResponseWriter, r *http.Request) {
	movement.Accelerate()
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) Decelerate(w http.ResponseWriter, r *http.Request) {
	movement.Decelerate()
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) Right(w http.ResponseWriter, r *http.Request) {
	movement.Right()
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) Left(w http.ResponseWriter, r *http.Request) {
	movement.Left()
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) NorthWest(w http.ResponseWriter, r *http.Request) {
	movement.NorthWest()
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) NorthEast(w http.ResponseWriter, r *http.Request) {
	movement.NorthEast()
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) SouthWest(w http.ResponseWriter, r *http.Request) {
	movement.SouthWest()
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) SouthEast(w http.ResponseWriter, r *http.Request) {
	movement.SouthEast()
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) Flip(w http.ResponseWriter, r *http.Request) {
	movement.Flip()
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) ToggleMode(w http.ResponseWriter, r *http.Request) {
	movement.ToggleMode()
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) Walk(w http.ResponseWriter, r *http.Request) {
	movement.SetMode(movement.WALK)
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) Jog(w http.ResponseWriter, r *http.Request) {
	movement.SetMode(movement.JOG)
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) Bicycle(w http.ResponseWriter, r *http.Request) {
	movement.SetMode(movement.BICYCLE)
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) Car(w http.ResponseWriter, r *http.Request) {
	movement.SetMode(movement.CAR)
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) ResetModifier(w http.ResponseWriter, r *http.Request) {
	movement.ResetModifier()
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) IncreaseModifier(w http.ResponseWriter, r *http.Request) {
	movement.IncreaseModifier()
	self.APIRespond(nil, w, r, 200)
}
func (self *Server) DecreaseModifier(w http.ResponseWriter, r *http.Request) {
	movement.DecreaseModifier()
	self.APIRespond(nil, w, r, 200)
}
