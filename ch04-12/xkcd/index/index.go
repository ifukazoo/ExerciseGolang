package index

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const URLPrefix = "https://xkcd.com/"
const URLSuffix = "/info.0.json"

const MaxIndex = 2000

func CreateIndex(dir string) {
	for i := 1; i < MaxIndex; i++ {
		path := dir + "/" + strconv.Itoa(i)
		var (
			bytes []byte
			f     *os.File
			err   error
		)
		if _, err = os.Stat(path); os.IsNotExist(err) {
			if bytes, err = getInfo(i); err != nil {
				goto fail
			}
			if f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644); err != nil {
				goto fail
			}
			if _, err = f.Write(bytes); err != nil {
				f.Close()
				goto fail
			}
			f.Close()
		} else {
			fmt.Printf("%04d exits.\n", i)
		}
	fail:
	}
}

func getInfo(num int) ([]byte, error) {
	url := URLPrefix + strconv.Itoa(num) + URLSuffix
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getinfo failed: %s", resp.Status)
	}
	var bytes []byte
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
