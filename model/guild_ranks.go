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


CREATE TABLE `guild_ranks` (
  `guild_id` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `rank` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `title` varchar(128) NOT NULL DEFAULT '',
  `can_hear` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `can_speak` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `can_invite` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `can_remove` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `can_promote` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `can_demote` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `can_motd` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `can_warpeace` tinyint(3) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`guild_id`,`rank`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "can_motd": 17,    "can_warpeace": 22,    "guild_id": 3,    "rank": 19,    "can_hear": 72,    "can_promote": 46,    "can_demote": 12,    "title": "UrxFoqhGNisvCLlLfELyOopgj",    "can_speak": 73,    "can_invite": 3,    "can_remove": 29}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned



*/

// GuildRanks struct is a row record of the guild_ranks table in the eqemu database
type GuildRanks struct {
	//[ 0] guild_id                                       umediumint           null: false  primary: true   isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	GuildID uint32 `gorm:"primary_key;column:guild_id;type:umediumint;default:0;" json:"guild_id"`
	//[ 1] rank                                           utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Rank uint32 `gorm:"primary_key;column:rank;type:utinyint;default:0;" json:"rank"`
	//[ 2] title                                          varchar(128)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 128     default: ['']
	Title string `gorm:"column:title;type:varchar;size:128;default:'';" json:"title"`
	//[ 3] can_hear                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	CanHear uint32 `gorm:"column:can_hear;type:utinyint;default:0;" json:"can_hear"`
	//[ 4] can_speak                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	CanSpeak uint32 `gorm:"column:can_speak;type:utinyint;default:0;" json:"can_speak"`
	//[ 5] can_invite                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	CanInvite uint32 `gorm:"column:can_invite;type:utinyint;default:0;" json:"can_invite"`
	//[ 6] can_remove                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	CanRemove uint32 `gorm:"column:can_remove;type:utinyint;default:0;" json:"can_remove"`
	//[ 7] can_promote                                    utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	CanPromote uint32 `gorm:"column:can_promote;type:utinyint;default:0;" json:"can_promote"`
	//[ 8] can_demote                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	CanDemote uint32 `gorm:"column:can_demote;type:utinyint;default:0;" json:"can_demote"`
	//[ 9] can_motd                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	CanMotd uint32 `gorm:"column:can_motd;type:utinyint;default:0;" json:"can_motd"`
	//[10] can_warpeace                                   utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	CanWarpeace uint32 `gorm:"column:can_warpeace;type:utinyint;default:0;" json:"can_warpeace"`
}

var guild_ranksTableInfo = &TableInfo{
	Name: "guild_ranks",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "guild_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "umediumint",
			DatabaseTypePretty: "umediumint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "umediumint",
			ColumnLength:       -1,
			GoFieldName:        "GuildID",
			GoFieldType:        "uint32",
			JSONFieldName:      "guild_id",
			ProtobufFieldName:  "guild_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "rank",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Rank",
			GoFieldType:        "uint32",
			JSONFieldName:      "rank",
			ProtobufFieldName:  "rank",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "title",
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
			GoFieldName:        "Title",
			GoFieldType:        "string",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "can_hear",
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
			GoFieldName:        "CanHear",
			GoFieldType:        "uint32",
			JSONFieldName:      "can_hear",
			ProtobufFieldName:  "can_hear",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "can_speak",
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
			GoFieldName:        "CanSpeak",
			GoFieldType:        "uint32",
			JSONFieldName:      "can_speak",
			ProtobufFieldName:  "can_speak",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "can_invite",
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
			GoFieldName:        "CanInvite",
			GoFieldType:        "uint32",
			JSONFieldName:      "can_invite",
			ProtobufFieldName:  "can_invite",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "can_remove",
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
			GoFieldName:        "CanRemove",
			GoFieldType:        "uint32",
			JSONFieldName:      "can_remove",
			ProtobufFieldName:  "can_remove",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "can_promote",
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
			GoFieldName:        "CanPromote",
			GoFieldType:        "uint32",
			JSONFieldName:      "can_promote",
			ProtobufFieldName:  "can_promote",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "can_demote",
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
			GoFieldName:        "CanDemote",
			GoFieldType:        "uint32",
			JSONFieldName:      "can_demote",
			ProtobufFieldName:  "can_demote",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "can_motd",
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
			GoFieldName:        "CanMotd",
			GoFieldType:        "uint32",
			JSONFieldName:      "can_motd",
			ProtobufFieldName:  "can_motd",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "can_warpeace",
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
			GoFieldName:        "CanWarpeace",
			GoFieldType:        "uint32",
			JSONFieldName:      "can_warpeace",
			ProtobufFieldName:  "can_warpeace",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *GuildRanks) TableName() string {
	return "guild_ranks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GuildRanks) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GuildRanks) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GuildRanks) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GuildRanks) TableInfo() *TableInfo {
	return guild_ranksTableInfo
}
