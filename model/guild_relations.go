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


CREATE TABLE `guild_relations` (
  `guild1` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `guild2` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `relation` tinyint(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`guild1`,`guild2`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "guild_1": 89,    "guild_2": 23,    "relation": 72}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// GuildRelations struct is a row record of the guild_relations table in the eqemu database
type GuildRelations struct {
	//[ 0] guild1                                         umediumint           null: false  primary: true   isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Guild1 uint32 `gorm:"primary_key;column:guild1;type:umediumint;default:0;" json:"guild_1"`
	//[ 1] guild2                                         umediumint           null: false  primary: true   isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Guild2 uint32 `gorm:"primary_key;column:guild2;type:umediumint;default:0;" json:"guild_2"`
	//[ 2] relation                                       tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Relation int32 `gorm:"column:relation;type:tinyint;default:0;" json:"relation"`
}

var guild_relationsTableInfo = &TableInfo{
	Name: "guild_relations",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "guild1",
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
			GoFieldName:        "Guild1",
			GoFieldType:        "uint32",
			JSONFieldName:      "guild_1",
			ProtobufFieldName:  "guild_1",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "guild2",
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
			GoFieldName:        "Guild2",
			GoFieldType:        "uint32",
			JSONFieldName:      "guild_2",
			ProtobufFieldName:  "guild_2",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "relation",
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
			GoFieldName:        "Relation",
			GoFieldType:        "int32",
			JSONFieldName:      "relation",
			ProtobufFieldName:  "relation",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *GuildRelations) TableName() string {
	return "guild_relations"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GuildRelations) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GuildRelations) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GuildRelations) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GuildRelations) TableInfo() *TableInfo {
	return guild_relationsTableInfo
}
