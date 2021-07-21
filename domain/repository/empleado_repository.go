package repository

import "github.com/yinethz/golang-ddd-example/domain/entity"

type EmpleadoRepository interface {
	Save(e entity.Empleado) entity.Empleado
	SearchById(id int) entity.Empleado
	SearchAll() []entity.Empleado
	Remove(id int)
	Update(id int, e entity.Empleado)
}
