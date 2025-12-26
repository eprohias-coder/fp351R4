package main

import (
        "fmt"
        "log"
        "net/http"
        "time"
)

func main() {
        // Servir archivos estáticos bajo /static
        fs := http.FileServer(http.Dir("./static"))
        http.Handle("/static/", http.StripPrefix("/static/", fs))

        // Ruta raíz que muestra la imagen
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                // LOG: registrar petición con timestamp
                log.Printf("[%s] %s %s %s\n", time.Now().Format(time.RFC3339), r.RemoteAddr, r.Method, r.URL.Path)

                w.Header().Set("Content-Type", "text/html; charset=utf-8")
                fmt.Fprintln(w, `
                        <!DOCTYPE html>
                        <html>
                        <head><title>Aplicación Go</title></head>
                        <body>
                                <h1>Aplicación Go desplegada</h1>
                                <img src="/static/uocimatge.jpeg" alt="Imagen">
                        </body>
                        </html>
                `)
        })

        fmt.Println("Servidor seguro escuchando en el puerto 8080...")

        // CONFIGURACIÓN SEGURA
        server := &http.Server{
                Addr:         ":8080",
                Handler:      nil,
                ReadTimeout:  10 * time.Second,
                WriteTimeout: 10 * time.Second,
        }

        // LOG: registrar errores graves
        err := server.ListenAndServe()
        if err != nil {
                log.Fatalf("[%s] ERROR: %v\n", time.Now().Format(time.RFC3339), err)
        }
}
