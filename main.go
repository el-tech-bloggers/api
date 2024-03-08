package main

import (
	"context"
	"el-tech-bloggers/api/api"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbpool *pgxpool.Pool

// Runs by default before anything else
func init() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()
}

func main() {
	listenAddr := flag.String("listenaddr", os.Getenv("PORT"), "Server port")
	flag.Parse()
	// Registering a new handler (path,function)
	http.HandleFunc("GET /", api.RootHandler)
	fmt.Println("Server is up on running on http://localhost:5000")
	err := http.ListenAndServe(*listenAddr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
