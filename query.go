package jally_orm

import (
	"fmt"
	"reflect"

	"github.com/s4kibs4mi/jally-orm/helper"
)

type Query struct {
	tableName string
	spaceName string
	model     interface{}
}

func NewQuery() Query {
	return Query{}
}

func (q Query) Table(name string) Query {
	q.tableName = name
	return q
}

func (q Query) Space(name string) Query {
	q.spaceName = name
	return q
}

func (q Query) Model(v interface{}) Query {
	q.model = v
	return q
}

func (q *Query) Create() string {
	vOf := reflect.ValueOf(q.model)
	qField := ""
	pKey := ""
	for i := 0; i < vOf.NumField(); i++ {
		typeOfField := vOf.Type().Field(i)
		name := helper.ToCQLXName(typeOfField)
		if i == 0 {
			qField = fmt.Sprintf("%s%s %s", qField, name, helper.ToCQLXType(typeOfField))
		} else {
			qField = fmt.Sprintf("%s,%s %s", qField, name, helper.ToCQLXType(typeOfField))
		}
		if helper.IsPrimaryKey(typeOfField) {
			if pKey == "" {
				pKey = name
			} else {
				pKey = fmt.Sprintf("%s,%s", pKey, name)
			}
		}
	}
	if pKey == "" {
		return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s(%s);", q.spaceName, q.tableName, qField)
	}
	cQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s(%s, PRIMARY KEY(%s));",
		q.spaceName, q.tableName, qField, pKey)
	return cQuery
}

func (q *Query) Insert() (string, []interface{}) {
	vOf := reflect.ValueOf(q.model)
	qField := ""
	qVal := ""
	var values []interface{}
	for i := 0; i < vOf.NumField(); i++ {
		valOfField := vOf.Field(i)
		typeOfField := vOf.Type().Field(i)

		if i == 0 {
			qField += helper.ToCQLXName(typeOfField)
			qVal += "?"
		} else {
			qField += "," + helper.ToCQLXName(typeOfField)
			qVal += ",?"
		}
		values = append(values, valOfField.Interface())
	}
	return fmt.Sprintf("INSERT INTO %s.%s(%s) VALUES(%s)", q.spaceName, q.tableName, qField, qVal), values
}

func (q *Query) FindByID() string {
	cQuery := fmt.Sprintf("SELECT * FROM %s.%s WHERE id = ?;", q.spaceName, q.tableName)
	return cQuery
}
