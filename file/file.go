package file

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"bytes"
	"io"
	"os"
	"sabey.co/spoofgo/log"
)

func Open(
	file string,
) (
	[]byte,
	error,
) {
	f, err := os.Open(file)
	if err != nil {
		log.Log(err.Error())
		return nil, err
	}
	defer f.Close()
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, f)
	if err != nil {
		log.Log(err.Error())
		return nil, err
	}
	return buf.Bytes(), nil
}
func Save(
	file string,
	bs []byte,
) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Log(err.Error())
		return err
	}
	defer f.Close()
	_, err = f.Write(bs)
	if err != nil {
		log.Log(err.Error())
		return err
	}
	return nil
}
