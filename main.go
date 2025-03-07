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
	index := client.Index("campaigns")
	index.UpdateSearchableAttributes(&[]string{"id", "status", "name", "code", "brand", "country"})

	campaigns := []Campaign{
		NewCampaign("1", "active", "Summer Sale", "SUM2024", "PeYa", "AR", "GE12345", "user1"),
		NewCampaign("2", "inactive", "Winter Discount", "WIN2024", "Hunger", "EC", "GE67890", "user2"),
	}

	var documents []interface{}
	for _, campaign := range campaigns {
		documents = append(documents, campaign)
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
