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


CREATE TABLE `tribute_levels` (
  `tribute_id` int(10) unsigned NOT NULL DEFAULT 0,
  `level` int(10) unsigned NOT NULL DEFAULT 0,
  `cost` int(10) unsigned NOT NULL DEFAULT 0,
  `item_id` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`tribute_id`,`level`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "tribute_id": 54,    "level": 58,    "cost": 51,    "item_id": 11}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned



*/

// TributeLevels struct is a row record of the tribute_levels table in the eqemu database
type TributeLevels struct {
	//[ 0] tribute_id                                     uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	TributeID uint32 `gorm:"primary_key;column:tribute_id;type:uint;default:0;" json:"tribute_id"`
	//[ 1] level                                          uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Level uint32 `gorm:"primary_key;column:level;type:uint;default:0;" json:"level"`
	//[ 2] cost                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Cost uint32 `gorm:"column:cost;type:uint;default:0;" json:"cost"`
	//[ 3] item_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ItemID uint32 `gorm:"column:item_id;type:uint;default:0;" json:"item_id"`
}

var tribute_levelsTableInfo = &TableInfo{
	Name: "tribute_levels",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "tribute_id",
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
			GoFieldName:        "TributeID",
			GoFieldType:        "uint32",
			JSONFieldName:      "tribute_id",
			ProtobufFieldName:  "tribute_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "level",
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
			GoFieldName:        "Level",
			GoFieldType:        "uint32",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "cost",
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
			GoFieldName:        "Cost",
			GoFieldType:        "uint32",
			JSONFieldName:      "cost",
			ProtobufFieldName:  "cost",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "item_id",
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
			GoFieldName:        "ItemID",
			GoFieldType:        "uint32",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TributeLevels) TableName() string {
	return "tribute_levels"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TributeLevels) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TributeLevels) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TributeLevels) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TributeLevels) TableInfo() *TableInfo {
	return tribute_levelsTableInfo
}
