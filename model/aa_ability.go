package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
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


CREATE TABLE `aa_ability` (
  `id` int(10) unsigned NOT NULL,
  `name` text NOT NULL,
  `category` int(10) NOT NULL DEFAULT -1,
  `classes` int(10) NOT NULL DEFAULT 131070,
  `races` int(10) NOT NULL DEFAULT 65535,
  `drakkin_heritage` int(10) NOT NULL DEFAULT 127,
  `deities` int(10) NOT NULL DEFAULT 131071,
  `status` int(10) NOT NULL DEFAULT 0,
  `type` int(10) NOT NULL DEFAULT 0,
  `charges` int(11) NOT NULL DEFAULT 0,
  `grant_only` tinyint(4) NOT NULL DEFAULT 0,
  `first_rank_id` int(10) NOT NULL DEFAULT -1,
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT 1,
  `reset_on_death` tinyint(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "grant_only": 65,    "id": 56,    "name": "MZuCKvZsiouVTwmIPbiQZxVjL",    "drakkin_heritage": 53,    "status": 24,    "type": 78,    "classes": 14,    "races": 38,    "first_rank_id": 68,    "category": 24,    "deities": 16,    "charges": 37,    "enabled": 92,    "reset_on_death": 46}


Comments
-------------------------------------
[ 0] column is set for unsigned
[12] column is set for unsigned



*/

// AaAbility struct is a row record of the aa_ability table in the eqemu database
type AaAbility struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;column:id;type:uint;" json:"id"`
	//[ 1] name                                           text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Name string `gorm:"column:name;type:text;size:65535;" json:"name"`
	//[ 2] category                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	Category int32 `gorm:"column:category;type:int;default:-1;" json:"category"`
	//[ 3] classes                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [131070]
	Classes int32 `gorm:"column:classes;type:int;default:131070;" json:"classes"`
	//[ 4] races                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [65535]
	Races int32 `gorm:"column:races;type:int;default:65535;" json:"races"`
	//[ 5] drakkin_heritage                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [127]
	DrakkinHeritage int32 `gorm:"column:drakkin_heritage;type:int;default:127;" json:"drakkin_heritage"`
	//[ 6] deities                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [131071]
	Deities int32 `gorm:"column:deities;type:int;default:131071;" json:"deities"`
	//[ 7] status                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Status int32 `gorm:"column:status;type:int;default:0;" json:"status"`
	//[ 8] type                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Type int32 `gorm:"column:type;type:int;default:0;" json:"type"`
	//[ 9] charges                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Charges int32 `gorm:"column:charges;type:int;default:0;" json:"charges"`
	//[10] grant_only                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	GrantOnly int32 `gorm:"column:grant_only;type:tinyint;default:0;" json:"grant_only"`
	//[11] first_rank_id                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	FirstRankID int32 `gorm:"column:first_rank_id;type:int;default:-1;" json:"first_rank_id"`
	//[12] enabled                                        utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	Enabled uint32 `gorm:"column:enabled;type:utinyint;default:1;" json:"enabled"`
	//[13] reset_on_death                                 tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	ResetOnDeath int32 `gorm:"column:reset_on_death;type:tinyint;default:0;" json:"reset_on_death"`
}

var aa_abilityTableInfo = &TableInfo{
	Name: "aa_ability",
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
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "category",
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
			GoFieldName:        "Category",
			GoFieldType:        "int32",
			JSONFieldName:      "category",
			ProtobufFieldName:  "category",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "classes",
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
			GoFieldName:        "Classes",
			GoFieldType:        "int32",
			JSONFieldName:      "classes",
			ProtobufFieldName:  "classes",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "races",
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
			GoFieldName:        "Races",
			GoFieldType:        "int32",
			JSONFieldName:      "races",
			ProtobufFieldName:  "races",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "drakkin_heritage",
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
			GoFieldName:        "DrakkinHeritage",
			GoFieldType:        "int32",
			JSONFieldName:      "drakkin_heritage",
			ProtobufFieldName:  "drakkin_heritage",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "deities",
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
			GoFieldName:        "Deities",
			GoFieldType:        "int32",
			JSONFieldName:      "deities",
			ProtobufFieldName:  "deities",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "status",
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
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "charges",
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
			GoFieldName:        "Charges",
			GoFieldType:        "int32",
			JSONFieldName:      "charges",
			ProtobufFieldName:  "charges",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "grant_only",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "GrantOnly",
			GoFieldType:        "int32",
			JSONFieldName:      "grant_only",
			ProtobufFieldName:  "grant_only",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "first_rank_id",
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
			GoFieldName:        "FirstRankID",
			GoFieldType:        "int32",
			JSONFieldName:      "first_rank_id",
			ProtobufFieldName:  "first_rank_id",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "enabled",
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
			GoFieldName:        "Enabled",
			GoFieldType:        "uint32",
			JSONFieldName:      "enabled",
			ProtobufFieldName:  "enabled",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "reset_on_death",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "ResetOnDeath",
			GoFieldType:        "int32",
			JSONFieldName:      "reset_on_death",
			ProtobufFieldName:  "reset_on_death",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AaAbility) TableName() string {
	return "aa_ability"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AaAbility) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AaAbility) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AaAbility) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AaAbility) TableInfo() *TableInfo {
	return aa_abilityTableInfo
}
