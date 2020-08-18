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


CREATE TABLE `hackers` (
  `id` int(4) NOT NULL AUTO_INCREMENT,
  `account` text NOT NULL,
  `name` text NOT NULL,
  `hacked` text NOT NULL,
  `zone` text DEFAULT NULL,
  `date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "name": "HRycUPXHksxdpYKptSvGWrnHi",    "hacked": "fuhMGLSdJjuqwqDrNHTtcwEip",    "zone": "OsRLifknefVgSUyNwDXqPmWKU",    "date": "2102-08-18T11:29:32.63334672-07:00",    "id": 7,    "account": "cFaoFIWEyccgHvTLMjhWDslUV"}



*/

// Hackers struct is a row record of the hackers table in the eqemu database
type Hackers struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] account                                        text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Account string `gorm:"column:account;type:text;size:65535;" json:"account"`
	//[ 2] name                                           text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Name string `gorm:"column:name;type:text;size:65535;" json:"name"`
	//[ 3] hacked                                         text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Hacked string `gorm:"column:hacked;type:text;size:65535;" json:"hacked"`
	//[ 4] zone                                           text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: [NULL]
	Zone null.String `gorm:"column:zone;type:text;size:65535;" json:"zone"`
	//[ 5] date                                           timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [current_timestamp()]
	Date time.Time `gorm:"column:date;type:timestamp;" json:"date"`
}

var hackersTableInfo = &TableInfo{
	Name: "hackers",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "account",
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
			GoFieldName:        "Account",
			GoFieldType:        "string",
			JSONFieldName:      "account",
			ProtobufFieldName:  "account",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "hacked",
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
			GoFieldName:        "Hacked",
			GoFieldType:        "string",
			JSONFieldName:      "hacked",
			ProtobufFieldName:  "hacked",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "zone",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Zone",
			GoFieldType:        "null.String",
			JSONFieldName:      "zone",
			ProtobufFieldName:  "zone",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "date",
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
			GoFieldName:        "Date",
			GoFieldType:        "time.Time",
			JSONFieldName:      "date",
			ProtobufFieldName:  "date",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (h *Hackers) TableName() string {
	return "hackers"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (h *Hackers) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (h *Hackers) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (h *Hackers) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (h *Hackers) TableInfo() *TableInfo {
	return hackersTableInfo
}
