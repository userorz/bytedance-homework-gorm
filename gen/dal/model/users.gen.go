// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePerson = "users"

// Person mapped from table <users>
type Person struct {
	UUID    string `gorm:"column:uuid" json:"uuid"`
	Name    string `gorm:"column:name" json:"name"`
	Age     int64  `gorm:"column:age" json:"age"`
	Version int64  `gorm:"column:version" json:"version"`
}

// TableName Person's table name
func (*Person) TableName() string {
	return TableNamePerson
}
