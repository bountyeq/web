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


CREATE TABLE `character_pet_inventory` (
  `char_id` int(11) NOT NULL,
  `pet` int(11) NOT NULL,
  `slot` int(11) NOT NULL,
  `item_id` int(11) NOT NULL,
  PRIMARY KEY (`char_id`,`pet`,`slot`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "char_id": 10,    "pet": 92,    "slot": 64,    "item_id": 57}



*/

// CharacterPetInventory struct is a row record of the character_pet_inventory table in the eqemu database
type CharacterPetInventory struct {
	//[ 0] char_id                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	CharID int32 `gorm:"primary_key;column:char_id;type:int;" json:"char_id"`
	//[ 1] pet                                            int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	Pet int32 `gorm:"primary_key;column:pet;type:int;" json:"pet"`
	//[ 2] slot                                           int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	Slot int32 `gorm:"primary_key;column:slot;type:int;" json:"slot"`
	//[ 3] item_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ItemID int32 `gorm:"column:item_id;type:int;" json:"item_id"`
}

var character_pet_inventoryTableInfo = &TableInfo{
	Name: "character_pet_inventory",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "char_id",
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
			GoFieldName:        "CharID",
			GoFieldType:        "int32",
			JSONFieldName:      "char_id",
			ProtobufFieldName:  "char_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "pet",
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
			GoFieldName:        "Pet",
			GoFieldType:        "int32",
			JSONFieldName:      "pet",
			ProtobufFieldName:  "pet",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterPetInventory) TableName() string {
	return "character_pet_inventory"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterPetInventory) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterPetInventory) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterPetInventory) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterPetInventory) TableInfo() *TableInfo {
	return character_pet_inventoryTableInfo
}
