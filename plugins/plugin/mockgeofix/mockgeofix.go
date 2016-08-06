package mockgeofix

import (
	"bytes"
	"fmt"
	"net"
	"sabey.co/spoofgo/api"
	clog "sabey.co/spoofgo/log"
	p "sabey.co/spoofgo/plugins/plugin"
	"sync"
	"time"
)

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

var (
	mu  sync.Mutex
	con *connection
)

type connection struct {
	con net.Conn
}

func MOCKGEOFIX(
	plugin *p.Plugin,
	a *api.API,
) bool {
	var err error
	var c *connection
	mu.Lock()
	if con == nil {
		// new connection
		o := &connection{}
		state := plugin.GetState()
		o.con, err = net.DialTimeout("tcp", state.Addr, time.Second*5)
		clog.Log(fmt.Sprintf("./spoofgo/plugins/plugin/mockgeofix.MOCKGEOFIX(): dial \"%s\"", plugin.GetState().Addr))
		if err != nil {
			mu.Unlock()
			// failed to connect
			clog.Log(fmt.Sprintf("./spoofgo/plugins/plugin/mockgeofix.MOCKGEOFIX(): failed to dial \"%s\"", err))
			return false
		}
		// connected!
		clog.Log(fmt.Sprintf("./spoofgo/plugins/plugin/mockgeofix.MOCKGEOFIX(): connected \"%s\"", plugin.GetState().Addr))
		con = o
		c = con
		mu.Unlock()
		// send password if it exists
		clog.Log("./spoofgo/plugins/plugin/mockgeofix.MOCKGEOFIX(): sending password if it exists")
		c.password(plugin.GetState())
	} else {
		// old connection
		c = con
		mu.Unlock()
	}
	// send coords
	// the api seems to be "longitude latitude" not lat long
	// also may support altitude? we don't support that yet
	// also supports nmea but writing long/lat is fine
	if flip(plugin.GetState()) {
		// lat/long
		_, err = c.con.Write([]byte(fmt.Sprintf("geo fix %f %f\n", a.Latitude, a.Longitude)))
	} else {
		// long/lat
		// THIS IS THIS API'S DEFAULT POSITION
		_, err = c.con.Write([]byte(fmt.Sprintf("geo fix %f %f\n", a.Longitude, a.Latitude)))
	}
	if err != nil {
		clog.Log(fmt.Sprintf("./spoofgo/plugins/plugin/mockgeofix.MOCKGEOFIX(): failed to write \"%s\"", err))
		c.restart(err)
		return false
	}
	// read response
	reply := make([]byte, 256)
	_, err = c.con.Read(reply)
	if err != nil {
		clog.Log(fmt.Sprintf("./spoofgo/plugins/plugin/mockgeofix.MOCKGEOFIX(): failed to read \"%s\"", err))
		//c.restart(err)
		return false
	}
	// sent coords
	if bytes.HasPrefix(reply, []byte("OK")) {
		// success
		return true
	}
	if bytes.HasPrefix(reply, []byte("KO: password required")) {
		// send password again
		c.password(plugin.GetState())
		clog.Log("./spoofgo/plugins/plugin/mockgeofix.MOCKGEOFIX(): sending password again")
		// don't read response
		return false
	}
	if bytes.HasPrefix(reply, []byte("MockGeoFix:")) {
		// ignore error
		return false
	}
	// failed
	clog.Log(fmt.Sprintf("./spoofgo/plugins/plugin/mockgeofix.MOCKGEOFIX(): command failed, response: \"%s\"", reply))
	return false
}
func (self *connection) restart(
	err error,
) {
	// todo: this could be better
	clog.Log(fmt.Sprintf("./spoofgo/plugins/plugin/mockgeofix.MOCKGEOFIX(): restarting connection since we recieved an error: \"%s\"", err))
	// doing so in a goroutine incase this is misused
	go func() {
		mu.Lock()
		// kill connection
		go con.con.Close()
		con = nil
		mu.Unlock()
	}()
}
func (self *connection) password(
	state *p.State,
) {
	if v, ok := state.KeyValue["password"]; ok {
		if pass, ok := v.(string); ok {
			clog.Log(fmt.Sprintf("./spoofgo/plugins/plugin/mockgeofix.MOCKGEOFIX(): writing password: \"%s\"\n", pass))
			_, err := self.con.Write([]byte(fmt.Sprintf("password %s\n", pass)))
			if err != nil {
				clog.Log(fmt.Sprintf("./spoofgo/plugins/plugin/mockgeofix.MOCKGEOFIX(): failed to write password \"%s\"", err))
				return
			}
			// sent password
		} else {
			clog.Log(fmt.Sprintf("./spoofgo/plugins/plugin/mockgeofix.MOCKGEOFIX(): password set but not a string, ignoring! \"%v\"", v))
		}
	}
}
func flip(
	state *p.State,
) bool {
	// just incase this is needed in the future
	if v, ok := state.KeyValue["flip-coords"]; ok {
		if flip, ok := v.(bool); ok && flip {
			// flip, return lat/long
			return true
		}
	}
	// don't flip, return long/lat
	return false
}
