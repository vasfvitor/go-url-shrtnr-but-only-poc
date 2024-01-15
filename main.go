package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// base62 code source: https://ucarion.com/go-base62
type UUID = [16]byte

func toBase62(uuid UUID) string {
	var i big.Int
	i.SetBytes(uuid[:])
	return i.Text(62)
}

var urlMap = make(map[string]string)

func Shorten(url string) string {
	var url16 UUID
	copy(url16[:], url)
	val := toBase62(url16)
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
	response := struct {
		LongURL  string `json:"long_url"`
		ShortURL string `json:"short_url"`
	}{
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
