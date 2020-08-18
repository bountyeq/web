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


CREATE TABLE `completed_tasks` (
  `charid` int(11) unsigned NOT NULL DEFAULT 0,
  `completedtime` int(11) unsigned NOT NULL DEFAULT 0,
  `taskid` int(11) unsigned NOT NULL DEFAULT 0,
  `activityid` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`charid`,`completedtime`,`taskid`,`activityid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "charid": 28,    "completedtime": 91,    "taskid": 93,    "activityid": 64}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// CompletedTasks struct is a row record of the completed_tasks table in the eqemu database
type CompletedTasks struct {
	//[ 0] charid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Charid uint32 `gorm:"primary_key;column:charid;type:uint;default:0;" json:"charid"`
	//[ 1] completedtime                                  uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Completedtime uint32 `gorm:"primary_key;column:completedtime;type:uint;default:0;" json:"completedtime"`
	//[ 2] taskid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Taskid uint32 `gorm:"primary_key;column:taskid;type:uint;default:0;" json:"taskid"`
	//[ 3] activityid                                     int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Activityid int32 `gorm:"primary_key;column:activityid;type:int;default:0;" json:"activityid"`
}

var completed_tasksTableInfo = &TableInfo{
	Name: "completed_tasks",
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
			Name:               "completedtime",
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
			GoFieldName:        "Completedtime",
			GoFieldType:        "uint32",
			JSONFieldName:      "completedtime",
			ProtobufFieldName:  "completedtime",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "activityid",
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
			GoFieldName:        "Activityid",
			GoFieldType:        "int32",
			JSONFieldName:      "activityid",
			ProtobufFieldName:  "activityid",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CompletedTasks) TableName() string {
	return "completed_tasks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CompletedTasks) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CompletedTasks) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CompletedTasks) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CompletedTasks) TableInfo() *TableInfo {
	return completed_tasksTableInfo
}
