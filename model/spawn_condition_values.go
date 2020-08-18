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


CREATE TABLE `spawn_condition_values` (
  `id` int(10) unsigned NOT NULL,
  `value` tinyint(3) unsigned DEFAULT NULL,
  `zone` varchar(64) NOT NULL,
  `instance_id` int(10) unsigned NOT NULL,
  UNIQUE KEY `instance` (`id`,`instance_id`,`zone`),
  KEY `zoneinstance` (`zone`,`instance_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "zone": "BqUJhogqmlirUlAMPbnSoTrri",    "instance_id": 8,    "id": 46,    "value": 88}


Comments
-------------------------------------
[ 0] column is set for unsignedWarning table: spawn_condition_values does not have a primary key defined, setting col position 1 id as primary key

[ 1] column is set for unsigned
[ 3] column is set for unsigned



*/

// SpawnConditionValues struct is a row record of the spawn_condition_values table in the eqemu database
type SpawnConditionValues struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;column:id;type:uint;" json:"id"`
	//[ 1] value                                          utinyint             null: true   primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [NULL]
	Value null.Int `gorm:"column:value;type:utinyint;" json:"value"`
	//[ 2] zone                                           varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Zone string `gorm:"column:zone;type:varchar;size:64;" json:"zone"`
	//[ 3] instance_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	InstanceID uint32 `gorm:"column:instance_id;type:uint;" json:"instance_id"`
}

var spawn_condition_valuesTableInfo = &TableInfo{
	Name: "spawn_condition_values",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: ``,
			Notes: `column is set for unsignedWarning table: spawn_condition_values does not have a primary key defined, setting col position 1 id as primary key
`,
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
			Name:               "value",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Value",
			GoFieldType:        "null.Int",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "zone",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Zone",
			GoFieldType:        "string",
			JSONFieldName:      "zone",
			ProtobufFieldName:  "zone",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "instance_id",
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
			GoFieldName:        "InstanceID",
			GoFieldType:        "uint32",
			JSONFieldName:      "instance_id",
			ProtobufFieldName:  "instance_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SpawnConditionValues) TableName() string {
	return "spawn_condition_values"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SpawnConditionValues) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SpawnConditionValues) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SpawnConditionValues) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SpawnConditionValues) TableInfo() *TableInfo {
	return spawn_condition_valuesTableInfo
}
