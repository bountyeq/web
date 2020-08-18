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


CREATE TABLE `adventure_members` (
  `id` int(10) unsigned NOT NULL,
  `charid` int(10) unsigned NOT NULL,
  PRIMARY KEY (`charid`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "id": 85,    "charid": 83}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// AdventureMembers struct is a row record of the adventure_members table in the eqemu database
type AdventureMembers struct {
	//[ 0] id                                             uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	ID uint32 `gorm:"column:id;type:uint;" json:"id"`
	//[ 1] charid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	Charid uint32 `gorm:"primary_key;column:charid;type:uint;" json:"charid"`
}

var adventure_membersTableInfo = &TableInfo{
	Name: "adventure_members",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AdventureMembers) TableName() string {
	return "adventure_members"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AdventureMembers) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AdventureMembers) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AdventureMembers) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AdventureMembers) TableInfo() *TableInfo {
	return adventure_membersTableInfo
}
