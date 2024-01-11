package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"go-fiber-service/initializers"
	"go-fiber-service/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatePokemonHandler(c *fiber.Ctx) error {
	var payload *models.CreatePokemonSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newPokemon := models.Pokemon{
		Num:           payload.Num,
		Name:          payload.Name,
		Img:           payload.Img,
		Type:          payload.Type,
		Height:        payload.Height,
		Weight:        payload.Weight,
		Candy:         payload.Candy,
		Egg:           payload.Egg,
		Multipliers:   payload.Multipliers,
		Weaknesses:    payload.Weaknesses,
		CandyCount:    payload.CandyCount,
		SpawnChance:   payload.SpawnChance,
		AvgSpawns:     payload.AvgSpawns,
		SpawnTime:     payload.SpawnTime,
		PrevEvolution: payload.PrevEvolution,
		NextEvolution: payload.NextEvolution,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	result := initializers.DB.Create(&newPokemon)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Title already exist, please use another title"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"pokemon": newPokemon}})
}

func FindPokemons(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var pokemons []models.Pokemon
	results := initializers.DB.Limit(intLimit).Offset(offset).Find(&pokemons)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(pokemons), "data": pokemons})
}

func UpdatePokemon(c *fiber.Ctx) error {
	pokemonId := c.Params("pokemonId")
	var payload *models.UpdatePokemonSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	var pokemon models.Pokemon
	result := initializers.DB.First(&pokemon, "id = ?", pokemonId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No pokemon with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	fmt.Println("UPDATEING")
	updates := mapToUpdatePokemon(payload)

	initializers.DB.Model(&pokemon).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"pokemon": pokemon}})
}

func FindPokemonById(c *fiber.Ctx) error {
	pokemonId := c.Params("pokemonId")

	var pokemon models.Pokemon
	result := initializers.DB.First(&pokemon, "id = ?", pokemonId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"pokemon": pokemon}})
}

func DeletePokemon(c *fiber.Ctx) error {
	pokemonId := c.Params("pokemonId")

	result := initializers.DB.Delete(&models.Pokemon{}, "id = ?", pokemonId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Delete success"})
}

func mapToUpdatePokemon(payload *models.UpdatePokemonSchema) models.UpdatePokemonSchema {
	update := models.UpdatePokemonSchema{
		Num:           payload.Num,
		Name:          payload.Name,
		Img:           payload.Name,
		Type:          payload.Type,
		Height:        payload.Height,
		Weight:        payload.Weight,
		Candy:         payload.Candy,
		Egg:           payload.Egg,
		Multipliers:   payload.Multipliers,
		Weaknesses:    payload.Weaknesses,
		CandyCount:    payload.CandyCount,
		SpawnChance:   payload.SpawnChance,
		AvgSpawns:     payload.AvgSpawns,
		SpawnTime:     payload.SpawnTime,
		PrevEvolution: payload.PrevEvolution,
		NextEvolution: payload.NextEvolution,
		UpdatedAt:     time.Now(),
	}

	return update
}
