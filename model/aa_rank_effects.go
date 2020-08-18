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


CREATE TABLE `aa_rank_effects` (
  `rank_id` int(10) unsigned NOT NULL,
  `slot` int(10) unsigned NOT NULL DEFAULT 1,
  `effect_id` int(10) NOT NULL DEFAULT 0,
  `base1` int(10) NOT NULL DEFAULT 0,
  `base2` int(10) NOT NULL DEFAULT 0,
  PRIMARY KEY (`rank_id`,`slot`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "base_2": 81,    "rank_id": 16,    "slot": 29,    "effect_id": 74,    "base_1": 15}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// AaRankEffects struct is a row record of the aa_rank_effects table in the eqemu database
type AaRankEffects struct {
	//[ 0] rank_id                                        uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	RankID uint32 `gorm:"primary_key;column:rank_id;type:uint;" json:"rank_id"`
	//[ 1] slot                                           uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [1]
	Slot uint32 `gorm:"primary_key;column:slot;type:uint;default:1;" json:"slot"`
	//[ 2] effect_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EffectID int32 `gorm:"column:effect_id;type:int;default:0;" json:"effect_id"`
	//[ 3] base1                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Base1 int32 `gorm:"column:base1;type:int;default:0;" json:"base_1"`
	//[ 4] base2                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Base2 int32 `gorm:"column:base2;type:int;default:0;" json:"base_2"`
}

var aa_rank_effectsTableInfo = &TableInfo{
	Name: "aa_rank_effects",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "rank_id",
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
			GoFieldName:        "RankID",
			GoFieldType:        "uint32",
			JSONFieldName:      "rank_id",
			ProtobufFieldName:  "rank_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "slot",
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
			GoFieldName:        "Slot",
			GoFieldType:        "uint32",
			JSONFieldName:      "slot",
			ProtobufFieldName:  "slot",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "effect_id",
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
			GoFieldName:        "EffectID",
			GoFieldType:        "int32",
			JSONFieldName:      "effect_id",
			ProtobufFieldName:  "effect_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "base1",
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
			GoFieldName:        "Base1",
			GoFieldType:        "int32",
			JSONFieldName:      "base_1",
			ProtobufFieldName:  "base_1",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "base2",
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
			GoFieldName:        "Base2",
			GoFieldType:        "int32",
			JSONFieldName:      "base_2",
			ProtobufFieldName:  "base_2",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AaRankEffects) TableName() string {
	return "aa_rank_effects"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AaRankEffects) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AaRankEffects) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AaRankEffects) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AaRankEffects) TableInfo() *TableInfo {
	return aa_rank_effectsTableInfo
}
