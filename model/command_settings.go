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


CREATE TABLE `command_settings` (
  `command` varchar(128) NOT NULL DEFAULT '',
  `access` int(11) NOT NULL DEFAULT 0,
  `aliases` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`command`),
  UNIQUE KEY `UK_command_settings_1` (`command`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "command": "YbgotuIidYRErkmdBMSaBylYZ",    "access": 23,    "aliases": "AvFfXsdHwYlvplfQcJEHSVwwL"}



*/

// CommandSettings struct is a row record of the command_settings table in the eqemu database
type CommandSettings struct {
	//[ 0] command                                        varchar(128)         null: false  primary: true   isArray: false  auto: false  col: varchar         len: 128     default: ['']
	Command string `gorm:"primary_key;column:command;type:varchar;size:128;default:'';" json:"command"`
	//[ 1] access                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Access int32 `gorm:"column:access;type:int;default:0;" json:"access"`
	//[ 2] aliases                                        varchar(256)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 256     default: ['']
	Aliases string `gorm:"column:aliases;type:varchar;size:256;default:'';" json:"aliases"`
}

var command_settingsTableInfo = &TableInfo{
	Name: "command_settings",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "command",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(128)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       128,
			GoFieldName:        "Command",
			GoFieldType:        "string",
			JSONFieldName:      "command",
			ProtobufFieldName:  "command",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "access",
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
			GoFieldName:        "Access",
			GoFieldType:        "int32",
			JSONFieldName:      "access",
			ProtobufFieldName:  "access",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "aliases",
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
			GoFieldName:        "Aliases",
			GoFieldType:        "string",
			JSONFieldName:      "aliases",
			ProtobufFieldName:  "aliases",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CommandSettings) TableName() string {
	return "command_settings"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CommandSettings) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CommandSettings) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CommandSettings) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CommandSettings) TableInfo() *TableInfo {
	return command_settingsTableInfo
}
