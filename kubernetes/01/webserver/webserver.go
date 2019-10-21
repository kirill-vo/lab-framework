package main

import (
    "fmt"
    "log"
    "net/http"

    "io/ioutil"
    "strings"
)

func check(){
    return new page || current page
}

func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")

    switch r.Method {
        case "GET":
            fmt.Printf("Getting GET...\n")
            http.ServeFile(w, r, "index.html")
        case "POST":
            if err := r.ParseForm(); err != nil {
                fmt.Fprintf(w, "ParseForm() err: %v", err)
                return
            }
            fmt.Printf("Getting POST...\n")
            // http.ServeFile(w, r, check())
    }
}

func main() {
    input, err := ioutil.ReadFile("/var/_data/intro.md")
    if err != nil {
            log.Fatalln(err)
    }

    lines := strings.Split(string(input), "\n")

    for i, line := range lines {
            if strings.Contains(line, "]") {
                    lines[i] = "LOL"
            }
    }
    output := strings.Join(lines, "\n")
    err = ioutil.WriteFile("/var/www/html/index.html", []byte(output), 0644)
    if err != nil {
            log.Fatalln(err)
    }



    http.HandleFunc("/", hello)

    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
        log.Fatal(err)
    }
}
