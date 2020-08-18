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


CREATE TABLE `lfguild` (
  `type` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `name` varchar(32) NOT NULL,
  `comment` varchar(256) NOT NULL,
  `fromlevel` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `tolevel` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `classes` int(10) unsigned NOT NULL DEFAULT 0,
  `aacount` int(10) unsigned NOT NULL DEFAULT 0,
  `timezone` int(10) unsigned NOT NULL DEFAULT 0,
  `timeposted` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`type`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "aacount": 81,    "timeposted": 5,    "type": 21,    "name": "YLylWPLDMtUDxWdYwAJHdkORO",    "comment": "ViFQNoKSLYwYMagAhqtIltBWx",    "fromlevel": 21,    "tolevel": 52,    "classes": 74,    "timezone": 76}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned



*/

// Lfguild struct is a row record of the lfguild table in the eqemu database
type Lfguild struct {
	//[ 0] type                                           utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Type uint32 `gorm:"primary_key;column:type;type:utinyint;default:0;" json:"type"`
	//[ 1] name                                           varchar(32)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 32      default: []
	Name string `gorm:"primary_key;column:name;type:varchar;size:32;" json:"name"`
	//[ 2] comment                                        varchar(256)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	Comment string `gorm:"column:comment;type:varchar;size:256;" json:"comment"`
	//[ 3] fromlevel                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Fromlevel uint32 `gorm:"column:fromlevel;type:utinyint;default:0;" json:"fromlevel"`
	//[ 4] tolevel                                        utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Tolevel uint32 `gorm:"column:tolevel;type:utinyint;default:0;" json:"tolevel"`
	//[ 5] classes                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Classes uint32 `gorm:"column:classes;type:uint;default:0;" json:"classes"`
	//[ 6] aacount                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Aacount uint32 `gorm:"column:aacount;type:uint;default:0;" json:"aacount"`
	//[ 7] timezone                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Timezone uint32 `gorm:"column:timezone;type:uint;default:0;" json:"timezone"`
	//[ 8] timeposted                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Timeposted uint32 `gorm:"column:timeposted;type:uint;default:0;" json:"timeposted"`
}

var lfguildTableInfo = &TableInfo{
	Name: "lfguild",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "uint32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
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
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "comment",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "Comment",
			GoFieldType:        "string",
			JSONFieldName:      "comment",
			ProtobufFieldName:  "comment",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "fromlevel",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Fromlevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "fromlevel",
			ProtobufFieldName:  "fromlevel",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "tolevel",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Tolevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "tolevel",
			ProtobufFieldName:  "tolevel",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "classes",
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
			GoFieldName:        "Classes",
			GoFieldType:        "uint32",
			JSONFieldName:      "classes",
			ProtobufFieldName:  "classes",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "aacount",
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
			GoFieldName:        "Aacount",
			GoFieldType:        "uint32",
			JSONFieldName:      "aacount",
			ProtobufFieldName:  "aacount",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "timezone",
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
			GoFieldName:        "Timezone",
			GoFieldType:        "uint32",
			JSONFieldName:      "timezone",
			ProtobufFieldName:  "timezone",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "timeposted",
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
			GoFieldName:        "Timeposted",
			GoFieldType:        "uint32",
			JSONFieldName:      "timeposted",
			ProtobufFieldName:  "timeposted",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *Lfguild) TableName() string {
	return "lfguild"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *Lfguild) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *Lfguild) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *Lfguild) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *Lfguild) TableInfo() *TableInfo {
	return lfguildTableInfo
}
