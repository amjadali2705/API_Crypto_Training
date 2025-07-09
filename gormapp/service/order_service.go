package service

import (
	"gormapp/dao"
	"gormapp/model"
)

func GetOrders() ([]model.Order, error) {
	return dao.GetAllOrders()
}

func GetOrder(id int) (model.Order, error) {
	return dao.GetOrderByID(id)
}

func CreateOrder(Order model.Order) error {
	return dao.CreateOrder(Order)
}

func UpdateOrder(Order model.Order) error {
	return dao.UpdateOrder(Order)
}

func DeleteOrder(id int) error {
	return dao.DeleteOrder(id)
}
