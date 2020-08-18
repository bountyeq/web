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


CREATE TABLE `bug_reports` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `zone` varchar(32) NOT NULL DEFAULT 'Unknown',
  `client_version_id` int(11) unsigned NOT NULL DEFAULT 0,
  `client_version_name` varchar(24) NOT NULL DEFAULT 'Unknown',
  `account_id` int(11) unsigned NOT NULL DEFAULT 0,
  `character_id` int(11) unsigned NOT NULL DEFAULT 0,
  `character_name` varchar(64) NOT NULL DEFAULT 'Unknown',
  `reporter_spoof` tinyint(1) NOT NULL DEFAULT 1,
  `category_id` int(11) unsigned NOT NULL DEFAULT 0,
  `category_name` varchar(64) NOT NULL DEFAULT 'Other',
  `reporter_name` varchar(64) NOT NULL DEFAULT 'Unknown',
  `ui_path` varchar(128) NOT NULL DEFAULT 'Unknown',
  `pos_x` float NOT NULL DEFAULT 0,
  `pos_y` float NOT NULL DEFAULT 0,
  `pos_z` float NOT NULL DEFAULT 0,
  `heading` int(11) unsigned NOT NULL DEFAULT 0,
  `time_played` int(11) unsigned NOT NULL DEFAULT 0,
  `target_id` int(11) unsigned NOT NULL DEFAULT 0,
  `target_name` varchar(64) NOT NULL DEFAULT 'Unknown',
  `optional_info_mask` int(11) unsigned NOT NULL DEFAULT 0,
  `_can_duplicate` tinyint(1) NOT NULL DEFAULT 0,
  `_crash_bug` tinyint(1) NOT NULL DEFAULT 0,
  `_target_info` tinyint(1) NOT NULL DEFAULT 0,
  `_character_flags` tinyint(1) NOT NULL DEFAULT 0,
  `_unknown_value` tinyint(1) NOT NULL DEFAULT 0,
  `bug_report` varchar(1024) NOT NULL DEFAULT '',
  `system_info` varchar(1024) NOT NULL DEFAULT '',
  `report_datetime` datetime NOT NULL DEFAULT current_timestamp(),
  `bug_status` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `last_review` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `last_reviewer` varchar(64) NOT NULL DEFAULT 'None',
  `reviewer_notes` varchar(1024) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "_unknown_value": 50,    "reviewer_notes": "ZyPqxwXlUaKUHduKEcbtEwEeN",    "client_version_name": "GKcTajlxBIRftnGoYEqOkjUZF",    "character_name": "EJKXarwmdhOAjTyyLbNlapkdS",    "target_name": "iEKtXkQKkIvcejGctaekTAwQG",    "pos_y": 0.63110477,    "pos_z": 0.14464252,    "system_info": "RbHtDGAHIWUVhmOSNfandUUUG",    "report_datetime": "2124-04-26T00:25:50.337782015-07:00",    "category_id": 67,    "ui_path": "tFZgpWogidTsVkRqEBvXeDwlm",    "last_review": "2229-07-25T08:01:09.705830987-07:00",    "character_id": 47,    "reporter_name": "pEpTlaedRWUkvLuMUtxkfIrxc",    "optional_info_mask": 92,    "_crash_bug": 63,    "_character_flags": 79,    "bug_status": 86,    "id": 3,    "zone": "NaMuWIMVntnhiQjGgRxuKsPpg",    "_can_duplicate": 16,    "_target_info": 13,    "client_version_id": 42,    "account_id": 4,    "time_played": 90,    "bug_report": "KKNjQFvcQdLqfcuiWaYuTLOkS",    "reporter_spoof": 33,    "category_name": "xjLSDGXaKnnyHkUsMricJTgQA",    "pos_x": 0.46308246,    "heading": 15,    "target_id": 30,    "last_reviewer": "GnoWRvqQoUNlYFtoUrfwGYWZk"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 8] column is set for unsigned
[15] column is set for unsigned
[16] column is set for unsigned
[17] column is set for unsigned
[19] column is set for unsigned
[28] column is set for unsigned



*/

// BugReports struct is a row record of the bug_reports table in the eqemu database
type BugReports struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] zone                                           varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['Unknown']
	Zone string `gorm:"column:zone;type:varchar;size:32;default:'Unknown';" json:"zone"`
	//[ 2] client_version_id                              uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ClientVersionID uint32 `gorm:"column:client_version_id;type:uint;default:0;" json:"client_version_id"`
	//[ 3] client_version_name                            varchar(24)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 24      default: ['Unknown']
	ClientVersionName string `gorm:"column:client_version_name;type:varchar;size:24;default:'Unknown';" json:"client_version_name"`
	//[ 4] account_id                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AccountID uint32 `gorm:"column:account_id;type:uint;default:0;" json:"account_id"`
	//[ 5] character_id                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CharacterID uint32 `gorm:"column:character_id;type:uint;default:0;" json:"character_id"`
	//[ 6] character_name                                 varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['Unknown']
	CharacterName string `gorm:"column:character_name;type:varchar;size:64;default:'Unknown';" json:"character_name"`
	//[ 7] reporter_spoof                                 tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	ReporterSpoof int32 `gorm:"column:reporter_spoof;type:tinyint;default:1;" json:"reporter_spoof"`
	//[ 8] category_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CategoryID uint32 `gorm:"column:category_id;type:uint;default:0;" json:"category_id"`
	//[ 9] category_name                                  varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['Other']
	CategoryName string `gorm:"column:category_name;type:varchar;size:64;default:'Other';" json:"category_name"`
	//[10] reporter_name                                  varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['Unknown']
	ReporterName string `gorm:"column:reporter_name;type:varchar;size:64;default:'Unknown';" json:"reporter_name"`
	//[11] ui_path                                        varchar(128)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 128     default: ['Unknown']
	UIPath string `gorm:"column:ui_path;type:varchar;size:128;default:'Unknown';" json:"ui_path"`
	//[12] pos_x                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	PosX float32 `gorm:"column:pos_x;type:float;default:0;" json:"pos_x"`
	//[13] pos_y                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	PosY float32 `gorm:"column:pos_y;type:float;default:0;" json:"pos_y"`
	//[14] pos_z                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	PosZ float32 `gorm:"column:pos_z;type:float;default:0;" json:"pos_z"`
	//[15] heading                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Heading uint32 `gorm:"column:heading;type:uint;default:0;" json:"heading"`
	//[16] time_played                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TimePlayed uint32 `gorm:"column:time_played;type:uint;default:0;" json:"time_played"`
	//[17] target_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TargetID uint32 `gorm:"column:target_id;type:uint;default:0;" json:"target_id"`
	//[18] target_name                                    varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['Unknown']
	TargetName string `gorm:"column:target_name;type:varchar;size:64;default:'Unknown';" json:"target_name"`
	//[19] optional_info_mask                             uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	OptionalInfoMask uint32 `gorm:"column:optional_info_mask;type:uint;default:0;" json:"optional_info_mask"`
	//[20] _can_duplicate                                 tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	CanDuplicate int32 `gorm:"column:_can_duplicate;type:tinyint;default:0;" json:"_can_duplicate"`
	//[21] _crash_bug                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	CrashBug int32 `gorm:"column:_crash_bug;type:tinyint;default:0;" json:"_crash_bug"`
	//[22] _target_info                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	TargetInfo int32 `gorm:"column:_target_info;type:tinyint;default:0;" json:"_target_info"`
	//[23] _character_flags                               tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	CharacterFlags int32 `gorm:"column:_character_flags;type:tinyint;default:0;" json:"_character_flags"`
	//[24] _unknown_value                                 tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	UnknownValue int32 `gorm:"column:_unknown_value;type:tinyint;default:0;" json:"_unknown_value"`
	//[25] bug_report                                     varchar(1024)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1024    default: ['']
	BugReport string `gorm:"column:bug_report;type:varchar;size:1024;default:'';" json:"bug_report"`
	//[26] system_info                                    varchar(1024)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1024    default: ['']
	SystemInfo string `gorm:"column:system_info;type:varchar;size:1024;default:'';" json:"system_info"`
	//[27] report_datetime                                datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [current_timestamp()]
	ReportDatetime time.Time `gorm:"column:report_datetime;type:datetime;" json:"report_datetime"`
	//[28] bug_status                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	BugStatus uint32 `gorm:"column:bug_status;type:utinyint;default:0;" json:"bug_status"`
	//[29] last_review                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [current_timestamp()]
	LastReview time.Time `gorm:"column:last_review;type:datetime;" json:"last_review"`
	//[30] last_reviewer                                  varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['None']
	LastReviewer string `gorm:"column:last_reviewer;type:varchar;size:64;default:'None';" json:"last_reviewer"`
	//[31] reviewer_notes                                 varchar(1024)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1024    default: ['']
	ReviewerNotes string `gorm:"column:reviewer_notes;type:varchar;size:1024;default:'';" json:"reviewer_notes"`
}

var bug_reportsTableInfo = &TableInfo{
	Name: "bug_reports",
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
			Name:               "zone",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Zone",
			GoFieldType:        "string",
			JSONFieldName:      "zone",
			ProtobufFieldName:  "zone",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "client_version_id",
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
			GoFieldName:        "ClientVersionID",
			GoFieldType:        "uint32",
			JSONFieldName:      "client_version_id",
			ProtobufFieldName:  "client_version_id",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "client_version_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(24)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       24,
			GoFieldName:        "ClientVersionName",
			GoFieldType:        "string",
			JSONFieldName:      "client_version_name",
			ProtobufFieldName:  "client_version_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "account_id",
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
			GoFieldName:        "AccountID",
			GoFieldType:        "uint32",
			JSONFieldName:      "account_id",
			ProtobufFieldName:  "account_id",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "character_id",
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
			GoFieldName:        "CharacterID",
			GoFieldType:        "uint32",
			JSONFieldName:      "character_id",
			ProtobufFieldName:  "character_id",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "character_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "CharacterName",
			GoFieldType:        "string",
			JSONFieldName:      "character_name",
			ProtobufFieldName:  "character_name",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "reporter_spoof",
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
			GoFieldName:        "ReporterSpoof",
			GoFieldType:        "int32",
			JSONFieldName:      "reporter_spoof",
			ProtobufFieldName:  "reporter_spoof",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "category_id",
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
			GoFieldName:        "CategoryID",
			GoFieldType:        "uint32",
			JSONFieldName:      "category_id",
			ProtobufFieldName:  "category_id",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "category_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "CategoryName",
			GoFieldType:        "string",
			JSONFieldName:      "category_name",
			ProtobufFieldName:  "category_name",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "reporter_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "ReporterName",
			GoFieldType:        "string",
			JSONFieldName:      "reporter_name",
			ProtobufFieldName:  "reporter_name",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "ui_path",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       128,
			GoFieldName:        "UIPath",
			GoFieldType:        "string",
			JSONFieldName:      "ui_path",
			ProtobufFieldName:  "ui_path",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "pos_x",
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
			GoFieldName:        "PosX",
			GoFieldType:        "float32",
			JSONFieldName:      "pos_x",
			ProtobufFieldName:  "pos_x",
			ProtobufType:       "float",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "pos_y",
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
			GoFieldName:        "PosY",
			GoFieldType:        "float32",
			JSONFieldName:      "pos_y",
			ProtobufFieldName:  "pos_y",
			ProtobufType:       "float",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "pos_z",
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
			GoFieldName:        "PosZ",
			GoFieldType:        "float32",
			JSONFieldName:      "pos_z",
			ProtobufFieldName:  "pos_z",
			ProtobufType:       "float",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "heading",
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
			GoFieldName:        "Heading",
			GoFieldType:        "uint32",
			JSONFieldName:      "heading",
			ProtobufFieldName:  "heading",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "time_played",
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
			GoFieldName:        "TimePlayed",
			GoFieldType:        "uint32",
			JSONFieldName:      "time_played",
			ProtobufFieldName:  "time_played",
			ProtobufType:       "uint32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "target_id",
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
			GoFieldName:        "TargetID",
			GoFieldType:        "uint32",
			JSONFieldName:      "target_id",
			ProtobufFieldName:  "target_id",
			ProtobufType:       "uint32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "target_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "TargetName",
			GoFieldType:        "string",
			JSONFieldName:      "target_name",
			ProtobufFieldName:  "target_name",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "optional_info_mask",
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
			GoFieldName:        "OptionalInfoMask",
			GoFieldType:        "uint32",
			JSONFieldName:      "optional_info_mask",
			ProtobufFieldName:  "optional_info_mask",
			ProtobufType:       "uint32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "_can_duplicate",
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
			GoFieldName:        "CanDuplicate",
			GoFieldType:        "int32",
			JSONFieldName:      "_can_duplicate",
			ProtobufFieldName:  "_can_duplicate",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "_crash_bug",
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
			GoFieldName:        "CrashBug",
			GoFieldType:        "int32",
			JSONFieldName:      "_crash_bug",
			ProtobufFieldName:  "_crash_bug",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "_target_info",
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
			GoFieldName:        "TargetInfo",
			GoFieldType:        "int32",
			JSONFieldName:      "_target_info",
			ProtobufFieldName:  "_target_info",
			ProtobufType:       "int32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "_character_flags",
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
			GoFieldName:        "CharacterFlags",
			GoFieldType:        "int32",
			JSONFieldName:      "_character_flags",
			ProtobufFieldName:  "_character_flags",
			ProtobufType:       "int32",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "_unknown_value",
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
			GoFieldName:        "UnknownValue",
			GoFieldType:        "int32",
			JSONFieldName:      "_unknown_value",
			ProtobufFieldName:  "_unknown_value",
			ProtobufType:       "int32",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "bug_report",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       1024,
			GoFieldName:        "BugReport",
			GoFieldType:        "string",
			JSONFieldName:      "bug_report",
			ProtobufFieldName:  "bug_report",
			ProtobufType:       "string",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "system_info",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       1024,
			GoFieldName:        "SystemInfo",
			GoFieldType:        "string",
			JSONFieldName:      "system_info",
			ProtobufFieldName:  "system_info",
			ProtobufType:       "string",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "report_datetime",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "ReportDatetime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "report_datetime",
			ProtobufFieldName:  "report_datetime",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "bug_status",
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
			GoFieldName:        "BugStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "bug_status",
			ProtobufFieldName:  "bug_status",
			ProtobufType:       "uint32",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "last_review",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "LastReview",
			GoFieldType:        "time.Time",
			JSONFieldName:      "last_review",
			ProtobufFieldName:  "last_review",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "last_reviewer",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "LastReviewer",
			GoFieldType:        "string",
			JSONFieldName:      "last_reviewer",
			ProtobufFieldName:  "last_reviewer",
			ProtobufType:       "string",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "reviewer_notes",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       1024,
			GoFieldName:        "ReviewerNotes",
			GoFieldType:        "string",
			JSONFieldName:      "reviewer_notes",
			ProtobufFieldName:  "reviewer_notes",
			ProtobufType:       "string",
			ProtobufPos:        32,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BugReports) TableName() string {
	return "bug_reports"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BugReports) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BugReports) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BugReports) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BugReports) TableInfo() *TableInfo {
	return bug_reportsTableInfo
}
