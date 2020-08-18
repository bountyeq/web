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


CREATE TABLE `data_buckets` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(100) DEFAULT NULL,
  `value` text DEFAULT NULL,
  `expires` int(11) unsigned DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `key_index` (`key`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

JSON Sample
-------------------------------------
{    "id": 52,    "key": "IFcHEPfxKbADhlqulDnufKfCC",    "value": "ZoOpcfwkKeKBHYcRgfZkbFEGt",    "expires": 28}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 3] column is set for unsigned



*/

// DataBuckets struct is a row record of the data_buckets table in the eqemu database
type DataBuckets struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] key                                            varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	Key null.String `gorm:"column:key;type:varchar;size:100;" json:"key"`
	//[ 2] value                                          text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: [NULL]
	Value null.String `gorm:"column:value;type:text;size:65535;" json:"value"`
	//[ 3] expires                                        uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Expires null.Int `gorm:"column:expires;type:uint;default:0;" json:"expires"`
}

var data_bucketsTableInfo = &TableInfo{
	Name: "data_buckets",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
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

		&ColumnInfo{
			Index:              3,
			Name:               "expires",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Expires",
			GoFieldType:        "null.Int",
			JSONFieldName:      "expires",
			ProtobufFieldName:  "expires",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *DataBuckets) TableName() string {
	return "data_buckets"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DataBuckets) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DataBuckets) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DataBuckets) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *DataBuckets) TableInfo() *TableInfo {
	return data_bucketsTableInfo
}
