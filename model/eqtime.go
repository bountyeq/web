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


CREATE TABLE `eqtime` (
  `minute` tinyint(4) NOT NULL DEFAULT 0,
  `hour` tinyint(4) NOT NULL DEFAULT 0,
  `day` tinyint(4) NOT NULL DEFAULT 0,
  `month` tinyint(4) NOT NULL DEFAULT 0,
  `year` int(4) NOT NULL DEFAULT 0,
  `realtime` int(11) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "month": 74,    "year": 30,    "realtime": 49,    "minute": 36,    "hour": 92,    "day": 91}


Comments
-------------------------------------
[ 0] Warning table: eqtime does not have a primary key defined, setting col position 1 minute as primary key




*/

// Eqtime struct is a row record of the eqtime table in the eqemu database
type Eqtime struct {
	//[ 0] minute                                         tinyint              null: false  primary: true   isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Minute int32 `gorm:"primary_key;column:minute;type:tinyint;default:0;" json:"minute"`
	//[ 1] hour                                           tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Hour int32 `gorm:"column:hour;type:tinyint;default:0;" json:"hour"`
	//[ 2] day                                            tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Day int32 `gorm:"column:day;type:tinyint;default:0;" json:"day"`
	//[ 3] month                                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Month int32 `gorm:"column:month;type:tinyint;default:0;" json:"month"`
	//[ 4] year                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Year int32 `gorm:"column:year;type:int;default:0;" json:"year"`
	//[ 5] realtime                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Realtime int32 `gorm:"column:realtime;type:int;default:0;" json:"realtime"`
}

var eqtimeTableInfo = &TableInfo{
	Name: "eqtime",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "minute",
			Comment: ``,
			Notes: `Warning table: eqtime does not have a primary key defined, setting col position 1 minute as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Minute",
			GoFieldType:        "int32",
			JSONFieldName:      "minute",
			ProtobufFieldName:  "minute",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "hour",
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
			GoFieldName:        "Hour",
			GoFieldType:        "int32",
			JSONFieldName:      "hour",
			ProtobufFieldName:  "hour",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "day",
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
			GoFieldName:        "Day",
			GoFieldType:        "int32",
			JSONFieldName:      "day",
			ProtobufFieldName:  "day",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "month",
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
			GoFieldName:        "Month",
			GoFieldType:        "int32",
			JSONFieldName:      "month",
			ProtobufFieldName:  "month",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "year",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Year",
			GoFieldType:        "int32",
			JSONFieldName:      "year",
			ProtobufFieldName:  "year",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "realtime",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Realtime",
			GoFieldType:        "int32",
			JSONFieldName:      "realtime",
			ProtobufFieldName:  "realtime",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Eqtime) TableName() string {
	return "eqtime"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Eqtime) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Eqtime) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Eqtime) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Eqtime) TableInfo() *TableInfo {
	return eqtimeTableInfo
}
