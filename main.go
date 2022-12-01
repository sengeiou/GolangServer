package main

import (
	"main/database"
	// _ "main/docslocal"

	_ "main/docs"
	"main/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func main() {

	database.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		// AllowOrigins:     "http://puplicid:5000",
		AllowOrigins: "*",
		//44.202.36.198
	}))

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// 	// AllowOrigins:     "http://puplicid:5000",
	// 	AllowOrigins: "http://ec2-54-166-232-24.compute-1.amazonaws.com:5000",
	// 	//44.202.36.198
	// }))

	// app.Static("/", "./docs.md")
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	// end points
	// get products
	routes.Setup(app)
	// run server

	// app.ListenTLS(":5000", "certificate.crt", "private.key")
	app.Listen(":5000")

}
