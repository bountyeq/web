package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `name_filter` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `name_search_index` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=278 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 96,    "name": "IyqtxUHohYGgMcGlLHDjReDvH"}



*/

// NameFilter struct is a row record of the name_filter table in the eqemu database
type NameFilter struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] name                                           varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: ['']
	Name string `gorm:"column:name;type:varchar;size:30;default:'';" json:"name"`
}

var name_filterTableInfo = &TableInfo{
	Name: "name_filter",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (n *NameFilter) TableName() string {
	return "name_filter"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (n *NameFilter) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (n *NameFilter) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (n *NameFilter) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (n *NameFilter) TableInfo() *TableInfo {
	return name_filterTableInfo
}
