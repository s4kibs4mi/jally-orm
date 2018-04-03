package helper

import (
	"reflect"
	"testing"
	"time"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/assert"
	"fmt"
)

type Table struct {
	Tuuid      gocql.UUID
	Tvarchar   string
	Tbigint    int64
	Ttimestamp time.Time
	Tblob      []byte
	Tbool      bool
	Tfloat     float32
	Tdouble    float64
	Ttinyint   int8
	Tsmallint  int16
	Tint       int32
}

func TestToCQLXType(t *testing.T) {
	d := Table{
		Tuuid:      gocql.TimeUUID(),
		Tvarchar:   "custom varchar",
		Tbigint:    int64(984932984858),
		Ttimestamp: time.Now().Truncate(time.Millisecond).UTC(),
		Tblob:      []byte("blob field"),
		Tbool:      false,
		Tfloat:     float32(3.1416),
		Tdouble:    float64(3.14167890374624854),
		Ttinyint:   int8(10),
		Tsmallint:  int16(103),
		Tint:       int32(1038),
	}
	x := reflect.ValueOf(d)
	assert.Equal(t, "UUID", ToCQLXType(x.Type().Field(0)))
	assert.Equal(t, "TEXT", ToCQLXType(x.Type().Field(1)))
	assert.Equal(t, "BIGINT", ToCQLXType(x.Type().Field(2)))
	assert.Equal(t, "timestamp", ToCQLXType(x.Type().Field(3)))
	assert.Equal(t, "BOLB", ToCQLXType(x.Type().Field(4)))
	assert.Equal(t, "BOOLEAN", ToCQLXType(x.Type().Field(5)))
	assert.Equal(t, "FLOAT", ToCQLXType(x.Type().Field(6)))
	assert.Equal(t, "DOUBLE", ToCQLXType(x.Type().Field(7)))
	assert.Equal(t, "TINYINT", ToCQLXType(x.Type().Field(8)))
	assert.Equal(t, "SMALLINT", ToCQLXType(x.Type().Field(9)))
	assert.Equal(t, "INT", ToCQLXType(x.Type().Field(10)))
}

type Message struct {
	From string
	Body string
	When time.Time
}

func TestMapToStruct(t *testing.T) {
	m := map[string]interface{}{"From": "Sakib", "Body": "Hello World", "When": time.Now()}
	message := Message{}
	MapToStruct(&message, m)
	fmt.Println(message)
}
