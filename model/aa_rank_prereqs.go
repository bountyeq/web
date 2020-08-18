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


CREATE TABLE `aa_rank_prereqs` (
  `rank_id` int(10) unsigned NOT NULL,
  `aa_id` int(10) NOT NULL,
  `points` int(10) NOT NULL,
  PRIMARY KEY (`rank_id`,`aa_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "rank_id": 6,    "aa_id": 15,    "points": 50}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// AaRankPrereqs struct is a row record of the aa_rank_prereqs table in the eqemu database
type AaRankPrereqs struct {
	//[ 0] rank_id                                        uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	RankID uint32 `gorm:"primary_key;column:rank_id;type:uint;" json:"rank_id"`
	//[ 1] aa_id                                          int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	AaID int32 `gorm:"primary_key;column:aa_id;type:int;" json:"aa_id"`
	//[ 2] points                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Points int32 `gorm:"column:points;type:int;" json:"points"`
}

var aa_rank_prereqsTableInfo = &TableInfo{
	Name: "aa_rank_prereqs",
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
			Name:               "aa_id",
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
			GoFieldName:        "AaID",
			GoFieldType:        "int32",
			JSONFieldName:      "aa_id",
			ProtobufFieldName:  "aa_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "points",
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
			GoFieldName:        "Points",
			GoFieldType:        "int32",
			JSONFieldName:      "points",
			ProtobufFieldName:  "points",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AaRankPrereqs) TableName() string {
	return "aa_rank_prereqs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AaRankPrereqs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AaRankPrereqs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AaRankPrereqs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AaRankPrereqs) TableInfo() *TableInfo {
	return aa_rank_prereqsTableInfo
}
