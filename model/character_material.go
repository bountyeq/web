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


CREATE TABLE `character_material` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `slot` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `blue` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `green` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `red` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `use_tint` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `color` int(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`slot`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 47,    "slot": 84,    "blue": 97,    "green": 69,    "red": 48,    "use_tint": 95,    "color": 15}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned



*/

// CharacterMaterial struct is a row record of the character_material table in the eqemu database
type CharacterMaterial struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] slot                                           utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Slot uint32 `gorm:"primary_key;column:slot;type:utinyint;default:0;" json:"slot"`
	//[ 2] blue                                           utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Blue uint32 `gorm:"column:blue;type:utinyint;default:0;" json:"blue"`
	//[ 3] green                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Green uint32 `gorm:"column:green;type:utinyint;default:0;" json:"green"`
	//[ 4] red                                            utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Red uint32 `gorm:"column:red;type:utinyint;default:0;" json:"red"`
	//[ 5] use_tint                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	UseTint uint32 `gorm:"column:use_tint;type:utinyint;default:0;" json:"use_tint"`
	//[ 6] color                                          uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Color uint32 `gorm:"column:color;type:uint;default:0;" json:"color"`
}

var character_materialTableInfo = &TableInfo{
	Name: "character_material",
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
			Name:               "slot",
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
			GoFieldName:        "Slot",
			GoFieldType:        "uint32",
			JSONFieldName:      "slot",
			ProtobufFieldName:  "slot",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "blue",
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
			GoFieldName:        "Blue",
			GoFieldType:        "uint32",
			JSONFieldName:      "blue",
			ProtobufFieldName:  "blue",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "green",
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
			GoFieldName:        "Green",
			GoFieldType:        "uint32",
			JSONFieldName:      "green",
			ProtobufFieldName:  "green",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "red",
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
			GoFieldName:        "Red",
			GoFieldType:        "uint32",
			JSONFieldName:      "red",
			ProtobufFieldName:  "red",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "use_tint",
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
			GoFieldName:        "UseTint",
			GoFieldType:        "uint32",
			JSONFieldName:      "use_tint",
			ProtobufFieldName:  "use_tint",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "color",
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
			GoFieldName:        "Color",
			GoFieldType:        "uint32",
			JSONFieldName:      "color",
			ProtobufFieldName:  "color",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterMaterial) TableName() string {
	return "character_material"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterMaterial) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterMaterial) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterMaterial) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterMaterial) TableInfo() *TableInfo {
	return character_materialTableInfo
}
