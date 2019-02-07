package datos

import (
	"database/sql"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> ab22ff565f46f4880635a9a49132e21d7b4cd959
	"log"

	//
	_ "github.com/lib/pq"
)

var (
	conexion *sql.DB
)

<<<<<<< HEAD
//Inicia : inicia la ejecución la conexion
func Inicia() (ErrDB error) {

	var Error error
	conexion, Error = obtenerConexion("192.168.13.15", 5432, "dev", "dev19", "alfa")
	return Error
}

func obtenerConexion(_DireccionIP string, _Puerto int, _User string, _Pass string, _BDname string) (Conexion *sql.DB, ErrOpen error) {

	CadenaConexion := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", _DireccionIP, _Puerto, _User, _Pass, _BDname)

=======
//Inicia inicializa la conexión a la BD
func Inicia() (ErrDB error) {

	var Error error
	conexion, Error = obtenerConexion("ip", 0)

	return Error
}

func obtenerConexion(_DireccionIP string, _Puerto int) (Conexion *sql.DB, ErrOpen error) {
	CadenaConexion := ""
>>>>>>> ab22ff565f46f4880635a9a49132e21d7b4cd959
	Conexion, ErrOpen = sql.Open("postgres", CadenaConexion)
	if ErrOpen != nil {
		log.Println(ErrOpen.Error())
	}

	ErrPing := Conexion.Ping()
	if ErrPing != nil {
		log.Println(ErrPing.Error())
	}

	return Conexion, ErrOpen
<<<<<<< HEAD
}

//Buscar Busca y muestra inforacion del usuario respecto al ID
func Buscar(id string) *sql.Row {

	sqlStatementSelect := `SELECT * FROM "Golang".users WHERE id=$1;`

	return conexion.QueryRow(sqlStatementSelect, id)
=======
>>>>>>> ab22ff565f46f4880635a9a49132e21d7b4cd959
}
