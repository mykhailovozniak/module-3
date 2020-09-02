/*
This package is simplified version of curl
$ cd module-3/gocurl

$ go build

$ ./gocurl https://google.com

$ ./gocurl https://google.com get
*/
package main

import (
	"fmt"
	"github.com/mykhailovozniak/module-3/arrstr"
	"github.com/mykhailovozniak/module-3/request"
	"os"
	"strings"
)

func checkUrl()  {
	url := os.Args[1]
	validProtocolScheme := strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")

	if !validProtocolScheme {
		fmt.Println("Usage: gocurl <url:required> <method:not required>- url should starts with http:// or https://")
		os.Exit(1)
	}
}

func checkMethodType()  {
	hasMethodArg := len(os.Args) == 3

	allowedMethods := []string{"get", "post", "head"}

	if hasMethodArg {
		method := os.Args[2]
		allowedMethod := arrstr.ContainsString(allowedMethods[:], method)

		if !allowedMethod {
			fmt.Println("Usage: gocurl <url> - only allowed these methods: ", allowedMethods)
			os.Exit(1)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gocurl <url:required> <method:not required>")
		os.Exit(1)
	}

	checkUrl()
	checkMethodType()

	if len(os.Args) == 2 {
		url := os.Args[1]
		res, err := request.Request("", url)

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("status code of site is: ", res.StatusCode)
		os.Exit(0)
	}

	url := os.Args[1]
	method := os.Args[2]

	res, err := request.Request(method, url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("status code of site is: ", res.StatusCode)
}
