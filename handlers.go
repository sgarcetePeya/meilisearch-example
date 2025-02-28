package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/meilisearch/meilisearch-go"
)

func CreateHandlers(client meilisearch.ServiceManager, index meilisearch.IndexManager) {
	http.HandleFunc("/search/movies", handleSearch(client))
	http.HandleFunc("/movies/all", handleGetMovies(client))
	http.HandleFunc("/", handleAddDocument(index))
	http.HandleFunc("/movies/delete/{id}", handleDeleteDocument(index))
	http.HandleFunc("/movies/{id}", handleGetDocument(index))
	http.HandleFunc("/movies/update/{id}", handleUpdateDocument(index))
}

func handleSearch(client meilisearch.ServiceManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("title")
		if query == "" {
			http.Error(w, "Missing query parameter 'title'", http.StatusBadRequest)
			return
		}

		searchRes, err := client.Index("movies").Search(query, &meilisearch.SearchRequest{Limit: 10})
		if err != nil {
			http.Error(w, fmt.Sprintf("Search error: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(searchRes.Hits)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		}
	}
}

func handleGetMovies(client meilisearch.ServiceManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		searchRes, err := client.Index("movies").Search("", &meilisearch.SearchRequest{})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving documents: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(searchRes.Hits)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		}
	}
}

func handleAddDocument(index meilisearch.IndexManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid HTTP method. Only POST is allowed.", http.StatusMethodNotAllowed)
			return
		}

		var movie Movie
		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error decoding document: %v", err), http.StatusBadRequest)
			return
		}

		movies := []interface{}{movie}
		task, err := index.AddDocuments(movies)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error adding document: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"task_uid": task.TaskUID,
			"document": movie,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func handleDeleteDocument(index meilisearch.IndexManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Invalid HTTP method. Only DELETE is allowed.", http.StatusMethodNotAllowed)
			return
		}

		path := r.URL.Path
		id := strings.TrimPrefix(path, "/movies/")
		if id == "" {
			http.Error(w, "Missing or invalid ID", http.StatusBadRequest)
			return
		}

		var movie Movie

		documentQuery := &meilisearch.DocumentQuery{}

		err := index.GetDocument(id, documentQuery, &movie)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving document: %v", err), http.StatusInternalServerError)
			return
		}

		task, err := index.DeleteDocument(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error deleting document: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"task_uid":         task.TaskUID,
			"document_id":      id,
			"deleted_document": movie,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func handleGetDocument(index meilisearch.IndexManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid HTTP method. Only GET is allowed.", http.StatusMethodNotAllowed)
			return
		}

		path := r.URL.Path
		id := strings.TrimPrefix(path, "/movies/")
		if id == "" {
			http.Error(w, "Missing or invalid ID", http.StatusBadRequest)
			return
		}

		var movie Movie

		documentQuery := &meilisearch.DocumentQuery{}

		err := index.GetDocument(id, documentQuery, &movie)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving document: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(movie)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		}
	}
}

func handleUpdateDocument(index meilisearch.IndexManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Invalid HTTP method. Only PUT is allowed.", http.StatusMethodNotAllowed)
			return
		}

		path := r.URL.Path
		id := strings.TrimPrefix(path, "/movies/update/")
		if id == "" {
			http.Error(w, "Missing or invalid ID", http.StatusBadRequest)
			return
		}

		documentID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid ID format: %v", err), http.StatusBadRequest)
			return
		}

		var movie Movie
		err = json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error decoding document: %v", err), http.StatusBadRequest)
			return
		}

		movie.ID = documentID

		movies := []interface{}{movie}
		task, err := index.UpdateDocuments(movies)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error updating document: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"task_uid":    task.TaskUID,
			"document_id": documentID,
			"document":    movie,
		}
		json.NewEncoder(w).Encode(response)
	}
}
