package myhttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {
	var (
		resp *http.Response
		err  error
	)
	if resp, err = http.Get(url); err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("httpGet failed: %s", resp.Status)
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	return bytes, nil
}
