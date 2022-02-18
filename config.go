//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2022 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of gocheck24.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package gocheck24

import (
	"bytes"
	"encoding/base64"
	"net/http"
)

const (
	baseUrl = "https://opentrans.shopping.check24.de/api"
)

// Config is to define config data
type Config struct {
	Path, Method string
	Body         []byte
}

// Request is to define the request data
type Request struct {
	Username, Password string
}

// Send is to send a new request
func (c *Config) Send(r Request) (*http.Response, error) {

	// Set url
	url := baseUrl + c.Path

	// Define client
	client := &http.Client{}

	// Request
	request, err := http.NewRequest(c.Method, url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	// Define basic auth
	auth := r.Username + ":" + r.Password
	basic := base64.StdEncoding.EncodeToString([]byte(auth))
	request.Header.Add("Authorization", "Basic "+basic)

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return data
	return response, nil

}
