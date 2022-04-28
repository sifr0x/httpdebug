package main

import (
	"fmt"
	"net/http"
)

func debug(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Method: %v\n", req.Method)
	fmt.Fprintf(w, "URL: %v\n", req.URL)
	fmt.Fprintf(w, "Host: %v\n", req.Host)

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/", debug)
	http.ListenAndServe(":8080", nil)
}
