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


CREATE TABLE `ldon_trap_entries` (
  `id` int(10) unsigned NOT NULL,
  `trap_id` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`trap_id`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "id": 74,    "trap_id": 75}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// LdonTrapEntries struct is a row record of the ldon_trap_entries table in the eqemu database
type LdonTrapEntries struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;column:id;type:uint;" json:"id"`
	//[ 1] trap_id                                        uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	TrapID uint32 `gorm:"primary_key;column:trap_id;type:uint;default:0;" json:"trap_id"`
}

var ldon_trap_entriesTableInfo = &TableInfo{
	Name: "ldon_trap_entries",
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
			Name:               "trap_id",
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
			GoFieldName:        "TrapID",
			GoFieldType:        "uint32",
			JSONFieldName:      "trap_id",
			ProtobufFieldName:  "trap_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LdonTrapEntries) TableName() string {
	return "ldon_trap_entries"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LdonTrapEntries) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LdonTrapEntries) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LdonTrapEntries) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LdonTrapEntries) TableInfo() *TableInfo {
	return ldon_trap_entriesTableInfo
}
