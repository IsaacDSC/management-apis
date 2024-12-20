package main

import (
	"bff/internal/bff/restapi"
	"bff/internal/management/web/middlewares"
	"bff/pkg/cherlog"
	"context"
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

			ctx := middlewares.WithRequestLogger(request.Context(), request)
			request = request.WithContext(ctx)

			lrw := &middlewares.LoggingResponseWriter{ResponseWriter: writer}
			l := cherlog.GetLogFromCtx(ctx)
			l.Info("Request")

			if err := handler(lrw, request); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}

			l.Info("Response",
				zap.Int("http.response.status", lrw.Status),
				zap.String("http.response.body", string(lrw.Body)),
			)

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
