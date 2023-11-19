package controllers

import (
	"strconv"

	"ecommerce/database"
	"ecommerce/models"

	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.Database.Db.Find(&permissions)

	return c.JSON(permissions)

}

func CreatePermission(c *fiber.Ctx) error {
	var permission models.Permission

	if err := c.BodyParser(&permission); err != nil {
		return err
	}

	database.Database.Db.Create(&permission)

	return c.JSON(permission)

}

func GetPermission(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	Permission := models.Permission{
		Id: uint(id),
	}
	database.Database.Db.Find(&Permission)
	return c.JSON(Permission)
}

func UpdatePermission(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	permission := models.Permission{
		Id: uint(id),
	}
	if err := c.BodyParser(&permission); err != nil {
		return err
	}
	database.Database.Db.Model(&permission).Updates(permission)
	return c.JSON(permission)
}

func DeletePermission(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	permission := models.Permission{
		Id: uint(id),
	}

	database.Database.Db.Delete(&permission)
	return nil
}
