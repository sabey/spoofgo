package standard

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sabey.co/spoofgo/api"
	clog "sabey.co/spoofgo/log"
	p "sabey.co/spoofgo/plugins/plugin"
)

// default is a reserved keyword in golang so we're calling this package standard
func HTTP(
	plugin *p.Plugin,
	a *api.API,
) bool {
	bs, _ := json.Marshal(a)
	request, err := http.NewRequest("POST", plugin.GetState().Addr, bytes.NewBuffer(bs))
	if err != nil {
		clog.Log(fmt.Sprintf("./spoofgo/plugins/plugin/standard.HTTP(): Failed to Create Request \"%s\"", err))
		return false
	}
	request.Header.Set("User-Agent", plugin.GetVersion().GetUserAgent())
	request.Header.Set("Content-Type", "application/json")
	response, err := plugin.GetState().HTTPClient.Do(request)
	if err != nil {
		clog.Log(fmt.Sprintf("./spoofgo/plugins/plugin/standard.HTTP(): Failed to Do Request \"%s\"", err))
		return false
	}
	defer response.Body.Close()
	return true
}
func Log(
	plugin *p.Plugin,
	a *api.API,
) bool {
	bs, _ := json.Marshal(a)
	clog.Log(fmt.Sprintf("./spoofgo/plugins/plugin/standard.Log(): \"%s\"", bs))
	return true
}
