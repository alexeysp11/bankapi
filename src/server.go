package main

import (
    "net/http"
    "io"
    "time"
	"log"
	"io/ioutil"
	"fmt"
)

type BankApiServer struct {
	mux *http.ServeMux
	config Configuration
}

func (s *BankApiServer) MyHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request Headers:", req.Header)
    body, _ := ioutil.ReadAll(req.Body)
    // fmt.Println("request URL:", string(req.URL.Path))
    fmt.Println("request Body:", string(body))
	// fmt.Println("request param:", req.FormValue("id"))

	// parse request 

	// send request to the core 
	sendToCoreServer(req.URL.Path, string(body))
	fmt.Println("BankCoreAddress:", s.config.BankCoreAddress)
	result := "Hello, world!\n" 

	io.WriteString(w, result)
}

func (s *BankApiServer) AddHandleFunc(urlPattern string, configIds []string, configPaths []string) {
	for i := 0; i < len(configIds); i++ {
		for j := 0; j < len(configPaths); j++ {
			s.mux.HandleFunc(urlPattern + configIds[i] + configPaths[j], s.MyHandler)
		}
	}
}

func (s *BankApiServer) StartListening() {
	// define a serveMux to handle routes
	s.mux = http.NewServeMux()

	// assign a routes to a handler 
	s.config = getConfig("server.json")
	s.AddHandleFunc("/banking/atm/", s.config.Atm, s.config.HttpPathsAtm)
	s.AddHandleFunc("/banking/eftpos/", s.config.Eftpos, s.config.HttpPathsEftpos)
 
	// assign a route to an anonymous func
	s.mux.HandleFunc("/banking/notes", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Some notes..."))
	})
 
	httpServer := &http.Server {
		Addr:           "127.0.0.1:8081",
		Handler:        s.mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
 
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start with error %v", err.Error())
	}
}
