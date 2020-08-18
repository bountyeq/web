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


CREATE TABLE `raid_leaders` (
  `gid` int(4) unsigned NOT NULL,
  `rid` int(4) unsigned NOT NULL,
  `marknpc` varchar(64) NOT NULL,
  `maintank` varchar(64) NOT NULL,
  `assist` varchar(64) NOT NULL,
  `puller` varchar(64) NOT NULL,
  `leadershipaa` tinyblob NOT NULL,
  `mentoree` varchar(64) NOT NULL,
  `mentor_percent` int(4) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "maintank": "PbrNQKokMoLXLlVxjClGZNYOQ",    "leadershipaa": "N0QWFFtWRzFjXE8+DEFeID9aKgo3QjQ3B0kYOyRCUkMOYRAKCDMEPTMKLlhXAEQsThkqXBNbKDxcYWEDJksNEEg=",    "mentoree": "NjfipXLcCNngOfVuiltTFnEUO",    "mentor_percent": 75,    "gid": 58,    "rid": 75,    "marknpc": "rkPrrOBXrfVLbUOvQhhRISMGL",    "assist": "qetXYPuToZMQukFrJMJUOlcmR",    "puller": "cGZyihZmFVtOgRKWZwXfDuDXQ"}


Comments
-------------------------------------
[ 0] column is set for unsignedWarning table: raid_leaders does not have a primary key defined, setting col position 1 gid as primary key

[ 1] column is set for unsigned



*/

// RaidLeaders struct is a row record of the raid_leaders table in the eqemu database
type RaidLeaders struct {
	//[ 0] gid                                            uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	Gid uint32 `gorm:"primary_key;column:gid;type:uint;" json:"gid"`
	//[ 1] rid                                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	Rid uint32 `gorm:"column:rid;type:uint;" json:"rid"`
	//[ 2] marknpc                                        varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Marknpc string `gorm:"column:marknpc;type:varchar;size:64;" json:"marknpc"`
	//[ 3] maintank                                       varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Maintank string `gorm:"column:maintank;type:varchar;size:64;" json:"maintank"`
	//[ 4] assist                                         varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Assist string `gorm:"column:assist;type:varchar;size:64;" json:"assist"`
	//[ 5] puller                                         varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Puller string `gorm:"column:puller;type:varchar;size:64;" json:"puller"`
	//[ 6] leadershipaa                                   blob                 null: false  primary: false  isArray: false  auto: false  col: blob            len: -1      default: []
	Leadershipaa []byte `gorm:"column:leadershipaa;type:blob;" json:"leadershipaa"`
	//[ 7] mentoree                                       varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Mentoree string `gorm:"column:mentoree;type:varchar;size:64;" json:"mentoree"`
	//[ 8] mentor_percent                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MentorPercent int32 `gorm:"column:mentor_percent;type:int;default:0;" json:"mentor_percent"`
}

var raid_leadersTableInfo = &TableInfo{
	Name: "raid_leaders",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "gid",
			Comment: ``,
			Notes: `column is set for unsignedWarning table: raid_leaders does not have a primary key defined, setting col position 1 gid as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Gid",
			GoFieldType:        "uint32",
			JSONFieldName:      "gid",
			ProtobufFieldName:  "gid",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "rid",
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
			GoFieldName:        "Rid",
			GoFieldType:        "uint32",
			JSONFieldName:      "rid",
			ProtobufFieldName:  "rid",
			ProtobufType:       "uint32",
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "leadershipaa",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
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
func (r *RaidLeaders) TableName() string {
	return "raid_leaders"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RaidLeaders) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RaidLeaders) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RaidLeaders) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RaidLeaders) TableInfo() *TableInfo {
	return raid_leadersTableInfo
}
