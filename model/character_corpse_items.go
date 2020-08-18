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


CREATE TABLE `character_corpse_items` (
  `corpse_id` int(11) unsigned NOT NULL,
  `equip_slot` int(11) unsigned NOT NULL,
  `item_id` int(11) unsigned DEFAULT NULL,
  `charges` int(11) unsigned DEFAULT NULL,
  `aug_1` int(11) unsigned DEFAULT 0,
  `aug_2` int(11) unsigned DEFAULT 0,
  `aug_3` int(11) unsigned DEFAULT 0,
  `aug_4` int(11) unsigned DEFAULT 0,
  `aug_5` int(11) unsigned DEFAULT 0,
  `aug_6` int(11) NOT NULL DEFAULT 0,
  `attuned` smallint(5) NOT NULL DEFAULT 0,
  PRIMARY KEY (`corpse_id`,`equip_slot`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "aug_1": 57,    "aug_6": 72,    "attuned": 92,    "item_id": 17,    "equip_slot": 0,    "charges": 64,    "aug_2": 14,    "aug_3": 22,    "aug_4": 73,    "aug_5": 73,    "corpse_id": 93}


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



*/

// CharacterCorpseItems struct is a row record of the character_corpse_items table in the eqemu database
type CharacterCorpseItems struct {
	//[ 0] corpse_id                                      uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	CorpseID uint32 `gorm:"primary_key;column:corpse_id;type:uint;" json:"corpse_id"`
	//[ 1] equip_slot                                     uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	EquipSlot uint32 `gorm:"primary_key;column:equip_slot;type:uint;" json:"equip_slot"`
	//[ 2] item_id                                        uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [NULL]
	ItemID null.Int `gorm:"column:item_id;type:uint;" json:"item_id"`
	//[ 3] charges                                        uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [NULL]
	Charges null.Int `gorm:"column:charges;type:uint;" json:"charges"`
	//[ 4] aug_1                                          uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Aug1 null.Int `gorm:"column:aug_1;type:uint;default:0;" json:"aug_1"`
	//[ 5] aug_2                                          uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Aug2 null.Int `gorm:"column:aug_2;type:uint;default:0;" json:"aug_2"`
	//[ 6] aug_3                                          uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Aug3 null.Int `gorm:"column:aug_3;type:uint;default:0;" json:"aug_3"`
	//[ 7] aug_4                                          uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Aug4 null.Int `gorm:"column:aug_4;type:uint;default:0;" json:"aug_4"`
	//[ 8] aug_5                                          uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Aug5 null.Int `gorm:"column:aug_5;type:uint;default:0;" json:"aug_5"`
	//[ 9] aug_6                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Aug6 int32 `gorm:"column:aug_6;type:int;default:0;" json:"aug_6"`
	//[10] attuned                                        smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Attuned int32 `gorm:"column:attuned;type:smallint;default:0;" json:"attuned"`
}

var character_corpse_itemsTableInfo = &TableInfo{
	Name: "character_corpse_items",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "corpse_id",
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
			GoFieldName:        "CorpseID",
			GoFieldType:        "uint32",
			JSONFieldName:      "corpse_id",
			ProtobufFieldName:  "corpse_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "equip_slot",
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
			GoFieldName:        "EquipSlot",
			GoFieldType:        "uint32",
			JSONFieldName:      "equip_slot",
			ProtobufFieldName:  "equip_slot",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "item_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ItemID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "charges",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Charges",
			GoFieldType:        "null.Int",
			JSONFieldName:      "charges",
			ProtobufFieldName:  "charges",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "aug_1",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Aug1",
			GoFieldType:        "null.Int",
			JSONFieldName:      "aug_1",
			ProtobufFieldName:  "aug_1",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "aug_2",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Aug2",
			GoFieldType:        "null.Int",
			JSONFieldName:      "aug_2",
			ProtobufFieldName:  "aug_2",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "aug_3",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Aug3",
			GoFieldType:        "null.Int",
			JSONFieldName:      "aug_3",
			ProtobufFieldName:  "aug_3",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "aug_4",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Aug4",
			GoFieldType:        "null.Int",
			JSONFieldName:      "aug_4",
			ProtobufFieldName:  "aug_4",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "aug_5",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Aug5",
			GoFieldType:        "null.Int",
			JSONFieldName:      "aug_5",
			ProtobufFieldName:  "aug_5",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "aug_6",
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
			GoFieldName:        "Aug6",
			GoFieldType:        "int32",
			JSONFieldName:      "aug_6",
			ProtobufFieldName:  "aug_6",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "attuned",
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
			GoFieldName:        "Attuned",
			GoFieldType:        "int32",
			JSONFieldName:      "attuned",
			ProtobufFieldName:  "attuned",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterCorpseItems) TableName() string {
	return "character_corpse_items"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterCorpseItems) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterCorpseItems) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterCorpseItems) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterCorpseItems) TableInfo() *TableInfo {
	return character_corpse_itemsTableInfo
}
