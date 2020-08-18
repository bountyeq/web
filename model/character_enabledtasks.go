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


CREATE TABLE `character_enabledtasks` (
  `charid` int(11) unsigned NOT NULL DEFAULT 0,
  `taskid` int(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`charid`,`taskid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "charid": 86,    "taskid": 67}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// CharacterEnabledtasks struct is a row record of the character_enabledtasks table in the eqemu database
type CharacterEnabledtasks struct {
	//[ 0] charid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Charid uint32 `gorm:"primary_key;column:charid;type:uint;default:0;" json:"charid"`
	//[ 1] taskid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Taskid uint32 `gorm:"primary_key;column:taskid;type:uint;default:0;" json:"taskid"`
}

var character_enabledtasksTableInfo = &TableInfo{
	Name: "character_enabledtasks",
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
			Name:               "taskid",
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
			GoFieldName:        "Taskid",
			GoFieldType:        "uint32",
			JSONFieldName:      "taskid",
			ProtobufFieldName:  "taskid",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterEnabledtasks) TableName() string {
	return "character_enabledtasks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterEnabledtasks) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterEnabledtasks) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterEnabledtasks) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterEnabledtasks) TableInfo() *TableInfo {
	return character_enabledtasksTableInfo
}
