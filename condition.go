package jally_orm

import (
	"fmt"
	"strings"
)

type Condition struct {
	qFields string
	qValues []interface{}
}

func (c *Condition) Eq(key string, val interface{}) {
	c.qFields = fmt.Sprintf("%s %s = ?", c.qFields, key)
	c.qValues = append(c.qValues, val)
}

func (c *Condition) Where() string {
	return strings.TrimSpace(c.qFields)
}
