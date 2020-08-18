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


CREATE TABLE `lootdrop_entries` (
  `lootdrop_id` int(11) unsigned NOT NULL DEFAULT 0,
  `item_id` int(11) NOT NULL DEFAULT 0,
  `item_charges` smallint(2) unsigned NOT NULL DEFAULT 1,
  `equip_item` tinyint(2) unsigned NOT NULL DEFAULT 0,
  `chance` float NOT NULL DEFAULT 1,
  `disabled_chance` float NOT NULL DEFAULT 0,
  `trivial_min_level` smallint(5) unsigned NOT NULL DEFAULT 0,
  `trivial_max_level` smallint(5) unsigned NOT NULL DEFAULT 0,
  `multiplier` tinyint(2) unsigned NOT NULL DEFAULT 1,
  `npc_min_level` smallint(5) unsigned NOT NULL DEFAULT 0,
  `npc_max_level` smallint(5) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`lootdrop_id`,`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "npc_min_level": 73,    "npc_max_level": 11,    "lootdrop_id": 19,    "item_id": 81,    "trivial_min_level": 88,    "trivial_max_level": 10,    "multiplier": 43,    "item_charges": 31,    "equip_item": 28,    "chance": 0.6139835,    "disabled_chance": 0.72154593}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned



*/

// LootdropEntries struct is a row record of the lootdrop_entries table in the eqemu database
type LootdropEntries struct {
	//[ 0] lootdrop_id                                    uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	LootdropID uint32 `gorm:"primary_key;column:lootdrop_id;type:uint;default:0;" json:"lootdrop_id"`
	//[ 1] item_id                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	ItemID int32 `gorm:"primary_key;column:item_id;type:int;default:0;" json:"item_id"`
	//[ 2] item_charges                                   usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [1]
	ItemCharges uint32 `gorm:"column:item_charges;type:usmallint;default:1;" json:"item_charges"`
	//[ 3] equip_item                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	EquipItem uint32 `gorm:"column:equip_item;type:utinyint;default:0;" json:"equip_item"`
	//[ 4] chance                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [1]
	Chance float32 `gorm:"column:chance;type:float;default:1;" json:"chance"`
	//[ 5] disabled_chance                                float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	DisabledChance float32 `gorm:"column:disabled_chance;type:float;default:0;" json:"disabled_chance"`
	//[ 6] trivial_min_level                              usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	TrivialMinLevel uint32 `gorm:"column:trivial_min_level;type:usmallint;default:0;" json:"trivial_min_level"`
	//[ 7] trivial_max_level                              usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	TrivialMaxLevel uint32 `gorm:"column:trivial_max_level;type:usmallint;default:0;" json:"trivial_max_level"`
	//[ 8] multiplier                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	Multiplier uint32 `gorm:"column:multiplier;type:utinyint;default:1;" json:"multiplier"`
	//[ 9] npc_min_level                                  usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	NpcMinLevel uint32 `gorm:"column:npc_min_level;type:usmallint;default:0;" json:"npc_min_level"`
	//[10] npc_max_level                                  usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	NpcMaxLevel uint32 `gorm:"column:npc_max_level;type:usmallint;default:0;" json:"npc_max_level"`
}

var lootdrop_entriesTableInfo = &TableInfo{
	Name: "lootdrop_entries",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "lootdrop_id",
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
			GoFieldName:        "LootdropID",
			GoFieldType:        "uint32",
			JSONFieldName:      "lootdrop_id",
			ProtobufFieldName:  "lootdrop_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "item_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ItemID",
			GoFieldType:        "int32",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "item_charges",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "ItemCharges",
			GoFieldType:        "uint32",
			JSONFieldName:      "item_charges",
			ProtobufFieldName:  "item_charges",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "equip_item",
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
			GoFieldName:        "EquipItem",
			GoFieldType:        "uint32",
			JSONFieldName:      "equip_item",
			ProtobufFieldName:  "equip_item",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "chance",
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
			GoFieldName:        "Chance",
			GoFieldType:        "float32",
			JSONFieldName:      "chance",
			ProtobufFieldName:  "chance",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "disabled_chance",
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
			GoFieldName:        "DisabledChance",
			GoFieldType:        "float32",
			JSONFieldName:      "disabled_chance",
			ProtobufFieldName:  "disabled_chance",
			ProtobufType:       "float",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "trivial_min_level",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "TrivialMinLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "trivial_min_level",
			ProtobufFieldName:  "trivial_min_level",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "trivial_max_level",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "TrivialMaxLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "trivial_max_level",
			ProtobufFieldName:  "trivial_max_level",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "multiplier",
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
			GoFieldName:        "Multiplier",
			GoFieldType:        "uint32",
			JSONFieldName:      "multiplier",
			ProtobufFieldName:  "multiplier",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "npc_min_level",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "NpcMinLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "npc_min_level",
			ProtobufFieldName:  "npc_min_level",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "npc_max_level",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "NpcMaxLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "npc_max_level",
			ProtobufFieldName:  "npc_max_level",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LootdropEntries) TableName() string {
	return "lootdrop_entries"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LootdropEntries) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LootdropEntries) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LootdropEntries) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LootdropEntries) TableInfo() *TableInfo {
	return lootdrop_entriesTableInfo
}
