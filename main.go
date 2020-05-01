package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id   int
	Name string
	City string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "gokeep"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}

	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)

		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}

	render(w, "index.html", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Id := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id = ?", Id)
	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}
	render(w, "show.html", res)
}

func New(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pageVars := []Employee{}
		render(w, "new.html", pageVars)
	} else {
		db := dbConn()
		r.ParseForm()

		if len(r.FormValue("name"))==0 || len(r.FormValue("city"))== 0{
			fmt.Println("Fake value!")
			http.Redirect(w, r, "/", 301)
			
		}else{
			name := r.FormValue("name")
			city := r.FormValue("city")
			insForm, err := db.Prepare("INSERT INTO Employee(name, city) VALUES(?, ?)")
			if err != nil {
				panic(err.Error())
			}
			insForm.Exec(name, city)
			log.Println("INSERT: Name: " + name + " | City: " + city)
			defer db.Close()
			http.Redirect(w, r, "/", 301)
		}
		
	}

}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	editForm, err := db.Query("SELECT * FROM Employee WHERE id = ? LIMIT 1", nId)
	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	res := []Employee{}
	for editForm.Next() {
		var id int
		var name, city string
		err = editForm.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		
		
	}

	res = append(res, emp)
	render(w, "edit.html", res)

	
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	//fmt.Println("Detected ", r.Method, "request")
	if r.Method == "POST" {

		
		id := r.URL.Query().Get("id")
		name := r.FormValue("name")
		city := r.FormValue("city")
		updateDB, err := db.Prepare("UPDATE Employee SET name = ?, city = ? WHERE id = ?")

		if err != nil {
			panic(err.Error())
		}
		updateDB.Exec(name, city, id)
		defer db.Close()
		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", Index)
	http.HandleFunc("/new", New)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8000", nil)
}

func render(w http.ResponseWriter, tmpl string, pageVars []Employee) {
	tmpl = fmt.Sprintf("form/%s", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	//fmt.Println(detailss)
	err = t.Execute(w, pageVars)
	if err != nil {
		log.Println("template executing error: ", err)
	}

}