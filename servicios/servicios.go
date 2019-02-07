package servicios

import (
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
func Inicia() {
	utilidades.Inicia()

	if Error := datos.Inicia(); Error != nil {
		log.Fatalln(Error.Error())
	}
	r := mux.NewRouter()
	r.HandleFunc("/", saluda).Methods("GET")
	r.HandleFunc("/user/{id}", busca).Methods("GET")
	r.HandleFunc("/user/", inserta).Methods("GET")
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
}

func inserta(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	t := usuario.User{}
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	//log.Println(t.Test)
	res := datos.Insertar(t.Edad, t.Nombre, t.Apellido, t.Email)

	log.Println("\tNuevo usuario registrado ID", res)
}
