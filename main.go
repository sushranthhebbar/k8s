package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Starting hello-world server...")
    http.HandleFunc("/", helloServer)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }

    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        duration := time.Now().Sub(started)
        if duration.Seconds() > 10 {
            w.WriteHeader(500)
            w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
        } else {
            w.WriteHeader(200)
            w.Write([]byte("ok"))
        }
    })

}

func helloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello world!")
}