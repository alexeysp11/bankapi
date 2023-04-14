package main 

import (
    "fmt"
    // "io/ioutil"
    "net/http"
    // "net/url"
    "encoding/json"
    "github.com/jmcvetta/napping"
)

func sendToCoreServer(urlPath string, urlValues string) {
    lUrl := "http://localhost:8080" + urlPath
    fmt.Println("URL:>", lUrl)
    fmt.Println("urlValues:", urlValues)

    // resp, err := http.Get(lUrl)
    // resp, err := http.PostForm(lUrl, url.Values{
    //     "ln": {"12"},
    //     "ip": {"4252"},
    //     "ua": {"342"}})
    // v := url.Values{}
    // v.Encode()

    s := napping.Session{}
    h := &http.Header{}
    h.Set("X-Custom-Header", "myvalue")
    s.Header = h
    var jsonStr = []byte("{ \"bankapi\": \"" + urlValues + "\" }")
    var data map[string]json.RawMessage
    err := json.Unmarshal(jsonStr, &data)
    if err != nil {
        fmt.Println(err)
    }
    resp, err := s.Post(lUrl, &data, nil, nil)

    // resp, err := http.PostForm(lUrl, v)
	if err != nil {
		panic(err)
	}
    
	fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    // body, _ := ioutil.ReadAll(resp.Body)
    // fmt.Println("response Body:", string(body))
}
