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


CREATE TABLE `inventory` (
  `charid` int(11) unsigned NOT NULL DEFAULT 0,
  `slotid` mediumint(7) unsigned NOT NULL DEFAULT 0,
  `itemid` int(11) unsigned DEFAULT 0,
  `charges` smallint(3) unsigned DEFAULT 0,
  `color` int(11) unsigned NOT NULL DEFAULT 0,
  `augslot1` mediumint(7) unsigned NOT NULL DEFAULT 0,
  `augslot2` mediumint(7) unsigned NOT NULL DEFAULT 0,
  `augslot3` mediumint(7) unsigned NOT NULL DEFAULT 0,
  `augslot4` mediumint(7) unsigned NOT NULL DEFAULT 0,
  `augslot5` mediumint(7) unsigned DEFAULT 0,
  `augslot6` mediumint(7) NOT NULL DEFAULT 0,
  `instnodrop` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `custom_data` text DEFAULT NULL,
  `ornamenticon` int(11) unsigned NOT NULL DEFAULT 0,
  `ornamentidfile` int(11) unsigned NOT NULL DEFAULT 0,
  `ornament_hero_model` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`charid`,`slotid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "augslot_1": 58,    "instnodrop": 48,    "charid": 13,    "slotid": 68,    "charges": 72,    "augslot_6": 48,    "custom_data": "VIlhvvVWtTJWkUCnbyLHyFRDg",    "ornament_hero_model": 41,    "augslot_5": 81,    "ornamentidfile": 71,    "itemid": 89,    "color": 14,    "augslot_2": 44,    "augslot_3": 52,    "augslot_4": 63,    "ornamenticon": 52}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[11] column is set for unsigned
[13] column is set for unsigned
[14] column is set for unsigned



*/

// Inventory struct is a row record of the inventory table in the eqemu database
type Inventory struct {
	//[ 0] charid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Charid uint32 `gorm:"primary_key;column:charid;type:uint;default:0;" json:"charid"`
	//[ 1] slotid                                         umediumint           null: false  primary: true   isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Slotid uint32 `gorm:"primary_key;column:slotid;type:umediumint;default:0;" json:"slotid"`
	//[ 2] itemid                                         uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Itemid null.Int `gorm:"column:itemid;type:uint;default:0;" json:"itemid"`
	//[ 3] charges                                        usmallint            null: true   primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	Charges null.Int `gorm:"column:charges;type:usmallint;default:0;" json:"charges"`
	//[ 4] color                                          uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Color uint32 `gorm:"column:color;type:uint;default:0;" json:"color"`
	//[ 5] augslot1                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot1 uint32 `gorm:"column:augslot1;type:umediumint;default:0;" json:"augslot_1"`
	//[ 6] augslot2                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot2 uint32 `gorm:"column:augslot2;type:umediumint;default:0;" json:"augslot_2"`
	//[ 7] augslot3                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot3 uint32 `gorm:"column:augslot3;type:umediumint;default:0;" json:"augslot_3"`
	//[ 8] augslot4                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot4 uint32 `gorm:"column:augslot4;type:umediumint;default:0;" json:"augslot_4"`
	//[ 9] augslot5                                       umediumint           null: true   primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot5 null.Int `gorm:"column:augslot5;type:umediumint;default:0;" json:"augslot_5"`
	//[10] augslot6                                       mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	Augslot6 int32 `gorm:"column:augslot6;type:mediumint;default:0;" json:"augslot_6"`
	//[11] instnodrop                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Instnodrop uint32 `gorm:"column:instnodrop;type:utinyint;default:0;" json:"instnodrop"`
	//[12] custom_data                                    text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: [NULL]
	CustomData null.String `gorm:"column:custom_data;type:text;size:65535;" json:"custom_data"`
	//[13] ornamenticon                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Ornamenticon uint32 `gorm:"column:ornamenticon;type:uint;default:0;" json:"ornamenticon"`
	//[14] ornamentidfile                                 uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Ornamentidfile uint32 `gorm:"column:ornamentidfile;type:uint;default:0;" json:"ornamentidfile"`
	//[15] ornament_hero_model                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	OrnamentHeroModel int32 `gorm:"column:ornament_hero_model;type:int;default:0;" json:"ornament_hero_model"`
}

var inventoryTableInfo = &TableInfo{
	Name: "inventory",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "charid",
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
			GoFieldName:        "Charid",
			GoFieldType:        "uint32",
			JSONFieldName:      "charid",
			ProtobufFieldName:  "charid",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "slotid",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "umediumint",
			DatabaseTypePretty: "umediumint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "umediumint",
			ColumnLength:       -1,
			GoFieldName:        "Slotid",
			GoFieldType:        "uint32",
			JSONFieldName:      "slotid",
			ProtobufFieldName:  "slotid",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "itemid",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Itemid",
			GoFieldType:        "null.Int",
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
			Nullable:           true,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "Charges",
			GoFieldType:        "null.Int",
			JSONFieldName:      "charges",
			ProtobufFieldName:  "charges",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "color",
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
			GoFieldName:        "Color",
			GoFieldType:        "uint32",
			JSONFieldName:      "color",
			ProtobufFieldName:  "color",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "augslot1",
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
			GoFieldName:        "Augslot1",
			GoFieldType:        "uint32",
			JSONFieldName:      "augslot_1",
			ProtobufFieldName:  "augslot_1",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "augslot2",
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
			GoFieldName:        "Augslot2",
			GoFieldType:        "uint32",
			JSONFieldName:      "augslot_2",
			ProtobufFieldName:  "augslot_2",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "augslot3",
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
			GoFieldName:        "Augslot3",
			GoFieldType:        "uint32",
			JSONFieldName:      "augslot_3",
			ProtobufFieldName:  "augslot_3",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "augslot4",
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
			GoFieldName:        "Augslot4",
			GoFieldType:        "uint32",
			JSONFieldName:      "augslot_4",
			ProtobufFieldName:  "augslot_4",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "augslot5",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "umediumint",
			DatabaseTypePretty: "umediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "umediumint",
			ColumnLength:       -1,
			GoFieldName:        "Augslot5",
			GoFieldType:        "null.Int",
			JSONFieldName:      "augslot_5",
			ProtobufFieldName:  "augslot_5",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "augslot6",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "mediumint",
			DatabaseTypePretty: "mediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "mediumint",
			ColumnLength:       -1,
			GoFieldName:        "Augslot6",
			GoFieldType:        "int32",
			JSONFieldName:      "augslot_6",
			ProtobufFieldName:  "augslot_6",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "instnodrop",
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
			GoFieldName:        "Instnodrop",
			GoFieldType:        "uint32",
			JSONFieldName:      "instnodrop",
			ProtobufFieldName:  "instnodrop",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "custom_data",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "CustomData",
			GoFieldType:        "null.String",
			JSONFieldName:      "custom_data",
			ProtobufFieldName:  "custom_data",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "ornamenticon",
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
			GoFieldName:        "Ornamenticon",
			GoFieldType:        "uint32",
			JSONFieldName:      "ornamenticon",
			ProtobufFieldName:  "ornamenticon",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "ornamentidfile",
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
			GoFieldName:        "Ornamentidfile",
			GoFieldType:        "uint32",
			JSONFieldName:      "ornamentidfile",
			ProtobufFieldName:  "ornamentidfile",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "ornament_hero_model",
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
			GoFieldName:        "OrnamentHeroModel",
			GoFieldType:        "int32",
			JSONFieldName:      "ornament_hero_model",
			ProtobufFieldName:  "ornament_hero_model",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (i *Inventory) TableName() string {
	return "inventory"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *Inventory) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *Inventory) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *Inventory) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *Inventory) TableInfo() *TableInfo {
	return inventoryTableInfo
}
