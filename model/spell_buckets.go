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


CREATE TABLE `spell_buckets` (
  `spellid` bigint(11) unsigned NOT NULL,
  `key` varchar(100) DEFAULT NULL,
  `value` text DEFAULT NULL,
  PRIMARY KEY (`spellid`),
  KEY `key_index` (`key`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

JSON Sample
-------------------------------------
{    "spellid": 24,    "key": "jMgmiUiiWhRXSwdPkEHhCcedG",    "value": "bUvLwbCxFrpvTAbIvTfFEJmci"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// SpellBuckets struct is a row record of the spell_buckets table in the eqemu database
type SpellBuckets struct {
	//[ 0] spellid                                        ubigint              null: false  primary: true   isArray: false  auto: false  col: ubigint         len: -1      default: []
	Spellid uint64 `gorm:"primary_key;column:spellid;type:ubigint;" json:"spellid"`
	//[ 1] key                                            varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	Key null.String `gorm:"column:key;type:varchar;size:100;" json:"key"`
	//[ 2] value                                          text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: [NULL]
	Value null.String `gorm:"column:value;type:text;size:65535;" json:"value"`
}

var spell_bucketsTableInfo = &TableInfo{
	Name: "spell_buckets",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "spellid",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "Spellid",
			GoFieldType:        "uint64",
			JSONFieldName:      "spellid",
			ProtobufFieldName:  "spellid",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "key",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "Key",
			GoFieldType:        "null.String",
			JSONFieldName:      "key",
			ProtobufFieldName:  "key",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Value",
			GoFieldType:        "null.String",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SpellBuckets) TableName() string {
	return "spell_buckets"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SpellBuckets) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SpellBuckets) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SpellBuckets) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SpellBuckets) TableInfo() *TableInfo {
	return spell_bucketsTableInfo
}
