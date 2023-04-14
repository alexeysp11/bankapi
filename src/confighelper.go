package main 

import (
    "encoding/json"
    "os"
)

type Configuration struct {
    Atm 	[]string
    Eftpos 	[]string

    HttpPathsAtm 	[]string
    HttpPathsEftpos []string

	BankCoreAddress string 
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
