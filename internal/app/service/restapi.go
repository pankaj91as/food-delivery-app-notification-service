package service

import "food-delivery-app-notification-service/internal/app/repository"

type IRestService interface{}
type RestService struct {
	restRepo repository.IRepository
}

func NewRestService(restRepo repository.IRepository) IRestService {
	return &RestService{
		restRepo: restRepo,
	}
}
