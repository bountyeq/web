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


CREATE TABLE `character_disciplines` (
  `id` int(11) unsigned NOT NULL DEFAULT 0,
  `slot_id` smallint(11) unsigned NOT NULL DEFAULT 0,
  `disc_id` smallint(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`slot_id`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 10,    "slot_id": 41,    "disc_id": 47}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// CharacterDisciplines struct is a row record of the character_disciplines table in the eqemu database
type CharacterDisciplines struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ID uint32 `gorm:"primary_key;column:id;type:uint;default:0;" json:"id"`
	//[ 1] slot_id                                        usmallint            null: false  primary: true   isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	SlotID uint32 `gorm:"primary_key;column:slot_id;type:usmallint;default:0;" json:"slot_id"`
	//[ 2] disc_id                                        usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	DiscID uint32 `gorm:"column:disc_id;type:usmallint;default:0;" json:"disc_id"`
}

var character_disciplinesTableInfo = &TableInfo{
	Name: "character_disciplines",
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
			Name:               "slot_id",
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
			GoFieldName:        "SlotID",
			GoFieldType:        "uint32",
			JSONFieldName:      "slot_id",
			ProtobufFieldName:  "slot_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "disc_id",
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
			GoFieldName:        "DiscID",
			GoFieldType:        "uint32",
			JSONFieldName:      "disc_id",
			ProtobufFieldName:  "disc_id",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterDisciplines) TableName() string {
	return "character_disciplines"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterDisciplines) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterDisciplines) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterDisciplines) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterDisciplines) TableInfo() *TableInfo {
	return character_disciplinesTableInfo
}
