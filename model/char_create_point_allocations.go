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


CREATE TABLE `char_create_point_allocations` (
  `id` int(10) unsigned NOT NULL,
  `base_str` int(10) unsigned NOT NULL,
  `base_sta` int(10) unsigned NOT NULL,
  `base_dex` int(10) unsigned NOT NULL,
  `base_agi` int(10) unsigned NOT NULL,
  `base_int` int(10) unsigned NOT NULL,
  `base_wis` int(10) unsigned NOT NULL,
  `base_cha` int(10) unsigned NOT NULL,
  `alloc_str` int(10) unsigned NOT NULL,
  `alloc_sta` int(10) unsigned NOT NULL,
  `alloc_dex` int(10) unsigned NOT NULL,
  `alloc_agi` int(10) unsigned NOT NULL,
  `alloc_int` int(10) unsigned NOT NULL,
  `alloc_wis` int(10) unsigned NOT NULL,
  `alloc_cha` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "base_int": 12,    "alloc_sta": 87,    "alloc_dex": 52,    "alloc_int": 88,    "alloc_str": 68,    "alloc_wis": 79,    "alloc_cha": 50,    "base_sta": 9,    "base_dex": 79,    "base_wis": 65,    "base_cha": 96,    "id": 76,    "base_str": 10,    "base_agi": 5,    "alloc_agi": 7}


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
[11] column is set for unsigned
[12] column is set for unsigned
[13] column is set for unsigned
[14] column is set for unsigned



*/

// CharCreatePointAllocations struct is a row record of the char_create_point_allocations table in the eqemu database
type CharCreatePointAllocations struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;column:id;type:uint;" json:"id"`
	//[ 1] base_str                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	BaseStr uint32 `gorm:"column:base_str;type:uint;" json:"base_str"`
	//[ 2] base_sta                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	BaseSta uint32 `gorm:"column:base_sta;type:uint;" json:"base_sta"`
	//[ 3] base_dex                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	BaseDex uint32 `gorm:"column:base_dex;type:uint;" json:"base_dex"`
	//[ 4] base_agi                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	BaseAgi uint32 `gorm:"column:base_agi;type:uint;" json:"base_agi"`
	//[ 5] base_int                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	BaseInt uint32 `gorm:"column:base_int;type:uint;" json:"base_int"`
	//[ 6] base_wis                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	BaseWis uint32 `gorm:"column:base_wis;type:uint;" json:"base_wis"`
	//[ 7] base_cha                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	BaseCha uint32 `gorm:"column:base_cha;type:uint;" json:"base_cha"`
	//[ 8] alloc_str                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AllocStr uint32 `gorm:"column:alloc_str;type:uint;" json:"alloc_str"`
	//[ 9] alloc_sta                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AllocSta uint32 `gorm:"column:alloc_sta;type:uint;" json:"alloc_sta"`
	//[10] alloc_dex                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AllocDex uint32 `gorm:"column:alloc_dex;type:uint;" json:"alloc_dex"`
	//[11] alloc_agi                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AllocAgi uint32 `gorm:"column:alloc_agi;type:uint;" json:"alloc_agi"`
	//[12] alloc_int                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AllocInt uint32 `gorm:"column:alloc_int;type:uint;" json:"alloc_int"`
	//[13] alloc_wis                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AllocWis uint32 `gorm:"column:alloc_wis;type:uint;" json:"alloc_wis"`
	//[14] alloc_cha                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AllocCha uint32 `gorm:"column:alloc_cha;type:uint;" json:"alloc_cha"`
}

var char_create_point_allocationsTableInfo = &TableInfo{
	Name: "char_create_point_allocations",
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
			IsAutoIncrement:    false,
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
			Name:               "base_str",
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
			GoFieldName:        "BaseStr",
			GoFieldType:        "uint32",
			JSONFieldName:      "base_str",
			ProtobufFieldName:  "base_str",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "base_sta",
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
			GoFieldName:        "BaseSta",
			GoFieldType:        "uint32",
			JSONFieldName:      "base_sta",
			ProtobufFieldName:  "base_sta",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "base_dex",
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
			GoFieldName:        "BaseDex",
			GoFieldType:        "uint32",
			JSONFieldName:      "base_dex",
			ProtobufFieldName:  "base_dex",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "base_agi",
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
			GoFieldName:        "BaseAgi",
			GoFieldType:        "uint32",
			JSONFieldName:      "base_agi",
			ProtobufFieldName:  "base_agi",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "base_int",
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
			GoFieldName:        "BaseInt",
			GoFieldType:        "uint32",
			JSONFieldName:      "base_int",
			ProtobufFieldName:  "base_int",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "base_wis",
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
			GoFieldName:        "BaseWis",
			GoFieldType:        "uint32",
			JSONFieldName:      "base_wis",
			ProtobufFieldName:  "base_wis",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "base_cha",
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
			GoFieldName:        "BaseCha",
			GoFieldType:        "uint32",
			JSONFieldName:      "base_cha",
			ProtobufFieldName:  "base_cha",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "alloc_str",
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
			GoFieldName:        "AllocStr",
			GoFieldType:        "uint32",
			JSONFieldName:      "alloc_str",
			ProtobufFieldName:  "alloc_str",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "alloc_sta",
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
			GoFieldName:        "AllocSta",
			GoFieldType:        "uint32",
			JSONFieldName:      "alloc_sta",
			ProtobufFieldName:  "alloc_sta",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "alloc_dex",
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
			GoFieldName:        "AllocDex",
			GoFieldType:        "uint32",
			JSONFieldName:      "alloc_dex",
			ProtobufFieldName:  "alloc_dex",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "alloc_agi",
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
			GoFieldName:        "AllocAgi",
			GoFieldType:        "uint32",
			JSONFieldName:      "alloc_agi",
			ProtobufFieldName:  "alloc_agi",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "alloc_int",
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
			GoFieldName:        "AllocInt",
			GoFieldType:        "uint32",
			JSONFieldName:      "alloc_int",
			ProtobufFieldName:  "alloc_int",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "alloc_wis",
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
			GoFieldName:        "AllocWis",
			GoFieldType:        "uint32",
			JSONFieldName:      "alloc_wis",
			ProtobufFieldName:  "alloc_wis",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "alloc_cha",
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
			GoFieldName:        "AllocCha",
			GoFieldType:        "uint32",
			JSONFieldName:      "alloc_cha",
			ProtobufFieldName:  "alloc_cha",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharCreatePointAllocations) TableName() string {
	return "char_create_point_allocations"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharCreatePointAllocations) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharCreatePointAllocations) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharCreatePointAllocations) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharCreatePointAllocations) TableInfo() *TableInfo {
	return char_create_point_allocationsTableInfo
}
