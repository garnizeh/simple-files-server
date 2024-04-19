package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed files
var assets embed.FS

func Assets() (fs.FS, error) {
	return fs.Sub(assets, "files")
}

const port = ":8080"

func main() {
	server := &http.Server{Addr: port}

	root, err := Assets()
	if err != nil {
		log.Fatalf("failed to get embed files system: %v\n", err)
	}

	// Use the file system to serve static files
	fs := http.FileServer(http.FS(root))
	http.Handle("/", http.StripPrefix("/", fs))

	// Serve the files using the default HTTP server
	log.Printf("Listening on %s...", port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
