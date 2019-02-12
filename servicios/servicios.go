package servicios

import (
	"database/sql"
	"elevenminds/web/webservice/datos"
	usuario "elevenminds/web/webservice/entidades"
	"elevenminds/web/webservice/utilidades"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	r.HandleFunc("/user/", inserta).Methods("POST")
	r.HandleFunc("/user/delete/{id}", elimina).Methods("GET")
	r.HandleFunc("/user/actualiza/{id}", actualiza).Methods("PATCH")
	r.HandleFunc("/user/usuarios/", getuserss).Methods("GET")

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
	w.Write([]byte("actualizado correctamente"))
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

var m map[int]usuario.User

func getuserss(w http.ResponseWriter, r *http.Request) {
	datos.Inicia()
	us, err := datos.GetUsers()
	if err != nil {
		log.Println(us)
	}
	m = make(map[int]usuario.User)
	for _, v := range us {
		m[v.ID] = v
	}

	//jsonString, err := json.Marshal(m)
	json.NewEncoder(w).Encode(us)
	if err != nil {
		panic(err)
	}
	//fmt.Println(m[1])
}
func inserta(w http.ResponseWriter, r *http.Request) {
	user := usuario.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	//log.Println(user.Test)
	res := datos.Insertar(user.Edad, user.Nombre, user.Apellido, user.Email)
	//datos.Insertar(user.Edad, user.Nombre, user.Apellido, user.Email)

	//userJson, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.Write([]byte("\tNuevo usuario registrado ID: " + strconv.Itoa(res)))
}
