package app

import (
	"net/http"

	"github.com/brad4au67/go-chi-microservice/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })

	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}

	router.Post("/", orderHandler.CreateOrder)
	router.Get("/", orderHandler.GetOrders)
	router.Get("/{id}", orderHandler.GetOrderByID)
	router.Put("/{id}", orderHandler.UpdateOrderByID)
	router.Delete("/{id}", orderHandler.DeleteOrderByID)

}
