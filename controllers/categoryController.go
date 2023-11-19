package controllers

import (
	"math"
	"strconv"

	"ecommerce/database"
	"ecommerce/models"

	"github.com/gofiber/fiber/v2"
)

func AllCategorys(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 10
	offset := (page - 1) * limit
	var total int64
	var categorys []models.Category

	database.Database.Db.Offset(offset).Limit(limit).Find(&categorys)
	database.Database.Db.Model(&models.Category{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": categorys,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})

}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category

	if err := c.BodyParser(&category); err != nil {
		return err
	}

	database.Database.Db.Create(&category)

	return c.JSON(category)

}

func GetCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	category := models.Category{
		Id: uint(id),
	}
	database.Database.Db.Find(&category)
	return c.JSON(category)
}

func UpdateCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	category := models.Category{
		Id: uint(id),
	}
	if err := c.BodyParser(&category); err != nil {
		return err
	}
	database.Database.Db.Model(&category).Updates(category)
	return c.JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	category := models.Category{
		Id: uint(id),
	}

	database.Database.Db.Delete(&category)
	return nil
}
