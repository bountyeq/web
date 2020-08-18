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


CREATE TABLE `pets_equipmentset_entries` (
  `set_id` int(11) NOT NULL,
  `slot` int(11) NOT NULL,
  `item_id` int(11) NOT NULL,
  PRIMARY KEY (`set_id`,`slot`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "set_id": 57,    "slot": 75,    "item_id": 30}



*/

// PetsEquipmentsetEntries struct is a row record of the pets_equipmentset_entries table in the eqemu database
type PetsEquipmentsetEntries struct {
	//[ 0] set_id                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SetID int32 `gorm:"primary_key;column:set_id;type:int;" json:"set_id"`
	//[ 1] slot                                           int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	Slot int32 `gorm:"primary_key;column:slot;type:int;" json:"slot"`
	//[ 2] item_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ItemID int32 `gorm:"column:item_id;type:int;" json:"item_id"`
}

var pets_equipmentset_entriesTableInfo = &TableInfo{
	Name: "pets_equipmentset_entries",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "set_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SetID",
			GoFieldType:        "int32",
			JSONFieldName:      "set_id",
			ProtobufFieldName:  "set_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "slot",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Slot",
			GoFieldType:        "int32",
			JSONFieldName:      "slot",
			ProtobufFieldName:  "slot",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "item_id",
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
			GoFieldName:        "ItemID",
			GoFieldType:        "int32",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PetsEquipmentsetEntries) TableName() string {
	return "pets_equipmentset_entries"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PetsEquipmentsetEntries) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PetsEquipmentsetEntries) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PetsEquipmentsetEntries) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PetsEquipmentsetEntries) TableInfo() *TableInfo {
	return pets_equipmentset_entriesTableInfo
}
