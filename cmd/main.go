package main

import (
    "log"
    "net/http"
    _ "net/http/pprof"
)

func fibo(a int) int {
    if a<=1 {
        return 1
    } else {
        return fibo(a-1) + fibo(a-2)
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    // w.Write([]byte("Hello, World"))
    w.Write([]byte(string(fibo(30))))
}

func main() {
    http.HandleFunc("/", handler)
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
}
