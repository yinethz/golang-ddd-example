package persistence

import (
	"database/sql"

	"github.com/yinethz/golang-ddd-example/domain/repository"

	_ "github.com/go-sql-driver/mysql"
)

type Repositories struct {
	EmpleadoRepo repository.EmpleadoRepository
}

func NewRepositories() *Repositories {
	conn := connection()
	return &Repositories{
		EmpleadoRepo: NewEmpleadoRepositoryImpl(conn),
	}
}

func connection() *sql.DB {
	Driver := "mysql"
	Usuario := "root"
	Contrasena := ""
	Nombre := "sistemago"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasena+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}
