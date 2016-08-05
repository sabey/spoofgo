package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sabey.co/spoofgo/version"
)

type APIVersion struct {
	Version      string             `json:"version"`
	VersionState string             `json:"version-state"`
	Major        int                `json:"major"`
	Minor        int                `json:"minor"`
	Build        int                `json:"build"`
	Versions     []*version.Version `json:"versions,omitempty"`
	Running      bool               `json:"running"`
	Stopping     bool               `json:"stopping"`
	Once         bool               `json:"once"`
}

func (self *Server) GetAPIVersion() *APIVersion {
	return &APIVersion{
		Version:      self.version.GetUserAgent(),
		VersionState: self.version.PrintVersionState(),
		Major:        self.version.GetMajor(),
		Minor:        self.version.GetMinor(),
		Build:        self.version.GetBuild(),
		Versions:     self.version.GetVersions(),
		Running:      self.version.IsRunning(),
		Stopping:     self.version.IsStop(),
		Once:         self.version.IsOnce(),
	}
}
func (self *Server) APIVersionRespond(
	api *APIVersion,
	w http.ResponseWriter,
	r *http.Request,
) {
	self.Headers(w, CONTENT_JSON, 200)
	if api == nil {
		api = self.GetAPIVersion()
	}
	bs, _ := json.MarshalIndent(api, "", "\t")
	fmt.Fprintf(w, "%s", bs)
}
