package main

import (
  // "html/template"
  "io/ioutil"
  // "os"
  "fmt"
)

const TemplateFile = "template.html"

// Struct to represent data injected into template.html
type PageData struct {
  Updated string
}

// Retrieve JSON for desired bookmarks from the filesystem
func getChromeJSON() string {
  return "TODO"
}

// Transform bookmark JSON into a PageData struct
func pageDataFromJSON(JSON string) PageData {
  return PageData{"TODO"}
}

// Write out a file generated from template.html using the provided PageData
func generatePage(pd PageData) {
  // Read in the template
  data, err := ioutil.ReadFile(TemplateFile)
  check(err)
  fmt.Print(string(data))

}

// Check an error
func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  bookmarks := getChromeJSON()
  pd := pageDataFromJSON(bookmarks)
  generatePage(pd)
}