package main

type Movie struct {
	ID     int64    `json:"id"`
	Title  string   `json:"title"`
	Genres []string `json:"genres"`
}

func NewMovie(id int64, title string, genres []string) Movie {
	return Movie{
		ID:     id,
		Title:  title,
		Genres: genres,
	}
}
