package helper

import (
	"reflect"
	"strings"
)

// ToCQLXType convert go struct type to CQL type
func ToCQLXType(f reflect.StructField) string {
	switch f.Type.String() {
	case "gocql.UUID":
		return "UUID"
	case "string":
		return "TEXT"
	case "int8":
		return "TINYINT"
	case "int16":
		return "SMALLINT"
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
	case "time.Time":
		return "timestamp"
	case "[]uint8":
		return "BOLB"
	}
	return ""
}

// ToCQLXName convert go struct field to CQL name
func ToCQLXName(f reflect.StructField) string {
	val, ok := f.Tag.Lookup("json")
	if ok {
		return val
	}
	return f.Name
}

// IsPrimaryKey checks if the field is primary key
func IsPrimaryKey(f reflect.StructField) bool {
	val := f.Tag.Get("jorm")
	return strings.Contains(strings.ToLower(val), "primary_key")
}

func MapToStruct(dst interface{}, values map[string]interface{}) {
	vOf := reflect.ValueOf(dst).Elem()
	tOf := reflect.TypeOf(dst).Elem()
	for i := 0; i < vOf.NumField(); i++ {
		field := tOf.Field(i)
		tag := ToCQLXName(field)
		vField := vOf.Field(i)
		if vField.IsValid() && vField.CanSet() {
			vField.Set(reflect.ValueOf(values[tag]))
		}
	}
}
