package main

import (
	"log"

	"github.com/LucasxTS/GoBackEnd/src/model"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
        AllowOrigins:     "*",
        AllowCredentials: false,
        AllowMethods:     "GET,POST,PUT,DELETE",
    }))

	app.Post("/verify", func(c fiber.Ctx) error {
		var scoreModel model.ScoreModel
		if err := 
	})
	log.Fatal(app.Listen(":8080"))
}
