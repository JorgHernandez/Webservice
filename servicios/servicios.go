package servicios

import (
	"elevenminds/ejemploweb/datos"
	"elevenminds/ejemploweb/servicios/usuario"
	"elevenminds/ejemploweb/utilidades"
	"fmt"
	"log"
	"net/http"
)

//Inicia incializa los servicios
func Inicia() {
	utilidades.Inicia()

	if Error := datos.Inicia(); Error != nil {
		log.Fatalln(Error.Error())
	}

	Servidor := http.NewServeMux()

	usuario.Inicia()

	if Error := http.ListenAndServe(fmt.Sprintf(":%d", 5050), Servidor); Error != nil {
		log.Println(Error.Error())
	}

}
