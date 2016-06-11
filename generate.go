package main

import (
  "encoding/json"
  "flag"
  "html/template"
  "io/ioutil"
  "log"
  "os"
  "time"
)

// MARK: Constants for later use

const PrimaryTemplateName = "main"
const TemplateFile = "templates.html"

// MARK: Config defaults and global state

const DefaultChromeBookmarksFile = "/Users/matt/Library/Application Support/Google/Chrome/Default/Bookmarks"
const DefaultOutputFile = "index.html"
const DefaultRootFolderName = "Bookmarks Bar"

var GlobalConfig Config
type Config struct {
  ChromeBookmarksFile string
  OutputFile string
  RootFolderName string
}

// MARK: Structs to represent page & bookmark data

type PageData struct {
  RootNodes []BookmarkNode
  Updated time.Time
}

type BookmarkNode struct {
  Title, URL string
  Children []BookmarkNode
}

// MARK: Generic types for JSON parsing, for easy reference

type JSON map[string]interface{}
type JSONArr []interface{}

// MARK: The meat & potatoes

// Retrieve JSON for desired bookmarks from the filesystem
func getChromeJSON() JSON {
  bookmarksFile, readError := ioutil.ReadFile(GlobalConfig.ChromeBookmarksFile)
  check(readError)
  var bookmarksJSON JSON
  unmarshalError := json.Unmarshal(bookmarksFile, &bookmarksJSON)
  check(unmarshalError)
  return bookmarksJSON
}

// Transform bookmark JSON into a PageData struct
func pageDataFromJSON(data JSON) PageData {
  nodes := bookmarkNodesFromJSON(data, GlobalConfig.RootFolderName, false)
  return PageData{nodes, time.Now()}
}

// Parse JSON into recursively defined bookmark nodes. If `rootFound` is
// false, this will traverse down to the folder with name `rootFolder`
// before starting to build the BookmarkNode struct.
func bookmarkNodesFromJSON(data JSON, rootFolder string, rootFound bool) []BookmarkNode {
  rootFoundOrIsRoot := rootFound || isBookmarkWithName(data, rootFolder)
  var name string
  var URL string
  children := []BookmarkNode{}
  for key, val := range data {
    switch key {
    case "roots", "bookmark_bar":
      // Dive down through Chrome's root bookmark file node(s)
      valJSON, isJSON := val.(map[string]interface{})
      if isJSON {
        return bookmarkNodesFromJSON(JSON(valJSON), rootFolder, rootFound)
      } else {
        log.Fatal("Failed to parse root node JSON")
      }
    case "name":
      // Set the name for this node (applies for both folders & link bookmarks)
      name = val.(string)
    case "url":
      // Set the URL for this node, if exists (applies only for link bookmarks)
      URL = val.(string)
    case "children":
      // Recursively parse the array of JSON children for this folder node
      valJSONArr, isJSONArr := val.([]interface{})
      if isJSONArr {
        nodeArr := JSONArr(valJSONArr)
        for i := range nodeArr {
          node := nodeArr[i]
          nodeJSON, isJSON := node.(map[string]interface{})
          if isJSON {
            newNodes := bookmarkNodesFromJSON(JSON(nodeJSON), rootFolder, rootFoundOrIsRoot)
            children = append(children, newNodes...)
          } else {
            log.Fatal("Failed to parse child node JSON")
          }
        }
      } else {
        log.Fatal("Failed to parse child array JSON")
      }
    }
  }
  // If the root folder has been found previously, return this node. Otherwise,
  // only return its children, such that traversal to root folder will continue
  // without yet starting to build the BookmarkNode struct for the page.
  if rootFound {
    return []BookmarkNode{BookmarkNode{name, URL, children}}
  } else {
    return children
  }
}

// Identify if the name of the top level JSON node is equal to `name`
func isBookmarkWithName(data JSON, name string) bool {
  for key, val := range data {
    switch key {
    case "name":
      return val == name
    }
  }
  return false
}

// Write out a file generated from template.html using the provided PageData
func generatePage(pageContents PageData) {
  // Setup the template
  pageTemplate, templateCreationError := template.ParseFiles(TemplateFile)
  check(templateCreationError)
  // Write out the templated HTML
  file, fileError := os.Create(GlobalConfig.OutputFile)
  check(fileError)
  defer file.Close()
  templateUseErr := pageTemplate.ExecuteTemplate(file, PrimaryTemplateName, pageContents)
  check(templateUseErr)
}

// Check an error
func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

// MARK: Main

func main() {
  // Parse flags into GlobalConfig
  bookmarksFilePtr := flag.String("bookmarks", DefaultChromeBookmarksFile, "Path to Chrome's bookmarks file")
  outputFilePtr := flag.String("pagename", DefaultOutputFile, "Name for output file, e.g. 'bookmarks.html'")
  rootFolderPtr := flag.String("root", DefaultRootFolderName, "Name of the root bookmark folder to parse")
  flag.Parse()
  GlobalConfig = Config{*bookmarksFilePtr, *outputFilePtr, *rootFolderPtr}

  // Generate the page
  bookmarks := getChromeJSON()
  pd := pageDataFromJSON(bookmarks)
  generatePage(pd)
}