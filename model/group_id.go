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


CREATE TABLE `group_id` (
  `groupid` int(4) NOT NULL,
  `charid` int(4) NOT NULL,
  `name` varchar(64) NOT NULL,
  `ismerc` tinyint(3) NOT NULL DEFAULT 0,
  PRIMARY KEY (`groupid`,`charid`,`ismerc`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "name": "bsjLruHJZSUCkGykYaksqSSQv",    "ismerc": 65,    "groupid": 17,    "charid": 74}



*/

// GroupID struct is a row record of the group_id table in the eqemu database
type GroupID struct {
	//[ 0] groupid                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	Groupid int32 `gorm:"primary_key;column:groupid;type:int;" json:"groupid"`
	//[ 1] charid                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	Charid int32 `gorm:"primary_key;column:charid;type:int;" json:"charid"`
	//[ 2] name                                           varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Name string `gorm:"column:name;type:varchar;size:64;" json:"name"`
	//[ 3] ismerc                                         tinyint              null: false  primary: true   isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Ismerc int32 `gorm:"primary_key;column:ismerc;type:tinyint;default:0;" json:"ismerc"`
}

var group_idTableInfo = &TableInfo{
	Name: "group_id",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "groupid",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Groupid",
			GoFieldType:        "int32",
			JSONFieldName:      "groupid",
			ProtobufFieldName:  "groupid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "charid",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Charid",
			GoFieldType:        "int32",
			JSONFieldName:      "charid",
			ProtobufFieldName:  "charid",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "ismerc",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Ismerc",
			GoFieldType:        "int32",
			JSONFieldName:      "ismerc",
			ProtobufFieldName:  "ismerc",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *GroupID) TableName() string {
	return "group_id"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GroupID) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GroupID) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GroupID) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GroupID) TableInfo() *TableInfo {
	return group_idTableInfo
}
