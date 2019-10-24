package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil" // copy with Asset
    "os/exec"
    "os"
    "bytes"
    "io" // for parsing env
    "encoding/json"
)

func Copy(src, dst string) bool {    
    if os.Getenv("DEV") == "" {
        // read data from Asset
        data, err := Asset(src)
        if err != nil {
            fmt.Printf("Asset was not found.\n")
            return false
        }

        // write to file
        err2 := ioutil.WriteFile(dst, data, 0644)
        if err2 != nil {
            fmt.Printf("File wasn't written.\n")
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

var current_step int = 0
var count_steps int = 9

func sendToELK() bool {
    log.Printf("env: %s\n", os.Getenv("ANALYTICS"))
    log.Printf("env: %s\n", os.Getenv("TRAINING"))
    log.Printf("env: %s\n", os.Getenv("STUDENT"))
    log.Printf("env: %s\n", os.Getenv("LAB"))

    url := fmt.Sprintf("http://%s:9880/%s", os.Getenv("ANALYTICS"), os.Getenv("TRAINING"))

    type Student struct {
        Student    string `json:"student"`
        Lab        string `json:"lab"`
        Task       int `json:"task"`
        Status     bool `json:"status"` 
    }

    // .....

    body := &Student{
        Student: os.Getenv("STUDENT"), 
        Lab: os.Getenv("LAB"),
        Task: current_step,
        Status: true,
    }

    buf := new(bytes.Buffer)
    json.NewEncoder(buf).Encode(body)
    req, _ := http.NewRequest("POST", url, buf)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    res, e := client.Do(req)
    if e != nil {
        log.Printf("Failure %v\n", e)
    }

    defer res.Body.Close()

    fmt.Println("response Status:", res.Status)
    // Print the body to the stdout
    io.Copy(os.Stdout, res.Body)
    return true
}

func verify() bool{
    if current_step == 0 || current_step == count_steps - 1 {
        return true
    }

    // try copy verify (if it doesn't exist - return true)
    isVerifyCopied := Copy(fmt.Sprintf("tasks/%d/verify.sh", current_step), "/tmp/verify.sh")
    if !isVerifyCopied {
        return true
    }

    cmd := exec.Command("bash", "/tmp/verify.sh")
    err := cmd.Run()

    exec.Command("rm", "-f", "/tmp/verify.sh").Run()

    if err == nil {
        log.Printf("You've complete task %d\n", current_step)
        sendToELK()
        return true
    } else {
        log.Printf("You haven't complete task %d\n", current_step)
        return false
    }
}


func go_step(step int){
    if step < 0 {
        current_step = 0
    } else if step >= count_steps {
        current_step = count_steps - 1
    } else {
        current_step = step
    }

    // copy task.md (if it doesn't exist - remove current.md (with old data))
    isTaskCopied := Copy(fmt.Sprintf("tasks/%d/task.md", current_step), "current.md")
    if !isTaskCopied {
        exec.Command("rm", "current.md").Run()
    }

    // tasks/##/index.html (if it doesn't exist - copy default tasks/index.html)
    isIndexCopied := Copy(fmt.Sprintf("tasks/%d/index.html", current_step), "index.html")
    if !isIndexCopied {
        Copy("tasks/index.html", "index.html")
    }

    isVerifyCopied := Copy(fmt.Sprintf("tasks/%d/courseData.sh", current_step), "/tmp/courseData.sh")
    if isVerifyCopied {
        cmd := exec.Command("/bin/bash", "/tmp/courseData.sh")
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        err := cmd.Run()

        log.Printf("Stdout: %v\n", cmd.Stdout)
        if err != nil {
            log.Printf("%v\n", err)
        }

        exec.Command("rm", "-f", "/tmp/courseData.sh").Run()
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
            http.Redirect(w, r, "/content", http.StatusSeeOther)
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
            http.Redirect(w, r, "/content", http.StatusSeeOther)
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

func WebHandlerContent(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/content" {
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
            http.ServeFile(w, r, "main.html")

    }
}

func main() {
    go_step(0)

    http.HandleFunc("/", WebHandlerRoot)
    http.HandleFunc("/content", WebHandlerContent)
    http.HandleFunc("/_data", WebHandlerData)
    http.HandleFunc("/_next", WebHandlerNext)
    http.HandleFunc("/_back", WebHandlerBack)
    http.HandleFunc("/_check", WebHandlerCheck)

    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe("0.0.0.0:8081", nil); err != nil {
        log.Fatal(err)
    }
}