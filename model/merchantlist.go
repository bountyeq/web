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


CREATE TABLE `merchantlist` (
  `merchantid` int(11) NOT NULL DEFAULT 0,
  `slot` int(11) NOT NULL DEFAULT 0,
  `item` int(11) NOT NULL DEFAULT 0,
  `faction_required` smallint(6) NOT NULL DEFAULT -100,
  `level_required` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `alt_currency_cost` smallint(5) unsigned NOT NULL DEFAULT 0,
  `classes_required` int(11) NOT NULL DEFAULT 65535,
  `probability` int(3) NOT NULL DEFAULT 100,
  `min_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `max_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `content_flags` varchar(100) DEFAULT NULL,
  `content_flags_disabled` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`merchantid`,`slot`),
  UNIQUE KEY `merchantid` (`merchantid`,`item`),
  KEY `item` (`item`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "merchantid": 51,    "item": 26,    "probability": 73,    "min_expansion": 74,    "content_flags": "ngebCHioicMWedDHYjFSJhHPT",    "content_flags_disabled": "MYUggbfQjdcbxkbAgJbvHfBXu",    "slot": 15,    "faction_required": 89,    "level_required": 74,    "alt_currency_cost": 31,    "classes_required": 13,    "max_expansion": 60}


Comments
-------------------------------------
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned



*/

// Merchantlist struct is a row record of the merchantlist table in the eqemu database
type Merchantlist struct {
	//[ 0] merchantid                                     int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Merchantid int32 `gorm:"primary_key;column:merchantid;type:int;default:0;" json:"merchantid"`
	//[ 1] slot                                           int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Slot int32 `gorm:"primary_key;column:slot;type:int;default:0;" json:"slot"`
	//[ 2] item                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Item int32 `gorm:"column:item;type:int;default:0;" json:"item"`
	//[ 3] faction_required                               smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [-100]
	FactionRequired int32 `gorm:"column:faction_required;type:smallint;default:-100;" json:"faction_required"`
	//[ 4] level_required                                 utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	LevelRequired uint32 `gorm:"column:level_required;type:utinyint;default:0;" json:"level_required"`
	//[ 5] alt_currency_cost                              usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	AltCurrencyCost uint32 `gorm:"column:alt_currency_cost;type:usmallint;default:0;" json:"alt_currency_cost"`
	//[ 6] classes_required                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [65535]
	ClassesRequired int32 `gorm:"column:classes_required;type:int;default:65535;" json:"classes_required"`
	//[ 7] probability                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [100]
	Probability int32 `gorm:"column:probability;type:int;default:100;" json:"probability"`
	//[ 8] min_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MinExpansion uint32 `gorm:"column:min_expansion;type:utinyint;default:0;" json:"min_expansion"`
	//[ 9] max_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MaxExpansion uint32 `gorm:"column:max_expansion;type:utinyint;default:0;" json:"max_expansion"`
	//[10] content_flags                                  varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlags null.String `gorm:"column:content_flags;type:varchar;size:100;" json:"content_flags"`
	//[11] content_flags_disabled                         varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlagsDisabled null.String `gorm:"column:content_flags_disabled;type:varchar;size:100;" json:"content_flags_disabled"`
}

var merchantlistTableInfo = &TableInfo{
	Name: "merchantlist",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "merchantid",
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
			GoFieldName:        "Merchantid",
			GoFieldType:        "int32",
			JSONFieldName:      "merchantid",
			ProtobufFieldName:  "merchantid",
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
			Name:               "item",
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
			GoFieldName:        "Item",
			GoFieldType:        "int32",
			JSONFieldName:      "item",
			ProtobufFieldName:  "item",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "faction_required",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "FactionRequired",
			GoFieldType:        "int32",
			JSONFieldName:      "faction_required",
			ProtobufFieldName:  "faction_required",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "level_required",
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
			GoFieldName:        "LevelRequired",
			GoFieldType:        "uint32",
			JSONFieldName:      "level_required",
			ProtobufFieldName:  "level_required",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "alt_currency_cost",
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
			GoFieldName:        "AltCurrencyCost",
			GoFieldType:        "uint32",
			JSONFieldName:      "alt_currency_cost",
			ProtobufFieldName:  "alt_currency_cost",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "classes_required",
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
			GoFieldName:        "ClassesRequired",
			GoFieldType:        "int32",
			JSONFieldName:      "classes_required",
			ProtobufFieldName:  "classes_required",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "probability",
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
			GoFieldName:        "Probability",
			GoFieldType:        "int32",
			JSONFieldName:      "probability",
			ProtobufFieldName:  "probability",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "min_expansion",
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
			GoFieldName:        "MinExpansion",
			GoFieldType:        "uint32",
			JSONFieldName:      "min_expansion",
			ProtobufFieldName:  "min_expansion",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "max_expansion",
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
			GoFieldName:        "MaxExpansion",
			GoFieldType:        "uint32",
			JSONFieldName:      "max_expansion",
			ProtobufFieldName:  "max_expansion",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "content_flags",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "ContentFlags",
			GoFieldType:        "null.String",
			JSONFieldName:      "content_flags",
			ProtobufFieldName:  "content_flags",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "content_flags_disabled",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "ContentFlagsDisabled",
			GoFieldType:        "null.String",
			JSONFieldName:      "content_flags_disabled",
			ProtobufFieldName:  "content_flags_disabled",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *Merchantlist) TableName() string {
	return "merchantlist"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *Merchantlist) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *Merchantlist) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *Merchantlist) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *Merchantlist) TableInfo() *TableInfo {
	return merchantlistTableInfo
}
