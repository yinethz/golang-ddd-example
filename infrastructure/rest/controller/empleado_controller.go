package controller

import (
	"strconv"

	"github.com/yinethz/golang-ddd-example/application"
	"github.com/yinethz/golang-ddd-example/domain/entity"

	"github.com/gofiber/fiber/v2"
)

type EmpleadoController struct {
	ap application.EmpleadoAppInterface
}

func NewRoutingEmpleadoController(app *fiber.App, ap application.EmpleadoAppInterface) {
	ctr := &EmpleadoController{
		ap: ap,
	}
	routing(app, ctr)
}

func routing(app *fiber.App, ctr *EmpleadoController) {

	route := app.Group("/empleados")
	route.Get("/", func(c *fiber.Ctx) error {
		return ctr.SearchAll(c)
	})
	route.Get("/:id", func(c *fiber.Ctx) error {
		return ctr.SearchById(c)
	})
	route.Post("/", func(c *fiber.Ctx) error {
		return ctr.Save(c)
	})
	route.Put("/:id", func(c *fiber.Ctx) error {
		return ctr.Update(c)
	})
	route.Delete("/:id", func(c *fiber.Ctx) error {
		return ctr.Remove(c)
	})
}

func (c *EmpleadoController) Save(ctx *fiber.Ctx) error {
	e := entity.Empleado{}
	if err := ctx.BodyParser(&e); err != nil {
		return err
	}

	res := c.ap.Save(e)
	if (entity.Empleado{}) == res {
		ctx.Status(400)
		return nil
	}

	return ctx.Status(201).JSON(res)
}

func (c *EmpleadoController) SearchById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	res := c.ap.SearchById(id)
	if (entity.Empleado{}) == res {
		ctx.Status(204)
		return nil
	}
	return ctx.Status(200).JSON(res)
}

func (c *EmpleadoController) SearchAll(ctx *fiber.Ctx) error {
	res := c.ap.SearchAll()
	if len(res) <= 0 {
		ctx.Status(204)
		return nil
	}
	return ctx.Status(200).JSON(res)
}

func (c *EmpleadoController) Remove(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	c.ap.Remove(id)
	ctx.Status(200)
	return nil
}

func (c *EmpleadoController) Update(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	e := entity.Empleado{Id: id}
	if err := ctx.BodyParser(&e); err != nil {
		return err
	}

	c.ap.Update(id, e)
	ctx.Status(201)

	return nil
}
