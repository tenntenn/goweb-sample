package main

import (
    "fmt"
    "text/template"
    "flag"
    "os"
    "encoding/json"
)

var (
    pattern string
    jsonFile string
    mainTemplate string
)
func init() {
    flag.StringVar(&pattern, "p", "templates/*.html", "template pattern like template/*.html")
    flag.StringVar(&jsonFile, "f", "", "JSON file which is embeded in templates.")
    flag.StringVar(&mainTemplate, "t", "index", "main template like index")
}

const SIZE = 1000000

func main() {

    flag.Parse()

    template, err := template.ParseGlob(pattern)
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        return
    }

    file,_ := os.Open(jsonFile)
    buff := make([]byte, SIZE)
    count,_ := file.Read(buff)
    jsonData := buff[:count]
    data := make(map[string]interface{})
    if string(jsonData) != "" {
        if err := json.Unmarshal(jsonData, &data); err != nil {
            fmt.Fprintln(os.Stderr, err.Error())
            return
        }
    }

    template.ExecuteTemplate(os.Stdout, mainTemplate, data)
}
