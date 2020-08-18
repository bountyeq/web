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


CREATE TABLE `timers` (
  `char_id` int(11) NOT NULL DEFAULT 0,
  `type` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `start` int(10) unsigned NOT NULL DEFAULT 0,
  `duration` int(10) unsigned NOT NULL DEFAULT 0,
  `enable` tinyint(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`char_id`,`type`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "char_id": 47,    "type": 46,    "start": 29,    "duration": 89,    "enable": 31}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned



*/

// Timers struct is a row record of the timers table in the eqemu database
type Timers struct {
	//[ 0] char_id                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	CharID int32 `gorm:"primary_key;column:char_id;type:int;default:0;" json:"char_id"`
	//[ 1] type                                           umediumint           null: false  primary: true   isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Type uint32 `gorm:"primary_key;column:type;type:umediumint;default:0;" json:"type"`
	//[ 2] start                                          uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Start uint32 `gorm:"column:start;type:uint;default:0;" json:"start"`
	//[ 3] duration                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Duration uint32 `gorm:"column:duration;type:uint;default:0;" json:"duration"`
	//[ 4] enable                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Enable int32 `gorm:"column:enable;type:tinyint;default:0;" json:"enable"`
}

var timersTableInfo = &TableInfo{
	Name: "timers",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "char_id",
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
			GoFieldName:        "CharID",
			GoFieldType:        "int32",
			JSONFieldName:      "char_id",
			ProtobufFieldName:  "char_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "type",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "umediumint",
			DatabaseTypePretty: "umediumint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "umediumint",
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
			Name:               "start",
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
			GoFieldName:        "Start",
			GoFieldType:        "uint32",
			JSONFieldName:      "start",
			ProtobufFieldName:  "start",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "duration",
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
			GoFieldName:        "Duration",
			GoFieldType:        "uint32",
			JSONFieldName:      "duration",
			ProtobufFieldName:  "duration",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "enable",
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
			GoFieldName:        "Enable",
			GoFieldType:        "int32",
			JSONFieldName:      "enable",
			ProtobufFieldName:  "enable",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *Timers) TableName() string {
	return "timers"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *Timers) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *Timers) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *Timers) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *Timers) TableInfo() *TableInfo {
	return timersTableInfo
}
