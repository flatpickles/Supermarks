package main

import (
  "html/template"
  "io/ioutil"
  "os"
)

const TemplateFile = "template.html"
const OutputFile = "index.html"

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
func generatePage(pageContents PageData) {
  // Read in the template
  pageHTML, readError := ioutil.ReadFile(TemplateFile)
  check(readError)
  // Setup the template
  pageTemplate, templateCreationError := template.New("bookmarks").Parse(string(pageHTML))
  check(templateCreationError)
  // Write out the templated HTML
  file, fileError := os.Create(OutputFile)
  check(fileError)
  defer file.Close()
  templateUseErr := pageTemplate.Execute(file, pageContents)
  check(templateUseErr)
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