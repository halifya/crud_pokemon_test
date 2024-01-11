package main

import (
	"log"

	"go-fiber-service/controllers"
	"go-fiber-service/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectMariaDB(&config)
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	v1 := micro.Group("/v1")

	v1.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and GORM",
		})
	})

	notes := v1.Group("/notes")
	notes.Post("/", controllers.CreateNoteHandler)
	notes.Get("/", controllers.FindNotes)
	notes.Delete("/:noteId", controllers.DeleteNote)
	notes.Get("/:noteId", controllers.FindNoteById)
	notes.Patch("/:noteId", controllers.UpdateNote)

	pokemons := v1.Group("/pokemons")
	pokemons.Post("/", controllers.CreatePokemonHandler)
	pokemons.Get("/", controllers.FindPokemons)
	pokemons.Delete("/:pokemonId", controllers.DeletePokemon)
	pokemons.Get("/:pokemonId", controllers.FindPokemonById)
	pokemons.Put("/:pokemonId", controllers.UpdatePokemon)
	log.Fatal(app.Listen(":3000"))
}
