package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	port := 5000
	route := mux.NewRouter()

	// store local folder on browser
	route.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	//routing endpoint to access link on browser
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/add-project", addProject).Methods("GET")
	route.HandleFunc("/contact-me", contactme).Methods("GET")

	fmt.Println("Server running on port", port)
	http.ListenAndServe("localhost:5000", route)
}

// function handling to render home page
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")

	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		w.Write([]byte("Message error : " + err.Error()))
		return
	}

	temp.Execute(w, nil)
}

// function handling to render add project page
func addProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")

	temp, err := template.ParseFiles("views/addProject.html")
	if err != nil {
		w.Write([]byte("Message error : " + err.Error()))
		return
	}

	temp.Execute(w, nil)
}

// function handling to render contact-me page
func contactme(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")

	temp, err := template.ParseFiles("views/contact-me.html")
	if err != nil {
		w.Write([]byte("Message error : " + err.Error()))
		return
	}

	temp.Execute(w, nil)
}
