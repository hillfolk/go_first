package main

import (
	"net/http"
)


func main() { // 관리자 권한 필요 
	mux := http.NewServeMux()

	mux.HandleFunc("/",func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world!"))
	})
	http.ListenAndServe(":80",mux)
}
