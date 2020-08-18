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


CREATE TABLE `character_potionbelt` (
  `id` int(11) unsigned NOT NULL DEFAULT 0,
  `potion_id` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `item_id` int(11) unsigned NOT NULL DEFAULT 0,
  `icon` int(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`potion_id`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 17,    "potion_id": 79,    "item_id": 31,    "icon": 66}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned



*/

// CharacterPotionbelt struct is a row record of the character_potionbelt table in the eqemu database
type CharacterPotionbelt struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ID uint32 `gorm:"primary_key;column:id;type:uint;default:0;" json:"id"`
	//[ 1] potion_id                                      utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	PotionID uint32 `gorm:"primary_key;column:potion_id;type:utinyint;default:0;" json:"potion_id"`
	//[ 2] item_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ItemID uint32 `gorm:"column:item_id;type:uint;default:0;" json:"item_id"`
	//[ 3] icon                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Icon uint32 `gorm:"column:icon;type:uint;default:0;" json:"icon"`
}

var character_potionbeltTableInfo = &TableInfo{
	Name: "character_potionbelt",
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
			Name:               "potion_id",
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
			GoFieldName:        "PotionID",
			GoFieldType:        "uint32",
			JSONFieldName:      "potion_id",
			ProtobufFieldName:  "potion_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "item_id",
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
			GoFieldName:        "ItemID",
			GoFieldType:        "uint32",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "icon",
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
			GoFieldName:        "Icon",
			GoFieldType:        "uint32",
			JSONFieldName:      "icon",
			ProtobufFieldName:  "icon",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterPotionbelt) TableName() string {
	return "character_potionbelt"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterPotionbelt) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterPotionbelt) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterPotionbelt) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterPotionbelt) TableInfo() *TableInfo {
	return character_potionbeltTableInfo
}
