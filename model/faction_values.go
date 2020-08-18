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


CREATE TABLE `faction_values` (
  `char_id` int(4) NOT NULL DEFAULT 0,
  `faction_id` int(4) NOT NULL DEFAULT 0,
  `current_value` smallint(6) NOT NULL DEFAULT 0,
  `temp` tinyint(3) NOT NULL DEFAULT 0,
  PRIMARY KEY (`char_id`,`faction_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "current_value": 41,    "temp": 7,    "char_id": 96,    "faction_id": 94}



*/

// FactionValues struct is a row record of the faction_values table in the eqemu database
type FactionValues struct {
	//[ 0] char_id                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	CharID int32 `gorm:"primary_key;column:char_id;type:int;default:0;" json:"char_id"`
	//[ 1] faction_id                                     int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	FactionID int32 `gorm:"primary_key;column:faction_id;type:int;default:0;" json:"faction_id"`
	//[ 2] current_value                                  smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	CurrentValue int32 `gorm:"column:current_value;type:smallint;default:0;" json:"current_value"`
	//[ 3] temp                                           tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Temp int32 `gorm:"column:temp;type:tinyint;default:0;" json:"temp"`
}

var faction_valuesTableInfo = &TableInfo{
	Name: "faction_values",
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
			Name:               "faction_id",
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
			GoFieldName:        "FactionID",
			GoFieldType:        "int32",
			JSONFieldName:      "faction_id",
			ProtobufFieldName:  "faction_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "current_value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "CurrentValue",
			GoFieldType:        "int32",
			JSONFieldName:      "current_value",
			ProtobufFieldName:  "current_value",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "temp",
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
			GoFieldName:        "Temp",
			GoFieldType:        "int32",
			JSONFieldName:      "temp",
			ProtobufFieldName:  "temp",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *FactionValues) TableName() string {
	return "faction_values"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FactionValues) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FactionValues) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FactionValues) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FactionValues) TableInfo() *TableInfo {
	return faction_valuesTableInfo
}
