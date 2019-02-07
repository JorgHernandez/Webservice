package datos

import (
	"database/sql"
	"fmt"
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

//Buscar Busca y muestra inforacion del usuario respecto al ID
func Buscar(id string) *sql.Row {

	sqlStatementSelect := `SELECT * FROM "Golang".users WHERE id=$1;`

	return conexion.QueryRow(sqlStatementSelect, id)
=======
//primer comentario
//segundo comentario
func main() {
	fmt.Println("Hola")
>>>>>>> da519d1c8bffd90a21953b59d09934b282f660d9
}
