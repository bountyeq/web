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


CREATE TABLE `goallists` (
  `listid` int(10) unsigned NOT NULL DEFAULT 0,
  `entry` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`listid`,`entry`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "listid": 95,    "entry": 12}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// Goallists struct is a row record of the goallists table in the eqemu database
type Goallists struct {
	//[ 0] listid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Listid uint32 `gorm:"primary_key;column:listid;type:uint;default:0;" json:"listid"`
	//[ 1] entry                                          uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Entry uint32 `gorm:"primary_key;column:entry;type:uint;default:0;" json:"entry"`
}

var goallistsTableInfo = &TableInfo{
	Name: "goallists",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "listid",
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
			GoFieldName:        "Listid",
			GoFieldType:        "uint32",
			JSONFieldName:      "listid",
			ProtobufFieldName:  "listid",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "entry",
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
			GoFieldName:        "Entry",
			GoFieldType:        "uint32",
			JSONFieldName:      "entry",
			ProtobufFieldName:  "entry",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *Goallists) TableName() string {
	return "goallists"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *Goallists) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *Goallists) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *Goallists) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *Goallists) TableInfo() *TableInfo {
	return goallistsTableInfo
}
