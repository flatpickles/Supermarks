package main

import (
  "html/template"
  "os"
  "time"
  "encoding/json"
  "io/ioutil"
)

const TemplateFile = "templates.html"
const PrimaryTemplateName = "main"
const OutputFile = "index.html"
const ChromeBookmarksLocation = "/Users/matt/Library/Application Support/Google/Chrome/Default/Bookmarks"

type PageData struct {
  RootNodes []BookmarkNode
  Updated time.Time
}

type BookmarkNode struct {
  Title, URL string
  Children []BookmarkNode
}

type JSON map[string]interface{}
type JSONArr []interface{}

// Retrieve JSON for desired bookmarks from the filesystem
func getChromeJSON() JSON {
  bookmarksFile, readError := ioutil.ReadFile(ChromeBookmarksLocation)
  check(readError)
  var bookmarksJSON JSON
  unmarshalError := json.Unmarshal(bookmarksFile, &bookmarksJSON)
  check(unmarshalError)
  return bookmarksJSON
}

// Transform bookmark JSON into a PageData struct
func pageDataFromJSON(data JSON) PageData {
  nodes := bookmarkNodesFromJSON(data)
  return PageData{nodes, time.Now()}
}

func bookmarkNodesFromJSON(data JSON) []BookmarkNode {
  var name string
  var URL string
  children := []BookmarkNode{}
  for key, val := range data {
    switch key {
    case "roots", "bookmark_bar":
      valJSON, isJSON := val.(map[string]interface{})
      if isJSON {
        return bookmarkNodesFromJSON(JSON(valJSON))
      }
    case "name":
      name = val.(string)
    case "url":
      URL = val.(string)
    case "children":
      valJSONArr, isJSONArr := val.([]interface{})
      if isJSONArr {
        nodeArr := JSONArr(valJSONArr)
        for i := range nodeArr {
          node := nodeArr[i]
          nodeJSON, isJSON := node.(map[string]interface{})
          if isJSON {
            children = append(children, bookmarkNodesFromJSON(JSON(nodeJSON))...)
          }
        }
      }
    }
  }
  return []BookmarkNode{BookmarkNode{name, URL, children}}
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