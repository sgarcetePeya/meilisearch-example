package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/meilisearch/meilisearch-go"
)

func main() {
	client := meilisearch.New("http://localhost:7700", meilisearch.WithAPIKey("my_master_key"))
	index := client.Index("movies")

	movies := []Movie{
		NewMovie(1, "Carol", []string{"Romance", "Drama"}),
		NewMovie(2, "Wonder Woman", []string{"Action", "Adventure"}),
		NewMovie(3, "Life of Pi", []string{"Adventure", "Drama"}),
		NewMovie(4, "Mad Max: Fury Road", []string{"Adventure", "Science Fiction"}),
		NewMovie(5, "Moana", []string{"Fantasy", "Action"}),
		NewMovie(6, "Philadelphia", []string{"Drama"}),
	}

	var documents []interface{}
	for _, movie := range movies {
		documents = append(documents, movie)
	}

	task, err := index.AddDocuments(documents)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Task UID:", task.TaskUID)

	CreateHandlers(client, index)

	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
