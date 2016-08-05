package news

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	clog "sabey.co/spoofgo/log"
	"sabey.co/spoofgo/version"
	"sabey.co/textblock"
	"sync"
	"time"
)

var (
	oldage = time.Duration(time.Minute * 5)
)

const (
	URL = "https://spoofgo.com/news.json/v1"
)

type Client struct {
	// safe
	client  *http.Client
	version *version.Version
	// unsafe
	last_success time.Time
	news         []*News
	newsblock    *textblock.TextBlock
	mu           sync.RWMutex
}
type News struct {
	Title   string    `json:"title,omitempty"`
	Time    time.Time `json:"time,omitempty"`
	Content string    `json:"content,omitempty"`
}

func Create(
	client *http.Client,
	version *version.Version,
) *Client {
	if client == nil {
		log.Fatalln("./spoofgo/news.Create(): Client was nil!")
		return nil
	}
	if version == nil {
		log.Fatalln("./spoofgo/news.Create(): Version was nil!")
		return nil
	}
	o := &Client{
		client:    client,
		version:   version,
		newsblock: textblock.Create([]byte("could not load news yet!")),
	}
	o.download()
	return o
}
func (self *Client) GetVersion() *version.Version {
	return self.version.Clone()
}
func (self *Client) GetLastSuccess() time.Time {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.last_success
}
func (self *Client) GetNews() []*News {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.download()
	// dereference
	news := make([]*News, len(self.news))
	copy(news, self.news)
	return news
}
func (self *Client) GetNewsBlock() *textblock.TextBlock {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.download()
	return self.newsblock
}
func (self *Client) download() {
	if self.last_success.IsZero() || // we haven't downloaded any news yet
		time.Now().Sub(self.last_success) > oldage { // news is old
		// download news
		// we want to know if we're using the latest version right away
		request, err := http.NewRequest("GET", URL, nil)
		if err != nil {
			clog.Log(fmt.Sprintf("./spoofgo/news.download(): failed to create new request: \"%s\"", err))
		} else {
			request.Header.Set("User-Agent", self.version.GetUserAgent())
			// make request
			response, err := self.client.Do(request)
			if err != nil {
				clog.Log(fmt.Sprintf("./spoofgo/news.download(): failed to request: \"%s\"", err))
			} else {
				// made request
				if response.StatusCode != 200 {
					clog.Log(fmt.Sprintf("./spoofgo/news.download(): response.StatusCode: %d != 200", response.StatusCode))
				} else {
					// read body
					body, err := ioutil.ReadAll(response.Body)
					if err != nil {
						clog.Log(fmt.Sprintf("./spoofgo/news.download(): failed to read resposne body: \"%s\"", err))
						response.Body.Close()
					} else {
						response.Body.Close()
						news := []*News{}
						if err := json.Unmarshal(body, &news); err != nil {
							clog.Log(fmt.Sprintf("./spoofgo/news.download(): failed to unmarshal news: \"%s\"", err))
						} else {
							// news downloaded
							self.last_success = time.Now()
							self.news = news
							// make a new textblock
							buff := &bytes.Buffer{}
							for i, n := range news {
								if i > 0 {
									buff.WriteString("\n\n###\n\n")
								}
								buff.WriteString(fmt.Sprintf("# Title: %s\n", n.Title))
								buff.WriteString(fmt.Sprintf("# Time: %s\n\n", n.Time.Format(time.RFC1123)))
								buff.WriteString(n.Content)
							}
							self.newsblock = textblock.Create(buff.Bytes())
						}
					}
				}
			}
		}
	}
}
