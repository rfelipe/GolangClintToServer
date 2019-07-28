package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
    "bytes"
)
type Mensagem struct {
    Texto    string `json:"texto"`
}

func main() {
    url := "http://localhost:8000/ola-json"
    fmt.Println("URL:>", url)

    values := &Mensagem{
					    Texto: "Ola",
			}

	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

    if err != nil {
        panic(err)
    }
    fmt.Println(err)
    defer resp.Body.Close()

    fmt.Println("response		:", resp)
    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}