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


CREATE TABLE `pets_equipmentset` (
  `set_id` int(11) NOT NULL,
  `setname` varchar(30) NOT NULL DEFAULT '',
  `nested_set` int(11) NOT NULL DEFAULT -1,
  PRIMARY KEY (`set_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "set_id": 75,    "setname": "pARlsIApfVpYellZrEjYDpgwW",    "nested_set": 11}



*/

// PetsEquipmentset struct is a row record of the pets_equipmentset table in the eqemu database
type PetsEquipmentset struct {
	//[ 0] set_id                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SetID int32 `gorm:"primary_key;column:set_id;type:int;" json:"set_id"`
	//[ 1] setname                                        varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: ['']
	Setname string `gorm:"column:setname;type:varchar;size:30;default:'';" json:"setname"`
	//[ 2] nested_set                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	NestedSet int32 `gorm:"column:nested_set;type:int;default:-1;" json:"nested_set"`
}

var pets_equipmentsetTableInfo = &TableInfo{
	Name: "pets_equipmentset",
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
			Name:               "setname",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "Setname",
			GoFieldType:        "string",
			JSONFieldName:      "setname",
			ProtobufFieldName:  "setname",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "nested_set",
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
			GoFieldName:        "NestedSet",
			GoFieldType:        "int32",
			JSONFieldName:      "nested_set",
			ProtobufFieldName:  "nested_set",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PetsEquipmentset) TableName() string {
	return "pets_equipmentset"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PetsEquipmentset) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PetsEquipmentset) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PetsEquipmentset) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PetsEquipmentset) TableInfo() *TableInfo {
	return pets_equipmentsetTableInfo
}
