package servicios

import (
<<<<<<< HEAD
	"database/sql"
	"elevenminds/web/webservice/datos"
	usuario "elevenminds/web/webservice/entidades"
	"elevenminds/web/webservice/utilidades"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tempemail, tempname, templastname string
var tempid int

// Inicia la ap
=======
	"elevenminds/ejemploweb/datos"
	"elevenminds/ejemploweb/servicios/usuario"
	"elevenminds/ejemploweb/utilidades"
	"fmt"
	"log"
	"net/http"
)

//Inicia incializa los servicios
>>>>>>> ab22ff565f46f4880635a9a49132e21d7b4cd959
func Inicia() {
	utilidades.Inicia()

	if Error := datos.Inicia(); Error != nil {
		log.Fatalln(Error.Error())
	}
<<<<<<< HEAD
	r := mux.NewRouter()
	r.HandleFunc("/", saluda).Methods("GET")
	r.HandleFunc("/user/{id}", busca).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func saluda(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundo!!.."))

}
func busca(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	userID := Vars["id"]
	u1 := usuario.User{}
	row := datos.Buscar(userID)

	switch err := row.Scan(&u1.ID, &u1.Edad, &u1.Nombre, &u1.Apellido, &u1.Email); err {
	case sql.ErrNoRows:
		w.Write([]byte("Sin resultados"))
	case nil:
		json.NewEncoder(w).Encode(u1)
	default:
		panic(err)
	}
=======

	Servidor := http.NewServeMux()

	usuario.Inicia()

	if Error := http.ListenAndServe(fmt.Sprintf(":%d", 5050), Servidor); Error != nil {
		log.Println(Error.Error())
	}

>>>>>>> ab22ff565f46f4880635a9a49132e21d7b4cd959
}
