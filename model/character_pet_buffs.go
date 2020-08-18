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


CREATE TABLE `character_pet_buffs` (
  `char_id` int(11) NOT NULL,
  `pet` int(11) NOT NULL,
  `slot` int(11) NOT NULL,
  `spell_id` int(11) NOT NULL,
  `caster_level` tinyint(4) NOT NULL DEFAULT 0,
  `castername` varchar(64) NOT NULL DEFAULT '',
  `ticsremaining` int(11) NOT NULL DEFAULT 0,
  `counters` int(11) NOT NULL DEFAULT 0,
  `numhits` int(11) NOT NULL DEFAULT 0,
  `rune` int(11) NOT NULL DEFAULT 0,
  `instrument_mod` tinyint(3) unsigned NOT NULL DEFAULT 10,
  PRIMARY KEY (`char_id`,`pet`,`slot`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "ticsremaining": 23,    "numhits": 83,    "instrument_mod": 99,    "pet": 45,    "spell_id": 51,    "caster_level": 65,    "castername": "kOhQEekKuvsdgbOmIkwiWTkdv",    "char_id": 93,    "slot": 93,    "counters": 6,    "rune": 33}


Comments
-------------------------------------
[10] column is set for unsigned



*/

// CharacterPetBuffs struct is a row record of the character_pet_buffs table in the eqemu database
type CharacterPetBuffs struct {
	//[ 0] char_id                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	CharID int32 `gorm:"primary_key;column:char_id;type:int;" json:"char_id"`
	//[ 1] pet                                            int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	Pet int32 `gorm:"primary_key;column:pet;type:int;" json:"pet"`
	//[ 2] slot                                           int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	Slot int32 `gorm:"primary_key;column:slot;type:int;" json:"slot"`
	//[ 3] spell_id                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SpellID int32 `gorm:"column:spell_id;type:int;" json:"spell_id"`
	//[ 4] caster_level                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	CasterLevel int32 `gorm:"column:caster_level;type:tinyint;default:0;" json:"caster_level"`
	//[ 5] castername                                     varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Castername string `gorm:"column:castername;type:varchar;size:64;default:'';" json:"castername"`
	//[ 6] ticsremaining                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Ticsremaining int32 `gorm:"column:ticsremaining;type:int;default:0;" json:"ticsremaining"`
	//[ 7] counters                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Counters int32 `gorm:"column:counters;type:int;default:0;" json:"counters"`
	//[ 8] numhits                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Numhits int32 `gorm:"column:numhits;type:int;default:0;" json:"numhits"`
	//[ 9] rune                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Rune int32 `gorm:"column:rune;type:int;default:0;" json:"rune"`
	//[10] instrument_mod                                 utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [10]
	InstrumentMod uint32 `gorm:"column:instrument_mod;type:utinyint;default:10;" json:"instrument_mod"`
}

var character_pet_buffsTableInfo = &TableInfo{
	Name: "character_pet_buffs",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "char_id",
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
			GoFieldName:        "CharID",
			GoFieldType:        "int32",
			JSONFieldName:      "char_id",
			ProtobufFieldName:  "char_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "pet",
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
			GoFieldName:        "Pet",
			GoFieldType:        "int32",
			JSONFieldName:      "pet",
			ProtobufFieldName:  "pet",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "slot",
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
			GoFieldName:        "Slot",
			GoFieldType:        "int32",
			JSONFieldName:      "slot",
			ProtobufFieldName:  "slot",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "spell_id",
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
			GoFieldName:        "SpellID",
			GoFieldType:        "int32",
			JSONFieldName:      "spell_id",
			ProtobufFieldName:  "spell_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "caster_level",
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
			GoFieldName:        "CasterLevel",
			GoFieldType:        "int32",
			JSONFieldName:      "caster_level",
			ProtobufFieldName:  "caster_level",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "castername",
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
			GoFieldName:        "Castername",
			GoFieldType:        "string",
			JSONFieldName:      "castername",
			ProtobufFieldName:  "castername",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "counters",
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
			GoFieldName:        "Counters",
			GoFieldType:        "int32",
			JSONFieldName:      "counters",
			ProtobufFieldName:  "counters",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "numhits",
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
			GoFieldName:        "Numhits",
			GoFieldType:        "int32",
			JSONFieldName:      "numhits",
			ProtobufFieldName:  "numhits",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "rune",
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
			GoFieldName:        "Rune",
			GoFieldType:        "int32",
			JSONFieldName:      "rune",
			ProtobufFieldName:  "rune",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "instrument_mod",
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
			GoFieldName:        "InstrumentMod",
			GoFieldType:        "uint32",
			JSONFieldName:      "instrument_mod",
			ProtobufFieldName:  "instrument_mod",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterPetBuffs) TableName() string {
	return "character_pet_buffs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterPetBuffs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterPetBuffs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterPetBuffs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterPetBuffs) TableInfo() *TableInfo {
	return character_pet_buffsTableInfo
}
