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


CREATE TABLE `rule_values` (
  `ruleset_id` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `rule_name` varchar(64) NOT NULL DEFAULT '',
  `rule_value` varchar(30) NOT NULL DEFAULT '',
  `notes` text DEFAULT NULL,
  PRIMARY KEY (`ruleset_id`,`rule_name`),
  KEY `ruleset_id` (`ruleset_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "notes": "fXhdWwiqAhQBVIXyMYYXgsJyq",    "ruleset_id": 50,    "rule_name": "KGcHXiafPnpTFHbBPXMcQqLkJ",    "rule_value": "pDoNpgYwrfAWLvUXPAbXwLwYr"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// RuleValues struct is a row record of the rule_values table in the eqemu database
type RuleValues struct {
	//[ 0] ruleset_id                                     utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	RulesetID uint32 `gorm:"primary_key;column:ruleset_id;type:utinyint;default:0;" json:"ruleset_id"`
	//[ 1] rule_name                                      varchar(64)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 64      default: ['']
	RuleName string `gorm:"primary_key;column:rule_name;type:varchar;size:64;default:'';" json:"rule_name"`
	//[ 2] rule_value                                     varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: ['']
	RuleValue string `gorm:"column:rule_value;type:varchar;size:30;default:'';" json:"rule_value"`
	//[ 3] notes                                          text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: [NULL]
	Notes null.String `gorm:"column:notes;type:text;size:65535;" json:"notes"`
}

var rule_valuesTableInfo = &TableInfo{
	Name: "rule_values",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ruleset_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "RulesetID",
			GoFieldType:        "uint32",
			JSONFieldName:      "ruleset_id",
			ProtobufFieldName:  "ruleset_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "rule_name",
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
			GoFieldName:        "RuleName",
			GoFieldType:        "string",
			JSONFieldName:      "rule_name",
			ProtobufFieldName:  "rule_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "rule_value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "RuleValue",
			GoFieldType:        "string",
			JSONFieldName:      "rule_value",
			ProtobufFieldName:  "rule_value",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "notes",
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
			GoFieldName:        "Notes",
			GoFieldType:        "null.String",
			JSONFieldName:      "notes",
			ProtobufFieldName:  "notes",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *RuleValues) TableName() string {
	return "rule_values"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RuleValues) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RuleValues) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RuleValues) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RuleValues) TableInfo() *TableInfo {
	return rule_valuesTableInfo
}
