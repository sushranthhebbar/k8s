package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    started := time.Now()
    fmt.Println("Starting hello-world server...")
    http.HandleFunc("/", helloServer)

    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        duration := time.Now().Sub(started)
        if duration.Seconds() > 20 {
            w.WriteHeader(200)
            w.Write([]byte("ok"))
            fmt.Println(">5")
        } else {
            w.WriteHeader(500)
            w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
            fmt.Println("<5")
        }
    })
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}

func helloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello world")
}