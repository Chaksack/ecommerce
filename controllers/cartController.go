package controllers

import (
	"ecommerce/models"

	"github.com/gofiber/fiber/v2"
)

var cart models.Cart
var products []models.Product

func GetProducts(c *fiber.Ctx) error {
	return c.JSON(products)
}

func GetCart(c *fiber.Ctx) error {
	return c.JSON(cart)
}

func AddToCart(c *fiber.Ctx) error {
	var item models.CartItem
	if err := c.BodyParser(&item); err != nil {
		return err
	}

	for i, product := range products {
		if int(product.Id) == item.ProductID {
			products[i].Quantity += item.Quantity
			cart.Items = append(cart.Items, item)
			return c.JSON(cart)
		}
	}

	return fiber.NewError(fiber.StatusNotFound, "Product not found")
}
