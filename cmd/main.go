package main

import (
	"Go-Fiber-API/pkg/common/config"
	"Go-Fiber-API/pkg/common/db"
	"Go-Fiber-API/pkg/products"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	products.RegisterRoutes(app, h)

	app.Listen(c.Port)
}
