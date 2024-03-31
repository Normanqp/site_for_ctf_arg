package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("login.html"))

func handleRequests(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, r)
}

func CheckRequests(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "human?" && password == "Fade" {
		http.ServeFile(w, r, "main.html") //сразу
	} else {
		fmt.Fprintf(w, "wrong username or password")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	img := http.FileServer(http.Dir("img"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.Handle("/img/", http.StripPrefix("/img/", img))

	mux.HandleFunc("/", handleRequests)
	mux.HandleFunc("/index", CheckRequests)
	http.ListenAndServe(":"+port, mux)
}
