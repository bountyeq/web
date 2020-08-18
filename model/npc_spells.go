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


CREATE TABLE `npc_spells` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` tinytext DEFAULT NULL,
  `parent_list` int(11) unsigned NOT NULL DEFAULT 0,
  `attack_proc` smallint(5) NOT NULL DEFAULT -1,
  `proc_chance` tinyint(3) NOT NULL DEFAULT 3,
  `range_proc` smallint(5) NOT NULL DEFAULT -1,
  `rproc_chance` smallint(5) NOT NULL DEFAULT 0,
  `defensive_proc` smallint(5) NOT NULL DEFAULT -1,
  `dproc_chance` smallint(5) NOT NULL DEFAULT 0,
  `fail_recast` int(11) unsigned NOT NULL DEFAULT 0,
  `engaged_no_sp_recast_min` int(11) unsigned NOT NULL DEFAULT 0,
  `engaged_no_sp_recast_max` int(11) unsigned NOT NULL DEFAULT 0,
  `engaged_b_self_chance` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `engaged_b_other_chance` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `engaged_d_chance` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `pursue_no_sp_recast_min` int(3) unsigned NOT NULL DEFAULT 0,
  `pursue_no_sp_recast_max` int(11) unsigned NOT NULL DEFAULT 0,
  `pursue_d_chance` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `idle_no_sp_recast_min` int(11) unsigned NOT NULL DEFAULT 0,
  `idle_no_sp_recast_max` int(11) unsigned NOT NULL DEFAULT 0,
  `idle_b_chance` tinyint(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3612 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "engaged_b_self_chance": 51,    "engaged_d_chance": 16,    "range_proc": 28,    "dproc_chance": 88,    "defensive_proc": 26,    "name": "GaptvrlfToXgaorPdqhnKXunW",    "proc_chance": 53,    "engaged_b_other_chance": 92,    "pursue_no_sp_recast_min": 62,    "pursue_no_sp_recast_max": 79,    "idle_no_sp_recast_min": 15,    "idle_b_chance": 52,    "id": 20,    "engaged_no_sp_recast_max": 70,    "rproc_chance": 65,    "fail_recast": 30,    "engaged_no_sp_recast_min": 50,    "pursue_d_chance": 84,    "idle_no_sp_recast_max": 39,    "parent_list": 48,    "attack_proc": 98}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned
[11] column is set for unsigned
[12] column is set for unsigned
[13] column is set for unsigned
[14] column is set for unsigned
[15] column is set for unsigned
[16] column is set for unsigned
[17] column is set for unsigned
[18] column is set for unsigned
[19] column is set for unsigned
[20] column is set for unsigned



*/

// NpcSpells struct is a row record of the npc_spells table in the eqemu database
type NpcSpells struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] name                                           text(255)            null: true   primary: false  isArray: false  auto: false  col: text            len: 255     default: [NULL]
	Name null.String `gorm:"column:name;type:text;size:255;" json:"name"`
	//[ 2] parent_list                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ParentList uint32 `gorm:"column:parent_list;type:uint;default:0;" json:"parent_list"`
	//[ 3] attack_proc                                    smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [-1]
	AttackProc int32 `gorm:"column:attack_proc;type:smallint;default:-1;" json:"attack_proc"`
	//[ 4] proc_chance                                    tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [3]
	ProcChance int32 `gorm:"column:proc_chance;type:tinyint;default:3;" json:"proc_chance"`
	//[ 5] range_proc                                     smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [-1]
	RangeProc int32 `gorm:"column:range_proc;type:smallint;default:-1;" json:"range_proc"`
	//[ 6] rproc_chance                                   smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	RprocChance int32 `gorm:"column:rproc_chance;type:smallint;default:0;" json:"rproc_chance"`
	//[ 7] defensive_proc                                 smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [-1]
	DefensiveProc int32 `gorm:"column:defensive_proc;type:smallint;default:-1;" json:"defensive_proc"`
	//[ 8] dproc_chance                                   smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	DprocChance int32 `gorm:"column:dproc_chance;type:smallint;default:0;" json:"dproc_chance"`
	//[ 9] fail_recast                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	FailRecast uint32 `gorm:"column:fail_recast;type:uint;default:0;" json:"fail_recast"`
	//[10] engaged_no_sp_recast_min                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EngagedNoSpRecastMin uint32 `gorm:"column:engaged_no_sp_recast_min;type:uint;default:0;" json:"engaged_no_sp_recast_min"`
	//[11] engaged_no_sp_recast_max                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EngagedNoSpRecastMax uint32 `gorm:"column:engaged_no_sp_recast_max;type:uint;default:0;" json:"engaged_no_sp_recast_max"`
	//[12] engaged_b_self_chance                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	EngagedBSelfChance uint32 `gorm:"column:engaged_b_self_chance;type:utinyint;default:0;" json:"engaged_b_self_chance"`
	//[13] engaged_b_other_chance                         utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	EngagedBOtherChance uint32 `gorm:"column:engaged_b_other_chance;type:utinyint;default:0;" json:"engaged_b_other_chance"`
	//[14] engaged_d_chance                               utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	EngagedDChance uint32 `gorm:"column:engaged_d_chance;type:utinyint;default:0;" json:"engaged_d_chance"`
	//[15] pursue_no_sp_recast_min                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PursueNoSpRecastMin uint32 `gorm:"column:pursue_no_sp_recast_min;type:uint;default:0;" json:"pursue_no_sp_recast_min"`
	//[16] pursue_no_sp_recast_max                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PursueNoSpRecastMax uint32 `gorm:"column:pursue_no_sp_recast_max;type:uint;default:0;" json:"pursue_no_sp_recast_max"`
	//[17] pursue_d_chance                                utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	PursueDChance uint32 `gorm:"column:pursue_d_chance;type:utinyint;default:0;" json:"pursue_d_chance"`
	//[18] idle_no_sp_recast_min                          uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	IdleNoSpRecastMin uint32 `gorm:"column:idle_no_sp_recast_min;type:uint;default:0;" json:"idle_no_sp_recast_min"`
	//[19] idle_no_sp_recast_max                          uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	IdleNoSpRecastMax uint32 `gorm:"column:idle_no_sp_recast_max;type:uint;default:0;" json:"idle_no_sp_recast_max"`
	//[20] idle_b_chance                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	IdleBChance uint32 `gorm:"column:idle_b_chance;type:utinyint;default:0;" json:"idle_b_chance"`
}

var npc_spellsTableInfo = &TableInfo{
	Name: "npc_spells",
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
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "null.String",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "parent_list",
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
			GoFieldName:        "ParentList",
			GoFieldType:        "uint32",
			JSONFieldName:      "parent_list",
			ProtobufFieldName:  "parent_list",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "attack_proc",
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
			GoFieldName:        "AttackProc",
			GoFieldType:        "int32",
			JSONFieldName:      "attack_proc",
			ProtobufFieldName:  "attack_proc",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "proc_chance",
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
			GoFieldName:        "ProcChance",
			GoFieldType:        "int32",
			JSONFieldName:      "proc_chance",
			ProtobufFieldName:  "proc_chance",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "range_proc",
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
			GoFieldName:        "RangeProc",
			GoFieldType:        "int32",
			JSONFieldName:      "range_proc",
			ProtobufFieldName:  "range_proc",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "rproc_chance",
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
			GoFieldName:        "RprocChance",
			GoFieldType:        "int32",
			JSONFieldName:      "rproc_chance",
			ProtobufFieldName:  "rproc_chance",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "defensive_proc",
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
			GoFieldName:        "DefensiveProc",
			GoFieldType:        "int32",
			JSONFieldName:      "defensive_proc",
			ProtobufFieldName:  "defensive_proc",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "dproc_chance",
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
			GoFieldName:        "DprocChance",
			GoFieldType:        "int32",
			JSONFieldName:      "dproc_chance",
			ProtobufFieldName:  "dproc_chance",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "fail_recast",
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
			GoFieldName:        "FailRecast",
			GoFieldType:        "uint32",
			JSONFieldName:      "fail_recast",
			ProtobufFieldName:  "fail_recast",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "engaged_no_sp_recast_min",
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
			GoFieldName:        "EngagedNoSpRecastMin",
			GoFieldType:        "uint32",
			JSONFieldName:      "engaged_no_sp_recast_min",
			ProtobufFieldName:  "engaged_no_sp_recast_min",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "engaged_no_sp_recast_max",
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
			GoFieldName:        "EngagedNoSpRecastMax",
			GoFieldType:        "uint32",
			JSONFieldName:      "engaged_no_sp_recast_max",
			ProtobufFieldName:  "engaged_no_sp_recast_max",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "engaged_b_self_chance",
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
			GoFieldName:        "EngagedBSelfChance",
			GoFieldType:        "uint32",
			JSONFieldName:      "engaged_b_self_chance",
			ProtobufFieldName:  "engaged_b_self_chance",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "engaged_b_other_chance",
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
			GoFieldName:        "EngagedBOtherChance",
			GoFieldType:        "uint32",
			JSONFieldName:      "engaged_b_other_chance",
			ProtobufFieldName:  "engaged_b_other_chance",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "engaged_d_chance",
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
			GoFieldName:        "EngagedDChance",
			GoFieldType:        "uint32",
			JSONFieldName:      "engaged_d_chance",
			ProtobufFieldName:  "engaged_d_chance",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "pursue_no_sp_recast_min",
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
			GoFieldName:        "PursueNoSpRecastMin",
			GoFieldType:        "uint32",
			JSONFieldName:      "pursue_no_sp_recast_min",
			ProtobufFieldName:  "pursue_no_sp_recast_min",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "pursue_no_sp_recast_max",
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
			GoFieldName:        "PursueNoSpRecastMax",
			GoFieldType:        "uint32",
			JSONFieldName:      "pursue_no_sp_recast_max",
			ProtobufFieldName:  "pursue_no_sp_recast_max",
			ProtobufType:       "uint32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "pursue_d_chance",
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
			GoFieldName:        "PursueDChance",
			GoFieldType:        "uint32",
			JSONFieldName:      "pursue_d_chance",
			ProtobufFieldName:  "pursue_d_chance",
			ProtobufType:       "uint32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "idle_no_sp_recast_min",
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
			GoFieldName:        "IdleNoSpRecastMin",
			GoFieldType:        "uint32",
			JSONFieldName:      "idle_no_sp_recast_min",
			ProtobufFieldName:  "idle_no_sp_recast_min",
			ProtobufType:       "uint32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "idle_no_sp_recast_max",
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
			GoFieldName:        "IdleNoSpRecastMax",
			GoFieldType:        "uint32",
			JSONFieldName:      "idle_no_sp_recast_max",
			ProtobufFieldName:  "idle_no_sp_recast_max",
			ProtobufType:       "uint32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "idle_b_chance",
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
			GoFieldName:        "IdleBChance",
			GoFieldType:        "uint32",
			JSONFieldName:      "idle_b_chance",
			ProtobufFieldName:  "idle_b_chance",
			ProtobufType:       "uint32",
			ProtobufPos:        21,
		},
	},
}

// TableName sets the insert table name for this struct type
func (n *NpcSpells) TableName() string {
	return "npc_spells"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (n *NpcSpells) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (n *NpcSpells) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (n *NpcSpells) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (n *NpcSpells) TableInfo() *TableInfo {
	return npc_spellsTableInfo
}
