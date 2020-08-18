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


CREATE TABLE `mail` (
  `msgid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `charid` int(10) unsigned NOT NULL DEFAULT 0,
  `timestamp` int(11) NOT NULL DEFAULT 0,
  `from` varchar(100) NOT NULL DEFAULT '',
  `subject` varchar(200) NOT NULL DEFAULT '',
  `body` text NOT NULL,
  `to` text NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`msgid`),
  KEY `charid` (`charid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "to": "QDCSWbnnGkrCxkbqtyACPlSKW",    "status": 41,    "msgid": 60,    "charid": 51,    "timestamp": 93,    "from": "hmqglohBhKYqDVXLCOaWgnmMY",    "subject": "MlncPDrgyhjHdJGAZpOodSAkk",    "body": "JpPebcyNVNrMxrfPDOfAQgtTw"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// Mail struct is a row record of the mail table in the eqemu database
type Mail struct {
	//[ 0] msgid                                          uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	Msgid uint32 `gorm:"primary_key;AUTO_INCREMENT;column:msgid;type:uint;" json:"msgid"`
	//[ 1] charid                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Charid uint32 `gorm:"column:charid;type:uint;default:0;" json:"charid"`
	//[ 2] timestamp                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Timestamp int32 `gorm:"column:timestamp;type:int;default:0;" json:"timestamp"`
	//[ 3] from                                           varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: ['']
	From string `gorm:"column:from;type:varchar;size:100;default:'';" json:"from"`
	//[ 4] subject                                        varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: ['']
	Subject string `gorm:"column:subject;type:varchar;size:200;default:'';" json:"subject"`
	//[ 5] body                                           text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Body string `gorm:"column:body;type:text;size:65535;" json:"body"`
	//[ 6] to                                             text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	To string `gorm:"column:to;type:text;size:65535;" json:"to"`
	//[ 7] status                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Status int32 `gorm:"column:status;type:tinyint;default:0;" json:"status"`
}

var mailTableInfo = &TableInfo{
	Name: "mail",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "msgid",
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
			GoFieldName:        "Msgid",
			GoFieldType:        "uint32",
			JSONFieldName:      "msgid",
			ProtobufFieldName:  "msgid",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "charid",
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
			GoFieldName:        "Charid",
			GoFieldType:        "uint32",
			JSONFieldName:      "charid",
			ProtobufFieldName:  "charid",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "timestamp",
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
			GoFieldName:        "Timestamp",
			GoFieldType:        "int32",
			JSONFieldName:      "timestamp",
			ProtobufFieldName:  "timestamp",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "from",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "From",
			GoFieldType:        "string",
			JSONFieldName:      "from",
			ProtobufFieldName:  "from",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "subject",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "Subject",
			GoFieldType:        "string",
			JSONFieldName:      "subject",
			ProtobufFieldName:  "subject",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "body",
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
			GoFieldName:        "Body",
			GoFieldType:        "string",
			JSONFieldName:      "body",
			ProtobufFieldName:  "body",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "to",
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
			GoFieldName:        "To",
			GoFieldType:        "string",
			JSONFieldName:      "to",
			ProtobufFieldName:  "to",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "status",
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
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *Mail) TableName() string {
	return "mail"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *Mail) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *Mail) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *Mail) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *Mail) TableInfo() *TableInfo {
	return mailTableInfo
}
