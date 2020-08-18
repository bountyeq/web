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


CREATE TABLE `launcher` (
  `name` varchar(64) NOT NULL DEFAULT '',
  `dynamics` tinyint(3) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "dynamics": 17,    "name": "ooJbdxJAToTbRRRLOKFqwDqub"}


Comments
-------------------------------------
[ 1] column is set for unsigned



*/

// Launcher struct is a row record of the launcher table in the eqemu database
type Launcher struct {
	//[ 0] name                                           varchar(64)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Name string `gorm:"primary_key;column:name;type:varchar;size:64;default:'';" json:"name"`
	//[ 1] dynamics                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Dynamics uint32 `gorm:"column:dynamics;type:utinyint;default:0;" json:"dynamics"`
}

var launcherTableInfo = &TableInfo{
	Name: "launcher",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
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
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "dynamics",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Dynamics",
			GoFieldType:        "uint32",
			JSONFieldName:      "dynamics",
			ProtobufFieldName:  "dynamics",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *Launcher) TableName() string {
	return "launcher"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *Launcher) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *Launcher) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *Launcher) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *Launcher) TableInfo() *TableInfo {
	return launcherTableInfo
}
