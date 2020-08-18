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


CREATE TABLE `damageshieldtypes` (
  `spellid` int(10) unsigned NOT NULL DEFAULT 0,
  `type` tinyint(3) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`spellid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "type": 93,    "spellid": 97}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// Damageshieldtypes struct is a row record of the damageshieldtypes table in the eqemu database
type Damageshieldtypes struct {
	//[ 0] spellid                                        uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Spellid uint32 `gorm:"primary_key;column:spellid;type:uint;default:0;" json:"spellid"`
	//[ 1] type                                           utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Type uint32 `gorm:"column:type;type:utinyint;default:0;" json:"type"`
}

var damageshieldtypesTableInfo = &TableInfo{
	Name: "damageshieldtypes",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "spellid",
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
			GoFieldName:        "Spellid",
			GoFieldType:        "uint32",
			JSONFieldName:      "spellid",
			ProtobufFieldName:  "spellid",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "uint32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *Damageshieldtypes) TableName() string {
	return "damageshieldtypes"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *Damageshieldtypes) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *Damageshieldtypes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *Damageshieldtypes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *Damageshieldtypes) TableInfo() *TableInfo {
	return damageshieldtypesTableInfo
}
