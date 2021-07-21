package main

import (
	"github.com/yinethz/golang-ddd-example/application"
	"github.com/yinethz/golang-ddd-example/infraestructure/persistence"
	"github.com/yinethz/golang-ddd-example/infraestructure/rest/controller"
	"github.com/yinethz/golang-ddd-example/infraestructure/views"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./template/empleado", ".tmpl")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	repos := persistence.NewRepositories()
	empleadoService := application.NewEmpleadoApp(repos.EmpleadoRepo)

	controller.NewRoutingEmpleadoController(app, empleadoService)

	views.NewRoutingEmpleadoView(app, empleadoService)

	app.Listen(":8080")
}
