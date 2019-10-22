package main

import (
    "fmt"
    "log"
    "net/http"

    // "io" // copy with files
    // "os" // copy with files
    "io/ioutil" // copy with Asset
    "gopkg.in/yaml.v2"
)

var current_step int = 0



// func Copy(src, dst string) bool {
//     in, err := os.Open(src)
//     if err != nil {
//         return false
//     }
//     defer in.Close()

//     out, err := os.Create(dst)
//     if err != nil {
//         return false
//     }
//     defer out.Close()

//     _, err = io.Copy(out, in)
//     if err != nil {
//         return false
//     }
//     out.Close()
//     return true
// }



func Copy(src, dst string) bool {
    // read data from Asset
    data, err := Asset(src)
    if err != nil {
        fmt.Printf("Asset was not found.")
        return false
    }

    // write to file
    err2 := ioutil.WriteFile(dst, data, 0644)
    if err2 != nil {
        fmt.Printf("File wasn't written.")
        return false
    }
    return true
}



func check() {
    // sh verify.sh
    if (true) {
        current_step = current_step + 1
        res := Copy(fmt.Sprintf("step%d.md", current_step), "current.md")
        if (!res){
            Copy("finish.md", "current.md")
            current_step = current_step - 1
        }    
    }
}

func check_back() {
    current_step = current_step - 1
    res := Copy(fmt.Sprintf("step%d.md", current_step), "current.md")
    if (!res){
        Copy("intro.md", "current.md")
        current_step = 0
    }
}

func data(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/_data" {
        http.Error(w, "404 not found./data", http.StatusNotFound)
        return
    }

    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")

    switch r.Method {
        case "GET":
            fmt.Printf("Getting GET...\n")
            http.ServeFile(w, r, "/var/_data/current.md")
    }
}

func next(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/_next" {
        http.Error(w, "404 not found./next", http.StatusNotFound)
        return
    }

    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")

    switch r.Method {
        case "GET":
            if err := r.ParseForm(); err != nil {
                fmt.Fprintf(w, "ParseForm() err: %v", err)
                return
            }
            fmt.Printf("Getting Next ...\n")

            check()
            http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}
func back(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/_back" {
        http.Error(w, "404 not found./back", http.StatusNotFound)
        return
    }

    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")

    switch r.Method {
        case "GET":
            if err := r.ParseForm(); err != nil {
                fmt.Fprintf(w, "ParseForm() err: %v", err)
                return
            }
            fmt.Printf("Getting Back...\n")

            check_back()
            http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}


func root(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found./", http.StatusNotFound)
        return
    }

    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")

    switch r.Method {
        case "GET":
            fmt.Printf("Getting GET...\n")
            http.ServeFile(w, r, "index.html")

    }
}

func main() {
    data, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        log.Fatalln(err)
    }

    var v interface{}
    err = yaml.Unmarshal(data, &v)
    if err != nil {
        log.Fatalln(err)
    }


    Copy("index.html", "index.html")
    Copy("intro.md", "current.md")

    http.HandleFunc("/", root)
    http.HandleFunc("/_data", data)
    http.HandleFunc("/_next", next)
    http.HandleFunc("/_back", back)

    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe("0.0.0.0:8081", nil); err != nil {
        log.Fatal(err)
    }
}
