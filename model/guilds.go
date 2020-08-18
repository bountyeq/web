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


CREATE TABLE `guilds` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `leader` int(11) NOT NULL DEFAULT 0,
  `minstatus` smallint(5) NOT NULL DEFAULT 0,
  `motd` text NOT NULL,
  `tribute` int(10) unsigned NOT NULL DEFAULT 0,
  `motd_setter` varchar(64) NOT NULL DEFAULT '',
  `channel` varchar(128) NOT NULL DEFAULT '',
  `url` varchar(512) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `leader` (`leader`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 45,    "name": "IgJiwcoSYZIEqJDXgJemURgSL",    "motd": "qUdsAoNKmajkwKwFZUZBbvJAR",    "tribute": 4,    "motd_setter": "GZsmUDbHIkyBrpnwMDWEpcvyS",    "channel": "JqFKYsVymslgLqXxvyUCHbBlj",    "url": "qrsujNkPZTUGpyQUUHvDgkDey",    "leader": 23,    "minstatus": 31}


Comments
-------------------------------------
[ 5] column is set for unsigned



*/

// Guilds struct is a row record of the guilds table in the eqemu database
type Guilds struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] name                                           varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['']
	Name string `gorm:"column:name;type:varchar;size:32;default:'';" json:"name"`
	//[ 2] leader                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Leader int32 `gorm:"column:leader;type:int;default:0;" json:"leader"`
	//[ 3] minstatus                                      smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Minstatus int32 `gorm:"column:minstatus;type:smallint;default:0;" json:"minstatus"`
	//[ 4] motd                                           text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Motd string `gorm:"column:motd;type:text;size:65535;" json:"motd"`
	//[ 5] tribute                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Tribute uint32 `gorm:"column:tribute;type:uint;default:0;" json:"tribute"`
	//[ 6] motd_setter                                    varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	MotdSetter string `gorm:"column:motd_setter;type:varchar;size:64;default:'';" json:"motd_setter"`
	//[ 7] channel                                        varchar(128)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 128     default: ['']
	Channel string `gorm:"column:channel;type:varchar;size:128;default:'';" json:"channel"`
	//[ 8] url                                            varchar(512)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 512     default: ['']
	URL string `gorm:"column:url;type:varchar;size:512;default:'';" json:"url"`
}

var guildsTableInfo = &TableInfo{
	Name: "guilds",
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
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "leader",
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
			GoFieldName:        "Leader",
			GoFieldType:        "int32",
			JSONFieldName:      "leader",
			ProtobufFieldName:  "leader",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "minstatus",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "Minstatus",
			GoFieldType:        "int32",
			JSONFieldName:      "minstatus",
			ProtobufFieldName:  "minstatus",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "motd",
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
			GoFieldName:        "Motd",
			GoFieldType:        "string",
			JSONFieldName:      "motd",
			ProtobufFieldName:  "motd",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "tribute",
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
			GoFieldName:        "Tribute",
			GoFieldType:        "uint32",
			JSONFieldName:      "tribute",
			ProtobufFieldName:  "tribute",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "motd_setter",
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
			GoFieldName:        "MotdSetter",
			GoFieldType:        "string",
			JSONFieldName:      "motd_setter",
			ProtobufFieldName:  "motd_setter",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "channel",
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
			GoFieldName:        "Channel",
			GoFieldType:        "string",
			JSONFieldName:      "channel",
			ProtobufFieldName:  "channel",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "url",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(512)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       512,
			GoFieldName:        "URL",
			GoFieldType:        "string",
			JSONFieldName:      "url",
			ProtobufFieldName:  "url",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *Guilds) TableName() string {
	return "guilds"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *Guilds) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *Guilds) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *Guilds) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *Guilds) TableInfo() *TableInfo {
	return guildsTableInfo
}
