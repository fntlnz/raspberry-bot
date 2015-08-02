package sysinfo

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func IPAddress() (string, error) {
	res, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "", fmt.Errorf("An error occurred determining external ip address: %s", err.Error())
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", fmt.Errorf("An error occurred getting content for external ip address: %s", err.Error())
	}
	return string(content), nil
}
