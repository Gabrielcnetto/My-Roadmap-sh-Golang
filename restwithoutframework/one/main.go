package main

import (
	"context"
	"fmt"
	"net/http"
)

type userHandler struct {
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		fmt.Println("GET")
		return
	case r.Method == http.MethodPost:
		id, ok := r.Context().Value("id").(int)
		if ok {
			fmt.Println(id)
		}
		fmt.Println("Post chamado")
		return
	default:
		fmt.Println("Default chamado")
		return
	}
}

// -------- middlewares --------
func Middleare_route(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GLOBAL:", r.Method, r.URL.Path)
		if r.Method == http.MethodGet {
			fmt.Println("Method cancelado")
			return
		}
		ctx := context.WithValue(r.Context(), "id", 10)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
func globalLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GLOBAL:", r.Method, r.URL.Path)
		if r.Method == http.MethodGet {
			fmt.Println("Method cancelado")
			return
		}
		ctx := context.WithValue(r.Context(), "id", 10)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
func main() {
	mux := http.NewServeMux()
	userHandler := &userHandler{}
	userHandlerWithMiddleware := Middleare_route(userHandler)
	mux.Handle("/users", userHandlerWithMiddleware)
	// http.ListenAndServe(":8081", globalLogger(mux))
	http.ListenAndServe(":8081", mux)
}
