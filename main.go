// TODO: fix tests
// TODO: fix base62

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jxskiss/base62"
)

type URLs = struct {
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}

// dummy database for the shortened urls
var urlMap = make(map[string]string)

func Shorten(url string) string {
	val := base62.EncodeToString([]byte(url))
	return val
}

func RegisterShortenedUrl(w http.ResponseWriter, r *http.Request) {
	var u struct {
		Url string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	inputURL := u.Url
	shortned := Shorten(inputURL)

	// if urlMap[shortned] exists {
	// maybe do smth?
	// }
	urlMap[shortned] = inputURL

	w.Header().Set("Content-Type", "application/json")
	response := URLs{
		LongURL:  inputURL,
		ShortURL: "https://" + r.Host + "/" + shortned,
	}
	json.NewEncoder(w).Encode(response)
}

func GetShortenedUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var shortURL, ok = urlMap[vars["url"]]
	if !ok {
		http.Error(w, "Error: URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, shortURL, http.StatusMovedPermanently)
}

func LogAllUrls(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "console")
	fmt.Println("map:", urlMap)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/new", RegisterShortenedUrl).Methods("POST")
	r.HandleFunc("/api/v1/{url}", GetShortenedUrl).Methods("GET")
	r.HandleFunc("/api/v1/debug/listall", LogAllUrls).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
