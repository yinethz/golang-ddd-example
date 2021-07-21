package views

import (
	"strconv"

	"github.com/yinethz/golang-ddd-example/application"
	"github.com/yinethz/golang-ddd-example/domain/entity"

	"github.com/gofiber/fiber/v2"
)

const (
	pathBase = "/tmpl/empleados"
)

type EmpleadoView struct {
	ap application.EmpleadoAppInterface
}

func NewRoutingEmpleadoView(app *fiber.App, ap application.EmpleadoAppInterface) {
	ctr := &EmpleadoView{
		ap: ap,
	}
	routing(app, ctr)
}
func routing(app *fiber.App, ctr *EmpleadoView) {
	routeTmpl := app.Group(pathBase)
	routeTmpl.Get("/", func(c *fiber.Ctx) error {
		return ctr.Index(c)
	})
	routeTmpl.Get("/edit", func(c *fiber.Ctx) error {
		return ctr.Edit(c)
	})
	routeTmpl.Get("/create", func(c *fiber.Ctx) error {
		return ctr.Create(c)
	})
	routeTmpl.Post("/insert", func(c *fiber.Ctx) error {
		return ctr.Insert(c)
	})
	routeTmpl.Post("/update", func(c *fiber.Ctx) error {
		return ctr.Update(c)
	})
	routeTmpl.Get("/delete", func(c *fiber.Ctx) error {
		return ctr.Delete(c)
	})
}

func (c *EmpleadoView) Index(ctx *fiber.Ctx) error {
	data := c.ap.SearchAll()
	return ctx.Render("index", data)
}
func (c *EmpleadoView) Create(ctx *fiber.Ctx) error {
	return ctx.Render("create", nil)
}
func (c *EmpleadoView) Insert(ctx *fiber.Ctx) error {
	e := entity.Empleado{
		Nombre: ctx.FormValue("nombre"),
		Correo: ctx.FormValue("correo"),
	}
	c.ap.Save(e)
	return ctx.Redirect(pathBase, 301)
}
func (c *EmpleadoView) Edit(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Query("id"))
	data := c.ap.SearchById(id)
	return ctx.Render("edit", data)
}
func (c *EmpleadoView) Update(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.FormValue("id"))
	e := entity.Empleado{
		Nombre: ctx.FormValue("nombre"),
		Correo: ctx.FormValue("correo"),
	}
	c.ap.Update(id, e)
	return ctx.Redirect(pathBase, 301)
}
func (c *EmpleadoView) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Query("id"))
	c.ap.Remove(id)
	return ctx.Redirect(pathBase, 301)
}
