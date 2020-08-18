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


CREATE TABLE `npc_emotes` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `emoteid` int(10) unsigned NOT NULL DEFAULT 0,
  `event_` tinyint(3) NOT NULL DEFAULT 0,
  `type` tinyint(3) NOT NULL DEFAULT 0,
  `text` varchar(512) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `emoteid` (`emoteid`,`event_`)
) ENGINE=InnoDB AUTO_INCREMENT=2561 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 62,    "emoteid": 46,    "event_": 22,    "type": 15,    "text": "DPvVmGxCuqjsthHEnBoXHsHkg"}


Comments
-------------------------------------
[ 1] column is set for unsigned



*/

// NpcEmotes struct is a row record of the npc_emotes table in the eqemu database
type NpcEmotes struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] emoteid                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Emoteid uint32 `gorm:"column:emoteid;type:uint;default:0;" json:"emoteid"`
	//[ 2] event_                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Event int32 `gorm:"column:event_;type:tinyint;default:0;" json:"event_"`
	//[ 3] type                                           tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Type int32 `gorm:"column:type;type:tinyint;default:0;" json:"type"`
	//[ 4] text                                           varchar(512)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 512     default: []
	Text string `gorm:"column:text;type:varchar;size:512;" json:"text"`
}

var npc_emotesTableInfo = &TableInfo{
	Name: "npc_emotes",
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
			Name:               "emoteid",
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
			GoFieldName:        "Emoteid",
			GoFieldType:        "uint32",
			JSONFieldName:      "emoteid",
			ProtobufFieldName:  "emoteid",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "event_",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Event",
			GoFieldType:        "int32",
			JSONFieldName:      "event_",
			ProtobufFieldName:  "event_",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "text",
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
			GoFieldName:        "Text",
			GoFieldType:        "string",
			JSONFieldName:      "text",
			ProtobufFieldName:  "text",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (n *NpcEmotes) TableName() string {
	return "npc_emotes"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (n *NpcEmotes) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (n *NpcEmotes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (n *NpcEmotes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (n *NpcEmotes) TableInfo() *TableInfo {
	return npc_emotesTableInfo
}
