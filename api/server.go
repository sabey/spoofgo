package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
	"log"
	"net/http"
	c "sabey.co/spoofgo/api/controllers/controller"
	"sabey.co/spoofgo/news"
	p "sabey.co/spoofgo/plugins/plugin"
	"sabey.co/spoofgo/version"
	"time"
)

type Server struct {
	// safe
	server      *http.Server
	version     *version.Client
	news        *news.Client
	plugin      *p.Plugin     // can be nil
	plugins     []string      // list of plugins, since we can't include plugins package
	controller  *c.Controller // can be nil
	controllers []string      // list of controllers, since we can't include controllers package
}
type Handlers struct {
	Pattern string
	Handler http.HandlerFunc
}

func Create(
	server *http.Server,
	version *version.Client,
	news *news.Client,
	plugin *p.Plugin, // can be nil
	plugins []string, // list of plugins, since we can't include plugins package
	controller *c.Controller, // can be nil
	controllers []string, // list of controllers, since we can't include controllers package
) *Server {
	if server == nil {
		log.Fatalln("./spoofgo/api.Create(): Server was nil!")
		return nil
	}
	if version == nil {
		log.Fatalln("./spoofgo/api.Create(): Version was nil!")
		return nil
	}
	if news == nil {
		log.Fatalln("./spoofgo/api.Create(): News was nil!")
		return nil
	}
	s := &Server{
		server:      server,
		version:     version,
		news:        news,
		plugin:      plugin,
		plugins:     plugins,
		controller:  controller,
		controllers: controllers,
	}
	// override server handler
	mux := http.NewServeMux()
	for _, h := range s.GetHandlers() {
		mux.HandleFunc(h.Pattern, h.Handler)
	}
	server.Handler = mux
	return s
}
func (self *Server) ListenAndServe() error {
	return self.server.ListenAndServe()
}
func (self *Server) ListenAndServeTLS(certFile, keyFile string) error {
	return self.server.ListenAndServeTLS(certFile, keyFile)
}
func (self *Server) ServerAddr() string {
	return self.server.Addr
}
func no_cache_query() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
func (self *Server) GetHandlers() []*Handlers {
	return []*Handlers{
		// default response
		&Handlers{
			Pattern: "/",
			Handler: self.Index,
		},
		&Handlers{
			Pattern: "/get",
			Handler: self.Get,
		},
		&Handlers{
			Pattern: "/controller",
			Handler: self.Controller,
		},
		&Handlers{
			Pattern: "/controllers",
			Handler: self.Controllers,
		},
		// POST
		&Handlers{
			Pattern: "/set",
			Handler: self.Set,
		},
		&Handlers{
			Pattern: "/angle/set",
			Handler: self.SetAngle,
		},
		&Handlers{
			Pattern: "/modifier/set",
			Handler: self.SetModifier,
		},
		&Handlers{
			Pattern: "/coords/set",
			Handler: self.SetCoords,
		},
		// ACTIONS
		&Handlers{
			Pattern: "/accelerate",
			Handler: self.Accelerate,
		},
		&Handlers{
			Pattern: "/decelerate",
			Handler: self.Decelerate,
		},
		&Handlers{
			Pattern: "/right",
			Handler: self.Right,
		},
		&Handlers{
			Pattern: "/left",
			Handler: self.Left,
		},
		&Handlers{
			Pattern: "/northwest",
			Handler: self.NorthWest,
		},
		&Handlers{
			Pattern: "/northeast",
			Handler: self.NorthEast,
		},
		&Handlers{
			Pattern: "/southwest",
			Handler: self.SouthWest,
		},
		&Handlers{
			Pattern: "/southeast",
			Handler: self.SouthEast,
		},
		&Handlers{
			Pattern: "/flip",
			Handler: self.Flip,
		},
		&Handlers{
			Pattern: "/mode/toggle",
			Handler: self.ToggleMode,
		},
		&Handlers{
			Pattern: "/mode/walk",
			Handler: self.Walk,
		},
		&Handlers{
			Pattern: "/mode/jog",
			Handler: self.Jog,
		},
		&Handlers{
			Pattern: "/mode/bicycle",
			Handler: self.Bicycle,
		},
		&Handlers{
			Pattern: "/mode/car",
			Handler: self.Car,
		},
		&Handlers{
			Pattern: "/modifier/reset",
			Handler: self.ResetModifier,
		},
		&Handlers{
			Pattern: "/modifier/increase",
			Handler: self.IncreaseModifier,
		},
		&Handlers{
			Pattern: "/modifier/decrease",
			Handler: self.DecreaseModifier,
		},
		// version
		&Handlers{
			Pattern: "/version",
			Handler: self.Version,
		},
		&Handlers{
			Pattern: "/version/start",
			Handler: self.VersionStart,
		},
		&Handlers{
			Pattern: "/version/once",
			Handler: self.VersionOnce,
		},
		&Handlers{
			Pattern: "/version/stop",
			Handler: self.VersionStop,
		},
		// news
		&Handlers{
			Pattern: "/news",
			Handler: self.News,
		},
		// plugin
		&Handlers{
			Pattern: "/plugin",
			Handler: self.Plugin,
		},
		// plugins
		&Handlers{
			Pattern: "/plugins",
			Handler: self.Plugins,
		},
		// coordinates
		&Handlers{
			Pattern: "/coordinates",
			Handler: self.Coordinates,
		},
		&Handlers{
			Pattern: "/coordinates/set",
			Handler: self.SetCoordinates,
		},
		&Handlers{
			Pattern: "/coordinates/add",
			Handler: self.AddCoordinates,
		},
		&Handlers{
			Pattern: "/coordinates/delete",
			Handler: self.DeleteCoordinates,
		},
	}
}
