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


CREATE TABLE `quest_globals` (
  `charid` int(11) NOT NULL DEFAULT 0,
  `npcid` int(11) NOT NULL DEFAULT 0,
  `zoneid` int(11) NOT NULL DEFAULT 0,
  `name` varchar(65) NOT NULL DEFAULT '',
  `value` varchar(128) NOT NULL DEFAULT '?',
  `expdate` int(11) DEFAULT NULL,
  PRIMARY KEY (`charid`,`npcid`,`zoneid`,`name`),
  UNIQUE KEY `qname` (`name`,`charid`,`npcid`,`zoneid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "name": "fShwrDTiGajDUkeXUKSTHLYRE",    "value": "nRLlGoneeiwPVcukOOJXeBeOM",    "expdate": 22,    "charid": 31,    "npcid": 72,    "zoneid": 33}



*/

// QuestGlobals struct is a row record of the quest_globals table in the eqemu database
type QuestGlobals struct {
	//[ 0] charid                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Charid int32 `gorm:"primary_key;column:charid;type:int;default:0;" json:"charid"`
	//[ 1] npcid                                          int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Npcid int32 `gorm:"primary_key;column:npcid;type:int;default:0;" json:"npcid"`
	//[ 2] zoneid                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Zoneid int32 `gorm:"primary_key;column:zoneid;type:int;default:0;" json:"zoneid"`
	//[ 3] name                                           varchar(65)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 65      default: ['']
	Name string `gorm:"primary_key;column:name;type:varchar;size:65;default:'';" json:"name"`
	//[ 4] value                                          varchar(128)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 128     default: ['?']
	Value string `gorm:"column:value;type:varchar;size:128;default:'?';" json:"value"`
	//[ 5] expdate                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	Expdate null.Int `gorm:"column:expdate;type:int;" json:"expdate"`
}

var quest_globalsTableInfo = &TableInfo{
	Name: "quest_globals",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "charid",
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
			GoFieldName:        "Charid",
			GoFieldType:        "int32",
			JSONFieldName:      "charid",
			ProtobufFieldName:  "charid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "npcid",
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
			GoFieldName:        "Npcid",
			GoFieldType:        "int32",
			JSONFieldName:      "npcid",
			ProtobufFieldName:  "npcid",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "zoneid",
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
			GoFieldName:        "Zoneid",
			GoFieldType:        "int32",
			JSONFieldName:      "zoneid",
			ProtobufFieldName:  "zoneid",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(65)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       65,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "value",
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
			GoFieldName:        "Value",
			GoFieldType:        "string",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "expdate",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Expdate",
			GoFieldType:        "null.Int",
			JSONFieldName:      "expdate",
			ProtobufFieldName:  "expdate",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (q *QuestGlobals) TableName() string {
	return "quest_globals"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *QuestGlobals) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *QuestGlobals) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *QuestGlobals) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *QuestGlobals) TableInfo() *TableInfo {
	return quest_globalsTableInfo
}
