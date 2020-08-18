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


CREATE TABLE `npc_spells_effects` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` tinytext DEFAULT NULL,
  `parent_list` int(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 73,    "name": "eLETBwnUJeefySHitWNPNwxLK",    "parent_list": 36}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned



*/

// NpcSpellsEffects struct is a row record of the npc_spells_effects table in the eqemu database
type NpcSpellsEffects struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] name                                           text(255)            null: true   primary: false  isArray: false  auto: false  col: text            len: 255     default: [NULL]
	Name null.String `gorm:"column:name;type:text;size:255;" json:"name"`
	//[ 2] parent_list                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ParentList uint32 `gorm:"column:parent_list;type:uint;default:0;" json:"parent_list"`
}

var npc_spells_effectsTableInfo = &TableInfo{
	Name: "npc_spells_effects",
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
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "null.String",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "parent_list",
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
			GoFieldName:        "ParentList",
			GoFieldType:        "uint32",
			JSONFieldName:      "parent_list",
			ProtobufFieldName:  "parent_list",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (n *NpcSpellsEffects) TableName() string {
	return "npc_spells_effects"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (n *NpcSpellsEffects) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (n *NpcSpellsEffects) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (n *NpcSpellsEffects) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (n *NpcSpellsEffects) TableInfo() *TableInfo {
	return npc_spells_effectsTableInfo
}
