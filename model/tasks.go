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


CREATE TABLE `tasks` (
  `id` int(11) unsigned NOT NULL DEFAULT 0,
  `type` tinyint(4) NOT NULL DEFAULT 0,
  `duration` int(11) unsigned NOT NULL DEFAULT 0,
  `duration_code` tinyint(4) NOT NULL DEFAULT 0,
  `title` varchar(100) NOT NULL DEFAULT '',
  `description` text NOT NULL,
  `reward` varchar(64) NOT NULL DEFAULT '',
  `rewardid` int(11) unsigned NOT NULL DEFAULT 0,
  `cashreward` int(11) unsigned NOT NULL DEFAULT 0,
  `xpreward` int(10) NOT NULL DEFAULT 0,
  `rewardmethod` tinyint(3) unsigned NOT NULL DEFAULT 2,
  `minlevel` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `maxlevel` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `repeatable` tinyint(1) unsigned NOT NULL DEFAULT 1,
  `faction_reward` int(10) NOT NULL DEFAULT 0,
  `completion_emote` varchar(128) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "type": 49,    "duration_code": 67,    "cashreward": 88,    "xpreward": 49,    "rewardmethod": 64,    "description": "QVhXAfjZpRhbJHuHFeBNEkCNy",    "rewardid": 94,    "maxlevel": 69,    "completion_emote": "ChGdtXsRYXhUXBPjblSWRYQOH",    "title": "qIqNNtUYhMAjwCFAHAZIAVyUe",    "reward": "KJQdplCFNIcNntchyXmlwTLpn",    "minlevel": 23,    "repeatable": 98,    "id": 4,    "duration": 67,    "faction_reward": 29}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[10] column is set for unsigned
[11] column is set for unsigned
[12] column is set for unsigned
[13] column is set for unsigned



*/

// Tasks struct is a row record of the tasks table in the eqemu database
type Tasks struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ID uint32 `gorm:"primary_key;column:id;type:uint;default:0;" json:"id"`
	//[ 1] type                                           tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Type int32 `gorm:"column:type;type:tinyint;default:0;" json:"type"`
	//[ 2] duration                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Duration uint32 `gorm:"column:duration;type:uint;default:0;" json:"duration"`
	//[ 3] duration_code                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	DurationCode int32 `gorm:"column:duration_code;type:tinyint;default:0;" json:"duration_code"`
	//[ 4] title                                          varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: ['']
	Title string `gorm:"column:title;type:varchar;size:100;default:'';" json:"title"`
	//[ 5] description                                    text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Description string `gorm:"column:description;type:text;size:65535;" json:"description"`
	//[ 6] reward                                         varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Reward string `gorm:"column:reward;type:varchar;size:64;default:'';" json:"reward"`
	//[ 7] rewardid                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Rewardid uint32 `gorm:"column:rewardid;type:uint;default:0;" json:"rewardid"`
	//[ 8] cashreward                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Cashreward uint32 `gorm:"column:cashreward;type:uint;default:0;" json:"cashreward"`
	//[ 9] xpreward                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Xpreward int32 `gorm:"column:xpreward;type:int;default:0;" json:"xpreward"`
	//[10] rewardmethod                                   utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [2]
	Rewardmethod uint32 `gorm:"column:rewardmethod;type:utinyint;default:2;" json:"rewardmethod"`
	//[11] minlevel                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Minlevel uint32 `gorm:"column:minlevel;type:utinyint;default:0;" json:"minlevel"`
	//[12] maxlevel                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Maxlevel uint32 `gorm:"column:maxlevel;type:utinyint;default:0;" json:"maxlevel"`
	//[13] repeatable                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	Repeatable uint32 `gorm:"column:repeatable;type:utinyint;default:1;" json:"repeatable"`
	//[14] faction_reward                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	FactionReward int32 `gorm:"column:faction_reward;type:int;default:0;" json:"faction_reward"`
	//[15] completion_emote                               varchar(128)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 128     default: ['']
	CompletionEmote string `gorm:"column:completion_emote;type:varchar;size:128;default:'';" json:"completion_emote"`
}

var tasksTableInfo = &TableInfo{
	Name: "tasks",
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
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "duration",
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
			GoFieldName:        "Duration",
			GoFieldType:        "uint32",
			JSONFieldName:      "duration",
			ProtobufFieldName:  "duration",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "duration_code",
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
			GoFieldName:        "DurationCode",
			GoFieldType:        "int32",
			JSONFieldName:      "duration_code",
			ProtobufFieldName:  "duration_code",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "title",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "Title",
			GoFieldType:        "string",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "description",
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
			GoFieldName:        "Description",
			GoFieldType:        "string",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "reward",
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
			GoFieldName:        "Reward",
			GoFieldType:        "string",
			JSONFieldName:      "reward",
			ProtobufFieldName:  "reward",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "rewardid",
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
			GoFieldName:        "Rewardid",
			GoFieldType:        "uint32",
			JSONFieldName:      "rewardid",
			ProtobufFieldName:  "rewardid",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "cashreward",
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
			GoFieldName:        "Cashreward",
			GoFieldType:        "uint32",
			JSONFieldName:      "cashreward",
			ProtobufFieldName:  "cashreward",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "xpreward",
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
			GoFieldName:        "Xpreward",
			GoFieldType:        "int32",
			JSONFieldName:      "xpreward",
			ProtobufFieldName:  "xpreward",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "rewardmethod",
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
			GoFieldName:        "Rewardmethod",
			GoFieldType:        "uint32",
			JSONFieldName:      "rewardmethod",
			ProtobufFieldName:  "rewardmethod",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "minlevel",
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
			GoFieldName:        "Minlevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "minlevel",
			ProtobufFieldName:  "minlevel",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "maxlevel",
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
			GoFieldName:        "Maxlevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "maxlevel",
			ProtobufFieldName:  "maxlevel",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "repeatable",
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
			GoFieldName:        "Repeatable",
			GoFieldType:        "uint32",
			JSONFieldName:      "repeatable",
			ProtobufFieldName:  "repeatable",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "faction_reward",
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
			GoFieldName:        "FactionReward",
			GoFieldType:        "int32",
			JSONFieldName:      "faction_reward",
			ProtobufFieldName:  "faction_reward",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "completion_emote",
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
			GoFieldName:        "CompletionEmote",
			GoFieldType:        "string",
			JSONFieldName:      "completion_emote",
			ProtobufFieldName:  "completion_emote",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *Tasks) TableName() string {
	return "tasks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *Tasks) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *Tasks) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *Tasks) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *Tasks) TableInfo() *TableInfo {
	return tasksTableInfo
}
