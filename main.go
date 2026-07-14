package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!!!")
	})

	type Person struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
	}

	app.Post("/", func(c *fiber.Ctx) error {
		p := new(Person)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		log.Println(p.Name)
		log.Println(p.Pass)
		str := p.Name + p.Pass
		return c.JSON(str)
	})

	app.Get("/user/:name", func(c *fiber.Ctx) error {
		str := "hello ==> " + c.Params("name")
		return c.JSON(str)
	})

	app.Listen(":3000")
}
