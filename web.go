package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/emicklei/go-restful"
    "fiscaluno-ws/routes/institutionrouter"
)

func main() {
    http.HandleFunc("/", hello)
    fmt.Println("listening...")
    restful.Add(institutionrouter.New())
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "hello, world")
}
