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


CREATE TABLE `character_alternate_abilities` (
  `id` int(11) unsigned NOT NULL DEFAULT 0,
  `aa_id` smallint(11) unsigned NOT NULL DEFAULT 0,
  `aa_value` smallint(11) unsigned NOT NULL DEFAULT 0,
  `charges` smallint(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`aa_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "aa_value": 13,    "charges": 92,    "id": 90,    "aa_id": 25}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned



*/

// CharacterAlternateAbilities struct is a row record of the character_alternate_abilities table in the eqemu database
type CharacterAlternateAbilities struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ID uint32 `gorm:"primary_key;column:id;type:uint;default:0;" json:"id"`
	//[ 1] aa_id                                          usmallint            null: false  primary: true   isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	AaID uint32 `gorm:"primary_key;column:aa_id;type:usmallint;default:0;" json:"aa_id"`
	//[ 2] aa_value                                       usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	AaValue uint32 `gorm:"column:aa_value;type:usmallint;default:0;" json:"aa_value"`
	//[ 3] charges                                        usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	Charges uint32 `gorm:"column:charges;type:usmallint;default:0;" json:"charges"`
}

var character_alternate_abilitiesTableInfo = &TableInfo{
	Name: "character_alternate_abilities",
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
			Name:               "aa_id",
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
			GoFieldName:        "AaID",
			GoFieldType:        "uint32",
			JSONFieldName:      "aa_id",
			ProtobufFieldName:  "aa_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "aa_value",
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
			GoFieldName:        "AaValue",
			GoFieldType:        "uint32",
			JSONFieldName:      "aa_value",
			ProtobufFieldName:  "aa_value",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "charges",
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
			GoFieldName:        "Charges",
			GoFieldType:        "uint32",
			JSONFieldName:      "charges",
			ProtobufFieldName:  "charges",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterAlternateAbilities) TableName() string {
	return "character_alternate_abilities"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterAlternateAbilities) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterAlternateAbilities) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterAlternateAbilities) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterAlternateAbilities) TableInfo() *TableInfo {
	return character_alternate_abilitiesTableInfo
}
