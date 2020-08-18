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


CREATE TABLE `npc_faction_entries` (
  `npc_faction_id` int(11) unsigned NOT NULL DEFAULT 0,
  `faction_id` int(11) unsigned NOT NULL DEFAULT 0,
  `value` int(11) NOT NULL DEFAULT 0,
  `npc_value` tinyint(3) NOT NULL DEFAULT 0,
  `temp` tinyint(3) NOT NULL DEFAULT 0,
  PRIMARY KEY (`npc_faction_id`,`faction_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "faction_id": 46,    "value": 4,    "npc_value": 14,    "temp": 32,    "npc_faction_id": 0}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// NpcFactionEntries struct is a row record of the npc_faction_entries table in the eqemu database
type NpcFactionEntries struct {
	//[ 0] npc_faction_id                                 uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	NpcFactionID uint32 `gorm:"primary_key;column:npc_faction_id;type:uint;default:0;" json:"npc_faction_id"`
	//[ 1] faction_id                                     uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	FactionID uint32 `gorm:"primary_key;column:faction_id;type:uint;default:0;" json:"faction_id"`
	//[ 2] value                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Value int32 `gorm:"column:value;type:int;default:0;" json:"value"`
	//[ 3] npc_value                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	NpcValue int32 `gorm:"column:npc_value;type:tinyint;default:0;" json:"npc_value"`
	//[ 4] temp                                           tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Temp int32 `gorm:"column:temp;type:tinyint;default:0;" json:"temp"`
}

var npc_faction_entriesTableInfo = &TableInfo{
	Name: "npc_faction_entries",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "npc_faction_id",
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
			GoFieldName:        "NpcFactionID",
			GoFieldType:        "uint32",
			JSONFieldName:      "npc_faction_id",
			ProtobufFieldName:  "npc_faction_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "faction_id",
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
			GoFieldName:        "FactionID",
			GoFieldType:        "uint32",
			JSONFieldName:      "faction_id",
			ProtobufFieldName:  "faction_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "value",
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
			GoFieldName:        "Value",
			GoFieldType:        "int32",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "npc_value",
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
			GoFieldName:        "NpcValue",
			GoFieldType:        "int32",
			JSONFieldName:      "npc_value",
			ProtobufFieldName:  "npc_value",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "temp",
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
			GoFieldName:        "Temp",
			GoFieldType:        "int32",
			JSONFieldName:      "temp",
			ProtobufFieldName:  "temp",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (n *NpcFactionEntries) TableName() string {
	return "npc_faction_entries"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (n *NpcFactionEntries) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (n *NpcFactionEntries) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (n *NpcFactionEntries) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (n *NpcFactionEntries) TableInfo() *TableInfo {
	return npc_faction_entriesTableInfo
}
