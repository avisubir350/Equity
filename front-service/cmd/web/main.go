package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func loginPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/login.html")
	if err != nil {
		log.Println("Error parsing Login page:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Panic("Error while loading Login page")
		return
	}
}

func validateLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	pw := r.FormValue("password")

	fmt.Println("username:", username)
	fmt.Println("password:", pw)
}

func main() {
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/validateLogin", validateLogin)

	log.Println("Listening...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Panic(err)
	}
}
