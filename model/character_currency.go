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


CREATE TABLE `character_currency` (
  `id` int(11) unsigned NOT NULL DEFAULT 0,
  `platinum` int(11) unsigned NOT NULL DEFAULT 0,
  `gold` int(11) unsigned NOT NULL DEFAULT 0,
  `silver` int(11) unsigned NOT NULL DEFAULT 0,
  `copper` int(11) unsigned NOT NULL DEFAULT 0,
  `platinum_bank` int(11) unsigned NOT NULL DEFAULT 0,
  `gold_bank` int(11) unsigned NOT NULL DEFAULT 0,
  `silver_bank` int(11) unsigned NOT NULL DEFAULT 0,
  `copper_bank` int(11) unsigned NOT NULL DEFAULT 0,
  `platinum_cursor` int(11) unsigned NOT NULL DEFAULT 0,
  `gold_cursor` int(11) unsigned NOT NULL DEFAULT 0,
  `silver_cursor` int(11) unsigned NOT NULL DEFAULT 0,
  `copper_cursor` int(11) unsigned NOT NULL DEFAULT 0,
  `radiant_crystals` int(11) unsigned NOT NULL DEFAULT 0,
  `career_radiant_crystals` int(11) unsigned NOT NULL DEFAULT 0,
  `ebon_crystals` int(11) unsigned NOT NULL DEFAULT 0,
  `career_ebon_crystals` int(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "platinum_cursor": 32,    "id": 42,    "silver": 54,    "gold_bank": 3,    "platinum_bank": 50,    "copper_bank": 26,    "gold_cursor": 91,    "copper_cursor": 40,    "career_radiant_crystals": 0,    "platinum": 86,    "gold": 38,    "copper": 9,    "ebon_crystals": 55,    "silver_bank": 4,    "silver_cursor": 27,    "radiant_crystals": 56,    "career_ebon_crystals": 88}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned
[11] column is set for unsigned
[12] column is set for unsigned
[13] column is set for unsigned
[14] column is set for unsigned
[15] column is set for unsigned
[16] column is set for unsigned



*/

// CharacterCurrency struct is a row record of the character_currency table in the eqemu database
type CharacterCurrency struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ID uint32 `gorm:"primary_key;column:id;type:uint;default:0;" json:"id"`
	//[ 1] platinum                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Platinum uint32 `gorm:"column:platinum;type:uint;default:0;" json:"platinum"`
	//[ 2] gold                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Gold uint32 `gorm:"column:gold;type:uint;default:0;" json:"gold"`
	//[ 3] silver                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Silver uint32 `gorm:"column:silver;type:uint;default:0;" json:"silver"`
	//[ 4] copper                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Copper uint32 `gorm:"column:copper;type:uint;default:0;" json:"copper"`
	//[ 5] platinum_bank                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PlatinumBank uint32 `gorm:"column:platinum_bank;type:uint;default:0;" json:"platinum_bank"`
	//[ 6] gold_bank                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	GoldBank uint32 `gorm:"column:gold_bank;type:uint;default:0;" json:"gold_bank"`
	//[ 7] silver_bank                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	SilverBank uint32 `gorm:"column:silver_bank;type:uint;default:0;" json:"silver_bank"`
	//[ 8] copper_bank                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CopperBank uint32 `gorm:"column:copper_bank;type:uint;default:0;" json:"copper_bank"`
	//[ 9] platinum_cursor                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PlatinumCursor uint32 `gorm:"column:platinum_cursor;type:uint;default:0;" json:"platinum_cursor"`
	//[10] gold_cursor                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	GoldCursor uint32 `gorm:"column:gold_cursor;type:uint;default:0;" json:"gold_cursor"`
	//[11] silver_cursor                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	SilverCursor uint32 `gorm:"column:silver_cursor;type:uint;default:0;" json:"silver_cursor"`
	//[12] copper_cursor                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CopperCursor uint32 `gorm:"column:copper_cursor;type:uint;default:0;" json:"copper_cursor"`
	//[13] radiant_crystals                               uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	RadiantCrystals uint32 `gorm:"column:radiant_crystals;type:uint;default:0;" json:"radiant_crystals"`
	//[14] career_radiant_crystals                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CareerRadiantCrystals uint32 `gorm:"column:career_radiant_crystals;type:uint;default:0;" json:"career_radiant_crystals"`
	//[15] ebon_crystals                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EbonCrystals uint32 `gorm:"column:ebon_crystals;type:uint;default:0;" json:"ebon_crystals"`
	//[16] career_ebon_crystals                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CareerEbonCrystals uint32 `gorm:"column:career_ebon_crystals;type:uint;default:0;" json:"career_ebon_crystals"`
}

var character_currencyTableInfo = &TableInfo{
	Name: "character_currency",
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
			Name:               "platinum",
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
			GoFieldName:        "Platinum",
			GoFieldType:        "uint32",
			JSONFieldName:      "platinum",
			ProtobufFieldName:  "platinum",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "gold",
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
			GoFieldName:        "Gold",
			GoFieldType:        "uint32",
			JSONFieldName:      "gold",
			ProtobufFieldName:  "gold",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "silver",
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
			GoFieldName:        "Silver",
			GoFieldType:        "uint32",
			JSONFieldName:      "silver",
			ProtobufFieldName:  "silver",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "copper",
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
			GoFieldName:        "Copper",
			GoFieldType:        "uint32",
			JSONFieldName:      "copper",
			ProtobufFieldName:  "copper",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "platinum_bank",
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
			GoFieldName:        "PlatinumBank",
			GoFieldType:        "uint32",
			JSONFieldName:      "platinum_bank",
			ProtobufFieldName:  "platinum_bank",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "gold_bank",
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
			GoFieldName:        "GoldBank",
			GoFieldType:        "uint32",
			JSONFieldName:      "gold_bank",
			ProtobufFieldName:  "gold_bank",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "silver_bank",
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
			GoFieldName:        "SilverBank",
			GoFieldType:        "uint32",
			JSONFieldName:      "silver_bank",
			ProtobufFieldName:  "silver_bank",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "copper_bank",
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
			GoFieldName:        "CopperBank",
			GoFieldType:        "uint32",
			JSONFieldName:      "copper_bank",
			ProtobufFieldName:  "copper_bank",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "platinum_cursor",
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
			GoFieldName:        "PlatinumCursor",
			GoFieldType:        "uint32",
			JSONFieldName:      "platinum_cursor",
			ProtobufFieldName:  "platinum_cursor",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "gold_cursor",
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
			GoFieldName:        "GoldCursor",
			GoFieldType:        "uint32",
			JSONFieldName:      "gold_cursor",
			ProtobufFieldName:  "gold_cursor",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "silver_cursor",
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
			GoFieldName:        "SilverCursor",
			GoFieldType:        "uint32",
			JSONFieldName:      "silver_cursor",
			ProtobufFieldName:  "silver_cursor",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "copper_cursor",
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
			GoFieldName:        "CopperCursor",
			GoFieldType:        "uint32",
			JSONFieldName:      "copper_cursor",
			ProtobufFieldName:  "copper_cursor",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "radiant_crystals",
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
			GoFieldName:        "RadiantCrystals",
			GoFieldType:        "uint32",
			JSONFieldName:      "radiant_crystals",
			ProtobufFieldName:  "radiant_crystals",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "career_radiant_crystals",
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
			GoFieldName:        "CareerRadiantCrystals",
			GoFieldType:        "uint32",
			JSONFieldName:      "career_radiant_crystals",
			ProtobufFieldName:  "career_radiant_crystals",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "ebon_crystals",
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
			GoFieldName:        "EbonCrystals",
			GoFieldType:        "uint32",
			JSONFieldName:      "ebon_crystals",
			ProtobufFieldName:  "ebon_crystals",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "career_ebon_crystals",
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
			GoFieldName:        "CareerEbonCrystals",
			GoFieldType:        "uint32",
			JSONFieldName:      "career_ebon_crystals",
			ProtobufFieldName:  "career_ebon_crystals",
			ProtobufType:       "uint32",
			ProtobufPos:        17,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterCurrency) TableName() string {
	return "character_currency"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterCurrency) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterCurrency) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterCurrency) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterCurrency) TableInfo() *TableInfo {
	return character_currencyTableInfo
}
