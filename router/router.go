package router

import "github.com/gofiber/fiber"

func SetupRouter(app *fiber.App)  {
	router:=app.Group("/api/v1")

}