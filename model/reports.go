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


CREATE TABLE `reports` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) DEFAULT NULL,
  `reported` varchar(64) DEFAULT NULL,
  `reported_text` text DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "id": 13,    "name": "wdCLJajMyHdwhagrCKhyMDGpd",    "reported": "GTdmgPImkceFoaDSfpcjsGaBC",    "reported_text": "FQhRYQWWiyDjcbiGuIGkCEeSE"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// Reports struct is a row record of the reports table in the eqemu database
type Reports struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] name                                           varchar(64)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 64      default: [NULL]
	Name null.String `gorm:"column:name;type:varchar;size:64;" json:"name"`
	//[ 2] reported                                       varchar(64)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 64      default: [NULL]
	Reported null.String `gorm:"column:reported;type:varchar;size:64;" json:"reported"`
	//[ 3] reported_text                                  text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: [NULL]
	ReportedText null.String `gorm:"column:reported_text;type:text;size:65535;" json:"reported_text"`
}

var reportsTableInfo = &TableInfo{
	Name: "reports",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
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
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Name",
			GoFieldType:        "null.String",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "reported",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Reported",
			GoFieldType:        "null.String",
			JSONFieldName:      "reported",
			ProtobufFieldName:  "reported",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "reported_text",
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
			GoFieldName:        "ReportedText",
			GoFieldType:        "null.String",
			JSONFieldName:      "reported_text",
			ProtobufFieldName:  "reported_text",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *Reports) TableName() string {
	return "reports"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *Reports) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *Reports) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *Reports) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *Reports) TableInfo() *TableInfo {
	return reportsTableInfo
}
