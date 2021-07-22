package persistence

import (
	"database/sql"

	"github.com/yinethz/golang-ddd-example/domain/repository"

	_ "github.com/go-sql-driver/mysql"
)

type Repositories struct {
	EmpleadoRepo repository.EmpleadoRepository
}

//Está se conecta a ala bbdd y esa conexión se la envía a la creación de las implementaciones de los repositorios, devuelve una estructura con los repositorios
func NewRepositories() *Repositories {
	conn := connection()
	return &Repositories{
		//instancia la implementación del repositorio de empleado pasando la conexión a la bbdd
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
