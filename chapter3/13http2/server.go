/*
HTTP/1.1とHTTP/2(2015~)
→HTTP/2は複数のリクエストを同時に送ることができる
そろそろ知っておきたいHTTP/2の話
https://qiita.com/mogamin3/items/7698ee3336c70a482843#:~:text=HTTP%2F2%E3%81%AF1%E3%81%A4%E3%81%AE,%E4%B8%8A%E3%81%A7%E3%82%84%E3%82%8A%E5%8F%96%E3%82%8A%E3%81%95%E3%82%8C%E3%81%BE%E3%81%99%E3%80%82
*/

package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
