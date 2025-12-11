package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Servir archivos estáticos (imagen) desde la carpeta "static"
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Servidor escuchando en el puerto 8080...")

	// Iniciar el servidor
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
