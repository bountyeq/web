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


CREATE TABLE `spawn_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `zone` varchar(32) DEFAULT NULL,
  `cond_id` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `name` varchar(255) NOT NULL DEFAULT '',
  `period` int(10) unsigned NOT NULL DEFAULT 0,
  `next_minute` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `next_hour` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `next_day` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `next_month` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `next_year` int(10) unsigned NOT NULL DEFAULT 0,
  `enabled` tinyint(4) NOT NULL DEFAULT 1,
  `action` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `argument` mediumint(9) NOT NULL DEFAULT 0,
  `strict` tinyint(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "strict": 99,    "id": 25,    "enabled": 15,    "next_minute": 12,    "zone": "mEOftqIYeCIoGxfJGbPOUiIbg",    "cond_id": 65,    "next_year": 9,    "next_hour": 55,    "next_day": 5,    "next_month": 80,    "action": 15,    "argument": 25,    "name": "IaqtcIktJVOUidIZYakYgLFSA",    "period": 17}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[11] column is set for unsigned



*/

// SpawnEvents struct is a row record of the spawn_events table in the eqemu database
type SpawnEvents struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] zone                                           varchar(32)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 32      default: [NULL]
	Zone null.String `gorm:"column:zone;type:varchar;size:32;" json:"zone"`
	//[ 2] cond_id                                        umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	CondID uint32 `gorm:"column:cond_id;type:umediumint;default:0;" json:"cond_id"`
	//[ 3] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: ['']
	Name string `gorm:"column:name;type:varchar;size:255;default:'';" json:"name"`
	//[ 4] period                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Period uint32 `gorm:"column:period;type:uint;default:0;" json:"period"`
	//[ 5] next_minute                                    utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	NextMinute uint32 `gorm:"column:next_minute;type:utinyint;default:0;" json:"next_minute"`
	//[ 6] next_hour                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	NextHour uint32 `gorm:"column:next_hour;type:utinyint;default:0;" json:"next_hour"`
	//[ 7] next_day                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	NextDay uint32 `gorm:"column:next_day;type:utinyint;default:0;" json:"next_day"`
	//[ 8] next_month                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	NextMonth uint32 `gorm:"column:next_month;type:utinyint;default:0;" json:"next_month"`
	//[ 9] next_year                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	NextYear uint32 `gorm:"column:next_year;type:uint;default:0;" json:"next_year"`
	//[10] enabled                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Enabled int32 `gorm:"column:enabled;type:tinyint;default:1;" json:"enabled"`
	//[11] action                                         utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Action uint32 `gorm:"column:action;type:utinyint;default:0;" json:"action"`
	//[12] argument                                       mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	Argument int32 `gorm:"column:argument;type:mediumint;default:0;" json:"argument"`
	//[13] strict                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Strict int32 `gorm:"column:strict;type:tinyint;default:0;" json:"strict"`
}

var spawn_eventsTableInfo = &TableInfo{
	Name: "spawn_events",
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
			IsAutoIncrement:    true,
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
			Name:               "zone",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Zone",
			GoFieldType:        "null.String",
			JSONFieldName:      "zone",
			ProtobufFieldName:  "zone",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "cond_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "umediumint",
			DatabaseTypePretty: "umediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "umediumint",
			ColumnLength:       -1,
			GoFieldName:        "CondID",
			GoFieldType:        "uint32",
			JSONFieldName:      "cond_id",
			ProtobufFieldName:  "cond_id",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "period",
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
			GoFieldName:        "Period",
			GoFieldType:        "uint32",
			JSONFieldName:      "period",
			ProtobufFieldName:  "period",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "next_minute",
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
			GoFieldName:        "NextMinute",
			GoFieldType:        "uint32",
			JSONFieldName:      "next_minute",
			ProtobufFieldName:  "next_minute",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "next_hour",
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
			GoFieldName:        "NextHour",
			GoFieldType:        "uint32",
			JSONFieldName:      "next_hour",
			ProtobufFieldName:  "next_hour",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "next_day",
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
			GoFieldName:        "NextDay",
			GoFieldType:        "uint32",
			JSONFieldName:      "next_day",
			ProtobufFieldName:  "next_day",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "next_month",
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
			GoFieldName:        "NextMonth",
			GoFieldType:        "uint32",
			JSONFieldName:      "next_month",
			ProtobufFieldName:  "next_month",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "next_year",
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
			GoFieldName:        "NextYear",
			GoFieldType:        "uint32",
			JSONFieldName:      "next_year",
			ProtobufFieldName:  "next_year",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "enabled",
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
			GoFieldName:        "Enabled",
			GoFieldType:        "int32",
			JSONFieldName:      "enabled",
			ProtobufFieldName:  "enabled",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "action",
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
			GoFieldName:        "Action",
			GoFieldType:        "uint32",
			JSONFieldName:      "action",
			ProtobufFieldName:  "action",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "argument",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "mediumint",
			DatabaseTypePretty: "mediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "mediumint",
			ColumnLength:       -1,
			GoFieldName:        "Argument",
			GoFieldType:        "int32",
			JSONFieldName:      "argument",
			ProtobufFieldName:  "argument",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "strict",
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
			GoFieldName:        "Strict",
			GoFieldType:        "int32",
			JSONFieldName:      "strict",
			ProtobufFieldName:  "strict",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SpawnEvents) TableName() string {
	return "spawn_events"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SpawnEvents) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SpawnEvents) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SpawnEvents) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SpawnEvents) TableInfo() *TableInfo {
	return spawn_eventsTableInfo
}
