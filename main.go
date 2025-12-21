package main

import (
	"fmt"
	"log"
	"net/http"
	"time" // Importamos la librería de tiempo
)

func main() {
	// Servir archivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Servidor seguro escuchando en el puerto 8080...")

	// CONFIGURACIÓN SEGURA
	// En lugar de usar ListenAndServe directamente, configuramos un servidor con límites
	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  10 * time.Second, // Máximo 10 seg para leer
		WriteTimeout: 10 * time.Second, // Máximo 10 seg para responder
		// Esto evita ataques de denegación de servicio (Slowloris)
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
