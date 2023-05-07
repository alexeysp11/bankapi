package main

import (
    // "net/http"
    // "io"
    // "time"
	// "log"
	// "io/ioutil"
	"fmt"
)

type BankApiGrpcServer struct {
	// mux *http.ServeMux
	config Configuration
}

func (s *BankApiGrpcServer) StartListening() {
	// 
	fmt.Println(s.config.ServerType)
}
