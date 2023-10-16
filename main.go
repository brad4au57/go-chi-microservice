package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Server is running at port 3000")) })
	router.Get("/hello", basicHandler)

	// http.ListenAndServe(":3000", router)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		// Output the error
		fmt.Println("Failed to listen to server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))

}
