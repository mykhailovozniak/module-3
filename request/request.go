// This package will contains method for http
package request

import (
	"fmt"
	"net/http"
)

// This method perform different http method requests
func Request(method string, url string) (resp *http.Response, err error) {
	fmt.Printf("METHOD: %s, URL: %s\n", method, url)

	switch method {
	case "get":
		return http.Get(url)
	case "post":
		return http.Post(url, "", nil)
	case "head":
		return  http.Head(url)
	}

	return http.Head(url)
}
