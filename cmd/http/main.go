package main

import (
	"bff/internal/management/infra/containers"
	"bff/internal/management/web"
	"bff/internal/management/web/middlewares"
	"bff/pkg/cherlog"
	"database/sql"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
)

func main() {
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error on connection database: %v", err)
	}

	defer conn.Close()

	rc := containers.NewRepositoriesContainer(conn)
	gw := containers.NewGatewaysContainer()
	svc := containers.NewServicesContainer(gw, rc)

	hd := web.NewHandler(svc)
	for route, handler := range hd.GetRoutes() {
		http.HandleFunc(route, func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "application/json")

			ctx := middlewares.WithRequestLogger(request.Context(), request)
			request = request.WithContext(ctx)

			lrw := &middlewares.LoggingResponseWriter{ResponseWriter: writer}
			l := cherlog.GetLogFromCtx(ctx)
			l.Info("Request")

			if err := handler(lrw, request); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}

			l.Info("Response",
				zap.Int("status", lrw.Status),
				zap.String("body", string(lrw.Body)),
			)
		})
	}

	log.Println("[*] Server listen in port 3333")
	log.Println("[*] http://localhost:3333")
	log.Println("[*] Press CTRL+C to exit")

	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal(err)
	}

}
