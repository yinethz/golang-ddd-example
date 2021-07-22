package main

import (
	"github.com/yinethz/golang-ddd-example/application"
	"github.com/yinethz/golang-ddd-example/infrastructure/persistence"
	"github.com/yinethz/golang-ddd-example/infrastructure/rest/controller"
	"github.com/yinethz/golang-ddd-example/infrastructure/views"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./template/empleado", ".tmpl")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	//crea instancias de repositorios a utilizar
	repos := persistence.NewRepositories()
	//crea la instancia de la aplicación a la cual se le pasa el repositorio que va a utilizar
	empleadoService := application.NewEmpleadoApp(repos.EmpleadoRepo)
	//se crea el controlador con su enrutamiento, el cual utiliza la aplicación de empleado, así hacemos que el controlador pueda interactuar con la infraestructura
	controller.NewRoutingEmpleadoController(app, empleadoService)
	//se crea el enrutamiento que reenderiza a las vistas que utiiza la aplicación para obtener los datos que va a mostrar
	views.NewRoutingEmpleadoView(app, empleadoService)

	app.Listen(":8080")
}
