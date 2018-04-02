package jally_orm

import (
	"reflect"
	"strings"
)

func toCQLXType(f reflect.StructField) string {
	switch f.Type.String() {
	case "gocql.UUID":
		return "UUID"
	case "string":
		return "TEXT"
	case "int":
		return "INT"
	case "int32":
		return "INT"
	case "int64":
		return "BIGINT"
	case "float32":
		return "FLOAT"
	case "float64":
		return "DOUBLE"
	case "bool":
		return "BOOLEAN"
	case "time":
		return "timestamp"
	}
	return ""
}

func toCQLXName(f reflect.StructField) string {
	val, ok := f.Tag.Lookup("json")
	if ok {
		return val
	}
	return f.Name
}

func isPrimaryKey(f reflect.StructField) bool {
	val := f.Tag.Get("jorm")
	return strings.Contains(strings.ToLower(val), "primary_key")
}

func isUUID(f reflect.StructField) bool {
	if f.Type.Name() == "gocql.UUID" {
		return true
	}
	return false
}
