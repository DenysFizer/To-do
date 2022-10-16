package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "fmt"
	"go_rutines/database"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type ToDo struct {
	ID   uint64 `json:"ID"`
	Code string `json:"Code"`
	Type bool   `json:"Type"`
}

func home_page(w http.ResponseWriter, r *http.Request) {
	arrDo := database.Get_Values()
	tmpl, _ := template.ParseFiles("tamplate/my.html")
	tmpl.Execute(w, arrDo)
}
func delete(w http.ResponseWriter, r *http.Request) {
	var del_todo ToDo
	err := json.NewDecoder(r.Body).Decode(&del_todo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(del_todo.ID, del_todo.Code, del_todo.Type)
	database.DB.Delete(&del_todo)

}

func save(w http.ResponseWriter, r *http.Request) {
	Code := r.FormValue("Code")
	addTodo := ToDo{
		Code: Code,
		Type: false,
	}
	database.DB.Create(&addTodo)
	http.Redirect(w, r, "/", 301)
}
func add(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("tamplate/add.html")
	tmpl.ExecuteTemplate(w, "add.html", nil)
}
func edit(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("tamplate/test.html")
	tmpl.ExecuteTemplate(w, "test.html", nil)
}
func click(w http.ResponseWriter, r *http.Request) {
	var del_todo ToDo
	err := json.NewDecoder(r.Body).Decode(&del_todo)
	if err != nil {
		log.Fatal(err)
	}
	if del_todo.Type == true {
		database.DB.Model(&del_todo).Update("Type", false)
	} else {
		database.DB.Model(&del_todo).Update("Type", true)
	}
	http.Redirect(w, r, "/", 301)
}
func editdb(w http.ResponseWriter, r *http.Request) {
	type newval struct {
		NewCode string `json:"NewCode"`
	}
	raw, err := ioutil.ReadAll(r.Body)
	var edit_todo ToDo
	err = json.NewDecoder(bytes.NewReader(raw)).Decode(&edit_todo)
	var new newval
	err = json.NewDecoder(bytes.NewReader(raw)).Decode(&new)
	if err != nil {
		log.Fatal(err)
	}
	database.DB.Model(&edit_todo).Update("Code", new.NewCode)
	http.Redirect(w, r, "/", 301)
}
func handleRequest() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", home_page)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/add", add)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/save", save)
	http.HandleFunc("/click", click)
	http.HandleFunc("/editdb", editdb)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
func main() {
	database.InitsilizeDb()
	handleRequest()
}
