package jally_orm

import (
	"github.com/gocql/gocql"
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

// Create creates table of model if not exists
func (orm *JallyORM) Create(q Query) error {
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

func (orm *JallyORM) Find(v interface{}, q Query) error {
	return nil
}

func (orm *JallyORM) Count(v interface{}, q Query) (int, error) {
	return -1, nil
}
