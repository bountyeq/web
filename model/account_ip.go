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


CREATE TABLE `account_ip` (
  `accid` int(11) NOT NULL DEFAULT 0,
  `ip` varchar(32) NOT NULL DEFAULT '',
  `count` int(11) NOT NULL DEFAULT 1,
  `lastused` timestamp NOT NULL DEFAULT current_timestamp(),
  UNIQUE KEY `ip` (`accid`,`ip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "lastused": "2214-07-27T03:58:28.776935611-07:00",    "accid": 34,    "ip": "kXquXnbrjdwVhWhebkOPFAdLF",    "count": 0}


Comments
-------------------------------------
[ 0] Warning table: account_ip does not have a primary key defined, setting col position 1 accid as primary key




*/

// AccountIP struct is a row record of the account_ip table in the eqemu database
type AccountIP struct {
	//[ 0] accid                                          int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Accid int32 `gorm:"primary_key;column:accid;type:int;default:0;" json:"accid"`
	//[ 1] ip                                             varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['']
	IP string `gorm:"column:ip;type:varchar;size:32;default:'';" json:"ip"`
	//[ 2] count                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	Count int32 `gorm:"column:count;type:int;default:1;" json:"count"`
	//[ 3] lastused                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [current_timestamp()]
	Lastused time.Time `gorm:"column:lastused;type:timestamp;" json:"lastused"`
}

var account_ipTableInfo = &TableInfo{
	Name: "account_ip",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "accid",
			Comment: ``,
			Notes: `Warning table: account_ip does not have a primary key defined, setting col position 1 accid as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Accid",
			GoFieldType:        "int32",
			JSONFieldName:      "accid",
			ProtobufFieldName:  "accid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "ip",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "IP",
			GoFieldType:        "string",
			JSONFieldName:      "ip",
			ProtobufFieldName:  "ip",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "count",
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
			GoFieldName:        "Count",
			GoFieldType:        "int32",
			JSONFieldName:      "count",
			ProtobufFieldName:  "count",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "lastused",
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
			GoFieldName:        "Lastused",
			GoFieldType:        "time.Time",
			JSONFieldName:      "lastused",
			ProtobufFieldName:  "lastused",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AccountIP) TableName() string {
	return "account_ip"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AccountIP) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AccountIP) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AccountIP) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AccountIP) TableInfo() *TableInfo {
	return account_ipTableInfo
}
