package main

import (
	// "fmt"
	// "strings"
	// "html/template"
	// "net/http"
	"html/template"

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
    tmpl :=
    template.Must(template.ParseFiles("template/index.html"))
    db, err := gorm.Open(sqlite.Open("test.db"),
&gorm.Config{})
    if err != nil{
        panic("failed to connect to database")
    }
    db.AutoMigrate(&Todo{})
}
