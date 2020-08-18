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


CREATE TABLE `character_leadership_abilities` (
  `id` int(11) unsigned NOT NULL DEFAULT 0,
  `slot` smallint(11) unsigned NOT NULL DEFAULT 0,
  `rank` smallint(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`slot`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 58,    "slot": 82,    "rank": 98}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// CharacterLeadershipAbilities struct is a row record of the character_leadership_abilities table in the eqemu database
type CharacterLeadershipAbilities struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ID uint32 `gorm:"primary_key;column:id;type:uint;default:0;" json:"id"`
	//[ 1] slot                                           usmallint            null: false  primary: true   isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	Slot uint32 `gorm:"primary_key;column:slot;type:usmallint;default:0;" json:"slot"`
	//[ 2] rank                                           usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	Rank uint32 `gorm:"column:rank;type:usmallint;default:0;" json:"rank"`
}

var character_leadership_abilitiesTableInfo = &TableInfo{
	Name: "character_leadership_abilities",
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
			Name:               "slot",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
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
			Name:               "rank",
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
			GoFieldName:        "Rank",
			GoFieldType:        "uint32",
			JSONFieldName:      "rank",
			ProtobufFieldName:  "rank",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterLeadershipAbilities) TableName() string {
	return "character_leadership_abilities"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterLeadershipAbilities) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterLeadershipAbilities) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterLeadershipAbilities) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterLeadershipAbilities) TableInfo() *TableInfo {
	return character_leadership_abilitiesTableInfo
}
