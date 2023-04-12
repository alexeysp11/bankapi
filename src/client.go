package main 

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func main() {
    lUrl := "http://127.0.0.1:8080/banking/atm/v1"
    fmt.Println("URL:>", lUrl)

    // resp, err := http.Get(lUrl)
    resp, err := http.PostForm(lUrl, url.Values{
        "ln": {"12"},
        "ip": {"4252"},
        "ua": {"342"}})
	if err != nil {
		panic(err)
	}
	fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}