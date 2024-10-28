package services

import (
	"fmt"
)

// GenerateReport executes a user-specified SQL query and returns the results
func GenerateReport(query string) ([]map[string]interface{}, error) {
	if DB == nil {
		return nil, fmt.Errorf("database connection is not established")
	}

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	results := []map[string]interface{}{}
	for rows.Next() {
		row := make([]interface{}, len(columns))
		rowPointers := make([]interface{}, len(columns))
		for i := range row {
			rowPointers[i] = &row[i]
		}

		if err := rows.Scan(rowPointers...); err != nil {
			return nil, err
		}

		result := map[string]interface{}{}
		for i, col := range columns {
			val := row[i]
			b, ok := val.([]byte)
			if ok {
				val = string(b)
			}
			result[col] = val
		}

		results = append(results, result)
	}

	return results, nil
}
