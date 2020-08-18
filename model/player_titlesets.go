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


CREATE TABLE `player_titlesets` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `char_id` int(11) unsigned NOT NULL,
  `title_set` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "id": 74,    "char_id": 68,    "title_set": 65}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// PlayerTitlesets struct is a row record of the player_titlesets table in the eqemu database
type PlayerTitlesets struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] char_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	CharID uint32 `gorm:"column:char_id;type:uint;" json:"char_id"`
	//[ 2] title_set                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	TitleSet uint32 `gorm:"column:title_set;type:uint;" json:"title_set"`
}

var player_titlesetsTableInfo = &TableInfo{
	Name: "player_titlesets",
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
			Name:               "char_id",
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
			GoFieldName:        "CharID",
			GoFieldType:        "uint32",
			JSONFieldName:      "char_id",
			ProtobufFieldName:  "char_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "title_set",
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
			GoFieldName:        "TitleSet",
			GoFieldType:        "uint32",
			JSONFieldName:      "title_set",
			ProtobufFieldName:  "title_set",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PlayerTitlesets) TableName() string {
	return "player_titlesets"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PlayerTitlesets) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PlayerTitlesets) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PlayerTitlesets) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PlayerTitlesets) TableInfo() *TableInfo {
	return player_titlesetsTableInfo
}
