package main 

import (
	"fmt" 
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("Start the server\n"))

	server := new(BankApiServer)
	server.StartListening()
}
