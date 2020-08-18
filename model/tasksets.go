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


CREATE TABLE `tasksets` (
  `id` int(11) unsigned NOT NULL DEFAULT 0,
  `taskid` int(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`taskid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 21,    "taskid": 96}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// Tasksets struct is a row record of the tasksets table in the eqemu database
type Tasksets struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ID uint32 `gorm:"primary_key;column:id;type:uint;default:0;" json:"id"`
	//[ 1] taskid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Taskid uint32 `gorm:"primary_key;column:taskid;type:uint;default:0;" json:"taskid"`
}

var tasksetsTableInfo = &TableInfo{
	Name: "tasksets",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
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
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
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
func (t *Tasksets) TableName() string {
	return "tasksets"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *Tasksets) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *Tasksets) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *Tasksets) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *Tasksets) TableInfo() *TableInfo {
	return tasksetsTableInfo
}
