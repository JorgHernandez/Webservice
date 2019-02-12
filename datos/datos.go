package datos

import (
	//marcos "elevenminds/web/Webservice/entidades"
	"database/sql"
	usuario "elevenminds/web/webservice/entidades"
	"fmt"
	"log"

	//
	_ "github.com/lib/pq"
)

var (
	conexion *sql.DB
)

//Inicia : inicia la ejecución la conexion
func Inicia() (ErrDB error) {

	var Error error
	conexion, Error = obtenerConexion("localhost", 5432, "postgres", "hola", "Golang")
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

	sqlStatementSelect := `SELECT * FROM "public".users WHERE id=$1;`

	return conexion.QueryRow(sqlStatementSelect, id)
}

// fun para mostrar ussers en jason

//Insertar registros a bd
func Insertar(age int, firstName string, lastName string, email string) int {
	//Prepara la insercion y devuelve el id creado
	sqlStatementInsert := ` INSERT INTO "Golang".users (age, first_name, last_name, email) VALUES ($1, $2, $3, $4) RETURNING id`
	id := 0
	//Ejecuta y comprueba si hay errores
	err := conexion.QueryRow(sqlStatementInsert, age, firstName, lastName, email).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

//Actualiza la inforacion del usuario respecto al ID
func Actualiza(id string, u1 usuario.User) error {
	sqlStatementUpdate := `UPDATE "public".users SET first_name = $2, last_name = $3, email =$4,age=$5 WHERE id = $1;`
	_, err := conexion.Exec(sqlStatementUpdate, id, u1.Nombre, u1.Apellido, u1.Email, u1.Edad)
	return err
}

//Eliminar la inforacion del usuario respecto al ID
func Eliminar(id string) error {
	sqlStatement := `DELETE FROM "public".users WHERE id = $1;`
	_, err := conexion.Exec(sqlStatement, id)

	return err
}

//GetUsers consulta de datos
func GetUsers() ([]usuario.User, error) {
	rows, err := conexion.Query(
		`SELECT * FROM public.users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	usr := []usuario.User{}
	for rows.Next() {
		var u1 usuario.User
		if err := rows.Scan(&u1.ID, &u1.Edad, &u1.Nombre, &u1.Apellido, &u1.Email); err != nil {
			return nil, err
		}
		usr = append(usr, u1)
	}
	return usr, nil
}
