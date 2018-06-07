package main

import (
	"net/http"
)

func main() { // 권한 필요
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, World!"))
		
	})

	http.ListenAndServe(":80",nil)
}
