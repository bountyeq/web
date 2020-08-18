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


CREATE TABLE `merchantlist_temp` (
  `npcid` int(10) unsigned NOT NULL DEFAULT 0,
  `slot` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `itemid` int(10) unsigned NOT NULL DEFAULT 0,
  `charges` int(10) unsigned NOT NULL DEFAULT 1,
  PRIMARY KEY (`npcid`,`slot`),
  KEY `npcid_2` (`npcid`,`itemid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "npcid": 85,    "slot": 6,    "itemid": 48,    "charges": 51}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned



*/

// MerchantlistTemp struct is a row record of the merchantlist_temp table in the eqemu database
type MerchantlistTemp struct {
	//[ 0] npcid                                          uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Npcid uint32 `gorm:"primary_key;column:npcid;type:uint;default:0;" json:"npcid"`
	//[ 1] slot                                           utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Slot uint32 `gorm:"primary_key;column:slot;type:utinyint;default:0;" json:"slot"`
	//[ 2] itemid                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Itemid uint32 `gorm:"column:itemid;type:uint;default:0;" json:"itemid"`
	//[ 3] charges                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	Charges uint32 `gorm:"column:charges;type:uint;default:1;" json:"charges"`
}

var merchantlist_tempTableInfo = &TableInfo{
	Name: "merchantlist_temp",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "npcid",
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
			GoFieldName:        "Npcid",
			GoFieldType:        "uint32",
			JSONFieldName:      "npcid",
			ProtobufFieldName:  "npcid",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "slot",
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
			GoFieldName:        "Slot",
			GoFieldType:        "uint32",
			JSONFieldName:      "slot",
			ProtobufFieldName:  "slot",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "itemid",
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
			GoFieldName:        "Itemid",
			GoFieldType:        "uint32",
			JSONFieldName:      "itemid",
			ProtobufFieldName:  "itemid",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "charges",
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
			GoFieldName:        "Charges",
			GoFieldType:        "uint32",
			JSONFieldName:      "charges",
			ProtobufFieldName:  "charges",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *MerchantlistTemp) TableName() string {
	return "merchantlist_temp"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MerchantlistTemp) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MerchantlistTemp) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MerchantlistTemp) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MerchantlistTemp) TableInfo() *TableInfo {
	return merchantlist_tempTableInfo
}
