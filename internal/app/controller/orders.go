package controller

import (
	"encoding/json"
	"food-delivery-app-notification-service/pkg/model"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (c *RestController) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := c.restService.GetOrders(ctx)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func (c *RestController) GetOrdersByID(w http.ResponseWriter, r *http.Request) {
	pathVar := mux.Vars(r)
	orderID := pathVar["orderid"]
	ctx := r.Context()
	data := c.restService.GetOrdersByID(ctx, orderID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func (c *RestController) UpdateOrderByID(w http.ResponseWriter, r *http.Request) {
	pathVar := mux.Vars(r)
	orderID := pathVar["orderid"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var order model.Orders
	err = json.Unmarshal(body, &order)
	if err != nil {
		panic(err)
	}

	ctx := r.Context()
	RowsAffected, _ := c.restService.UpdateOrderByID(ctx, orderID, order.OrderStatus)
	data := model.Response{
		Status:  http.StatusAccepted,
		Message: strconv.Itoa(int(RowsAffected)),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(data)
}
