package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "sort"
    "strings"
)

const IndexHTML = `
<h1>hello world!</h1>
`

func main() {
    http.HandleFunc('/', Index)
    http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
    url := r.FormValue("url")
    if url == "" {
        fmt.Fprint(w, IndexHTML)
        return
    }
    key := "Placeholder"
    fmt.Fprintf(w, "http://localhost:8080/%s", key)
}

type response struct {
    Message string   `json:"message"`
    EnvVars []string `json:"env"`
    Fib     []int    `json:"fib"`
}

func fib() func() int {
    a, b := 0, 1
    return func() int {
        a, b = b, a+b
        return a
    }
}
