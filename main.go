package main

import (
	"encoding/json"
	"fmt"
	"log"
	"maps"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type UUID = [16]byte

var urlMap = make(map[string]string)

type URLS struct {
	Url string `json:"url"`
}

func Shorten(url string) string {
	return strings.ToUpper(url)
}

func RegisterShortenedUrl(w http.ResponseWriter, r *http.Request) {
	var u URLS
	fmt.Println("map:", urlMap)

	vars := mux.Vars(r)
	var url = vars["url"]

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "urk: %s", u.Url)

	urlMap[url] = Shorten(url)
	fmt.Println("map:", u.Url)

	// maine()
}

func GetShortenedUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("map:", urlMap)
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "url: %v\n", vars["url"])

	var shortURL, ok = urlMap[vars["name"]]
	if !ok {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "shorten: %v\n", shortURL)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/new", RegisterShortenedUrl).Methods("POST")
	r.HandleFunc("/api/v1/{url}", GetShortenedUrl).Methods("GET")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// http://localhost:8000/articles/books/123

func maine() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
