// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

const TableNameMLMDEnv = "MLMDEnv"

// MLMDEnv mapped from table <MLMDEnv>
type MLMDEnv struct {
	SchemaVersion int64 `gorm:"column:schema_version;primaryKey" json:"-"`
}

// TableName MLMDEnv's table name
func (*MLMDEnv) TableName() string {
	return TableNameMLMDEnv
}