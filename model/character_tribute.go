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


CREATE TABLE `character_tribute` (
  `id` int(11) unsigned NOT NULL DEFAULT 0,
  `tier` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `tribute` int(11) unsigned NOT NULL DEFAULT 0,
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 61,    "tier": 60,    "tribute": 27}


Comments
-------------------------------------
[ 0] column is set for unsignedWarning table: character_tribute does not have a primary key defined, setting col position 1 id as primary key

[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// CharacterTribute struct is a row record of the character_tribute table in the eqemu database
type CharacterTribute struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ID uint32 `gorm:"primary_key;column:id;type:uint;default:0;" json:"id"`
	//[ 1] tier                                           utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Tier uint32 `gorm:"column:tier;type:utinyint;default:0;" json:"tier"`
	//[ 2] tribute                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Tribute uint32 `gorm:"column:tribute;type:uint;default:0;" json:"tribute"`
}

var character_tributeTableInfo = &TableInfo{
	Name: "character_tribute",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: ``,
			Notes: `column is set for unsignedWarning table: character_tribute does not have a primary key defined, setting col position 1 id as primary key
`,
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
			Name:               "tier",
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
			GoFieldName:        "Tier",
			GoFieldType:        "uint32",
			JSONFieldName:      "tier",
			ProtobufFieldName:  "tier",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "tribute",
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
			GoFieldName:        "Tribute",
			GoFieldType:        "uint32",
			JSONFieldName:      "tribute",
			ProtobufFieldName:  "tribute",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterTribute) TableName() string {
	return "character_tribute"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterTribute) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterTribute) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterTribute) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterTribute) TableInfo() *TableInfo {
	return character_tributeTableInfo
}
