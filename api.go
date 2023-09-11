package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/google/go-querystring/query"
	"golang.org/x/time/rate"
)

var rateLimiter *rate.Limiter

func init() {
	reqPerSec, err := strconv.Atoi(os.Getenv("TOGGL_REQ_PER_SEC"))

	if err != nil {
		log.Fatal(err)
	}

	rateLimiter = rate.NewLimiter(rate.Limit(reqPerSec), 1)
}

func request(url string, method string, body interface{}, response interface{}) error {
	err := rateLimiter.Wait(context.Background())

	if err != nil {
		return err
	}

	var bodyBytes []byte

	if body != nil {
		if method == http.MethodGet {
			v, err := query.Values(body)
			if err != nil {
				return err
			}
			url = url + "?" + v.Encode()
		} else {
			jsonBytes, err := json.Marshal(body)

			if err != nil {
				return err
			}

			bodyBytes = jsonBytes
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyBytes))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	req.SetBasicAuth(email, password)

	log.Print("Request URL: ", req.URL)
	log.Print("Request Method: ", req.Method)
	log.Print("Request Headers: ", req.Header)
	log.Print("Request Body: ", string(bodyBytes))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err != nil {
		return err
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	log.Print("Response Body: ", string(respBody))
	log.Print("Response Status: ", resp.StatusCode)
	log.Print("Response Headers: ", resp.Header)

	if response != nil {
		return json.Unmarshal(respBody, response)
	}

	return nil
}
