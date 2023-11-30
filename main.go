package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/seigaalghi/e-library/pkg/zaplog"
	"github.com/seigaalghi/e-library/transport"
)

func main() {

	logger := zaplog.WithContext(context.Background())
	defer logger.Sync()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	logger.Info("starting backend server at port: " + port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), initRouter()))
}

func initRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "hello world",
		})
	})

	v1 := r.PathPrefix("/api/v1").Subrouter()

	go transport.NewHttpHandler(v1)

	return r
}

func corsHandler() http.Handler {
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	return handlers.CORS(header, methods, origins)(initRouter())
}
