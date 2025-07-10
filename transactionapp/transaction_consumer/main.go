package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/go-resty/resty/v2"
)

type Transaction struct {
    ID        uint    `json:"id"`
    FromUser  string  `json:"from_user"`
    ToUser    string  `json:"to_user"`
    Amount    float64 `json:"amount"`
    Timestamp string  `json:"timestamp"`
}

func main() {
    app := fiber.New()

    app.Post("/api/consumer/send", func(c *fiber.Ctx) error {
        tx := new(Transaction)
        if err := c.BodyParser(tx); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
        }

        client := resty.New()
        resp, err := client.R().
            SetHeader("Content-Type", "application/json").
            SetBody(tx).
            Post("http://localhost:8080/api/producer/transactions")

        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Failed to call producer"})
        }

        return c.JSON(fiber.Map{
            "message": "Transaction sent to producer",
            "result":  resp.String(),
        })
    })

    app.Listen(":8081")
}