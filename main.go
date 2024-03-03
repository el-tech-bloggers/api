package main

import (
	"context"
	"el-tech-bloggers/api/api"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

// Runs by default before anything else
func init() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
}

func main() {
	listenAddr := flag.String("listenaddr", os.Getenv("PORT"), "Server port")
	flag.Parse()
	// Registering a new handler (path,function)
	http.HandleFunc("GET /", api.RootHandler)
	http.ListenAndServe(*listenAddr, nil)
}
