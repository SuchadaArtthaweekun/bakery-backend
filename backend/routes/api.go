package routes

import (
	"go-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/add", controllers.AddUser)
	app.Post("/api/login", controllers.Login)
	app.Delete("/api/delete/", controllers.Delete)
	app.Post("/api/update", controllers.UpdateUser)


	app.Post("/api/pro/add", controllers.AddProduct)
	app.Get("/api/pro/all", controllers.ProductAll)
	app.Delete("/api/pro/delete", controllers.ProductDelete)
	app.Put("/api/pro/edit/:id",controllers.EditProduct)
	app.Get("/api/pro/find/:id", controllers.FineProduct)

	app.Post("/api/pro/file", controllers.ImgUp)
	app.Static("api/pro/img", "./images")

}
