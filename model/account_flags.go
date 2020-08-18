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


CREATE TABLE `account_flags` (
  `p_accid` int(10) unsigned NOT NULL,
  `p_flag` varchar(50) NOT NULL,
  `p_value` varchar(80) NOT NULL,
  PRIMARY KEY (`p_accid`,`p_flag`),
  KEY `p_accid` (`p_accid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "p_accid": 30,    "p_flag": "VbSJaJWxmXyfXEIgOAVDGmpVe",    "p_value": "GABNHLveeQKxIQqtBNdsFhXQq"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// AccountFlags struct is a row record of the account_flags table in the eqemu database
type AccountFlags struct {
	//[ 0] p_accid                                        uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	PAccid uint32 `gorm:"primary_key;column:p_accid;type:uint;" json:"p_accid"`
	//[ 1] p_flag                                         varchar(50)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 50      default: []
	PFlag string `gorm:"primary_key;column:p_flag;type:varchar;size:50;" json:"p_flag"`
	//[ 2] p_value                                        varchar(80)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 80      default: []
	PValue string `gorm:"column:p_value;type:varchar;size:80;" json:"p_value"`
}

var account_flagsTableInfo = &TableInfo{
	Name: "account_flags",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "p_accid",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PAccid",
			GoFieldType:        "uint32",
			JSONFieldName:      "p_accid",
			ProtobufFieldName:  "p_accid",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "p_flag",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "PFlag",
			GoFieldType:        "string",
			JSONFieldName:      "p_flag",
			ProtobufFieldName:  "p_flag",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "p_value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(80)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       80,
			GoFieldName:        "PValue",
			GoFieldType:        "string",
			JSONFieldName:      "p_value",
			ProtobufFieldName:  "p_value",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AccountFlags) TableName() string {
	return "account_flags"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AccountFlags) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AccountFlags) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AccountFlags) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AccountFlags) TableInfo() *TableInfo {
	return account_flagsTableInfo
}
