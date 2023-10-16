package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) CreateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an order")
}

func (o *Order) GetOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all orders")
}

func (o *Order) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an order by ID")
}

func (o *Order) UpdateOrderByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an order by ID")
}

func (o *Order) DeleteOrderByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by ID")
}
