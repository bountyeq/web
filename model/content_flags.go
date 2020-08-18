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


CREATE TABLE `content_flags` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `flag_name` varchar(75) DEFAULT NULL,
  `enabled` tinyint(4) DEFAULT NULL,
  `notes` text DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 89,    "flag_name": "JGYkHfkNBDKBotpEcRBEySZkp",    "enabled": 16,    "notes": "ysGJWUCpkbAjPuGkYYHpEHcaJ"}



*/

// ContentFlags struct is a row record of the content_flags table in the eqemu database
type ContentFlags struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] flag_name                                      varchar(75)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 75      default: [NULL]
	FlagName null.String `gorm:"column:flag_name;type:varchar;size:75;" json:"flag_name"`
	//[ 2] enabled                                        tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [NULL]
	Enabled null.Int `gorm:"column:enabled;type:tinyint;" json:"enabled"`
	//[ 3] notes                                          text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: [NULL]
	Notes null.String `gorm:"column:notes;type:text;size:65535;" json:"notes"`
}

var content_flagsTableInfo = &TableInfo{
	Name: "content_flags",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "flag_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(75)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       75,
			GoFieldName:        "FlagName",
			GoFieldType:        "null.String",
			JSONFieldName:      "flag_name",
			ProtobufFieldName:  "flag_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "enabled",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Enabled",
			GoFieldType:        "null.Int",
			JSONFieldName:      "enabled",
			ProtobufFieldName:  "enabled",
			ProtobufType:       "int32",
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
func (c *ContentFlags) TableName() string {
	return "content_flags"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ContentFlags) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ContentFlags) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ContentFlags) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ContentFlags) TableInfo() *TableInfo {
	return content_flagsTableInfo
}
