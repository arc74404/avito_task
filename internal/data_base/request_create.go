package database

import (
	"avito_go/internal/models"
	"fmt"
	"strings"
)

func generateAddPlaceholders(us_id uint, arr []int) []string {
	placeholders := make([]string, len(arr))

	for i := range placeholders {
		placeholders[i] = fmt.Sprintf("(%d, %d)", us_id, arr[i])
	}
	return placeholders
}

func generateRemovePlaceholders(us_id uint, arr []int) []string {
	placeholders := make([]string, len(arr))

	for i := range placeholders {
		placeholders[i] = fmt.Sprintf("(user_id = %d and segment_id = %d)", us_id, arr[i])
	}
	return placeholders
}

func AddRemoveSegmentsCreateRequests(us_id uint, data *models.AddRemoveUserUpdJsonData) []string {
	var queries []string

	add_placeholders := generateAddPlaceholders(us_id, data.Add)
	add_query := fmt.Sprintf(
		"INSERT INTO user_segments (user_id, segment_id) VALUES %s ON CONFLICT (user_id, segment_id) DO NOTHING;",
		strings.Join(add_placeholders, ", "),
	)
	queries = append(queries, add_query)

	remove_placeholders := generateRemovePlaceholders(us_id, data.Remove)
	remove_query := fmt.Sprintf(
		"DELETE FROM user_segments WHERE %s",
		strings.Join(remove_placeholders, " OR "),
	)
	queries = append(queries, remove_query)

	return queries
}
