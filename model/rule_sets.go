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


CREATE TABLE `rule_sets` (
  `ruleset_id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`ruleset_id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "ruleset_id": 17,    "name": "dvMtHdSIUESdUSYEdnMevYpCl"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// RuleSets struct is a row record of the rule_sets table in the eqemu database
type RuleSets struct {
	//[ 0] ruleset_id                                     utinyint             null: false  primary: true   isArray: false  auto: true   col: utinyint        len: -1      default: []
	RulesetID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:ruleset_id;type:utinyint;" json:"ruleset_id"`
	//[ 1] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: ['']
	Name string `gorm:"column:name;type:varchar;size:255;default:'';" json:"name"`
}

var rule_setsTableInfo = &TableInfo{
	Name: "rule_sets",
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
			IsAutoIncrement:    true,
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
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *RuleSets) TableName() string {
	return "rule_sets"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RuleSets) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RuleSets) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RuleSets) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RuleSets) TableInfo() *TableInfo {
	return rule_setsTableInfo
}
