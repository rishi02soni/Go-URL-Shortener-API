package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type URL struct {
	Original string `json:"original"`
	Short    string `json:"short"`
}

var urlStore = map[string]string{}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	length := 6
	short := make([]byte, length)

	for i := range short {
		short[i] = charset[rand.Intn(len(charset))]
	}
	return string(short)
}

// Create short URL
func createShortURL(w http.ResponseWriter, r *http.Request) {
	var req URL
	_ = json.NewDecoder(r.Body).Decode(&req)

	short := generateShortURL()
	urlStore[short] = req.Original

	res := URL{
		Original: req.Original,
		Short:    "http://localhost:8000/" + short,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Redirect to original URL
func redirectURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	short := params["short"]

	if original, ok := urlStore[short]; ok {
		http.Redirect(w, r, original, http.StatusFound)
		return
	}

	http.Error(w, "URL not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/shorten", createShortURL).Methods("POST")
	r.HandleFunc("/{short}", redirectURL).Methods("GET")

	log.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
