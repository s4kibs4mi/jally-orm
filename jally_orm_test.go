package jally_orm

import (
	"testing"
	"github.com/gocql/gocql"
	"time"
	"github.com/stretchr/testify/assert"
)

type Student struct {
	ID        gocql.UUID `json:"id" jorm:"primary_key"`
	Name      string     `json:"name"`
	Roll      int        `json:"roll"`
	CGPA      float32    `json:"cgpa"`
	IsPresent bool       `json:"is_present"`
}

func TestJallyORM_Create(t *testing.T) {
	config := gocql.NewCluster("localhost:9042")
	config.Consistency = gocql.Quorum
	config.Keyspace = "test"

	config.Timeout = 5 * time.Second
	config.PoolConfig = gocql.PoolConfig{
		HostSelectionPolicy: gocql.RoundRobinHostPolicy(),
	}
	config.NumConns = 5
	orm, err := NewSession(config)
	if err != nil {
		panic(err)
	}
	s := Student{}
	q := NewQuery().Space("test").Table("students").Model(s)
	err = orm.Create(q)
	assert.Nil(t, err, "Something went wrong")
}

func TestJallyORM_Save(t *testing.T) {
	config := gocql.NewCluster("localhost:9042")
	config.Consistency = gocql.Quorum
	config.Keyspace = "test"

	config.Timeout = 5 * time.Second
	config.PoolConfig = gocql.PoolConfig{
		HostSelectionPolicy: gocql.RoundRobinHostPolicy(),
	}
	config.NumConns = 5
	orm, err := NewSession(config)
	if err != nil {
		panic(err)
	}
	s := Student{
		ID:   gocql.TimeUUID(),
		Name: "Sakib",
		Roll: 12345,
		CGPA: 3.50,
	}
	q := NewQuery().Space("test").Table("students").Model(s)
	err = orm.Save(q)
	assert.Nil(t, err)
}
