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


CREATE TABLE `npc_spells_effects_entries` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `npc_spells_effects_id` int(11) NOT NULL DEFAULT 0,
  `spell_effect_id` smallint(5) NOT NULL DEFAULT 0,
  `minlevel` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `maxlevel` tinyint(3) unsigned NOT NULL DEFAULT 255,
  `se_base` int(11) NOT NULL DEFAULT 0,
  `se_limit` int(11) NOT NULL DEFAULT 0,
  `se_max` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `spellsid_spellid` (`npc_spells_effects_id`,`spell_effect_id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "se_limit": 56,    "se_max": 86,    "id": 49,    "npc_spells_effects_id": 54,    "spell_effect_id": 12,    "minlevel": 85,    "maxlevel": 8,    "se_base": 36}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned



*/

// NpcSpellsEffectsEntries struct is a row record of the npc_spells_effects_entries table in the eqemu database
type NpcSpellsEffectsEntries struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] npc_spells_effects_id                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	NpcSpellsEffectsID int32 `gorm:"column:npc_spells_effects_id;type:int;default:0;" json:"npc_spells_effects_id"`
	//[ 2] spell_effect_id                                smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	SpellEffectID int32 `gorm:"column:spell_effect_id;type:smallint;default:0;" json:"spell_effect_id"`
	//[ 3] minlevel                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Minlevel uint32 `gorm:"column:minlevel;type:utinyint;default:0;" json:"minlevel"`
	//[ 4] maxlevel                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [255]
	Maxlevel uint32 `gorm:"column:maxlevel;type:utinyint;default:255;" json:"maxlevel"`
	//[ 5] se_base                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SeBase int32 `gorm:"column:se_base;type:int;default:0;" json:"se_base"`
	//[ 6] se_limit                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SeLimit int32 `gorm:"column:se_limit;type:int;default:0;" json:"se_limit"`
	//[ 7] se_max                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SeMax int32 `gorm:"column:se_max;type:int;default:0;" json:"se_max"`
}

var npc_spells_effects_entriesTableInfo = &TableInfo{
	Name: "npc_spells_effects_entries",
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
			Name:               "npc_spells_effects_id",
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
			GoFieldName:        "NpcSpellsEffectsID",
			GoFieldType:        "int32",
			JSONFieldName:      "npc_spells_effects_id",
			ProtobufFieldName:  "npc_spells_effects_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "spell_effect_id",
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
			GoFieldName:        "SpellEffectID",
			GoFieldType:        "int32",
			JSONFieldName:      "spell_effect_id",
			ProtobufFieldName:  "spell_effect_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "se_base",
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
			GoFieldName:        "SeBase",
			GoFieldType:        "int32",
			JSONFieldName:      "se_base",
			ProtobufFieldName:  "se_base",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "se_limit",
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
			GoFieldName:        "SeLimit",
			GoFieldType:        "int32",
			JSONFieldName:      "se_limit",
			ProtobufFieldName:  "se_limit",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "se_max",
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
			GoFieldName:        "SeMax",
			GoFieldType:        "int32",
			JSONFieldName:      "se_max",
			ProtobufFieldName:  "se_max",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (n *NpcSpellsEffectsEntries) TableName() string {
	return "npc_spells_effects_entries"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (n *NpcSpellsEffectsEntries) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (n *NpcSpellsEffectsEntries) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (n *NpcSpellsEffectsEntries) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (n *NpcSpellsEffectsEntries) TableInfo() *TableInfo {
	return npc_spells_effects_entriesTableInfo
}
