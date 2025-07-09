package dao

import (
	"gormapp/config"
	"gormapp/model"
)

func GetAllOrders() ([]model.Order, error) {
	var orders []model.Order
	err := config.DB.Find(&orders).Error
	return orders, err
}

func GetOrderByID(id int) (model.Order, error) {
	var order model.Order
	err := config.DB.First(&order, id).Error
	return order, err
}

func CreateOrder(order model.Order) error {
	return config.DB.Create(&order).Error //insert into users (name, email) values (?, ?)
}

func UpdateOrder(order model.Order) error {
	return config.DB.Save(&order).Error //update users set name = ?, email = ? where id = ?
}

func DeleteOrder(id int) error {
	return config.DB.Delete(&model.Order{}, id).Error //delete from users where id = ?
}
