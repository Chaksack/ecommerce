package controllers

import (
	"strconv"

	"ecommerce/database"
	"ecommerce/models"

	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.Database.Db.Find(&roles)

	return c.JSON(roles)

}

func CreateRole(c *fiber.Ctx) error {
	var roleDTO fiber.Map

	if err := c.BodyParser(&roleDTO); err != nil {
		return err
	}
	list := roleDTO["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:       roleDTO["name"].(string),
		Permission: permissions,
	}

	database.Database.Db.Create(&role)

	return c.JSON(role)

}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}
	database.Database.Db.Preload("Permissions").Find(&role)
	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDTO fiber.Map

	if err := c.BodyParser(&roleDTO); err != nil {
		return err
	}

	list := roleDTO["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))
	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}
	var result interface{}

	database.Database.Db.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := models.Role{
		Id:         uint(id),
		Name:       roleDTO["name"].(string),
		Permission: permissions,
	}

	database.Database.Db.Model(&role).Updates(role)
	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.Database.Db.Delete(&role)
	return nil
}
