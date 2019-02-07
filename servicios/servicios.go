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
	r.HandleFunc("/user/delete/{id}", elimina).Methods("GET")
	r.HandleFunc("/user/actualiza/{id}", actualiza).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func actualiza(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	userID := Vars["id"]
	u1 := jsonToUser(r)
	err := datos.Actualiza(userID, u1)

	if err != nil {
		panic(err)
	}
}

func jsonToUser(r *http.Request) usuario.User {
	u1 := usuario.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u1)
	if err != nil {
		log.Fatal(err)
	}
	return u1
}

func elimina(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	userID := Vars["id"]
	status := datos.Eliminar(userID)
	if status != nil {
		w.Write([]byte("Fallo al eliminar!!.."))
		panic(status)
	}
	w.Write([]byte("Eliminado correctamente"))

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
