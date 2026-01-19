package main

import "github.com/dev-danilocordeiro/pizza-shop-order/internal/models"

type Handler struct {
	orders              *models.OrderModel
	users               *models.UserModel
	notificationManager *NotificationManager
}

func NewHandler(dbModel *models.DBModel) *Handler {
	return &Handler{
		orders:              &dbModel.Order,
		users:               &dbModel.User,
		notificationManager: NewNotificationManager(),
	}
}
