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


CREATE TABLE `tributes` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `unknown` int(10) unsigned NOT NULL DEFAULT 0,
  `name` varchar(255) NOT NULL DEFAULT '',
  `descr` mediumtext NOT NULL,
  `isguild` tinyint(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`isguild`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "name": "gsbEIURoWocBqXUEOdNZevLKn",    "descr": "YvJtdqxHfdDAuTPXdUWOgXPQV",    "isguild": 7,    "id": 29,    "unknown": 92}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// Tributes struct is a row record of the tributes table in the eqemu database
type Tributes struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ID uint32 `gorm:"primary_key;column:id;type:uint;default:0;" json:"id"`
	//[ 1] unknown                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Unknown uint32 `gorm:"column:unknown;type:uint;default:0;" json:"unknown"`
	//[ 2] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: ['']
	Name string `gorm:"column:name;type:varchar;size:255;default:'';" json:"name"`
	//[ 3] descr                                          text(16777215)       null: false  primary: false  isArray: false  auto: false  col: text            len: 16777215 default: []
	Descr string `gorm:"column:descr;type:text;size:16777215;" json:"descr"`
	//[ 4] isguild                                        tinyint              null: false  primary: true   isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Isguild int32 `gorm:"primary_key;column:isguild;type:tinyint;default:0;" json:"isguild"`
}

var tributesTableInfo = &TableInfo{
	Name: "tributes",
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
			Name:               "unknown",
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
			GoFieldName:        "Unknown",
			GoFieldType:        "uint32",
			JSONFieldName:      "unknown",
			ProtobufFieldName:  "unknown",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "descr",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(16777215)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       16777215,
			GoFieldName:        "Descr",
			GoFieldType:        "string",
			JSONFieldName:      "descr",
			ProtobufFieldName:  "descr",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "isguild",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Isguild",
			GoFieldType:        "int32",
			JSONFieldName:      "isguild",
			ProtobufFieldName:  "isguild",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *Tributes) TableName() string {
	return "tributes"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *Tributes) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *Tributes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *Tributes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *Tributes) TableInfo() *TableInfo {
	return tributesTableInfo
}
