package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Jan Kowalski", Email: "jk@host.pl"},
	{ID: 2, Name: "Adam Nowak", Email: "an@host.pl"},
}

func main() {

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for _, user := range users {
			if fmt.Sprint(user.ID) == id {
				return c.JSON(user)
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("user not found")
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			fmt.Println(err.Error())
			return c.Status(fiber.StatusBadRequest).SendString("invalid request body")
		}
		user.ID = len(users) + 1
		users = append(users, user)
		return c.JSON(user)
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, user := range users {
			if fmt.Sprint(user.ID) == id {
				users = append(users[:i], users[i+1:]...)
				return c.SendString("user deleted")
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("can't find")
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var userUpdate User
		if err := c.BodyParser(&userUpdate); err != nil {
			fmt.Println(err.Error())
			return c.Status(fiber.StatusBadRequest).SendString("invalid request body")
		}

		for i, user := range users {
			if fmt.Sprint(user.ID) == id {
				userUpdate.ID = user.ID
				users[i] = userUpdate
				return c.JSON(userUpdate)
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("user not found")

	})

	app.Listen(":5555")

}
