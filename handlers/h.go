package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hunzo/go-fiber-http-request/services"
)

func Home(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "home",
	})
}

func GetRequest(c *fiber.Ctx) error {
	ret := services.GetRequest()
	return c.JSON(ret)

}

func PostWithHeader(c *fiber.Ctx) error {
	ret := services.PostWithHeader()
	return c.JSON(ret)
}

func CustomRequest(c *fiber.Ctx) error {
	ret := services.CustomRequest()
	return c.JSON(ret)
}

func UrlEncode(c *fiber.Ctx) error {
	ret := services.UrlEncode()
	return c.JSON(ret)
}
