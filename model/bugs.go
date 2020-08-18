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


CREATE TABLE `bugs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `zone` varchar(32) NOT NULL,
  `name` varchar(64) NOT NULL,
  `ui` varchar(128) NOT NULL,
  `x` float NOT NULL DEFAULT 0,
  `y` float NOT NULL DEFAULT 0,
  `z` float NOT NULL DEFAULT 0,
  `type` varchar(64) NOT NULL,
  `flag` tinyint(3) unsigned NOT NULL,
  `target` varchar(64) DEFAULT NULL,
  `bug` varchar(1024) NOT NULL,
  `date` date NOT NULL,
  `status` tinyint(3) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "bug": "CUeWAYeWwXPXERgfTJavEkjYZ",    "date": "2166-06-24T01:06:42.562380451-07:00",    "id": 57,    "zone": "JrsvbFQMiSylFyEMacxAYIpCH",    "ui": "laptcJsFDOMNKoMrBsHleOXgm",    "y": 0.77338654,    "type": "ThyOQQUnSubaJOgFAbTqXaBqr",    "target": "IODsYVVYPgFLxCawYLYEKJwME",    "name": "fpicFrQoCgxXvpyLEDRXDyxoa",    "x": 0.43788415,    "z": 0.73830575,    "flag": 94,    "status": 93}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 8] column is set for unsigned
[12] column is set for unsigned



*/

// Bugs struct is a row record of the bugs table in the eqemu database
type Bugs struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] zone                                           varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	Zone string `gorm:"column:zone;type:varchar;size:32;" json:"zone"`
	//[ 2] name                                           varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Name string `gorm:"column:name;type:varchar;size:64;" json:"name"`
	//[ 3] ui                                             varchar(128)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 128     default: []
	UI string `gorm:"column:ui;type:varchar;size:128;" json:"ui"`
	//[ 4] x                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	X float32 `gorm:"column:x;type:float;default:0;" json:"x"`
	//[ 5] y                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Y float32 `gorm:"column:y;type:float;default:0;" json:"y"`
	//[ 6] z                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Z float32 `gorm:"column:z;type:float;default:0;" json:"z"`
	//[ 7] type                                           varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Type string `gorm:"column:type;type:varchar;size:64;" json:"type"`
	//[ 8] flag                                           utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: []
	Flag uint32 `gorm:"column:flag;type:utinyint;" json:"flag"`
	//[ 9] target                                         varchar(64)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 64      default: [NULL]
	Target null.String `gorm:"column:target;type:varchar;size:64;" json:"target"`
	//[10] bug                                            varchar(1024)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1024    default: []
	Bug string `gorm:"column:bug;type:varchar;size:1024;" json:"bug"`
	//[11] date                                           date                 null: false  primary: false  isArray: false  auto: false  col: date            len: -1      default: []
	Date time.Time `gorm:"column:date;type:date;" json:"date"`
	//[12] status                                         utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Status uint32 `gorm:"column:status;type:utinyint;default:0;" json:"status"`
}

var bugsTableInfo = &TableInfo{
	Name: "bugs",
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
			Name:               "zone",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Zone",
			GoFieldType:        "string",
			JSONFieldName:      "zone",
			ProtobufFieldName:  "zone",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "ui",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       128,
			GoFieldName:        "UI",
			GoFieldType:        "string",
			JSONFieldName:      "ui",
			ProtobufFieldName:  "ui",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "x",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "X",
			GoFieldType:        "float32",
			JSONFieldName:      "x",
			ProtobufFieldName:  "x",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "y",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "Y",
			GoFieldType:        "float32",
			JSONFieldName:      "y",
			ProtobufFieldName:  "y",
			ProtobufType:       "float",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "z",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "Z",
			GoFieldType:        "float32",
			JSONFieldName:      "z",
			ProtobufFieldName:  "z",
			ProtobufType:       "float",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Type",
			GoFieldType:        "string",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "flag",
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
			GoFieldName:        "Flag",
			GoFieldType:        "uint32",
			JSONFieldName:      "flag",
			ProtobufFieldName:  "flag",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "target",
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
			GoFieldName:        "Target",
			GoFieldType:        "null.String",
			JSONFieldName:      "target",
			ProtobufFieldName:  "target",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "bug",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       1024,
			GoFieldName:        "Bug",
			GoFieldType:        "string",
			JSONFieldName:      "bug",
			ProtobufFieldName:  "bug",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "date",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "date",
			DatabaseTypePretty: "date",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "date",
			ColumnLength:       -1,
			GoFieldName:        "Date",
			GoFieldType:        "time.Time",
			JSONFieldName:      "date",
			ProtobufFieldName:  "date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "status",
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
			GoFieldName:        "Status",
			GoFieldType:        "uint32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *Bugs) TableName() string {
	return "bugs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *Bugs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *Bugs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *Bugs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *Bugs) TableInfo() *TableInfo {
	return bugsTableInfo
}
