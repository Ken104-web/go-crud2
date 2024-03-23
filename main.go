package main

import (
  "strings"
	"html/template"
	"net/http"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
    gorm.Model
    ID  uint
    Title string
    Done  bool
}

type TodoPageData struct {
    PageTitle string
    Todo     []Todo
}


func main(){
    tmpl := template.Must(template.ParseFiles("template/index.html"))
    db, err := gorm.Open(sqlite.Open("test.db"),
    &gorm.Config{})
    if err != nil{
        panic("failed to connect to database")
    }
    db.AutoMigrate(&Todo{})
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
  if (r.Method == "POST") {
// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
    if err := r.ParseForm(); err != nil {
      fmt.Fprintf(w, "ParseForm() err: %v", err)
      return
    }
    todo := r.FormValue("todo")
    db.Create(&Todo{Title: todo, Done: false})
  }
  //Request not POST 
    var todos []Todo
    db.Find(&todos)
    data := TodoPageData{
    PageTitle: "My TODO list",
    Todo: todos,
    }
    tmpl.Execute(w, data)
  })
  http.HandleFunc("/done/", func(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/done/")
    var todo Todo
    db.First(&todo, id)
    todo.Done = true
    db.Save(&todo)
    http.Redirect(w, r, "/", http.StatusSeeOther)
  })
  http.HandleFunc("/delete/",  func(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/delete/")
    db.Delete(&Todo{}, id) 
    http.Redirect(w, r, "/", http.StatusSeeOther)
  })

  http.ListenAndServe(":8090", nil)
}

