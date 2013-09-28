package main

import (
	"fmt"
	"net/http"
	"sync"
)

var keyChar = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func genKey(n int) string {
	if n == 0 {
		return string(keyChar[0])
	}
	l := len(keyChar)
	s := make([]byte, 20) // FIXME: will overflow. eventually.
	i := len(s)
	for n > 0 && i >= 0 {
		i--
		j := n % l
		n = (n - j) / l
		s[i] = keyChar[j]
	}
	return string(s[i:])
}


type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex

	// func (s *URLStore) Get(key string) string {
	// 	s.mu.Rlock()
	// 	url := s.urls[key]
	// 	s.mu.RUnlock()
	// 	return url
	// }

	// func (s *URLStore) Set(key,url string) bool{
	// 	s.mu.Lock()
	// 	_, present := s.urls[key]
	// 	if present {
	// 			s.mu.Unlock()
	// 			return false
	// 	}
	// 	s.urls[key] = url
	// 	s.mu.Unlock()
	// 	return true
	// }
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "8080 is avaliable,Master !")
}

func deafult(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,World"+genKey(5))
}

// func Add(w http.ResponseWriter, r *http.Request) {
// 	url := r.FormValue("url")
// 	key := store.Put(url)
// 	fmt.Fprintf(w, "http://localhost:8080/%s", key)
// }

func main() {
	http.HandleFunc("/", deafult)
	http.HandleFunc("/shabuwenzi", Hello)
	// http.HandleFunc("/add", Add)

	http.ListenAndServe(":8080", nil)
}
