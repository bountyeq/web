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


CREATE TABLE `loottable_entries` (
  `loottable_id` int(11) unsigned NOT NULL DEFAULT 0,
  `lootdrop_id` int(11) unsigned NOT NULL DEFAULT 0,
  `multiplier` tinyint(2) unsigned NOT NULL DEFAULT 1,
  `droplimit` tinyint(2) unsigned NOT NULL DEFAULT 0,
  `mindrop` tinyint(2) unsigned NOT NULL DEFAULT 0,
  `probability` float NOT NULL DEFAULT 100,
  PRIMARY KEY (`loottable_id`,`lootdrop_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "loottable_id": 70,    "lootdrop_id": 8,    "multiplier": 88,    "droplimit": 43,    "mindrop": 88,    "probability": 0.6389524}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned



*/

// LoottableEntries struct is a row record of the loottable_entries table in the eqemu database
type LoottableEntries struct {
	//[ 0] loottable_id                                   uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	LoottableID uint32 `gorm:"primary_key;column:loottable_id;type:uint;default:0;" json:"loottable_id"`
	//[ 1] lootdrop_id                                    uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	LootdropID uint32 `gorm:"primary_key;column:lootdrop_id;type:uint;default:0;" json:"lootdrop_id"`
	//[ 2] multiplier                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	Multiplier uint32 `gorm:"column:multiplier;type:utinyint;default:1;" json:"multiplier"`
	//[ 3] droplimit                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Droplimit uint32 `gorm:"column:droplimit;type:utinyint;default:0;" json:"droplimit"`
	//[ 4] mindrop                                        utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Mindrop uint32 `gorm:"column:mindrop;type:utinyint;default:0;" json:"mindrop"`
	//[ 5] probability                                    float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [100]
	Probability float32 `gorm:"column:probability;type:float;default:100;" json:"probability"`
}

var loottable_entriesTableInfo = &TableInfo{
	Name: "loottable_entries",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "loottable_id",
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
			GoFieldName:        "LoottableID",
			GoFieldType:        "uint32",
			JSONFieldName:      "loottable_id",
			ProtobufFieldName:  "loottable_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "lootdrop_id",
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
			GoFieldName:        "LootdropID",
			GoFieldType:        "uint32",
			JSONFieldName:      "lootdrop_id",
			ProtobufFieldName:  "lootdrop_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "multiplier",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Multiplier",
			GoFieldType:        "uint32",
			JSONFieldName:      "multiplier",
			ProtobufFieldName:  "multiplier",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "droplimit",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Droplimit",
			GoFieldType:        "uint32",
			JSONFieldName:      "droplimit",
			ProtobufFieldName:  "droplimit",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "mindrop",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Mindrop",
			GoFieldType:        "uint32",
			JSONFieldName:      "mindrop",
			ProtobufFieldName:  "mindrop",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "probability",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "Probability",
			GoFieldType:        "float32",
			JSONFieldName:      "probability",
			ProtobufFieldName:  "probability",
			ProtobufType:       "float",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LoottableEntries) TableName() string {
	return "loottable_entries"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LoottableEntries) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LoottableEntries) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LoottableEntries) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LoottableEntries) TableInfo() *TableInfo {
	return loottable_entriesTableInfo
}
