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


CREATE TABLE `character_buffs` (
  `character_id` int(10) unsigned NOT NULL,
  `slot_id` tinyint(3) unsigned NOT NULL,
  `spell_id` smallint(10) unsigned NOT NULL,
  `caster_level` tinyint(3) unsigned NOT NULL,
  `caster_name` varchar(64) NOT NULL,
  `ticsremaining` int(11) NOT NULL,
  `counters` int(10) unsigned NOT NULL,
  `numhits` int(10) unsigned NOT NULL,
  `melee_rune` int(10) unsigned NOT NULL,
  `magic_rune` int(10) unsigned NOT NULL,
  `persistent` tinyint(3) unsigned NOT NULL,
  `dot_rune` int(10) NOT NULL DEFAULT 0,
  `caston_x` int(10) NOT NULL DEFAULT 0,
  `caston_y` int(10) NOT NULL DEFAULT 0,
  `caston_z` int(10) NOT NULL DEFAULT 0,
  `ExtraDIChance` int(10) NOT NULL DEFAULT 0,
  `instrument_mod` int(10) NOT NULL DEFAULT 10,
  PRIMARY KEY (`character_id`,`slot_id`),
  KEY `character_id` (`character_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "caster_level": 57,    "caster_name": "qYDYqIyoDwSlkwdEmvCfNnuqT",    "melee_rune": 10,    "persistent": 74,    "dot_rune": 54,    "spell_id": 34,    "numhits": 23,    "magic_rune": 59,    "caston_y": 46,    "instrument_mod": 25,    "character_id": 78,    "extra_di_chance": 45,    "slot_id": 80,    "counters": 11,    "caston_x": 21,    "caston_z": 31,    "ticsremaining": 50}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned



*/

// CharacterBuffs struct is a row record of the character_buffs table in the eqemu database
type CharacterBuffs struct {
	//[ 0] character_id                                   uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	CharacterID uint32 `gorm:"primary_key;column:character_id;type:uint;" json:"character_id"`
	//[ 1] slot_id                                        utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: []
	SlotID uint32 `gorm:"primary_key;column:slot_id;type:utinyint;" json:"slot_id"`
	//[ 2] spell_id                                       usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: []
	SpellID uint32 `gorm:"column:spell_id;type:usmallint;" json:"spell_id"`
	//[ 3] caster_level                                   utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: []
	CasterLevel uint32 `gorm:"column:caster_level;type:utinyint;" json:"caster_level"`
	//[ 4] caster_name                                    varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	CasterName string `gorm:"column:caster_name;type:varchar;size:64;" json:"caster_name"`
	//[ 5] ticsremaining                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Ticsremaining int32 `gorm:"column:ticsremaining;type:int;" json:"ticsremaining"`
	//[ 6] counters                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	Counters uint32 `gorm:"column:counters;type:uint;" json:"counters"`
	//[ 7] numhits                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	Numhits uint32 `gorm:"column:numhits;type:uint;" json:"numhits"`
	//[ 8] melee_rune                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	MeleeRune uint32 `gorm:"column:melee_rune;type:uint;" json:"melee_rune"`
	//[ 9] magic_rune                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	MagicRune uint32 `gorm:"column:magic_rune;type:uint;" json:"magic_rune"`
	//[10] persistent                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: []
	Persistent uint32 `gorm:"column:persistent;type:utinyint;" json:"persistent"`
	//[11] dot_rune                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DotRune int32 `gorm:"column:dot_rune;type:int;default:0;" json:"dot_rune"`
	//[12] caston_x                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CastonX int32 `gorm:"column:caston_x;type:int;default:0;" json:"caston_x"`
	//[13] caston_y                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CastonY int32 `gorm:"column:caston_y;type:int;default:0;" json:"caston_y"`
	//[14] caston_z                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CastonZ int32 `gorm:"column:caston_z;type:int;default:0;" json:"caston_z"`
	//[15] ExtraDIChance                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ExtraDIChance int32 `gorm:"column:ExtraDIChance;type:int;default:0;" json:"extra_di_chance"`
	//[16] instrument_mod                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [10]
	InstrumentMod int32 `gorm:"column:instrument_mod;type:int;default:10;" json:"instrument_mod"`
}

var character_buffsTableInfo = &TableInfo{
	Name: "character_buffs",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "character_id",
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
			GoFieldName:        "CharacterID",
			GoFieldType:        "uint32",
			JSONFieldName:      "character_id",
			ProtobufFieldName:  "character_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "slot_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "SlotID",
			GoFieldType:        "uint32",
			JSONFieldName:      "slot_id",
			ProtobufFieldName:  "slot_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "spell_id",
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
			GoFieldName:        "SpellID",
			GoFieldType:        "uint32",
			JSONFieldName:      "spell_id",
			ProtobufFieldName:  "spell_id",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "caster_level",
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
			GoFieldName:        "CasterLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "caster_level",
			ProtobufFieldName:  "caster_level",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "caster_name",
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
			GoFieldName:        "CasterName",
			GoFieldType:        "string",
			JSONFieldName:      "caster_name",
			ProtobufFieldName:  "caster_name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "ticsremaining",
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
			GoFieldName:        "Ticsremaining",
			GoFieldType:        "int32",
			JSONFieldName:      "ticsremaining",
			ProtobufFieldName:  "ticsremaining",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "counters",
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
			GoFieldName:        "Counters",
			GoFieldType:        "uint32",
			JSONFieldName:      "counters",
			ProtobufFieldName:  "counters",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "numhits",
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
			GoFieldName:        "Numhits",
			GoFieldType:        "uint32",
			JSONFieldName:      "numhits",
			ProtobufFieldName:  "numhits",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "melee_rune",
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
			GoFieldName:        "MeleeRune",
			GoFieldType:        "uint32",
			JSONFieldName:      "melee_rune",
			ProtobufFieldName:  "melee_rune",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "magic_rune",
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
			GoFieldName:        "MagicRune",
			GoFieldType:        "uint32",
			JSONFieldName:      "magic_rune",
			ProtobufFieldName:  "magic_rune",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "persistent",
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
			GoFieldName:        "Persistent",
			GoFieldType:        "uint32",
			JSONFieldName:      "persistent",
			ProtobufFieldName:  "persistent",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "dot_rune",
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
			GoFieldName:        "DotRune",
			GoFieldType:        "int32",
			JSONFieldName:      "dot_rune",
			ProtobufFieldName:  "dot_rune",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "caston_x",
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
			GoFieldName:        "CastonX",
			GoFieldType:        "int32",
			JSONFieldName:      "caston_x",
			ProtobufFieldName:  "caston_x",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "caston_y",
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
			GoFieldName:        "CastonY",
			GoFieldType:        "int32",
			JSONFieldName:      "caston_y",
			ProtobufFieldName:  "caston_y",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "caston_z",
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
			GoFieldName:        "CastonZ",
			GoFieldType:        "int32",
			JSONFieldName:      "caston_z",
			ProtobufFieldName:  "caston_z",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "ExtraDIChance",
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
			GoFieldName:        "ExtraDIChance",
			GoFieldType:        "int32",
			JSONFieldName:      "extra_di_chance",
			ProtobufFieldName:  "extra_di_chance",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "instrument_mod",
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
			GoFieldName:        "InstrumentMod",
			GoFieldType:        "int32",
			JSONFieldName:      "instrument_mod",
			ProtobufFieldName:  "instrument_mod",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterBuffs) TableName() string {
	return "character_buffs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterBuffs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterBuffs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterBuffs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterBuffs) TableInfo() *TableInfo {
	return character_buffsTableInfo
}
