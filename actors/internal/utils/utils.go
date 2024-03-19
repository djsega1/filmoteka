package utils

import (
	"strings"
)

// formats SQL query for logging
func FormatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}
