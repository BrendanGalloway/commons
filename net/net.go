package net

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func GET(url string, args ...interface{}) ([]byte, error) {
	url = fmt.Sprintf(url, args...)

	resp, err := http.Get(url)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}

func Download(url, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	log.Debugf("GET %s -> %d\n", url, resp.StatusCode)

	if resp.StatusCode != 200 {
		return fmt.Errorf(resp.Status)
	}
	defer resp.Body.Close()
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
