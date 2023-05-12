package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("GitHubユーザー名を指定してください。")
        return
    }

    username := os.Args[1]
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
