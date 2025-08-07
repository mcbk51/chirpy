package main  

import (
	"os"
	"log"
	"net/http"
	"sync/atomic"
  _	"github.com/lib/pq"
	"github.com/mcbk51/chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	godotenv.Load()
	const filepathRoot = "."
	const port = "8080"

	dbURl := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURl)
	if err != nil {
	  log.Fatal("Failed to connect database:", err)	
	}
	defer db.Close()

	dbQueries := database.New(db)
	
	cfg := &apiConfig{
		fileserverHits: atomic.Int32{},
	}

	mux := http.NewServeMux()
	fsHandler :=  cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))

	mux.Handle("/app/", fsHandler)
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /api/validate_chirp", handlerValidate)
	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.handlerReset)
	
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	
	log.Printf("Serving files from %s on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}


