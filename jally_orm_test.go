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
	CreatedAt time.Time  `json:"created_at"`
}

func (s *Student) Clean() JallyORMModel {
	return &Student{}
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
	q := NewQuery().Space("test").Table("students").Model(&s)
	err = orm.CreateTable(q)
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
		ID:        gocql.TimeUUID(),
		Name:      "Nur",
		Roll:      12346,
		CGPA:      3.10,
		IsPresent: false,
		CreatedAt: time.Now(),
	}
	q := NewQuery().Space("test").Table("students").Model(&s)
	err = orm.Save(q)
	assert.Nil(t, err)
}

func TestQuery_FindByID(t *testing.T) {
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
	queryID, err := gocql.ParseUUID("cf2921d6-3680-11e8-897d-d0a637eb34d1")
	assert.Nil(t, err)
	s := Student{}
	q := NewQuery().Space("test").Table("students").Model(&s)
	v, err := orm.FindByID(queryID, q)
	assert.Nil(t, err)
	if queryID == v.(*Student).ID {
		t.Log(*v.(*Student))
	}
}

func TestJallyORM_Find(t *testing.T) {
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
	q := NewQuery().Space("test").Table("students").Model(&s)
	c := Condition{}
	//c.Eq("name", "Sakib")

	students := orm.Find(c, q)
	t.Log("Len : ", len(students))
	for _, v := range students {
		t.Log(*v.(*Student))
	}
}
