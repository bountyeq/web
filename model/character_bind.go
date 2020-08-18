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


CREATE TABLE `character_bind` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `slot` int(4) NOT NULL DEFAULT 0,
  `zone_id` smallint(11) unsigned NOT NULL DEFAULT 0,
  `instance_id` mediumint(11) unsigned NOT NULL DEFAULT 0,
  `x` float NOT NULL DEFAULT 0,
  `y` float NOT NULL DEFAULT 0,
  `z` float NOT NULL DEFAULT 0,
  `heading` float NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`slot`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "zone_id": 86,    "instance_id": 60,    "x": 0.18543644,    "y": 0.25970104,    "z": 0.52253723,    "heading": 0.5898753,    "id": 35,    "slot": 55}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned



*/

// CharacterBind struct is a row record of the character_bind table in the eqemu database
type CharacterBind struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] slot                                           int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Slot int32 `gorm:"primary_key;column:slot;type:int;default:0;" json:"slot"`
	//[ 2] zone_id                                        usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	ZoneID uint32 `gorm:"column:zone_id;type:usmallint;default:0;" json:"zone_id"`
	//[ 3] instance_id                                    umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	InstanceID uint32 `gorm:"column:instance_id;type:umediumint;default:0;" json:"instance_id"`
	//[ 4] x                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	X float32 `gorm:"column:x;type:float;default:0;" json:"x"`
	//[ 5] y                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Y float32 `gorm:"column:y;type:float;default:0;" json:"y"`
	//[ 6] z                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Z float32 `gorm:"column:z;type:float;default:0;" json:"z"`
	//[ 7] heading                                        float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Heading float32 `gorm:"column:heading;type:float;default:0;" json:"heading"`
}

var character_bindTableInfo = &TableInfo{
	Name: "character_bind",
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
			IsAutoIncrement:    true,
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
			Name:               "zone_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "ZoneID",
			GoFieldType:        "uint32",
			JSONFieldName:      "zone_id",
			ProtobufFieldName:  "zone_id",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "instance_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "umediumint",
			DatabaseTypePretty: "umediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "umediumint",
			ColumnLength:       -1,
			GoFieldName:        "InstanceID",
			GoFieldType:        "uint32",
			JSONFieldName:      "instance_id",
			ProtobufFieldName:  "instance_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "x",
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
			GoFieldName:        "X",
			GoFieldType:        "float32",
			JSONFieldName:      "x",
			ProtobufFieldName:  "x",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "y",
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
			GoFieldName:        "Y",
			GoFieldType:        "float32",
			JSONFieldName:      "y",
			ProtobufFieldName:  "y",
			ProtobufType:       "float",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "z",
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
			GoFieldName:        "Z",
			GoFieldType:        "float32",
			JSONFieldName:      "z",
			ProtobufFieldName:  "z",
			ProtobufType:       "float",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "heading",
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
			GoFieldName:        "Heading",
			GoFieldType:        "float32",
			JSONFieldName:      "heading",
			ProtobufFieldName:  "heading",
			ProtobufType:       "float",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterBind) TableName() string {
	return "character_bind"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterBind) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterBind) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterBind) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterBind) TableInfo() *TableInfo {
	return character_bindTableInfo
}
