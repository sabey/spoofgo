package version

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

const (
	UNSAFE = iota
	OLD_BOTH
	OLD_MINOR
	OLD_MAJOR
	SAFE
)

const (
	URL = "https://spoofgo.com/version.json/v1"
)

type Client struct {
	// safe
	client  *http.Client
	major   int
	minor   int
	build   int
	version *Version
	// unsafe
	running  bool
	stop     bool
	once     bool
	versions map[string]*Version
	mu       sync.RWMutex
}
type Version struct {
	Major  int  `json:"major,omitempty"`
	Minor  int  `json:"minor,omitempty"`
	Build  int  `json:"build,omitempty"`
	Unsafe bool `json:"unsafe,omitempty"`
}

func (self *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", self.Major, self.Minor, self.Build)
}
func Create(
	client *http.Client,
	major int,
	minor int,
	build int,
) *Client {
	if client == nil {
		log.Fatalln("./spoofgo/version.Create(): Client was nil!")
		return nil
	}
	if major < 0 {
		log.Fatalf("./spoofgo/version.Create(): Major Version: %d < 0!\n", major)
		return nil
	}
	if minor < 0 {
		log.Fatalf("./spoofgo/version.Create(): Minor Version: %d < 0!\n", minor)
		return nil
	}
	if build < 0 {
		log.Fatalf("./spoofgo/version.Create(): Build Version: %d < 0!\n", build)
		return nil
	}
	return &Client{
		client: client,
		version: &Version{
			Major: major,
			Minor: minor,
			Build: build,
		},
		versions: make(map[string]*Version),
	}
}
func (self *Client) GetVersions() []*Version {
	self.mu.RLock()
	defer self.mu.RUnlock()
	versions := make([]*Version, 0, len(self.versions))
	for _, v := range self.versions {
		versions = append(versions, v)
	}
	return versions
}
func (self *Client) GetMajor() int {
	return self.version.Major
}
func (self *Client) GetMinor() int {
	return self.version.Minor
}
func (self *Client) GetBuild() int {
	return self.version.Build
}
func (self *Client) GetVersion() *Version {
	return self.version.Clone()
}
func (self *Version) Clone() *Version {
	return &Version{
		Major: self.Major,
		Minor: self.Minor,
		Build: self.Build,
	}
}
func (self *Version) GetVersion() string {
	return fmt.Sprintf("%d.%d.%d", self.Major, self.Minor, self.Build)
}
func (self *Version) GetUserAgent() string {
	return fmt.Sprintf("Spoof.go %d.%d.%d", self.Major, self.Minor, self.Build)
}
func (self *Client) GetUserAgent() string {
	return self.version.GetUserAgent()
}
func (self *Client) GetVersionState() int {
	// default is unsafe
	// we're going to return if anything actually unsafe is found
	level := 0
	self.mu.RLock()
	defer self.mu.RUnlock()
	for _, v := range self.versions {
		if v.Major > self.version.Major {
			// there's a newer major version out
			if level == OLD_MINOR {
				// both are out of date
				level = OLD_BOTH
			} else {
				// just major is out of date
				level = OLD_MAJOR
			}
		} else if v.Major == self.version.Major {
			// current major version
			if v.Minor > self.version.Minor {
				// minor version is out of date
				if v.Unsafe {
					// older versions are considered unsafe
					return UNSAFE
				}
				// there's a newer minor version out
				if level == OLD_MAJOR {
					// both are out of date
					level = OLD_BOTH
				} else {
					// just minor is out of date
					level = OLD_MINOR
				}
			} else {
				// current major minor or older minor version
				if level == 0 {
					// nothing set yet, mark safe
					level = SAFE
				}
			}
		}
	}
	return level
}
func (self *Client) PrintVersionState() string {
	return PrintVersionState(self.GetVersionState())
}
func PrintVersionState(
	i int,
) string {
	if i == UNSAFE {
		return "Unsafe Version"
	}
	if i == OLD_BOTH {
		return "New Major/Minor Version"
	}
	if i == OLD_MINOR {
		return "New Minor Version"
	}
	if i == OLD_MAJOR {
		return "New Major Version"
	}
	return "Safe Version"
}
