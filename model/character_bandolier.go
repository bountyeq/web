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


CREATE TABLE `character_bandolier` (
  `id` int(11) unsigned NOT NULL DEFAULT 0,
  `bandolier_id` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `bandolier_slot` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `item_id` int(11) unsigned NOT NULL DEFAULT 0,
  `icon` int(11) unsigned NOT NULL DEFAULT 0,
  `bandolier_name` varchar(32) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`,`bandolier_id`,`bandolier_slot`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 43,    "bandolier_id": 91,    "bandolier_slot": 67,    "item_id": 60,    "icon": 53,    "bandolier_name": "jZZxQxVFaZtEwoWfYBFfQufdF"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned



*/

// CharacterBandolier struct is a row record of the character_bandolier table in the eqemu database
type CharacterBandolier struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ID uint32 `gorm:"primary_key;column:id;type:uint;default:0;" json:"id"`
	//[ 1] bandolier_id                                   utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	BandolierID uint32 `gorm:"primary_key;column:bandolier_id;type:utinyint;default:0;" json:"bandolier_id"`
	//[ 2] bandolier_slot                                 utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	BandolierSlot uint32 `gorm:"primary_key;column:bandolier_slot;type:utinyint;default:0;" json:"bandolier_slot"`
	//[ 3] item_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ItemID uint32 `gorm:"column:item_id;type:uint;default:0;" json:"item_id"`
	//[ 4] icon                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Icon uint32 `gorm:"column:icon;type:uint;default:0;" json:"icon"`
	//[ 5] bandolier_name                                 varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['0']
	BandolierName string `gorm:"column:bandolier_name;type:varchar;size:32;default:'0';" json:"bandolier_name"`
}

var character_bandolierTableInfo = &TableInfo{
	Name: "character_bandolier",
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
			Name:               "bandolier_id",
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
			GoFieldName:        "BandolierID",
			GoFieldType:        "uint32",
			JSONFieldName:      "bandolier_id",
			ProtobufFieldName:  "bandolier_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "bandolier_slot",
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
			GoFieldName:        "BandolierSlot",
			GoFieldType:        "uint32",
			JSONFieldName:      "bandolier_slot",
			ProtobufFieldName:  "bandolier_slot",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "bandolier_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "BandolierName",
			GoFieldType:        "string",
			JSONFieldName:      "bandolier_name",
			ProtobufFieldName:  "bandolier_name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterBandolier) TableName() string {
	return "character_bandolier"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterBandolier) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterBandolier) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterBandolier) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterBandolier) TableInfo() *TableInfo {
	return character_bandolierTableInfo
}
