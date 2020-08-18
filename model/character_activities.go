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


CREATE TABLE `character_activities` (
  `charid` int(11) unsigned NOT NULL DEFAULT 0,
  `taskid` int(11) unsigned NOT NULL DEFAULT 0,
  `activityid` int(11) unsigned NOT NULL DEFAULT 0,
  `donecount` int(11) unsigned NOT NULL DEFAULT 0,
  `completed` tinyint(1) DEFAULT 0,
  PRIMARY KEY (`charid`,`taskid`,`activityid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "charid": 13,    "taskid": 34,    "activityid": 18,    "donecount": 67,    "completed": 88}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned



*/

// CharacterActivities struct is a row record of the character_activities table in the eqemu database
type CharacterActivities struct {
	//[ 0] charid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Charid uint32 `gorm:"primary_key;column:charid;type:uint;default:0;" json:"charid"`
	//[ 1] taskid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Taskid uint32 `gorm:"primary_key;column:taskid;type:uint;default:0;" json:"taskid"`
	//[ 2] activityid                                     uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Activityid uint32 `gorm:"primary_key;column:activityid;type:uint;default:0;" json:"activityid"`
	//[ 3] donecount                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Donecount uint32 `gorm:"column:donecount;type:uint;default:0;" json:"donecount"`
	//[ 4] completed                                      tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Completed null.Int `gorm:"column:completed;type:tinyint;default:0;" json:"completed"`
}

var character_activitiesTableInfo = &TableInfo{
	Name: "character_activities",
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
			Name:               "activityid",
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
			GoFieldName:        "Activityid",
			GoFieldType:        "uint32",
			JSONFieldName:      "activityid",
			ProtobufFieldName:  "activityid",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "donecount",
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
			GoFieldName:        "Donecount",
			GoFieldType:        "uint32",
			JSONFieldName:      "donecount",
			ProtobufFieldName:  "donecount",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "completed",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Completed",
			GoFieldType:        "null.Int",
			JSONFieldName:      "completed",
			ProtobufFieldName:  "completed",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterActivities) TableName() string {
	return "character_activities"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterActivities) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterActivities) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterActivities) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterActivities) TableInfo() *TableInfo {
	return character_activitiesTableInfo
}
