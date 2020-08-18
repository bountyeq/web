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


CREATE TABLE `guild_members` (
  `char_id` int(11) NOT NULL DEFAULT 0,
  `guild_id` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `rank` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `tribute_enable` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `total_tribute` int(10) unsigned NOT NULL DEFAULT 0,
  `last_tribute` int(10) unsigned NOT NULL DEFAULT 0,
  `banker` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `public_note` text NOT NULL,
  `alt` tinyint(3) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`char_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "guild_id": 88,    "total_tribute": 96,    "last_tribute": 3,    "public_note": "PnZtRTEuoyCiQSXEWwWWSWphN",    "alt": 17,    "char_id": 69,    "rank": 14,    "tribute_enable": 19,    "banker": 83}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 8] column is set for unsigned



*/

// GuildMembers struct is a row record of the guild_members table in the eqemu database
type GuildMembers struct {
	//[ 0] char_id                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	CharID int32 `gorm:"primary_key;column:char_id;type:int;default:0;" json:"char_id"`
	//[ 1] guild_id                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	GuildID uint32 `gorm:"column:guild_id;type:umediumint;default:0;" json:"guild_id"`
	//[ 2] rank                                           utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Rank uint32 `gorm:"column:rank;type:utinyint;default:0;" json:"rank"`
	//[ 3] tribute_enable                                 utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	TributeEnable uint32 `gorm:"column:tribute_enable;type:utinyint;default:0;" json:"tribute_enable"`
	//[ 4] total_tribute                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TotalTribute uint32 `gorm:"column:total_tribute;type:uint;default:0;" json:"total_tribute"`
	//[ 5] last_tribute                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTribute uint32 `gorm:"column:last_tribute;type:uint;default:0;" json:"last_tribute"`
	//[ 6] banker                                         utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Banker uint32 `gorm:"column:banker;type:utinyint;default:0;" json:"banker"`
	//[ 7] public_note                                    text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	PublicNote string `gorm:"column:public_note;type:text;size:65535;" json:"public_note"`
	//[ 8] alt                                            utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Alt uint32 `gorm:"column:alt;type:utinyint;default:0;" json:"alt"`
}

var guild_membersTableInfo = &TableInfo{
	Name: "guild_members",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "char_id",
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
			GoFieldName:        "CharID",
			GoFieldType:        "int32",
			JSONFieldName:      "char_id",
			ProtobufFieldName:  "char_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "guild_id",
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
			GoFieldName:        "GuildID",
			GoFieldType:        "uint32",
			JSONFieldName:      "guild_id",
			ProtobufFieldName:  "guild_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "rank",
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
			GoFieldName:        "Rank",
			GoFieldType:        "uint32",
			JSONFieldName:      "rank",
			ProtobufFieldName:  "rank",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "tribute_enable",
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
			GoFieldName:        "TributeEnable",
			GoFieldType:        "uint32",
			JSONFieldName:      "tribute_enable",
			ProtobufFieldName:  "tribute_enable",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "total_tribute",
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
			GoFieldName:        "TotalTribute",
			GoFieldType:        "uint32",
			JSONFieldName:      "total_tribute",
			ProtobufFieldName:  "total_tribute",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "last_tribute",
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
			GoFieldName:        "LastTribute",
			GoFieldType:        "uint32",
			JSONFieldName:      "last_tribute",
			ProtobufFieldName:  "last_tribute",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "banker",
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
			GoFieldName:        "Banker",
			GoFieldType:        "uint32",
			JSONFieldName:      "banker",
			ProtobufFieldName:  "banker",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "public_note",
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
			GoFieldName:        "PublicNote",
			GoFieldType:        "string",
			JSONFieldName:      "public_note",
			ProtobufFieldName:  "public_note",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "alt",
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
			GoFieldName:        "Alt",
			GoFieldType:        "uint32",
			JSONFieldName:      "alt",
			ProtobufFieldName:  "alt",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *GuildMembers) TableName() string {
	return "guild_members"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GuildMembers) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GuildMembers) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GuildMembers) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GuildMembers) TableInfo() *TableInfo {
	return guild_membersTableInfo
}
