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


CREATE TABLE `adventure_template_entry` (
  `id` int(10) unsigned NOT NULL,
  `template_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`,`template_id`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "id": 3,    "template_id": 10}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// AdventureTemplateEntry struct is a row record of the adventure_template_entry table in the eqemu database
type AdventureTemplateEntry struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;column:id;type:uint;" json:"id"`
	//[ 1] template_id                                    uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	TemplateID uint32 `gorm:"primary_key;column:template_id;type:uint;" json:"template_id"`
}

var adventure_template_entryTableInfo = &TableInfo{
	Name: "adventure_template_entry",
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
			Name:               "template_id",
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
			GoFieldName:        "TemplateID",
			GoFieldType:        "uint32",
			JSONFieldName:      "template_id",
			ProtobufFieldName:  "template_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AdventureTemplateEntry) TableName() string {
	return "adventure_template_entry"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AdventureTemplateEntry) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AdventureTemplateEntry) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AdventureTemplateEntry) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AdventureTemplateEntry) TableInfo() *TableInfo {
	return adventure_template_entryTableInfo
}
