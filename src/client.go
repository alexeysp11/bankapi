package main 

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/jmcvetta/napping"
)

func sendToCoreServer(urlPath string, urlValues string) {
    lUrl := "http://localhost:8080" + urlPath
    fmt.Println("URL:>", lUrl)
    fmt.Println("urlValues:", urlValues)

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
	if err != nil {
		panic(err)
	}
    
	fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
}
