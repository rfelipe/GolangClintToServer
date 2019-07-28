package main

import (
    "encoding/json"
    "log"
    "net/http"
    "io/ioutil"
)

type testa_struct struct {
    Texto string
}

func testa(rw http.ResponseWriter, req *http.Request) {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        panic(err)
    }
    log.Println(string(body))
    var mensagemRecebida testa_struct
    err = json.Unmarshal(body, &mensagemRecebida)
    if err != nil {
        panic(err)
    }
    log.Println(mensagemRecebida.Texto)
}

func main() {
    http.HandleFunc("/ola-json", testa)
    log.Fatal(http.ListenAndServe(":8000", nil))
}