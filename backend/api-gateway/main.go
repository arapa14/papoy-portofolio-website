package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Server
	fmt.Println("Server running at port: 8080")
	http.ListenAndServe(":8080", nil)
}