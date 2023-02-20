package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type article struct {
	ID      int
	Title   string
	Content string
}

var data = []article{
	article{1, "lorem", "lorem"},
	article{2, "ipsum", "ipsum"},
	article{3, "data3", "data ke 3"},
	article{4, "data4", "data ke 4"},
}

func articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	} else if r.Method == "POST" {
		var response map[string]any = map[string]any{}
		response["message"] = "data berhasil ditambahkan"
		response["status"] = true
		var result, err = json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "method tidak ada", http.StatusBadRequest)
}

func users(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/articles", articles)
	http.HandleFunc("/users", users)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
