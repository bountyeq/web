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


CREATE TABLE `aa_ranks` (
  `id` int(10) unsigned NOT NULL,
  `upper_hotkey_sid` int(10) NOT NULL DEFAULT -1,
  `lower_hotkey_sid` int(10) NOT NULL DEFAULT -1,
  `title_sid` int(10) NOT NULL DEFAULT -1,
  `desc_sid` int(10) NOT NULL DEFAULT -1,
  `cost` int(10) NOT NULL DEFAULT 1,
  `level_req` int(10) NOT NULL DEFAULT 51,
  `spell` int(10) NOT NULL DEFAULT -1,
  `spell_type` int(10) NOT NULL DEFAULT 0,
  `recast_time` int(10) NOT NULL DEFAULT 0,
  `expansion` int(10) NOT NULL DEFAULT 0,
  `prev_id` int(10) NOT NULL DEFAULT -1,
  `next_id` int(10) NOT NULL DEFAULT -1,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "upper_hotkey_sid": 47,    "lower_hotkey_sid": 59,    "desc_sid": 20,    "spell": 2,    "spell_type": 3,    "recast_time": 24,    "expansion": 39,    "next_id": 74,    "id": 59,    "title_sid": 0,    "cost": 65,    "level_req": 19,    "prev_id": 26}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// AaRanks struct is a row record of the aa_ranks table in the eqemu database
type AaRanks struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;column:id;type:uint;" json:"id"`
	//[ 1] upper_hotkey_sid                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	UpperHotkeySid int32 `gorm:"column:upper_hotkey_sid;type:int;default:-1;" json:"upper_hotkey_sid"`
	//[ 2] lower_hotkey_sid                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	LowerHotkeySid int32 `gorm:"column:lower_hotkey_sid;type:int;default:-1;" json:"lower_hotkey_sid"`
	//[ 3] title_sid                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	TitleSid int32 `gorm:"column:title_sid;type:int;default:-1;" json:"title_sid"`
	//[ 4] desc_sid                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	DescSid int32 `gorm:"column:desc_sid;type:int;default:-1;" json:"desc_sid"`
	//[ 5] cost                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	Cost int32 `gorm:"column:cost;type:int;default:1;" json:"cost"`
	//[ 6] level_req                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [51]
	LevelReq int32 `gorm:"column:level_req;type:int;default:51;" json:"level_req"`
	//[ 7] spell                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	Spell int32 `gorm:"column:spell;type:int;default:-1;" json:"spell"`
	//[ 8] spell_type                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SpellType int32 `gorm:"column:spell_type;type:int;default:0;" json:"spell_type"`
	//[ 9] recast_time                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RecastTime int32 `gorm:"column:recast_time;type:int;default:0;" json:"recast_time"`
	//[10] expansion                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Expansion int32 `gorm:"column:expansion;type:int;default:0;" json:"expansion"`
	//[11] prev_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	PrevID int32 `gorm:"column:prev_id;type:int;default:-1;" json:"prev_id"`
	//[12] next_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	NextID int32 `gorm:"column:next_id;type:int;default:-1;" json:"next_id"`
}

var aa_ranksTableInfo = &TableInfo{
	Name: "aa_ranks",
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
			Name:               "upper_hotkey_sid",
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
			GoFieldName:        "UpperHotkeySid",
			GoFieldType:        "int32",
			JSONFieldName:      "upper_hotkey_sid",
			ProtobufFieldName:  "upper_hotkey_sid",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "lower_hotkey_sid",
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
			GoFieldName:        "LowerHotkeySid",
			GoFieldType:        "int32",
			JSONFieldName:      "lower_hotkey_sid",
			ProtobufFieldName:  "lower_hotkey_sid",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "title_sid",
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
			GoFieldName:        "TitleSid",
			GoFieldType:        "int32",
			JSONFieldName:      "title_sid",
			ProtobufFieldName:  "title_sid",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "desc_sid",
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
			GoFieldName:        "DescSid",
			GoFieldType:        "int32",
			JSONFieldName:      "desc_sid",
			ProtobufFieldName:  "desc_sid",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "cost",
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
			GoFieldName:        "Cost",
			GoFieldType:        "int32",
			JSONFieldName:      "cost",
			ProtobufFieldName:  "cost",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "level_req",
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
			GoFieldName:        "LevelReq",
			GoFieldType:        "int32",
			JSONFieldName:      "level_req",
			ProtobufFieldName:  "level_req",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "spell",
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
			GoFieldName:        "Spell",
			GoFieldType:        "int32",
			JSONFieldName:      "spell",
			ProtobufFieldName:  "spell",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "spell_type",
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
			GoFieldName:        "SpellType",
			GoFieldType:        "int32",
			JSONFieldName:      "spell_type",
			ProtobufFieldName:  "spell_type",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "recast_time",
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
			GoFieldName:        "RecastTime",
			GoFieldType:        "int32",
			JSONFieldName:      "recast_time",
			ProtobufFieldName:  "recast_time",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "expansion",
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
			GoFieldName:        "Expansion",
			GoFieldType:        "int32",
			JSONFieldName:      "expansion",
			ProtobufFieldName:  "expansion",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "prev_id",
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
			GoFieldName:        "PrevID",
			GoFieldType:        "int32",
			JSONFieldName:      "prev_id",
			ProtobufFieldName:  "prev_id",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "next_id",
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
			GoFieldName:        "NextID",
			GoFieldType:        "int32",
			JSONFieldName:      "next_id",
			ProtobufFieldName:  "next_id",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AaRanks) TableName() string {
	return "aa_ranks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AaRanks) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AaRanks) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AaRanks) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AaRanks) TableInfo() *TableInfo {
	return aa_ranksTableInfo
}
