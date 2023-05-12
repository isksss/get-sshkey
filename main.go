package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    username := "isksss"
    url := fmt.Sprintf("https://github.com/%s.keys", username)

    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(body))
}
