package main 

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    url := "http://127.0.0.1:8080/banking"
    fmt.Println("URL:>", url)

    resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}