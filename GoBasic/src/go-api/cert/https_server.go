package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	/*
		openssl genrsa -out server.key 2048
		set OPENSSL_CONF=D:\Application\OpenSSL-1.1.1h_win32\OpenSSL-1.1.1h_win32\openssl.cnf
		openssl req -new -x509 -key server.key -out server.crt -days 3650
		请求 https://127.0.0.1:8080/bar  显示不安全
	*/
	err := server.ListenAndServeTLS("d:/tmp/server.crt", "d:/tmp/server.key")
	if err != nil {
		panic(err)
	}

}
