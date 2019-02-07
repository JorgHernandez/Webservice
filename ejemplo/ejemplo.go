package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.13.15"
	port     = 5432
	user     = "dev"
	password = "dev19."
	dbname   = "alfa"
)

var tempemail, tempname, templastname string
var tempid int

func main() {

	//Conexion a la base de datos
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//Defer se ejecuta al final de la ultima sentencia
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Insertando información...")

	//Prepara la insercion y devuelve el id creado
	sqlStatementInsert := ` INSERT INTO "Golang".users (age, email, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id`
	id := 0
	//Ejecuta y comprueba si hay errores
	err = db.QueryRow(sqlStatementInsert, 30, "Prueba15@a.com", "Prueba", "Abeurp").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("\tNuevo usuario registrado ID", id)

	//Consultando
	sqlStatementSelect := `SELECT id, email,first_name,last_name FROM "Golang".users WHERE id=$1;`
	row := db.QueryRow(sqlStatementSelect, id)
	switch err := row.Scan(&tempid, &tempemail, &tempname, &templastname); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println("\t", tempid, tempemail, tempname, templastname)
	default:
		panic(err)
	}

	sqlStatementUpdate := `UPDATE "Golang".users SET first_name = $2, last_name = $3 WHERE id = $1;`
	_, err = db.Exec(sqlStatementUpdate, id, "NewFirst", "NewLast")
	if err != nil {
		panic(err)
	}
	fmt.Println("\tUsuario actualizado")

	//Consultando
	sqlStatementSelect = `SELECT id, email,first_name,last_name FROM "Golang".users WHERE id=$1;`

	row = db.QueryRow(sqlStatementSelect, id)
	switch err := row.Scan(&tempid, &tempemail, &tempname, &templastname); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println("\t", tempid, tempemail, tempname, templastname)
	default:
		panic(err)
	}

	sqlStatement := `DELETE FROM "Golang".users WHERE id = $1;`
	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("\tEliminando")

}
