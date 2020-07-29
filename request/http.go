package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	return string(b), err
}

func Post(url string, data string) (string, error) {
	u := ioutil.NopCloser(strings.NewReader(data))
	r, err := http.Post(url, "application/x-www-form-urlencoded", u)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	return string(b), err
}

func JsonPost(url string, data string) (string, error) {
	body := bytes.NewReader([]byte(data))
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	return string(b), err
}
