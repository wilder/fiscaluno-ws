package main

import (
    "fmt"
    "net/http"
    "os"
    "fiscaluno-ws/services/institution"
    "github.com/emicklei/go-restful"
)

func main() {
    http.HandleFunc("/", hello)
    fmt.Println("listening...")
    restful.Add(institution.New())
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "hello, world")
}
