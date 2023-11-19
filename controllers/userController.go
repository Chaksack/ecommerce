package controllers

import (
	"math"
	"strconv"

	"ecommerce/database"
	"ecommerce/models"

	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 10
	offset := (page - 1) * limit
	var total int64
	var users []models.User

	database.Database.Db.Preload("Role").Offset(offset).Limit(limit).Find(&users)
	database.Database.Db.Model(&models.User{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})

}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.Database.Db.Create(&user)

	return c.JSON(user)

}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}
	database.Database.Db.Preload("Role").Find(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	database.Database.Db.Model(&user).Updates(user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.Database.Db.Delete(&user)
	return nil
}
