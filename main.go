package main

import (
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Blog struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Content     string
	PublishDate time.Time
}

func setup() {
	//db, _ := gorm.Open(sqlite.Open("blogdata.db"), &gorm.Config{})

	//db.AutoMigrate(&Blog{}) Already migrated
	//db.Create(&Blog{Title: "Second title!", Content: "ABOBA :)", PublishDate: time.Now()})
}

var tpl *template.Template

func main() {
	//setup()
	Route()
	http.ListenAndServe(":8080", nil)
}

func getData() []map[string]interface{} {
	db, _ := gorm.Open(sqlite.Open("blogdata.db"), &gorm.Config{})
	var result []map[string]interface{}
	db.Table("Blogs").Find(&result)
	return result
}

func getDataById(id string) map[string]interface{} {
	db, _ := gorm.Open(sqlite.Open("blogdata.db"), &gorm.Config{})
	var result map[string]interface{}
	db.Table("Blogs").Where("id = " + id).Take(&result)
	return result
}

func Route() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexView)
	r.HandleFunc("/detail/{id}", detailView)
	http.Handle("/", r)
}

func detailView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogs := getDataById(vars["id"])
	tmpl := template.Must(template.ParseFiles("detail.html"))
	tmpl.Execute(w, blogs)
}

func indexView(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	blogs := getData()
	tmpl.Execute(w, blogs)
}
