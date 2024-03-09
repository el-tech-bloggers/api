package main

import (
	"context"
	"el-tech-bloggers/api/api"
	"el-tech-bloggers/api/pg"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	dbpool, err := pg.NewPG(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	listenAddr := flag.String("listenaddr", os.Getenv("PORT"), "Server port")
	flag.Parse()
	// Registering a new handler (path,function)
	http.HandleFunc("GET /", api.RootHandler)
	// INFO: Posts handlers
	http.HandleFunc("GET /posts", api.GetPostsHandler)
	http.HandleFunc("POST /posts", api.CreatePostHandler)
	http.HandleFunc("GET /posts/{id}", api.GetPostHandler)

	fmt.Println("Server is up on running on http://localhost:5000")
	err = http.ListenAndServe(*listenAddr, nil)
	log.Fatal(err)
}
