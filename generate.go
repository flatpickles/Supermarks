package main

import (
  "html/template"
  "os"
  "time"
)

const TemplateFile = "templates.html"
const PrimaryTemplateName = "main"
const OutputFile = "index.html"

// Struct to represent data injected into template.html
type PageData struct {
  RootNodes []BookmarkNode
  Updated time.Time
}

type BookmarkNode struct {
  Title string
  URL string
  Updated time.Time
  Children []BookmarkNode
}

// Retrieve JSON for desired bookmarks from the filesystem
func getChromeJSON() string {
  return "TODO"
}

// Transform bookmark JSON into a PageData struct
func pageDataFromJSON(JSON string) PageData {
  root := BookmarkNode{"Test", "http://www.man1.biz", time.Now(), []BookmarkNode{}}
  return PageData{[]BookmarkNode{root}, time.Now()}
}

// Write out a file generated from template.html using the provided PageData
func generatePage(pageContents PageData) {
  // Setup the template
  pageTemplate, templateCreationError := template.ParseFiles(TemplateFile)
  check(templateCreationError)
  // Write out the templated HTML
  file, fileError := os.Create(OutputFile)
  check(fileError)
  defer file.Close()
  templateUseErr := pageTemplate.ExecuteTemplate(file, PrimaryTemplateName, pageContents)
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