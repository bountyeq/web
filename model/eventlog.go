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


CREATE TABLE `eventlog` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `accountname` varchar(30) NOT NULL DEFAULT '',
  `accountid` int(10) unsigned DEFAULT 0,
  `status` int(5) NOT NULL DEFAULT 0,
  `charname` varchar(64) NOT NULL DEFAULT '',
  `target` varchar(64) DEFAULT 'None',
  `time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `descriptiontype` varchar(50) NOT NULL DEFAULT '',
  `description` text NOT NULL,
  `event_nid` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 60,    "event_nid": 92,    "charname": "lgCxATuitmfjAhCHIUwyliJwr",    "target": "ZoUGexbeoZqgNcAVQHtWvUhOr",    "time": "2191-04-30T04:45:59.687843429-07:00",    "descriptiontype": "XeXrnYGnddPTNqvtioKkHvDKS",    "description": "lempmKcWYrYKQromJKfvZBRSd",    "accountname": "KeMABRZwLjfDiNVUFLMWHFiER",    "accountid": 82,    "status": 47}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned



*/

// Eventlog struct is a row record of the eventlog table in the eqemu database
type Eventlog struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] accountname                                    varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: ['']
	Accountname string `gorm:"column:accountname;type:varchar;size:30;default:'';" json:"accountname"`
	//[ 2] accountid                                      uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Accountid null.Int `gorm:"column:accountid;type:uint;default:0;" json:"accountid"`
	//[ 3] status                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Status int32 `gorm:"column:status;type:int;default:0;" json:"status"`
	//[ 4] charname                                       varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Charname string `gorm:"column:charname;type:varchar;size:64;default:'';" json:"charname"`
	//[ 5] target                                         varchar(64)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['None']
	Target null.String `gorm:"column:target;type:varchar;size:64;default:'None';" json:"target"`
	//[ 6] time                                           timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [current_timestamp()]
	Time time.Time `gorm:"column:time;type:timestamp;" json:"time"`
	//[ 7] descriptiontype                                varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: ['']
	Descriptiontype string `gorm:"column:descriptiontype;type:varchar;size:50;default:'';" json:"descriptiontype"`
	//[ 8] description                                    text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Description string `gorm:"column:description;type:text;size:65535;" json:"description"`
	//[ 9] event_nid                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EventNid int32 `gorm:"column:event_nid;type:int;default:0;" json:"event_nid"`
}

var eventlogTableInfo = &TableInfo{
	Name: "eventlog",
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
			Name:               "accountname",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "Accountname",
			GoFieldType:        "string",
			JSONFieldName:      "accountname",
			ProtobufFieldName:  "accountname",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "accountid",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Accountid",
			GoFieldType:        "null.Int",
			JSONFieldName:      "accountid",
			ProtobufFieldName:  "accountid",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "status",
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
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "charname",
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
			GoFieldName:        "Charname",
			GoFieldType:        "string",
			JSONFieldName:      "charname",
			ProtobufFieldName:  "charname",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "target",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Target",
			GoFieldType:        "null.String",
			JSONFieldName:      "target",
			ProtobufFieldName:  "target",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "Time",
			GoFieldType:        "time.Time",
			JSONFieldName:      "time",
			ProtobufFieldName:  "time",
			ProtobufType:       "uint64",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "descriptiontype",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "Descriptiontype",
			GoFieldType:        "string",
			JSONFieldName:      "descriptiontype",
			ProtobufFieldName:  "descriptiontype",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "event_nid",
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
			GoFieldName:        "EventNid",
			GoFieldType:        "int32",
			JSONFieldName:      "event_nid",
			ProtobufFieldName:  "event_nid",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Eventlog) TableName() string {
	return "eventlog"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Eventlog) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Eventlog) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Eventlog) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Eventlog) TableInfo() *TableInfo {
	return eventlogTableInfo
}
