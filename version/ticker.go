package version

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	clog "sabey.co/spoofgo/log"
	"time"
)

var (
	tick = time.Duration(time.Minute * 5)
)

func (self *Client) ticker() {
	clog.Log("./spoofgo/version.ticker(): starting")
	self.mu.Lock()
	if self.running {
		// ticker is running
		clog.Log("./spoofgo/version.ticker(): ticker is running")
		self.mu.Unlock()
		return
	}
	if self.stop {
		// ticker is stopped
		clog.Log("./spoofgo/version.ticker(): ticker is stopped")
		self.mu.Unlock()
		return
	}
	self.running = true
	self.mu.Unlock()
	clog.Log("./spoofgo/version.ticker(): started")
	defer func() {
		clog.Log("./spoofgo/version.ticker(): stopping")
		self.mu.Lock()
		self.stop = false
		self.running = false
		self.once = false
		self.mu.Unlock()
		clog.Log("./spoofgo/version.ticker(): stopped")
	}()
	// check for version
	for {
		// we want to know if we're using the latest version right away
		request, err := http.NewRequest("GET", URL, nil)
		if err != nil {
			clog.Log(fmt.Sprintf("./spoofgo/version.ticker(): failed to create new request: \"%s\"", err))
		} else {
			request.Header.Set("User-Agent", self.GetUserAgent())
			// make request
			response, err := self.client.Do(request)
			if err != nil {
				clog.Log(fmt.Sprintf("./spoofgo/version.ticker(): failed to request: \"%s\"", err))
			} else {
				// made request
				if response.StatusCode != 200 {
					clog.Log(fmt.Sprintf("./spoofgo/version.ticker(): response.StatusCode: %d != 200", response.StatusCode))
				} else {
					// read body
					body, err := ioutil.ReadAll(response.Body)
					if err != nil {
						clog.Log(fmt.Sprintf("./spoofgo/version.ticker(): failed to read resposne body: \"%s\"", err))
						response.Body.Close()
					} else {
						response.Body.Close()
						versions := []*Version{}
						if err := json.Unmarshal(body, &versions); err != nil {
							clog.Log(fmt.Sprintf("./spoofgo/version.ticker(): failed to unmarshal versions: \"%s\"", err))
						} else {
							// we don't want to remove old versions
							self.mu.Lock()
							for _, version := range versions {
								s := version.String()
								if _, ok := self.versions[s]; !ok {
									// new version
									clog.Log(fmt.Sprintf("./spoofgo/version.ticker(): Version Found: \"%s\"", s))
								}
								// update existing(incase of unsafe)/add new version
								self.versions[s] = version
							}
							self.mu.Unlock()
						}
					}
				}
			}
		}
		// should we stop looking?
		self.mu.RLock()
		if self.stop ||
			self.once {
			self.mu.RUnlock()
			// stopped
			clog.Log("./spoofgo/version.ticker(): tick - stopping")
			return
		}
		self.mu.RUnlock()
		// wait
		<-time.After(tick)
		// should we stop looking?
		self.mu.RLock()
		if self.stop {
			self.mu.RUnlock()
			// stopped
			clog.Log("./spoofgo/version.ticker(): waited - stopping")
			return
		}
		self.mu.RUnlock()
	}
}
func (self *Client) IsRunning() bool {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.running
}
func (self *Client) IsStop() bool {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.stop
}
func (self *Client) IsOnce() bool {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.once
}
func (self *Client) Stop() bool {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.stop = true // stop
	self.once = false
	if !self.running {
		// ticker not running
		clog.Log("./spoofgo/version.Stop(): ticker not running")
		return false
	}
	clog.Log("./spoofgo/version.Stop(): stopping on next tick")
	return true
}
func (self *Client) Start() bool {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.stop = false // don't stop if we triggered a stop
	self.once = false
	if self.running {
		// ticker running
		clog.Log("./spoofgo/version.Start(): ticker running")
		return false
	}
	clog.Log("./spoofgo/version.Start(): starting now")
	go self.ticker()
	return true
}
func (self *Client) StartOnce() bool {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.once = true // check version only once
	self.stop = false
	if self.running {
		// ticker running
		clog.Log("./spoofgo/version.StartOnce(): ticker running")
		return false
	}
	clog.Log("./spoofgo/version.StartOnce(): starting now")
	go self.ticker()
	return true
}
