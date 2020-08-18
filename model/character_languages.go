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


CREATE TABLE `character_languages` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `lang_id` smallint(11) unsigned NOT NULL DEFAULT 0,
  `value` smallint(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`lang_id`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 75,    "lang_id": 29,    "value": 66}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// CharacterLanguages struct is a row record of the character_languages table in the eqemu database
type CharacterLanguages struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] lang_id                                        usmallint            null: false  primary: true   isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	LangID uint32 `gorm:"primary_key;column:lang_id;type:usmallint;default:0;" json:"lang_id"`
	//[ 2] value                                          usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	Value uint32 `gorm:"column:value;type:usmallint;default:0;" json:"value"`
}

var character_languagesTableInfo = &TableInfo{
	Name: "character_languages",
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
			Name:               "lang_id",
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
			GoFieldName:        "LangID",
			GoFieldType:        "uint32",
			JSONFieldName:      "lang_id",
			ProtobufFieldName:  "lang_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "value",
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
			GoFieldName:        "Value",
			GoFieldType:        "uint32",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterLanguages) TableName() string {
	return "character_languages"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterLanguages) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterLanguages) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterLanguages) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterLanguages) TableInfo() *TableInfo {
	return character_languagesTableInfo
}
