package application

import (
	"github.com/yinethz/golang-ddd-example/domain/entity"
	"github.com/yinethz/golang-ddd-example/domain/repository"
)

type empleadoApp struct {
	rep repository.EmpleadoRepository
}

var _ EmpleadoAppInterface = &empleadoApp{}

type EmpleadoAppInterface interface {
	repository.EmpleadoRepository
	CreateOrUpdate(e entity.Empleado) entity.Empleado
}

func NewEmpleadoApp(rep repository.EmpleadoRepository) EmpleadoAppInterface {
	return &empleadoApp{rep}
}

func (app *empleadoApp) Save(e entity.Empleado) entity.Empleado {
	return app.rep.Save(e)
}
func (app *empleadoApp) SearchById(id int) entity.Empleado {
	return app.rep.SearchById(id)
}
func (app *empleadoApp) SearchAll() []entity.Empleado {
	return app.rep.SearchAll()
}
func (app *empleadoApp) Remove(id int) {
	app.rep.Remove(id)
}
func (app *empleadoApp) Update(id int, e entity.Empleado) {
	app.rep.Update(id, e)
}

func (app *empleadoApp) CreateOrUpdate(e entity.Empleado) entity.Empleado {
	re := app.rep.SearchById(e.Id)
	if re.Id != 0 {
		app.rep.Update(re.Id, e)
		return e
	} else {
		return app.rep.Save(e)
	}
}
