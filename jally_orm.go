package jally_orm

import (
	"github.com/gocql/gocql"
	"github.com/s4kibs4mi/jally-orm/helper"
)

// JallyORM represents JallyORM instance
type JallyORM struct {
	*gocql.Session
}

// NewSession create new database session & returns instance of JallyORM
func NewSession(cfg *gocql.ClusterConfig) (*JallyORM, error) {
	orm := &JallyORM{}
	session, err := cfg.CreateSession()
	if err != nil {
		return nil, err
	}
	orm.Session = session
	return orm, nil
}

// CreateTable creates table of model if not exists
func (orm *JallyORM) CreateTable(q Query) error {
	return orm.Query(q.Create()).Exec()
}

// Save inserts data of model to db
func (orm *JallyORM) Save(q Query) error {
	qTxt, values := q.Insert()
	return orm.Query(qTxt, values...).Exec()
}

func (orm *JallyORM) Update(v interface{}, u Updater) error {
	return nil
}

func (orm *JallyORM) Delete(v interface{}, u Updater) error {
	return nil
}

// FindByID queries by id and returns interface, error
func (orm *JallyORM) FindByID(ID interface{}, q Query) (interface{}, error) {
	qTxt := q.FindByID()
	m := map[string]interface{}{}
	ok := orm.Query(qTxt, ID).Iter().MapScan(m)
	if !ok {
		return nil, gocql.ErrNotFound
	}
	var dst interface{}
	helper.MapToStruct(dst, m)
	return dst, nil
}

// FindByID queries by id and returns list of interface
func (orm *JallyORM) Find(c Condition, q Query) []interface{} {
	qTxt := q.Find(c)
	var items []interface{}
	m := map[string]interface{}{}
	it := orm.Query(qTxt, c.qValues...).Iter()
	for it.MapScan(m) {
		helper.MapToStruct(q.model, m)
		items = append(items, q.model)
	}
	return items
}

func (orm *JallyORM) Count(v interface{}, q Query) (int, error) {
	return -1, nil
}
