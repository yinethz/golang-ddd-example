package persistence

import (
	"database/sql"
	"log"

	"github.com/yinethz/golang-ddd-example/domain/entity"
	"github.com/yinethz/golang-ddd-example/domain/repository"
)

var _ repository.EmpleadoRepository = &EmpleadoRepositoryImpl{}

type EmpleadoRepositoryImpl struct {
	conexion *sql.DB
}

func NewEmpleadoRepositoryImpl(conexion *sql.DB) *EmpleadoRepositoryImpl {
	return &EmpleadoRepositoryImpl{
		conexion: conexion,
	}
}

func (r *EmpleadoRepositoryImpl) Save(e entity.Empleado) entity.Empleado {
	insertarRegistros, err := r.conexion.Prepare("INSERT INTO empleados(nombre,correo)VALUES(?,?)")

	if err != nil {
		panic(err.Error())
	}
	registro, err := insertarRegistros.Exec(e.Nombre, e.Correo)
	id, err := registro.LastInsertId()
	e.Id = int(id)
	log.Println("id ", id, err)
	return e
}
func (r *EmpleadoRepositoryImpl) SearchById(id int) entity.Empleado {
	registro, err := r.conexion.Query("SELECT * FROM empleados WHERE id =?", id)

	e := entity.Empleado{}
	for registro.Next() {

		err = registro.Scan(&e.Id, &e.Nombre, &e.Correo)
		if err != nil {
			panic(err.Error())
		}
	}
	return e
}
func (r *EmpleadoRepositoryImpl) SearchAll() []entity.Empleado {
	registros, err := r.conexion.Query("SELECT * FROM empleados")
	//Validar error o margen de error
	if err != nil {
		panic(err.Error())
	}
	//Declarar valiables y arreglo de variables
	arregloEmpleado := []entity.Empleado{}

	//Recorrer todos los datos
	for registros.Next() {
		var e entity.Empleado
		err = registros.Scan(&e.Id, &e.Nombre, &e.Correo)
		if err != nil {
			panic(err.Error())
		}
		//Asignar valores y unirlos a la variable
		arregloEmpleado = append(arregloEmpleado, e)
	}
	return arregloEmpleado
}
func (r *EmpleadoRepositoryImpl) Remove(id int) {

	db, err := r.conexion.Prepare("DELETE FROM empleados WHERE id=?")

	if err != nil {
		panic(err.Error())
	}
	db.Exec(id)
}
func (r *EmpleadoRepositoryImpl) Update(id int, e entity.Empleado) {
	modificarRegistros, err := r.conexion.Prepare("UPDATE empleados SET nombre=?, correo=? WHERE id=?")

	if err != nil {
		panic(err.Error())
	}
	modificarRegistros.Exec(e.Nombre, e.Correo, id)
}
