package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))

	// api v1
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", c.HelloTest)
	v1.Post("/", c.InputName)
	v1.Get("/user/:name", c.ParamName)
	v1.Post("/inet", c.QueryName)
	v1.Post("/valid", c.ValidateTest)
	v1.Get("/fact/:number", c.Factorial)
	v1.Post("/register", c.Register)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJson)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)

	v3 := api.Group("/v3")
	v3.Get("/bas", c.Ascii)
}
