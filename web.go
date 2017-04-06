package main

import (
    "fmt"
    "net/http"
    "os"
    "fiscaluno-ws/institutionservice"
    "github.com/emicklei/go-restful"
)

func main() {
    http.HandleFunc("/", hello)
    fmt.Println("listening...")
    restful.Add(institutionservice.New())
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "hello, world")
}
