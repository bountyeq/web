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


CREATE TABLE `group_leaders` (
  `gid` int(4) NOT NULL DEFAULT 0,
  `leadername` varchar(64) NOT NULL DEFAULT '',
  `marknpc` varchar(64) NOT NULL DEFAULT '',
  `leadershipaa` tinyblob DEFAULT NULL,
  `maintank` varchar(64) NOT NULL DEFAULT '',
  `assist` varchar(64) NOT NULL DEFAULT '',
  `puller` varchar(64) NOT NULL DEFAULT '',
  `mentoree` varchar(64) NOT NULL,
  `mentor_percent` int(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`gid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "leadername": "RVAtmNEyXlXMbYxPMKpOQhbpG",    "marknpc": "fCQrTTvjJYJPdnxfoqsLAwklb",    "leadershipaa": "QVI1CCgUFwctKygUHSEzETxYWFAHEAVVMhs8JmEiQERTJ2MTJGERBkcyHjIFACEXIVNcOzkED0IDAQUHSgoYYjAYTzFXTzopASRaADleV08ACBorF0kK",    "maintank": "cLaKusdUyYrwNOxohkcmSTMkb",    "mentor_percent": 26,    "gid": 92,    "assist": "bYSLQVTgyhFtVrrpTOloWUBso",    "puller": "CMsDCuoPGdUppQPucpZCqwMEU",    "mentoree": "ectOjqTxSumZYeCePiQDXcdDb"}



*/

// GroupLeaders struct is a row record of the group_leaders table in the eqemu database
type GroupLeaders struct {
	//[ 0] gid                                            int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Gid int32 `gorm:"primary_key;column:gid;type:int;default:0;" json:"gid"`
	//[ 1] leadername                                     varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Leadername string `gorm:"column:leadername;type:varchar;size:64;default:'';" json:"leadername"`
	//[ 2] marknpc                                        varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Marknpc string `gorm:"column:marknpc;type:varchar;size:64;default:'';" json:"marknpc"`
	//[ 3] leadershipaa                                   blob                 null: true   primary: false  isArray: false  auto: false  col: blob            len: -1      default: [NULL]
	Leadershipaa []byte `gorm:"column:leadershipaa;type:blob;" json:"leadershipaa"`
	//[ 4] maintank                                       varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Maintank string `gorm:"column:maintank;type:varchar;size:64;default:'';" json:"maintank"`
	//[ 5] assist                                         varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Assist string `gorm:"column:assist;type:varchar;size:64;default:'';" json:"assist"`
	//[ 6] puller                                         varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Puller string `gorm:"column:puller;type:varchar;size:64;default:'';" json:"puller"`
	//[ 7] mentoree                                       varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Mentoree string `gorm:"column:mentoree;type:varchar;size:64;" json:"mentoree"`
	//[ 8] mentor_percent                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MentorPercent int32 `gorm:"column:mentor_percent;type:int;default:0;" json:"mentor_percent"`
}

var group_leadersTableInfo = &TableInfo{
	Name: "group_leaders",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "gid",
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
			GoFieldName:        "Gid",
			GoFieldType:        "int32",
			JSONFieldName:      "gid",
			ProtobufFieldName:  "gid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "leadername",
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
			GoFieldName:        "Leadername",
			GoFieldType:        "string",
			JSONFieldName:      "leadername",
			ProtobufFieldName:  "leadername",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "marknpc",
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
			GoFieldName:        "Marknpc",
			GoFieldType:        "string",
			JSONFieldName:      "marknpc",
			ProtobufFieldName:  "marknpc",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "leadershipaa",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "blob",
			DatabaseTypePretty: "blob",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "blob",
			ColumnLength:       -1,
			GoFieldName:        "Leadershipaa",
			GoFieldType:        "[]byte",
			JSONFieldName:      "leadershipaa",
			ProtobufFieldName:  "leadershipaa",
			ProtobufType:       "bytes",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "maintank",
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
			GoFieldName:        "Maintank",
			GoFieldType:        "string",
			JSONFieldName:      "maintank",
			ProtobufFieldName:  "maintank",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "assist",
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
			GoFieldName:        "Assist",
			GoFieldType:        "string",
			JSONFieldName:      "assist",
			ProtobufFieldName:  "assist",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "puller",
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
			GoFieldName:        "Puller",
			GoFieldType:        "string",
			JSONFieldName:      "puller",
			ProtobufFieldName:  "puller",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "mentoree",
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
			GoFieldName:        "Mentoree",
			GoFieldType:        "string",
			JSONFieldName:      "mentoree",
			ProtobufFieldName:  "mentoree",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "mentor_percent",
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
			GoFieldName:        "MentorPercent",
			GoFieldType:        "int32",
			JSONFieldName:      "mentor_percent",
			ProtobufFieldName:  "mentor_percent",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *GroupLeaders) TableName() string {
	return "group_leaders"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GroupLeaders) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GroupLeaders) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GroupLeaders) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GroupLeaders) TableInfo() *TableInfo {
	return group_leadersTableInfo
}
