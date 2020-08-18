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


CREATE TABLE `variables` (
  `varname` varchar(25) NOT NULL DEFAULT '',
  `value` text NOT NULL,
  `information` text NOT NULL,
  `ts` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`varname`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "varname": "ANSltBOkAgmZNpTlSCfZSjViZ",    "value": "FxEqJSCBMuXHCjwFIUYpJNWUR",    "information": "jtHhqQqqAluHCmCLNFyjLQnHy",    "ts": "2059-08-28T17:10:46.224128235-07:00"}



*/

// Variables struct is a row record of the variables table in the eqemu database
type Variables struct {
	//[ 0] varname                                        varchar(25)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 25      default: ['']
	Varname string `gorm:"primary_key;column:varname;type:varchar;size:25;default:'';" json:"varname"`
	//[ 1] value                                          text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Value string `gorm:"column:value;type:text;size:65535;" json:"value"`
	//[ 2] information                                    text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Information string `gorm:"column:information;type:text;size:65535;" json:"information"`
	//[ 3] ts                                             timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [current_timestamp()]
	Ts time.Time `gorm:"column:ts;type:timestamp;" json:"ts"`
}

var variablesTableInfo = &TableInfo{
	Name: "variables",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "varname",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(25)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       25,
			GoFieldName:        "Varname",
			GoFieldType:        "string",
			JSONFieldName:      "varname",
			ProtobufFieldName:  "varname",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Value",
			GoFieldType:        "string",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "information",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Information",
			GoFieldType:        "string",
			JSONFieldName:      "information",
			ProtobufFieldName:  "information",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "ts",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "Ts",
			GoFieldType:        "time.Time",
			JSONFieldName:      "ts",
			ProtobufFieldName:  "ts",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (v *Variables) TableName() string {
	return "variables"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (v *Variables) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (v *Variables) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (v *Variables) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (v *Variables) TableInfo() *TableInfo {
	return variablesTableInfo
}
