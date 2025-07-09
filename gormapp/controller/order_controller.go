package controller

import (
	"gormapp/model"
	"gormapp/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetOrders(c *fiber.Ctx) error {
	orders, err := service.GetOrders()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(orders)
}

func GetOrder(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	order, err := service.GetOrder(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(order)
}

func CreateOrder(c *fiber.Ctx) error {
	var order model.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := service.CreateOrder(order); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Order created"})
}

func UpdateOrder(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var order model.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	order.ID = id
	if err := service.UpdateOrder(order); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Order updated"})
}

func DeleteOrder(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := service.DeleteOrder(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Order deleted"})
}
