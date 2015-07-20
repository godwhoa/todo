package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

//Delete tasks
func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	db, err := sql.Open("sqlite3", "task.db")
	perror(err)
	defer db.Close()

	stmt, err := db.Prepare("delete from todo where Id=?")
	perror(err)
	_, err = stmt.Exec(id)
	perror(err)

	//Log
	ip := strings.Split(r.RemoteAddr, ":")[0]
	fmt.Printf("DELETED ID: %s IP: %s\n", id, ip)
}

//Add tasks
func Task(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "task.db")
	perror(err)
	defer db.Close()

	//Escape html tags before storing
	id := template.HTMLEscapeString(r.FormValue("id"))
	item := template.HTMLEscapeString(r.FormValue("item"))
	fmt.Println(r.Form)

	//Store in db
	stmt, err := db.Prepare("INSERT INTO todo (Id,Item) VALUES(?,?)")
	perror(err)
	_, err = stmt.Exec(id, item)
	perror(err)

	//Log
	ip := strings.Split(r.RemoteAddr, ":")[0]
	fmt.Printf("ADDED ID: %s ITEM: %s IP: %s\n", id, item, ip)
}

//Query tasks
func List(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "task.db")
	perror(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM todo")
	perror(err)
	//Query and append to a map
	m := make(map[string]interface{})
	for rows.Next() {
		var Id string
		var Item string
		err := rows.Scan(&Id, &Item)
		perror(err)
		m[Item] = Id
	}
	//Map to json
	mjson, _ := json.Marshal(m)
	fmt.Fprintf(w, string(mjson))
}

func main() {
	http.HandleFunc("/task", Task)
	http.HandleFunc("/list", List)
	http.HandleFunc("/delete", Delete)
	Home := http.FileServer(http.Dir("./public"))
	http.Handle("/", Home)

	fmt.Println("Running on :8080")
	http.ListenAndServe(":8080", nil)

}

//Error handling
func perror(err error) {
	if err != nil {
		panic(err)
	}
}
