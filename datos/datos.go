package datos

import (
	"database/sql"
	"log"

	//
	_ "github.com/lib/pq"
)

var (
	conexion *sql.DB
)

//Inicia inicializa la conexión a la BD
func Inicia() (ErrDB error) {

	var Error error
	conexion, Error = obtenerConexion("ip", 0)

	return Error
}

func obtenerConexion(_DireccionIP string, _Puerto int) (Conexion *sql.DB, ErrOpen error) {

	CadenaConexion := ""

	Conexion, ErrOpen = sql.Open("postgres", CadenaConexion)
	if ErrOpen != nil {
		log.Println(ErrOpen.Error())
	}

	ErrPing := Conexion.Ping()
	if ErrPing != nil {
		log.Println(ErrPing.Error())
	}

	return Conexion, ErrOpen
}
