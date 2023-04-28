package main 

import (
    "encoding/json"
    "os"
    "time"
)

type Configuration struct {
	Environment string 

	ServerHost string 
	ServerPort string 
	ServerReadTimeout time.Duration 
	ServerWriteTimeout time.Duration 

	BankCoreAddress string 

    Atm 	[]string
    Eftpos 	[]string

    HttpPathsAtm 	[]string
    HttpPathsEftpos []string
}

func getConfig(aFilename string) Configuration {
	file, _ := os.Open(aFilename)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}
