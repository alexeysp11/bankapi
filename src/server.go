package main

import (
    "net/http"
    "io"
    "time"
	"log"
)


func myHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	// define a serveMux to handle routes
	mux := http.NewServeMux()
	// assign a route/todo to a handler myHandler
	mux.HandleFunc("/banking", myHandler)
	mux.HandleFunc("/banking/atm", myHandler)
	mux.HandleFunc("/banking/serviceapp", myHandler)
 
	// assign a route/banking/notes to an anonymous func
	mux.HandleFunc("/banking/notes", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello again."))
	})
 
	s := &http.Server {
	   Addr:           "127.0.0.1:8080",
	   Handler:        mux,
	   ReadTimeout:    10 * time.Second,
	   WriteTimeout:   10 * time.Second,
	   MaxHeaderBytes: 1 << 20,
	}
 
	if err := s.ListenAndServe(); err != nil {
	   log.Fatalf("server failed to start with error %v", err.Error())
	}
}