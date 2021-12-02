// Implement a http and https request
package curl

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	ContentType                   = "Content-Type"
	ApplicationJson               = "application/json"
	ApplicationXml                = "application/xml"
	ApplicationForm               = "application/x-www-form-urlencoded"
	Timeout         time.Duration = 5
)

// GET
func Get(url string, header url.Values) (res string, err error) {
	ts := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: ts, Timeout: Timeout * time.Second}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Set the header
	if header != nil {
		for key, val := range header {
			request.Header.Add(key, val[0])
		}
	}

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// POST
func Post(url string, params url.Values, header url.Values) (res string, err error) {
	ts := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: ts, Timeout: Timeout * time.Second}

	var request *http.Request
	if params != nil {
		buffer := bytes.NewBuffer([]byte(params.Encode()))
		request, err = http.NewRequest("POST", url, buffer)
	} else {
		request, err = http.NewRequest("POST", url, nil)
	}

	if err != nil {
		return "", err
	}

	// Set the header
	if header != nil {
		for key, val := range header {
			request.Header.Add(key, val[0])
		}
	}

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
