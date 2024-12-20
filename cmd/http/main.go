package main

import (
	"bff/internal/management/infra/containers"
	"bff/internal/management/web"
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

	rc := containers.NewRepositoriesContainer(conn)
	gw := containers.NewGatewaysContainer()
	svc := containers.NewServicesContainer(gw, rc)

	hd := web.NewHandler(svc)
	for route, handler := range hd.GetRoutes() {
		http.HandleFunc(route, func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "application/json")
			if err := handler(writer, request); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}

			log.Println("[*] Request: ", request.Method, request.URL.Path)
		})
	}

	log.Println("[*] Server listen in port 3333")
	log.Println("[*] http://localhost:3333")
	log.Println("[*] Press CTRL+C to exit")

	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal(err)
	}

}
