package main

import "net/http"

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
		case "POST":
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
}
