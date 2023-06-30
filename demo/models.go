package main

import (
	"database/sql/driver"
	"github.com/archsh/go.xql/dialects/postgres"
	_ "github.com/lib/pq"
	"time"
)

type School struct {
	Id          int                  `json:"id" xql:"type=serial,pk"`
	Name        string               `json:"name" xql:"size=24,unique,index"`
	Tags        postgres.StringArray `json:"tags" xql:"size=32,nullable"`
	Description string               `json:"description"  xql:"name=desc,type=text,size=24,nullable=false,default=''"`
}

func (c School) TableName() string {
	return "schools"
}

type Character struct {
	postgres.JSON
	Attitude string
	Height   int
	Weight   int
}

//func (j Character) Declare(props xql.PropertySet) string {
//    return "JSONB"
//}

func (j *Character) Scan(value interface{}) error {
	return postgres.JsonbScan(j, value)
}

func (j Character) Value() (driver.Value, error) {
	return postgres.JsonbValue(j)
}

type People struct {
	Id          int        `json:"id" xql:"type=serial,pk"`
	FullName    string     `json:"fullName" xql:"size=80,unique=true,index=true"`
	FirstName   string     `json:"firstName" xql:"size=24,default=''"`
	MiddleName  string     `json:"middleName" xql:"size=24,default=''"`
	LastName    string     `json:"lastName" xql:"size=24,default=''"`
	Region      string     `json:"region"  xql:"size=24,nullable=true"`
	Age         int        `json:"age" xql:"check=(age>18)"`
	SchoolId    int        `json:"schoolId"  xql:"type=integer,fk=schools.id,ondelete=CASCADE"`
	Description string     `json:"description"  xql:"name=desc,type=text,size=24,default=''"`
	Created     *time.Time `json:"created"  xql:"type=timestamp,default=Now()"`
	Updated     *time.Time `json:"Updated"  xql:"type=timestamp,default=Now()"`
}

type Teacher struct {
	People
	Degree string `json:"degree" xql:"size=64,default=''"`
}

func (t Teacher) TableName() string {
	return "teachers"
}

type Student struct {
	People
	Grade      string               `json:"grade" xql:"size=32,default=''"`
	Attributes postgres.HSTORE      `json:"attributes" xql:"nullable"`
	Scores     postgres.DoubleArray `json:"scores" xql:"nullable"`
	Character  Character            `json:"character" xql:"nullable"`
}

func (c Student) TableName() string {
	return "students"
}
