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


CREATE TABLE `spawn_conditions` (
  `zone` varchar(32) NOT NULL DEFAULT '',
  `id` mediumint(8) unsigned NOT NULL DEFAULT 1,
  `value` mediumint(9) NOT NULL DEFAULT 0,
  `onchange` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `name` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`zone`,`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "zone": "bIgWvZUoJUdLHFIWUymxNXFbE",    "id": 38,    "value": 3,    "onchange": 62,    "name": "XexiloBtyNFBjXgsvjEuvhVqo"}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 3] column is set for unsigned



*/

// SpawnConditions struct is a row record of the spawn_conditions table in the eqemu database
type SpawnConditions struct {
	//[ 0] zone                                           varchar(32)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 32      default: ['']
	Zone string `gorm:"primary_key;column:zone;type:varchar;size:32;default:'';" json:"zone"`
	//[ 1] id                                             umediumint           null: false  primary: true   isArray: false  auto: false  col: umediumint      len: -1      default: [1]
	ID uint32 `gorm:"primary_key;column:id;type:umediumint;default:1;" json:"id"`
	//[ 2] value                                          mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	Value int32 `gorm:"column:value;type:mediumint;default:0;" json:"value"`
	//[ 3] onchange                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Onchange uint32 `gorm:"column:onchange;type:utinyint;default:0;" json:"onchange"`
	//[ 4] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: ['']
	Name string `gorm:"column:name;type:varchar;size:255;default:'';" json:"name"`
}

var spawn_conditionsTableInfo = &TableInfo{
	Name: "spawn_conditions",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "zone",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Zone",
			GoFieldType:        "string",
			JSONFieldName:      "zone",
			ProtobufFieldName:  "zone",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "id",
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
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "value",
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
			GoFieldName:        "Value",
			GoFieldType:        "int32",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "onchange",
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
			GoFieldName:        "Onchange",
			GoFieldType:        "uint32",
			JSONFieldName:      "onchange",
			ProtobufFieldName:  "onchange",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SpawnConditions) TableName() string {
	return "spawn_conditions"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SpawnConditions) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SpawnConditions) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SpawnConditions) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SpawnConditions) TableInfo() *TableInfo {
	return spawn_conditionsTableInfo
}
