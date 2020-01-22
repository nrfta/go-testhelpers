package testhelpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func ExecutePost(path string, inputs ...interface{}) ([]byte, *http.Response, error) {
	var input interface{}

	if len(inputs) > 0 {
		input = inputs[0]
	}

	b, err := json.Marshal(input)
	if err != nil {
		return nil, nil, err
	}

	resp, err := http.Post(path, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()

	if err != nil {
		return nil, resp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	return body, resp, nil
}

func ExecutePostForm(path string, form url.Values) ([]byte, *http.Response, error) {
	resp, err := http.PostForm(path, form)

	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()

	if err != nil {
		return nil, resp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	return body, resp, nil
}

func ExecuteGet(path string) ([]byte, *http.Response, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	return body, resp, nil
}

func ExecuteDelete(path string, inputs ...interface{}) ([]byte, *http.Response, error) {
	var input interface{}

	if len(inputs) > 0 {
		input = inputs[0]
	}

	b, err := json.Marshal(input)
	if err != nil {
		return nil, nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", path, bytes.NewBuffer(b))
	if err != nil {
		return nil, nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()

	if err != nil {
		return nil, resp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	return body, resp, nil
}
