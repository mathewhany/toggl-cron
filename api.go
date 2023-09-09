package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

func request(url string, method string, dataStruct interface{}) (string, error) {
	var req *http.Request
	var err error
	
	if dataStruct == nil {
		req, err = http.NewRequest(method, url, nil)
	} else if method == http.MethodGet {
		v, err := query.Values(dataStruct)
		
		if err != nil {
			return "", err
		}
		
		req, err = http.NewRequest(method, url + "?" + v.Encode(), nil)
	} else {
		data, err := json.Marshal(dataStruct)
		
		if err != nil {
			return "", err
		}
		
		log.Print(string(data))
			
		req, err = http.NewRequest(method, url, bytes.NewBuffer(data))
	}
		
	
	if err != nil {
		return "", err
	}
	
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	
	req.SetBasicAuth(email, password)
	
	client := &http.Client{}
	resp, err := client.Do(req)
	
	if err != nil {
		return "", err
	}
	
	defer resp.Body.Close()
		
	if err != nil {
		return "", err
	}

	respBody, err := io.ReadAll(resp.Body)
	
	if err != nil {
		return "", err
	}
		
	return string(respBody), nil
}
