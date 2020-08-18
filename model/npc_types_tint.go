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


CREATE TABLE `npc_types_tint` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `tint_set_name` text NOT NULL,
  `red1h` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `grn1h` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `blu1h` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `red2c` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `grn2c` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `blu2c` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `red3a` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `grn3a` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `blu3a` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `red4b` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `grn4b` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `blu4b` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `red5g` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `grn5g` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `blu5g` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `red6l` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `grn6l` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `blu6l` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `red7f` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `grn7f` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `blu7f` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `red8x` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `grn8x` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `blu8x` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `red9x` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `grn9x` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `blu9x` tinyint(3) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC

JSON Sample
-------------------------------------
{    "blu_6_l": 37,    "red_7_f": 7,    "red_8_x": 61,    "id": 81,    "red_1_h": 24,    "blu_1_h": 33,    "blu_3_a": 43,    "grn_6_l": 6,    "tint_set_name": "hXjRHPTVVmbSXRYDywZgkotat",    "grn_1_h": 87,    "red_2_c": 74,    "blu_8_x": 55,    "grn_5_g": 92,    "blu_5_g": 36,    "blu_2_c": 23,    "red_3_a": 97,    "blu_4_b": 40,    "blu_7_f": 57,    "grn_8_x": 59,    "red_9_x": 60,    "red_4_b": 41,    "grn_4_b": 77,    "red_5_g": 79,    "red_6_l": 19,    "blu_9_x": 7,    "grn_9_x": 68,    "grn_2_c": 90,    "grn_3_a": 91,    "grn_7_f": 58}


Comments
-------------------------------------
[ 0] column is set for unsigned
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
[15] column is set for unsigned
[16] column is set for unsigned
[17] column is set for unsigned
[18] column is set for unsigned
[19] column is set for unsigned
[20] column is set for unsigned
[21] column is set for unsigned
[22] column is set for unsigned
[23] column is set for unsigned
[24] column is set for unsigned
[25] column is set for unsigned
[26] column is set for unsigned
[27] column is set for unsigned
[28] column is set for unsigned



*/

// NpcTypesTint struct is a row record of the npc_types_tint table in the eqemu database
type NpcTypesTint struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ID uint32 `gorm:"primary_key;column:id;type:uint;default:0;" json:"id"`
	//[ 1] tint_set_name                                  text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	TintSetName string `gorm:"column:tint_set_name;type:text;size:65535;" json:"tint_set_name"`
	//[ 2] red1h                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Red1h uint32 `gorm:"column:red1h;type:utinyint;default:0;" json:"red_1_h"`
	//[ 3] grn1h                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Grn1h uint32 `gorm:"column:grn1h;type:utinyint;default:0;" json:"grn_1_h"`
	//[ 4] blu1h                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Blu1h uint32 `gorm:"column:blu1h;type:utinyint;default:0;" json:"blu_1_h"`
	//[ 5] red2c                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Red2c uint32 `gorm:"column:red2c;type:utinyint;default:0;" json:"red_2_c"`
	//[ 6] grn2c                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Grn2c uint32 `gorm:"column:grn2c;type:utinyint;default:0;" json:"grn_2_c"`
	//[ 7] blu2c                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Blu2c uint32 `gorm:"column:blu2c;type:utinyint;default:0;" json:"blu_2_c"`
	//[ 8] red3a                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Red3a uint32 `gorm:"column:red3a;type:utinyint;default:0;" json:"red_3_a"`
	//[ 9] grn3a                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Grn3a uint32 `gorm:"column:grn3a;type:utinyint;default:0;" json:"grn_3_a"`
	//[10] blu3a                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Blu3a uint32 `gorm:"column:blu3a;type:utinyint;default:0;" json:"blu_3_a"`
	//[11] red4b                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Red4b uint32 `gorm:"column:red4b;type:utinyint;default:0;" json:"red_4_b"`
	//[12] grn4b                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Grn4b uint32 `gorm:"column:grn4b;type:utinyint;default:0;" json:"grn_4_b"`
	//[13] blu4b                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Blu4b uint32 `gorm:"column:blu4b;type:utinyint;default:0;" json:"blu_4_b"`
	//[14] red5g                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Red5g uint32 `gorm:"column:red5g;type:utinyint;default:0;" json:"red_5_g"`
	//[15] grn5g                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Grn5g uint32 `gorm:"column:grn5g;type:utinyint;default:0;" json:"grn_5_g"`
	//[16] blu5g                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Blu5g uint32 `gorm:"column:blu5g;type:utinyint;default:0;" json:"blu_5_g"`
	//[17] red6l                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Red6l uint32 `gorm:"column:red6l;type:utinyint;default:0;" json:"red_6_l"`
	//[18] grn6l                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Grn6l uint32 `gorm:"column:grn6l;type:utinyint;default:0;" json:"grn_6_l"`
	//[19] blu6l                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Blu6l uint32 `gorm:"column:blu6l;type:utinyint;default:0;" json:"blu_6_l"`
	//[20] red7f                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Red7f uint32 `gorm:"column:red7f;type:utinyint;default:0;" json:"red_7_f"`
	//[21] grn7f                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Grn7f uint32 `gorm:"column:grn7f;type:utinyint;default:0;" json:"grn_7_f"`
	//[22] blu7f                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Blu7f uint32 `gorm:"column:blu7f;type:utinyint;default:0;" json:"blu_7_f"`
	//[23] red8x                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Red8x uint32 `gorm:"column:red8x;type:utinyint;default:0;" json:"red_8_x"`
	//[24] grn8x                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Grn8x uint32 `gorm:"column:grn8x;type:utinyint;default:0;" json:"grn_8_x"`
	//[25] blu8x                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Blu8x uint32 `gorm:"column:blu8x;type:utinyint;default:0;" json:"blu_8_x"`
	//[26] red9x                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Red9x uint32 `gorm:"column:red9x;type:utinyint;default:0;" json:"red_9_x"`
	//[27] grn9x                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Grn9x uint32 `gorm:"column:grn9x;type:utinyint;default:0;" json:"grn_9_x"`
	//[28] blu9x                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Blu9x uint32 `gorm:"column:blu9x;type:utinyint;default:0;" json:"blu_9_x"`
}

var npc_types_tintTableInfo = &TableInfo{
	Name: "npc_types_tint",
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
			Name:               "tint_set_name",
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
			GoFieldName:        "TintSetName",
			GoFieldType:        "string",
			JSONFieldName:      "tint_set_name",
			ProtobufFieldName:  "tint_set_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "red1h",
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
			GoFieldName:        "Red1h",
			GoFieldType:        "uint32",
			JSONFieldName:      "red_1_h",
			ProtobufFieldName:  "red_1_h",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "grn1h",
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
			GoFieldName:        "Grn1h",
			GoFieldType:        "uint32",
			JSONFieldName:      "grn_1_h",
			ProtobufFieldName:  "grn_1_h",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "blu1h",
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
			GoFieldName:        "Blu1h",
			GoFieldType:        "uint32",
			JSONFieldName:      "blu_1_h",
			ProtobufFieldName:  "blu_1_h",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "red2c",
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
			GoFieldName:        "Red2c",
			GoFieldType:        "uint32",
			JSONFieldName:      "red_2_c",
			ProtobufFieldName:  "red_2_c",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "grn2c",
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
			GoFieldName:        "Grn2c",
			GoFieldType:        "uint32",
			JSONFieldName:      "grn_2_c",
			ProtobufFieldName:  "grn_2_c",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "blu2c",
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
			GoFieldName:        "Blu2c",
			GoFieldType:        "uint32",
			JSONFieldName:      "blu_2_c",
			ProtobufFieldName:  "blu_2_c",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "red3a",
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
			GoFieldName:        "Red3a",
			GoFieldType:        "uint32",
			JSONFieldName:      "red_3_a",
			ProtobufFieldName:  "red_3_a",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "grn3a",
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
			GoFieldName:        "Grn3a",
			GoFieldType:        "uint32",
			JSONFieldName:      "grn_3_a",
			ProtobufFieldName:  "grn_3_a",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "blu3a",
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
			GoFieldName:        "Blu3a",
			GoFieldType:        "uint32",
			JSONFieldName:      "blu_3_a",
			ProtobufFieldName:  "blu_3_a",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "red4b",
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
			GoFieldName:        "Red4b",
			GoFieldType:        "uint32",
			JSONFieldName:      "red_4_b",
			ProtobufFieldName:  "red_4_b",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "grn4b",
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
			GoFieldName:        "Grn4b",
			GoFieldType:        "uint32",
			JSONFieldName:      "grn_4_b",
			ProtobufFieldName:  "grn_4_b",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "blu4b",
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
			GoFieldName:        "Blu4b",
			GoFieldType:        "uint32",
			JSONFieldName:      "blu_4_b",
			ProtobufFieldName:  "blu_4_b",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "red5g",
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
			GoFieldName:        "Red5g",
			GoFieldType:        "uint32",
			JSONFieldName:      "red_5_g",
			ProtobufFieldName:  "red_5_g",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "grn5g",
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
			GoFieldName:        "Grn5g",
			GoFieldType:        "uint32",
			JSONFieldName:      "grn_5_g",
			ProtobufFieldName:  "grn_5_g",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "blu5g",
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
			GoFieldName:        "Blu5g",
			GoFieldType:        "uint32",
			JSONFieldName:      "blu_5_g",
			ProtobufFieldName:  "blu_5_g",
			ProtobufType:       "uint32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "red6l",
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
			GoFieldName:        "Red6l",
			GoFieldType:        "uint32",
			JSONFieldName:      "red_6_l",
			ProtobufFieldName:  "red_6_l",
			ProtobufType:       "uint32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "grn6l",
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
			GoFieldName:        "Grn6l",
			GoFieldType:        "uint32",
			JSONFieldName:      "grn_6_l",
			ProtobufFieldName:  "grn_6_l",
			ProtobufType:       "uint32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "blu6l",
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
			GoFieldName:        "Blu6l",
			GoFieldType:        "uint32",
			JSONFieldName:      "blu_6_l",
			ProtobufFieldName:  "blu_6_l",
			ProtobufType:       "uint32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "red7f",
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
			GoFieldName:        "Red7f",
			GoFieldType:        "uint32",
			JSONFieldName:      "red_7_f",
			ProtobufFieldName:  "red_7_f",
			ProtobufType:       "uint32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "grn7f",
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
			GoFieldName:        "Grn7f",
			GoFieldType:        "uint32",
			JSONFieldName:      "grn_7_f",
			ProtobufFieldName:  "grn_7_f",
			ProtobufType:       "uint32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "blu7f",
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
			GoFieldName:        "Blu7f",
			GoFieldType:        "uint32",
			JSONFieldName:      "blu_7_f",
			ProtobufFieldName:  "blu_7_f",
			ProtobufType:       "uint32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "red8x",
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
			GoFieldName:        "Red8x",
			GoFieldType:        "uint32",
			JSONFieldName:      "red_8_x",
			ProtobufFieldName:  "red_8_x",
			ProtobufType:       "uint32",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "grn8x",
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
			GoFieldName:        "Grn8x",
			GoFieldType:        "uint32",
			JSONFieldName:      "grn_8_x",
			ProtobufFieldName:  "grn_8_x",
			ProtobufType:       "uint32",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "blu8x",
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
			GoFieldName:        "Blu8x",
			GoFieldType:        "uint32",
			JSONFieldName:      "blu_8_x",
			ProtobufFieldName:  "blu_8_x",
			ProtobufType:       "uint32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "red9x",
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
			GoFieldName:        "Red9x",
			GoFieldType:        "uint32",
			JSONFieldName:      "red_9_x",
			ProtobufFieldName:  "red_9_x",
			ProtobufType:       "uint32",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "grn9x",
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
			GoFieldName:        "Grn9x",
			GoFieldType:        "uint32",
			JSONFieldName:      "grn_9_x",
			ProtobufFieldName:  "grn_9_x",
			ProtobufType:       "uint32",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "blu9x",
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
			GoFieldName:        "Blu9x",
			GoFieldType:        "uint32",
			JSONFieldName:      "blu_9_x",
			ProtobufFieldName:  "blu_9_x",
			ProtobufType:       "uint32",
			ProtobufPos:        29,
		},
	},
}

// TableName sets the insert table name for this struct type
func (n *NpcTypesTint) TableName() string {
	return "npc_types_tint"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (n *NpcTypesTint) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (n *NpcTypesTint) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (n *NpcTypesTint) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (n *NpcTypesTint) TableInfo() *TableInfo {
	return npc_types_tintTableInfo
}
