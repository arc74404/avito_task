package handlers

import (
	// "avito_go/internal/models"
	"avito_go/internal/models"
	"database/sql"
	"encoding/json"
	"strconv"

	// "encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func UpdateUserSegments(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		user_id, err := strconv.Atoi(vars["user_id"])

		var data models.AddRemoveUserUpdJsonData

		json.NewDecoder(r.Body).Decode(&data)
	}
}
