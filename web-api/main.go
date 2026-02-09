package main

import (
	"fmt"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Shapes API"))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SHU-WEI"))
}

func currentTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	w.Write([]byte(fmt.Sprintf("Current server time: %s", now)))
}

// Custom route: Greeting
func greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SHU-WEI’s Shapes API!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/time", currentTime)
	mux.HandleFunc("/greeting", greeting)

	fmt.Println("Starting server on :4000")
	http.ListenAndServe(":4000", mux)
}
