package main

import (
	"fmt"
	"net/http"
	"os"
)

// API entry point
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello")) })
	fmt.Println("API starts on http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Start API error: ", err.Error())
		os.Exit(1)
	}
}
