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


CREATE TABLE `char_create_combinations` (
  `allocation_id` int(10) unsigned NOT NULL,
  `race` int(10) unsigned NOT NULL,
  `class` int(10) unsigned NOT NULL,
  `deity` int(10) unsigned NOT NULL,
  `start_zone` int(10) unsigned NOT NULL,
  `expansions_req` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`race`,`class`,`deity`,`start_zone`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "deity": 74,    "start_zone": 78,    "expansions_req": 15,    "allocation_id": 64,    "race": 79,    "class": 64}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned



*/

// CharCreateCombinations struct is a row record of the char_create_combinations table in the eqemu database
type CharCreateCombinations struct {
	//[ 0] allocation_id                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AllocationID uint32 `gorm:"column:allocation_id;type:uint;" json:"allocation_id"`
	//[ 1] race                                           uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	Race uint32 `gorm:"primary_key;column:race;type:uint;" json:"race"`
	//[ 2] class                                          uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	Class uint32 `gorm:"primary_key;column:class;type:uint;" json:"class"`
	//[ 3] deity                                          uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	Deity uint32 `gorm:"primary_key;column:deity;type:uint;" json:"deity"`
	//[ 4] start_zone                                     uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	StartZone uint32 `gorm:"primary_key;column:start_zone;type:uint;" json:"start_zone"`
	//[ 5] expansions_req                                 uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ExpansionsReq uint32 `gorm:"column:expansions_req;type:uint;default:0;" json:"expansions_req"`
}

var char_create_combinationsTableInfo = &TableInfo{
	Name: "char_create_combinations",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "allocation_id",
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
			GoFieldName:        "AllocationID",
			GoFieldType:        "uint32",
			JSONFieldName:      "allocation_id",
			ProtobufFieldName:  "allocation_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "race",
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
			GoFieldName:        "Race",
			GoFieldType:        "uint32",
			JSONFieldName:      "race",
			ProtobufFieldName:  "race",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "class",
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
			GoFieldName:        "Class",
			GoFieldType:        "uint32",
			JSONFieldName:      "class",
			ProtobufFieldName:  "class",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "deity",
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
			GoFieldName:        "Deity",
			GoFieldType:        "uint32",
			JSONFieldName:      "deity",
			ProtobufFieldName:  "deity",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "start_zone",
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
			GoFieldName:        "StartZone",
			GoFieldType:        "uint32",
			JSONFieldName:      "start_zone",
			ProtobufFieldName:  "start_zone",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "expansions_req",
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
			GoFieldName:        "ExpansionsReq",
			GoFieldType:        "uint32",
			JSONFieldName:      "expansions_req",
			ProtobufFieldName:  "expansions_req",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharCreateCombinations) TableName() string {
	return "char_create_combinations"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharCreateCombinations) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharCreateCombinations) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharCreateCombinations) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharCreateCombinations) TableInfo() *TableInfo {
	return char_create_combinationsTableInfo
}
