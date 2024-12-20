package main

import (
	"bff/internal/bff/restapi"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	defer conn.Close()

	ctx := context.Background()
	routers, err := restapi.GetRouters(ctx, conn)
	if err != nil {
		log.Fatalf("Error on get routers: %v", err)
	}

	for path, handler := range routers {
		log.Println("[*] registry routers: ", path)
		http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "application/json")
			if err := handler(writer, request); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}

			log.Println("[*] request: ", request.URL.Path)
		})
	}

	log.Println("[*] Server listen in port 8081")
	log.Println("[*] http://localhost:8081")
	log.Println("[*] Press CTRL+C to exit")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Error on start server: %v", err)
	}
}
