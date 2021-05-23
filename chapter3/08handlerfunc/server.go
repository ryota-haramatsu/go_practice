package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// Handlerは指定しない -> DefaultServeMuxをハンドラとして利用
	}
	// HandleFuncで関数をハンドラに変換
	// /foo というURLに対する関数を割り当てることができる
	// HandleFunc(http.HandleFunc()のこと)とHandlerFunc(型の名前)は別物なので注意
	http.HandleFunc("/hello", hello) // 関数をハンドラに変換して、DefaultServeMuxに登録
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
