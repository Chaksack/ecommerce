package routes

import (
	"ecommerce/controllers"
	"ecommerce/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//controllers endpoints
	app.Post("api/register", controllers.Register)
	app.Post("api/login", controllers.Login)

	app.Use(middleware.Isauthenticated)

	app.Put("api/users/info", controllers.UpdateInfo)
	app.Put("api/users/password", controllers.UpdatePassword)

	app.Get("api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)

	app.Get("api/users", controllers.AllUsers)
	app.Post("api/users", controllers.CreateUser)
	app.Get("api/users/:id", controllers.GetUser)
	app.Put("api/users/:id", controllers.UpdateUser)
	app.Delete("api/users/:id", controllers.DeleteUser)

	app.Get("api/roles", controllers.AllRoles)
	app.Post("api/roles", controllers.CreateRole)
	app.Get("api/roles/:id", controllers.GetRole)
	app.Put("api/roles/:id", controllers.UpdateRole)
	app.Delete("api/roles/:id", controllers.DeleteRole)

	app.Get("api/permissions", controllers.AllPermissions)
	app.Post("api/permissions", controllers.CreatePermission)
	app.Get("api/permissions/:id", controllers.GetPermission)
	app.Put("api/permissions/:id", controllers.UpdatePermission)
	app.Delete("api/permissions/:id", controllers.DeletePermission)

	app.Get("api/categorys", controllers.AllCategorys)
	app.Post("api/categorys", controllers.CreateCategory)
	app.Get("api/categorys/:id", controllers.GetCategory)
	app.Put("api/categorys/:id", controllers.UpdateCategory)
	app.Delete("api/categorys/:id", controllers.DeleteCategory)

	app.Get("api/products", controllers.AllProducts)
	app.Post("api/products", controllers.CreateProduct)
	app.Get("api/products/:id", controllers.GetProduct)
	app.Put("api/products/:id", controllers.UpdateProduct)
	app.Delete("api/products/:id", controllers.DeleteProduct)

	app.Get("api/cart", controllers.GetCart)
	app.Post("api/addtocart", controllers.AddToCart)

	app.Get("api/initiate", controllers.InitiatePayment)
	app.Post("api/payments", controllers.InitiatePayment)
}
