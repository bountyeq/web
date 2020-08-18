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


CREATE TABLE `friends` (
  `charid` int(10) unsigned NOT NULL,
  `type` tinyint(1) unsigned NOT NULL DEFAULT 1 COMMENT '1 = Friend, 0 = Ignore',
  `name` varchar(64) NOT NULL,
  PRIMARY KEY (`charid`,`type`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "name": "SIEoFZVvxHJsaZqmpMeHtEmMT",    "charid": 10,    "type": 93}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// Friends struct is a row record of the friends table in the eqemu database
type Friends struct {
	//[ 0] charid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	Charid uint32 `gorm:"primary_key;column:charid;type:uint;" json:"charid"`
	//[ 1] type                                           utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	Type uint32 `gorm:"primary_key;column:type;type:utinyint;default:1;" json:"type"` // 1 = Friend, 0 = Ignore
	//[ 2] name                                           varchar(64)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 64      default: []
	Name string `gorm:"primary_key;column:name;type:varchar;size:64;" json:"name"`
}

var friendsTableInfo = &TableInfo{
	Name: "friends",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "charid",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Charid",
			GoFieldType:        "uint32",
			JSONFieldName:      "charid",
			ProtobufFieldName:  "charid",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "type",
			Comment:            `1 = Friend, 0 = Ignore`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "uint32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "uint32",
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
			IsPrimaryKey:       true,
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
	},
}

// TableName sets the insert table name for this struct type
func (f *Friends) TableName() string {
	return "friends"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *Friends) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *Friends) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *Friends) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *Friends) TableInfo() *TableInfo {
	return friendsTableInfo
}
