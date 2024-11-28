package controller

import (
	"encoding/json"
	"food-delivery-app-notification-service/internal/app/config"
	"food-delivery-app-notification-service/internal/app/handler"
	"food-delivery-app-notification-service/pkg/model"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (c *RestController) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := c.restService.GetOrders(ctx)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (c *RestController) GetOrdersByID(w http.ResponseWriter, r *http.Request) {
	pathVar := mux.Vars(r)
	orderID := pathVar["orderid"]
	ctx := r.Context()
	data := c.restService.GetOrdersByID(ctx, orderID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
	if int(RowsAffected) > 0 {
		PriorityQue := config.Environment.CONF.PriorityQue
		PriorityQueSlice := strings.Split(*PriorityQue, ",")

		if slices.Contains(PriorityQueSlice, order.OrderStatus) {
			handler.Publish(*config.Environment.CONF.PriorityQueueName, "Priority Message!")
		} else {
			handler.Publish(*config.Environment.CONF.PramotionalQueueName, "Pramotional Message!")
		}

		// return response
		data := model.Response{
			Status:  http.StatusAccepted,
			Message: strconv.Itoa(int(RowsAffected)),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(data)
		return
	}

	// return response
	data := model.Response{
		Status:  http.StatusNotModified,
		Message: "No Change",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(data)
}
