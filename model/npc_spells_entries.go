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


CREATE TABLE `npc_spells_entries` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `npc_spells_id` int(11) NOT NULL DEFAULT 0,
  `spellid` smallint(5) NOT NULL DEFAULT 0,
  `type` int(10) unsigned NOT NULL DEFAULT 0,
  `minlevel` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `maxlevel` tinyint(3) unsigned NOT NULL DEFAULT 255,
  `manacost` smallint(5) NOT NULL DEFAULT -1,
  `recast_delay` int(11) NOT NULL DEFAULT -1,
  `priority` smallint(5) NOT NULL DEFAULT 0,
  `resist_adjust` int(11) DEFAULT NULL,
  `min_hp` smallint(5) DEFAULT 0,
  `max_hp` smallint(5) DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `spellsid_spellid` (`npc_spells_id`,`spellid`)
) ENGINE=InnoDB AUTO_INCREMENT=21250 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "spellid": 63,    "type": 69,    "minlevel": 18,    "priority": 3,    "min_hp": 10,    "max_hp": 93,    "npc_spells_id": 28,    "maxlevel": 17,    "manacost": 13,    "recast_delay": 89,    "resist_adjust": 56,    "id": 10}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned



*/

// NpcSpellsEntries struct is a row record of the npc_spells_entries table in the eqemu database
type NpcSpellsEntries struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] npc_spells_id                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	NpcSpellsID int32 `gorm:"column:npc_spells_id;type:int;default:0;" json:"npc_spells_id"`
	//[ 2] spellid                                        smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Spellid int32 `gorm:"column:spellid;type:smallint;default:0;" json:"spellid"`
	//[ 3] type                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Type uint32 `gorm:"column:type;type:uint;default:0;" json:"type"`
	//[ 4] minlevel                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Minlevel uint32 `gorm:"column:minlevel;type:utinyint;default:0;" json:"minlevel"`
	//[ 5] maxlevel                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [255]
	Maxlevel uint32 `gorm:"column:maxlevel;type:utinyint;default:255;" json:"maxlevel"`
	//[ 6] manacost                                       smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [-1]
	Manacost int32 `gorm:"column:manacost;type:smallint;default:-1;" json:"manacost"`
	//[ 7] recast_delay                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	RecastDelay int32 `gorm:"column:recast_delay;type:int;default:-1;" json:"recast_delay"`
	//[ 8] priority                                       smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Priority int32 `gorm:"column:priority;type:smallint;default:0;" json:"priority"`
	//[ 9] resist_adjust                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	ResistAdjust null.Int `gorm:"column:resist_adjust;type:int;" json:"resist_adjust"`
	//[10] min_hp                                         smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	MinHp null.Int `gorm:"column:min_hp;type:smallint;default:0;" json:"min_hp"`
	//[11] max_hp                                         smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	MaxHp null.Int `gorm:"column:max_hp;type:smallint;default:0;" json:"max_hp"`
}

var npc_spells_entriesTableInfo = &TableInfo{
	Name: "npc_spells_entries",
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
			Name:               "npc_spells_id",
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
			GoFieldName:        "NpcSpellsID",
			GoFieldType:        "int32",
			JSONFieldName:      "npc_spells_id",
			ProtobufFieldName:  "npc_spells_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "spellid",
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
			GoFieldName:        "Spellid",
			GoFieldType:        "int32",
			JSONFieldName:      "spellid",
			ProtobufFieldName:  "spellid",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "uint32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "minlevel",
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
			GoFieldName:        "Minlevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "minlevel",
			ProtobufFieldName:  "minlevel",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "maxlevel",
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
			GoFieldName:        "Maxlevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "maxlevel",
			ProtobufFieldName:  "maxlevel",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "manacost",
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
			GoFieldName:        "Manacost",
			GoFieldType:        "int32",
			JSONFieldName:      "manacost",
			ProtobufFieldName:  "manacost",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "recast_delay",
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
			GoFieldName:        "RecastDelay",
			GoFieldType:        "int32",
			JSONFieldName:      "recast_delay",
			ProtobufFieldName:  "recast_delay",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "priority",
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
			GoFieldName:        "Priority",
			GoFieldType:        "int32",
			JSONFieldName:      "priority",
			ProtobufFieldName:  "priority",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "resist_adjust",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ResistAdjust",
			GoFieldType:        "null.Int",
			JSONFieldName:      "resist_adjust",
			ProtobufFieldName:  "resist_adjust",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "min_hp",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "MinHp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "min_hp",
			ProtobufFieldName:  "min_hp",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "max_hp",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "MaxHp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "max_hp",
			ProtobufFieldName:  "max_hp",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (n *NpcSpellsEntries) TableName() string {
	return "npc_spells_entries"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (n *NpcSpellsEntries) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (n *NpcSpellsEntries) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (n *NpcSpellsEntries) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (n *NpcSpellsEntries) TableInfo() *TableInfo {
	return npc_spells_entriesTableInfo
}
