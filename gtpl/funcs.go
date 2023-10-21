package gtpl

import (
	"strings"
)

const (
	TYPE_DB_INT      = "INT"
	TYPE_DB_SMALLINT = "SMALLINT"
	TYPE_DB_BIGINT   = "BIGINT"
	TYPE_DB_FLOAT    = "FLOAT"
	TYPE_DB_STRING   = "STRING"
	TYPE_DB_TEXT     = "TEXT"
)

func gotype(t string) string {
	result := "string"
	switch strings.ToUpper(t) {
	case TYPE_DB_FLOAT:
		result = "float64"
	case TYPE_DB_BIGINT:
		result = "int64"
	case TYPE_DB_INT, TYPE_DB_SMALLINT:
		result = "int"
	case TYPE_DB_TEXT, TYPE_DB_STRING:
		result = "string"
	}
	return result
}
