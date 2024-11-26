package controller

import (
	"encoding/json"
	"net/http"
)

func (c *RestController) GetCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := c.restService.GetCustomers(ctx)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
