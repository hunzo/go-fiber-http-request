package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hunzo/go-fiber-http-request/routers"
)

func main() {
	s := fiber.New()

	routers.SetupRouters(s)

	s.Listen(":8080")

}
