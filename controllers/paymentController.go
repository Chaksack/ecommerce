package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	// "ecommerce/models"
)

const stripePublicKey = "your_stripe_public_key"
const stripeSecretKey = "your_stripe_secret_key"

func InitiatePayment(c *fiber.Ctx) error {
	domain := "http://localhost:3000" // Update with your domain

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Product Name"),
					},
					UnitAmount: stripe.Int64(1999), // Amount in cents
				},
				Quantity: stripe.Int64(1),
			},
			// Add more line items as needed
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(fmt.Sprintf("%s/success", domain)),
		CancelURL:  stripe.String(fmt.Sprintf("%s/cancel", domain)),
	}

	session, err := session.New(params)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"id": session.ID,
	})
}
