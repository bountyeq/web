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


CREATE TABLE `character_item_recast` (
  `id` int(11) unsigned NOT NULL DEFAULT 0,
  `recast_type` smallint(11) unsigned NOT NULL DEFAULT 0,
  `timestamp` int(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`recast_type`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "recast_type": 52,    "timestamp": 53,    "id": 76}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// CharacterItemRecast struct is a row record of the character_item_recast table in the eqemu database
type CharacterItemRecast struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ID uint32 `gorm:"primary_key;column:id;type:uint;default:0;" json:"id"`
	//[ 1] recast_type                                    usmallint            null: false  primary: true   isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	RecastType uint32 `gorm:"primary_key;column:recast_type;type:usmallint;default:0;" json:"recast_type"`
	//[ 2] timestamp                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Timestamp uint32 `gorm:"column:timestamp;type:uint;default:0;" json:"timestamp"`
}

var character_item_recastTableInfo = &TableInfo{
	Name: "character_item_recast",
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
			Name:               "recast_type",
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
			GoFieldName:        "RecastType",
			GoFieldType:        "uint32",
			JSONFieldName:      "recast_type",
			ProtobufFieldName:  "recast_type",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "timestamp",
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
			GoFieldName:        "Timestamp",
			GoFieldType:        "uint32",
			JSONFieldName:      "timestamp",
			ProtobufFieldName:  "timestamp",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterItemRecast) TableName() string {
	return "character_item_recast"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterItemRecast) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterItemRecast) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterItemRecast) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterItemRecast) TableInfo() *TableInfo {
	return character_item_recastTableInfo
}
