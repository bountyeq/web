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


CREATE TABLE `inventory_snapshots` (
  `time_index` int(11) unsigned NOT NULL DEFAULT 0,
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
  PRIMARY KEY (`time_index`,`charid`,`slotid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "augslot_4": 21,    "custom_data": "DCNfmCyujFggcIekEmeeROLCv",    "charges": 52,    "charid": 57,    "itemid": 13,    "augslot_5": 38,    "augslot_6": 0,    "ornament_hero_model": 14,    "time_index": 57,    "color": 25,    "augslot_1": 72,    "augslot_2": 44,    "augslot_3": 46,    "instnodrop": 2,    "ornamenticon": 25,    "ornamentidfile": 67,    "slotid": 40}


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
[10] column is set for unsigned
[12] column is set for unsigned
[14] column is set for unsigned
[15] column is set for unsigned



*/

// InventorySnapshots struct is a row record of the inventory_snapshots table in the eqemu database
type InventorySnapshots struct {
	//[ 0] time_index                                     uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	TimeIndex uint32 `gorm:"primary_key;column:time_index;type:uint;default:0;" json:"time_index"`
	//[ 1] charid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Charid uint32 `gorm:"primary_key;column:charid;type:uint;default:0;" json:"charid"`
	//[ 2] slotid                                         umediumint           null: false  primary: true   isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Slotid uint32 `gorm:"primary_key;column:slotid;type:umediumint;default:0;" json:"slotid"`
	//[ 3] itemid                                         uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Itemid null.Int `gorm:"column:itemid;type:uint;default:0;" json:"itemid"`
	//[ 4] charges                                        usmallint            null: true   primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	Charges null.Int `gorm:"column:charges;type:usmallint;default:0;" json:"charges"`
	//[ 5] color                                          uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Color uint32 `gorm:"column:color;type:uint;default:0;" json:"color"`
	//[ 6] augslot1                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot1 uint32 `gorm:"column:augslot1;type:umediumint;default:0;" json:"augslot_1"`
	//[ 7] augslot2                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot2 uint32 `gorm:"column:augslot2;type:umediumint;default:0;" json:"augslot_2"`
	//[ 8] augslot3                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot3 uint32 `gorm:"column:augslot3;type:umediumint;default:0;" json:"augslot_3"`
	//[ 9] augslot4                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot4 uint32 `gorm:"column:augslot4;type:umediumint;default:0;" json:"augslot_4"`
	//[10] augslot5                                       umediumint           null: true   primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot5 null.Int `gorm:"column:augslot5;type:umediumint;default:0;" json:"augslot_5"`
	//[11] augslot6                                       mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	Augslot6 int32 `gorm:"column:augslot6;type:mediumint;default:0;" json:"augslot_6"`
	//[12] instnodrop                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Instnodrop uint32 `gorm:"column:instnodrop;type:utinyint;default:0;" json:"instnodrop"`
	//[13] custom_data                                    text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: [NULL]
	CustomData null.String `gorm:"column:custom_data;type:text;size:65535;" json:"custom_data"`
	//[14] ornamenticon                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Ornamenticon uint32 `gorm:"column:ornamenticon;type:uint;default:0;" json:"ornamenticon"`
	//[15] ornamentidfile                                 uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Ornamentidfile uint32 `gorm:"column:ornamentidfile;type:uint;default:0;" json:"ornamentidfile"`
	//[16] ornament_hero_model                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	OrnamentHeroModel int32 `gorm:"column:ornament_hero_model;type:int;default:0;" json:"ornament_hero_model"`
}

var inventory_snapshotsTableInfo = &TableInfo{
	Name: "inventory_snapshots",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "time_index",
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
			GoFieldName:        "TimeIndex",
			GoFieldType:        "uint32",
			JSONFieldName:      "time_index",
			ProtobufFieldName:  "time_index",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},
	},
}

// TableName sets the insert table name for this struct type
func (i *InventorySnapshots) TableName() string {
	return "inventory_snapshots"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *InventorySnapshots) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *InventorySnapshots) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *InventorySnapshots) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *InventorySnapshots) TableInfo() *TableInfo {
	return inventory_snapshotsTableInfo
}
