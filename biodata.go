package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type biodata struct {
	ID      int
	Name    string
	Address string
	Title   string
	Reason  string
}

func main() {
	http.HandleFunc("/", routeRootGet)
	http.HandleFunc("/data", routeSubmitPost)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)

}

func routeRootGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.New("form").ParseFiles("view.html"))
		err := tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tmpl := template.Must(template.New("result").ParseFiles("view.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		nameForm := r.FormValue("name")

		names := []string{"bagus", "kurnia", "ahmad", "jokowi", "agus"}
		biodatas := []biodata{
			{ID: 1, Name: "Bagus", Address: "Bengkulu", Title: "Farmer", Reason: "live suck"},
			{ID: 2, Name: "Kurnia", Address: "Tangerang", Title: "Teacher", Reason: "live suck"},
			{ID: 3, Name: "Ahmad", Address: "Jakarta", Title: "Student", Reason: "live suck"},
			{ID: 4, Name: "Jokowi", Address: "Solo", Title: "President", Reason: "----"},
			{ID: 5, Name: "Agus", Address: "Bekasi", Title: "CEO", Reason: "live suck"},
		}

		var whichIndex int = -1

		for i, name := range names {
			if name == strings.ToLower(nameForm) {
				whichIndex = i
			}
		}

		for i, data := range biodatas {
			if i == whichIndex {
				if err := tmpl.Execute(w, data); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}

		if whichIndex == -1 {
			http.Error(w, "", http.StatusBadRequest)
		}
	}
}
