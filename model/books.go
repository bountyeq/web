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


CREATE TABLE `books` (
  `name` varchar(30) NOT NULL DEFAULT '',
  `txtfile` text NOT NULL,
  `language` int(11) NOT NULL DEFAULT 0,
  UNIQUE KEY `id` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "language": 90,    "name": "VmbPdYaVmEcNnteDFknRIahqj",    "txtfile": "aXVivSLhfincGJkNdgfwNHgeG"}


Comments
-------------------------------------
[ 0] Warning table: books does not have a primary key defined, setting col position 1 name as primary key




*/

// Books struct is a row record of the books table in the eqemu database
type Books struct {
	//[ 0] name                                           varchar(30)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 30      default: ['']
	Name string `gorm:"primary_key;column:name;type:varchar;size:30;default:'';" json:"name"`
	//[ 1] txtfile                                        text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Txtfile string `gorm:"column:txtfile;type:text;size:65535;" json:"txtfile"`
	//[ 2] language                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Language int32 `gorm:"column:language;type:int;default:0;" json:"language"`
}

var booksTableInfo = &TableInfo{
	Name: "books",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "name",
			Comment: ``,
			Notes: `Warning table: books does not have a primary key defined, setting col position 1 name as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "txtfile",
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
			GoFieldName:        "Txtfile",
			GoFieldType:        "string",
			JSONFieldName:      "txtfile",
			ProtobufFieldName:  "txtfile",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "language",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Language",
			GoFieldType:        "int32",
			JSONFieldName:      "language",
			ProtobufFieldName:  "language",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *Books) TableName() string {
	return "books"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *Books) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *Books) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *Books) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *Books) TableInfo() *TableInfo {
	return booksTableInfo
}
