package servicios

import (
	usuario "elevenminds/web/webservice/entidades"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Inicia la ap
func Inicia() {
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
	u1 := usuario.User{
		ID:       0,
		Edad:     30,
		Nombre:   "Juan",
		Apellido: "Perez",
		Email:    "juan.Perez@gmail.com",
	}
	w.Write([]byte(userID))
	json.NewEncoder(w).Encode(u1)
}
