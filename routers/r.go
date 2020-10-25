package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hunzo/go-fiber-http-request/handlers"
)

func SetupRouters(r *fiber.App) {

	r.Get("/", handlers.Home)
	r.Get("/get", handlers.GetRequest)
	r.Get("/post", handlers.PostWithHeader)
	r.Get("/custom", handlers.CustomRequest)
	r.Get("/urlencode", handlers.UrlEncode)

}
