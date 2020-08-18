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


CREATE TABLE `character_tasks` (
  `charid` int(11) unsigned NOT NULL DEFAULT 0,
  `taskid` int(11) unsigned NOT NULL DEFAULT 0,
  `slot` int(11) unsigned NOT NULL DEFAULT 0,
  `type` tinyint(4) NOT NULL DEFAULT 0,
  `acceptedtime` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`charid`,`taskid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "charid": 57,    "taskid": 54,    "slot": 67,    "type": 44,    "acceptedtime": 36}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 4] column is set for unsigned



*/

// CharacterTasks struct is a row record of the character_tasks table in the eqemu database
type CharacterTasks struct {
	//[ 0] charid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Charid uint32 `gorm:"primary_key;column:charid;type:uint;default:0;" json:"charid"`
	//[ 1] taskid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Taskid uint32 `gorm:"primary_key;column:taskid;type:uint;default:0;" json:"taskid"`
	//[ 2] slot                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Slot uint32 `gorm:"column:slot;type:uint;default:0;" json:"slot"`
	//[ 3] type                                           tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Type int32 `gorm:"column:type;type:tinyint;default:0;" json:"type"`
	//[ 4] acceptedtime                                   uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [NULL]
	Acceptedtime null.Int `gorm:"column:acceptedtime;type:uint;" json:"acceptedtime"`
}

var character_tasksTableInfo = &TableInfo{
	Name: "character_tasks",
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

		&ColumnInfo{
			Index:              2,
			Name:               "slot",
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
			GoFieldName:        "Slot",
			GoFieldType:        "uint32",
			JSONFieldName:      "slot",
			ProtobufFieldName:  "slot",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "acceptedtime",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Acceptedtime",
			GoFieldType:        "null.Int",
			JSONFieldName:      "acceptedtime",
			ProtobufFieldName:  "acceptedtime",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterTasks) TableName() string {
	return "character_tasks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterTasks) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterTasks) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterTasks) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterTasks) TableInfo() *TableInfo {
	return character_tasksTableInfo
}
