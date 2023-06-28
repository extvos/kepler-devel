package main

import xql "github.com/archsh/go.xql"

type Student struct {
	Id   xql.BigSerial `xql:"pk"`
	Name string        `xql:"type=varchar,size=64,unique,nullable=false"`
	Age  int           `xql:"type=Integer,nullable=false,default=0"`
}

func (Student) TableName() string {
	return "students"
}
