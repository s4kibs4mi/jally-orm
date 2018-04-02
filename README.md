# JallyORM
An ORM library for Cassandra and ScyllaDB written in Golang.

Status : Development ongoing

Example,

Connect,
```go
config := gocql.NewCluster("localhost:9042")
config.Consistency = gocql.Quorum
config.Keyspace = "test"

config.Timeout = 5 * time.Second
config.PoolConfig = gocql.PoolConfig{
	HostSelectionPolicy: gocql.RoundRobinHostPolicy(),
}
config.NumConns = 5
orm, err := NewSession(config)
```

Create Table,
```go
orm, err := NewSession(config)
if err != nil {
	panic(err)
}
s := Student{}
q := NewQuery().Space("test").Table("students").Model(s)
err = orm.Create(q)
```

Insert value,
```go
orm, err := NewSession(config)
if err != nil {
	panic(err)
}
s := Student{
	ID:        gocql.TimeUUID(),
	Name:      "Sakib",
	Roll:      12345,
	CGPA:      3.50,
	IsPresent: true,
}
q := NewQuery().Space("test").Table("students").Model(s)
err = orm.Save(q)
```

#### License
Copyright © Sakib Sami

Distributed under [MIT](https://github.com/s4kibs4mi/jally-orm/blob/master/LICENSE) license
