package main

import (
    "fmt"
    "github.com/smallfish/simpleyaml"
    "io/ioutil"
    "os"
)

func main() {
    filename := os.Args[1]
    source, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }
    yaml, err := simpleyaml.NewYaml(source)
    if err != nil {
        panic(err)
    }
    // bar, err := yaml.Get("bar").GetIndex(0).String()
    // if err != nil {
    //     panic(err)
    // }
    // fmt.Printf("Value: %#v\n", bar)

    fmt.Printf("%s\n", yaml.Get("courses").GetIndex(0).Get("task"))

}