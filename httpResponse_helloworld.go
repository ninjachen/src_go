package main

import (
	"fmt"
	"net/http"
	"sync"
)

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex

	func (s *URLStore) Get(key string) string {
		s.mu.Rlock()
		url := s.urls[key]
		s.mu.RUnlock()
		return url
	}
	func (s *URLStore) Set(key,url string) bool{
		s.mu.Lock()
		_, present := s.urls[key]
		if present {
				s.mu.Unlock()
				return false
		}
		s.urls[key] = url
		s.mu.Unlock()
		return true
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "8080 is avaliable,Master !")
}

func deafult(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,World")
}

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	key := store.Put(url)
	fmt.Fprintf(w, "http://localhost:8080/%s", key)
}

func main() {
	http.HandleFunc("/", deafult)
	http.HandleFunc("/shabuwenzi", Hello)
	http.HandleFunc("/add", Add)

	http.ListenAndServe(":8080", nil)
}
