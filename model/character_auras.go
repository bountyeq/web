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


CREATE TABLE `character_auras` (
  `id` int(10) NOT NULL,
  `slot` tinyint(10) NOT NULL,
  `spell_id` int(10) NOT NULL,
  PRIMARY KEY (`id`,`slot`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 38,    "slot": 19,    "spell_id": 53}



*/

// CharacterAuras struct is a row record of the character_auras table in the eqemu database
type CharacterAuras struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;column:id;type:int;" json:"id"`
	//[ 1] slot                                           tinyint              null: false  primary: true   isArray: false  auto: false  col: tinyint         len: -1      default: []
	Slot int32 `gorm:"primary_key;column:slot;type:tinyint;" json:"slot"`
	//[ 2] spell_id                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SpellID int32 `gorm:"column:spell_id;type:int;" json:"spell_id"`
}

var character_aurasTableInfo = &TableInfo{
	Name: "character_auras",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
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
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "slot",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Slot",
			GoFieldType:        "int32",
			JSONFieldName:      "slot",
			ProtobufFieldName:  "slot",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterAuras) TableName() string {
	return "character_auras"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterAuras) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterAuras) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterAuras) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterAuras) TableInfo() *TableInfo {
	return character_aurasTableInfo
}
