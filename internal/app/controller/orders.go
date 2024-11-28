package controller

import (
	"encoding/json"
	"fmt"
	"food-delivery-app-notification-service/internal/app/config"
	"food-delivery-app-notification-service/internal/app/handler"
	"food-delivery-app-notification-service/pkg/model"
	"io"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/exp/rand"
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
	PriorityQue := config.Environment.CONF.PriorityQue
	PriorityQueSlice := strings.Split(*PriorityQue, ",")
	fmt.Println(PriorityQueSlice, order.OrderStatus, slices.Contains(PriorityQueSlice, order.OrderStatus))
	if slices.Contains(PriorityQueSlice, order.OrderStatus) {
		fmt.Printf("%s queue triggered!\n", *config.Environment.CONF.PramotionalQueueName)

		// Notification Template
		notificationType := []string{"sms", "email", "push"}
		notificationTemplate, err := config.GetNotificationTemplate(*config.Environment.CONF.PriorityQueueName, notificationType[rand.Intn(len(notificationType))])
		if err != nil {
			log.Panicf("Unable to get template from %s/%s", *config.Environment.CONF.PriorityQueueName, notificationType[rand.Intn(len(notificationType))])
		}

		// Publish payload into message queue
		handler.Publish(*config.Environment.CONF.PriorityQueueName, notificationTemplate)

		RowsAffected, _ := c.restService.UpdateOrderByID(ctx, orderID, order.OrderStatus)
		if int(RowsAffected) > 0 {
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
	} else {
		fmt.Printf("%s queue triggered!\n", *config.Environment.CONF.PramotionalQueueName)

		// Notification Template
		notificationType := []string{"sms", "email", "push"}
		notificationTemplate, err := config.GetNotificationTemplate(*config.Environment.CONF.PramotionalQueueName, notificationType[rand.Intn(len(notificationType))])
		if err != nil {
			log.Panicf("Unable to get template from %s/%s", *config.Environment.CONF.PramotionalQueueName, notificationType[rand.Intn(len(notificationType))])
		}

		// Publish payload into message queue
		handler.Publish(*config.Environment.CONF.PramotionalQueueName, notificationTemplate)
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
