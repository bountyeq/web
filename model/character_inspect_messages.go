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


CREATE TABLE `character_inspect_messages` (
  `id` int(11) unsigned NOT NULL DEFAULT 0,
  `inspect_message` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 66,    "inspect_message": "jsklByDCMkiCdkoYZNPExbIir"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// CharacterInspectMessages struct is a row record of the character_inspect_messages table in the eqemu database
type CharacterInspectMessages struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ID uint32 `gorm:"primary_key;column:id;type:uint;default:0;" json:"id"`
	//[ 1] inspect_message                                varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: ['']
	InspectMessage string `gorm:"column:inspect_message;type:varchar;size:255;default:'';" json:"inspect_message"`
}

var character_inspect_messagesTableInfo = &TableInfo{
	Name: "character_inspect_messages",
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
			Name:               "inspect_message",
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
			GoFieldName:        "InspectMessage",
			GoFieldType:        "string",
			JSONFieldName:      "inspect_message",
			ProtobufFieldName:  "inspect_message",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterInspectMessages) TableName() string {
	return "character_inspect_messages"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterInspectMessages) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterInspectMessages) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterInspectMessages) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterInspectMessages) TableInfo() *TableInfo {
	return character_inspect_messagesTableInfo
}
