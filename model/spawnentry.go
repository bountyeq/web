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


CREATE TABLE `spawnentry` (
  `spawngroupID` int(11) NOT NULL DEFAULT 0,
  `npcID` int(11) NOT NULL DEFAULT 0,
  `chance` smallint(4) NOT NULL DEFAULT 0,
  `condition_value_filter` mediumint(9) NOT NULL DEFAULT 1,
  PRIMARY KEY (`spawngroupID`,`npcID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "spawngroup_id": 43,    "npc_id": 30,    "chance": 83,    "condition_value_filter": 31}



*/

// Spawnentry struct is a row record of the spawnentry table in the eqemu database
type Spawnentry struct {
	//[ 0] spawngroupID                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	SpawngroupID int32 `gorm:"primary_key;column:spawngroupID;type:int;default:0;" json:"spawngroup_id"`
	//[ 1] npcID                                          int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	NpcID int32 `gorm:"primary_key;column:npcID;type:int;default:0;" json:"npc_id"`
	//[ 2] chance                                         smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Chance int32 `gorm:"column:chance;type:smallint;default:0;" json:"chance"`
	//[ 3] condition_value_filter                         mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [1]
	ConditionValueFilter int32 `gorm:"column:condition_value_filter;type:mediumint;default:1;" json:"condition_value_filter"`
}

var spawnentryTableInfo = &TableInfo{
	Name: "spawnentry",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "spawngroupID",
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
			GoFieldName:        "SpawngroupID",
			GoFieldType:        "int32",
			JSONFieldName:      "spawngroup_id",
			ProtobufFieldName:  "spawngroup_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "npcID",
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
			GoFieldName:        "NpcID",
			GoFieldType:        "int32",
			JSONFieldName:      "npc_id",
			ProtobufFieldName:  "npc_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "chance",
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
			GoFieldName:        "Chance",
			GoFieldType:        "int32",
			JSONFieldName:      "chance",
			ProtobufFieldName:  "chance",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "condition_value_filter",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "mediumint",
			DatabaseTypePretty: "mediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "mediumint",
			ColumnLength:       -1,
			GoFieldName:        "ConditionValueFilter",
			GoFieldType:        "int32",
			JSONFieldName:      "condition_value_filter",
			ProtobufFieldName:  "condition_value_filter",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *Spawnentry) TableName() string {
	return "spawnentry"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *Spawnentry) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *Spawnentry) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *Spawnentry) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *Spawnentry) TableInfo() *TableInfo {
	return spawnentryTableInfo
}
