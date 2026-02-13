package handlers

import (
	// "avito_go/internal/models"
	"database/sql"
	// "encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)



func PutUserSegments(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slug := vars["slug"]

		if slug == "" {
			http.Error(w, `{"error": "Slug is required"}`, http.StatusBadRequest)
		}

		
	}
}
