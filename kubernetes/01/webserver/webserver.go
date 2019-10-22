package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil" // copy with Asset
    "github.com/smallfish/simpleyaml"
    // "os/exec"
)

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

var current_step int = 0


var res bool = Copy("course.yaml", "course.yaml")

var source, _ = ioutil.ReadFile("course.yaml")
var yaml, _ = simpleyaml.NewYaml(source)
var tasks_number, _ = yaml.Get("courses").GetArraySize()


func check_next() {    
    // sh verify.sh
    // verify_path, _ := yaml.Get("courses").GetIndex(current_step).Get("verify").String()
    // Copy(verify_path, "/tmp/verify.sh")
    // out, err := exec.Command("date").Output()
    // if err != nil {
    //     log.Fatal(err)
    // }
    // fmt.Printf("The date is %s\n", out)


    if (true) {
        current_step = current_step + 1

        if(current_step < tasks_number){
            task_path, _ := yaml.Get("courses").GetIndex(current_step).Get("task").String()
            Copy(task_path, "current.md")
        } else {
            Copy("finish.md", "current.md")
            current_step = current_step - 1
        }  
    }
}

func check_back() {
    current_step = current_step - 1
    if(current_step >= 0){
        task_path, _ := yaml.Get("courses").GetIndex(current_step).Get("task").String()
        Copy(task_path, "current.md")
    } else {
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
            http.ServeFile(w, r, "current.md")
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

            check_next()
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

    // yaml as global variable




    // courses_number, err := yaml.Get("courses").GetArraySize()
    // courseData_path, err := yaml.Get("courses").GetIndex(0).Get("courseData").String()
    task_path, _ := yaml.Get("courses").GetIndex(0).Get("task").String()

    fmt.Printf("%s\n", task_path)

    fmt.Printf("%d\n", tasks_number)



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
