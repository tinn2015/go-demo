package main

import (
	"BookStore/pkg/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("11111")
	err := http.ListenAndServe("localhost:8005", r) // 只写端口号会有防火墙弹出
	fmt.Println("error", err)
	if err != nil {
		fmt.Println("server error:", err)
		return
	}
	fmt.Println("start server on port: 8005")
}
