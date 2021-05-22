package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{} // handlerはハンドラ。
	server := http.Server{
		Addr: "127.0.0.1:8080",
		//		Addr:    ":8080", // これでもOK

		// 指定すると単一のハンドラに全てのリクエストを処理
		// 指定しないとDefaultServeMuxををハンドラとして使用する
		// handler to invoke, http.DefaultServeMux if nil
		Handler: &handler,
	}
	server.ListenAndServe()
}
