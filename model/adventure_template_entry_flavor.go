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


CREATE TABLE `adventure_template_entry_flavor` (
  `id` int(10) unsigned NOT NULL,
  `text` varchar(512) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `id_2` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "id": 76,    "text": "CHxojVwixKKQejUDcDSsdrOTZ"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// AdventureTemplateEntryFlavor struct is a row record of the adventure_template_entry_flavor table in the eqemu database
type AdventureTemplateEntryFlavor struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;column:id;type:uint;" json:"id"`
	//[ 1] text                                           varchar(512)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 512     default: []
	Text string `gorm:"column:text;type:varchar;size:512;" json:"text"`
}

var adventure_template_entry_flavorTableInfo = &TableInfo{
	Name: "adventure_template_entry_flavor",
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
			Name:               "text",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(512)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       512,
			GoFieldName:        "Text",
			GoFieldType:        "string",
			JSONFieldName:      "text",
			ProtobufFieldName:  "text",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AdventureTemplateEntryFlavor) TableName() string {
	return "adventure_template_entry_flavor"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AdventureTemplateEntryFlavor) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AdventureTemplateEntryFlavor) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AdventureTemplateEntryFlavor) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AdventureTemplateEntryFlavor) TableInfo() *TableInfo {
	return adventure_template_entry_flavorTableInfo
}
