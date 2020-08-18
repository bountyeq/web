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


CREATE TABLE `trader` (
  `char_id` int(10) unsigned NOT NULL DEFAULT 0,
  `item_id` int(10) unsigned NOT NULL DEFAULT 0,
  `serialnumber` int(10) unsigned NOT NULL DEFAULT 0,
  `charges` int(11) NOT NULL DEFAULT 0,
  `item_cost` int(10) unsigned NOT NULL DEFAULT 0,
  `slot_id` tinyint(3) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`char_id`,`slot_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "char_id": 67,    "item_id": 64,    "serialnumber": 77,    "charges": 2,    "item_cost": 62,    "slot_id": 41}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned



*/

// Trader struct is a row record of the trader table in the eqemu database
type Trader struct {
	//[ 0] char_id                                        uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	CharID uint32 `gorm:"primary_key;column:char_id;type:uint;default:0;" json:"char_id"`
	//[ 1] item_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ItemID uint32 `gorm:"column:item_id;type:uint;default:0;" json:"item_id"`
	//[ 2] serialnumber                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Serialnumber uint32 `gorm:"column:serialnumber;type:uint;default:0;" json:"serialnumber"`
	//[ 3] charges                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Charges int32 `gorm:"column:charges;type:int;default:0;" json:"charges"`
	//[ 4] item_cost                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ItemCost uint32 `gorm:"column:item_cost;type:uint;default:0;" json:"item_cost"`
	//[ 5] slot_id                                        utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	SlotID uint32 `gorm:"primary_key;column:slot_id;type:utinyint;default:0;" json:"slot_id"`
}

var traderTableInfo = &TableInfo{
	Name: "trader",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "char_id",
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
			GoFieldName:        "CharID",
			GoFieldType:        "uint32",
			JSONFieldName:      "char_id",
			ProtobufFieldName:  "char_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "serialnumber",
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
			GoFieldName:        "Serialnumber",
			GoFieldType:        "uint32",
			JSONFieldName:      "serialnumber",
			ProtobufFieldName:  "serialnumber",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "charges",
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
			GoFieldName:        "Charges",
			GoFieldType:        "int32",
			JSONFieldName:      "charges",
			ProtobufFieldName:  "charges",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "item_cost",
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
			GoFieldName:        "ItemCost",
			GoFieldType:        "uint32",
			JSONFieldName:      "item_cost",
			ProtobufFieldName:  "item_cost",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "slot_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "SlotID",
			GoFieldType:        "uint32",
			JSONFieldName:      "slot_id",
			ProtobufFieldName:  "slot_id",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *Trader) TableName() string {
	return "trader"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *Trader) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *Trader) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *Trader) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *Trader) TableInfo() *TableInfo {
	return traderTableInfo
}
