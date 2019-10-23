package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil" // copy with Asset
    "github.com/smallfish/simpleyaml"
    "os/exec"
    "os"
    "io" // for parsing env
)

func Copy(src, dst string) bool {    
    if os.Getenv("DEV") == "" {
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
    } else {
        in, err := os.Open(src)
        if err != nil {
            return false
        }
        defer in.Close()

        out, err := os.Create(dst)
        if err != nil {
            return false
        }
        defer out.Close()

        _, err = io.Copy(out, in)
        if err != nil {
            return false
        }
        out.Close()
        return true
    }
}

var current_step int = 0 // on intro.md; task 1 - [0]
var _ bool = Copy("course.yaml", "course.yaml")
var source, _ = ioutil.ReadFile("course.yaml")
var yaml, _ = simpleyaml.NewYaml(source)
var tasks_number, _ = yaml.Get("courses").GetArraySize()


func verify() bool{
    // sh verify.sh
    if current_step == 0 || current_step == tasks_number - 1 {
        return true
    }
    verify_path, _ := yaml.Get("courses").GetIndex(current_step).Get("verify").String()
    Copy(verify_path, "/tmp/verify.sh")
    cmd := exec.Command("bash", "/tmp/verify.sh")
    err := cmd.Run()
    cmd_rm := exec.Command("rm", "/tmp/verify.sh")
    cmd_rm.Run()

    if err == nil {
        log.Printf("You've complete task %d\n", current_step)
        return true
    } else {
        log.Printf("You haven't complete task %d\n", current_step)
        return false
    }
}


func go_step(step int){
    if step < 0 {
        current_step = 0
    } else if step >= tasks_number {
        current_step = tasks_number - 1
    } else {
        current_step = step
    }
    task_path, _ := yaml.Get("courses").GetIndex(current_step).Get("task").String()
    Copy(task_path, "current.md")
    isIndexCopied := Copy(fmt.Sprintf("tasks/%d/index.html", current_step), "index.html")
    if !isIndexCopied {
        Copy("tasks/index.html", "index.html")
    }
}

func WebHandlerData(w http.ResponseWriter, r *http.Request){
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
            http.ServeFile(w, r, "current.md")
    }
}

func WebHandlerNext(w http.ResponseWriter, r *http.Request){
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
            go_step(current_step + 1)
            http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func WebHandlerBack(w http.ResponseWriter, r *http.Request){
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
            go_step(current_step - 1)
            http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func WebHandlerCheck(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/_check" {
        http.Error(w, "404 not found /_check", http.StatusNotFound)
        return
    }

    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")

    switch r.Method {
        case "POST":
            fmt.Printf("Getting POST (check) ...\n")
            if verify() {
                fmt.Printf("all good\n")
            } else {
                http.Error(w, "501", 501)
            }
    }
}

func WebHandlerRoot(w http.ResponseWriter, r *http.Request) {
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

    Copy("tasks/index.html", "index.html")
    go_step(0)

    http.HandleFunc("/", WebHandlerRoot)
    http.HandleFunc("/_data", WebHandlerData)
    http.HandleFunc("/_next", WebHandlerNext)
    http.HandleFunc("/_back", WebHandlerBack)
    http.HandleFunc("/_check", WebHandlerCheck)

    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe("0.0.0.0:8081", nil); err != nil {
        log.Fatal(err)
    }
}