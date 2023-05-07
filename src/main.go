package main 

import (
	"fmt" 
	"strings"
)

func main() {
	// var server IServer;
	fmt.Println(strings.ToUpper("Start the server\n"))
	
	config := getConfig("configs/appsettings.json")
	if (config.ServerType == "grpc") {
		server := BankApiGrpcServer {config: config}
		server.config = config
		server.StartListening()
	} else {
		server := BankApiHttpServer {config: config}
		server.config = config
		server.StartListening()
	}
}
