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


CREATE TABLE `gm_ips` (
  `name` varchar(64) NOT NULL,
  `account_id` int(11) NOT NULL,
  `ip_address` varchar(15) NOT NULL,
  UNIQUE KEY `account_id` (`account_id`,`ip_address`),
  UNIQUE KEY `account_id_2` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "name": "PKiVdJDEocfTjBwlIsttKcZUu",    "account_id": 69,    "ip_address": "EmCbYDBaAvkoGrPbaLgukeQyF"}


Comments
-------------------------------------
[ 0] Warning table: gm_ips does not have a primary key defined, setting col position 1 name as primary key




*/

// GmIps struct is a row record of the gm_ips table in the eqemu database
type GmIps struct {
	//[ 0] name                                           varchar(64)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 64      default: []
	Name string `gorm:"primary_key;column:name;type:varchar;size:64;" json:"name"`
	//[ 1] account_id                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AccountID int32 `gorm:"column:account_id;type:int;" json:"account_id"`
	//[ 2] ip_address                                     varchar(15)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 15      default: []
	IPAddress string `gorm:"column:ip_address;type:varchar;size:15;" json:"ip_address"`
}

var gm_ipsTableInfo = &TableInfo{
	Name: "gm_ips",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "name",
			Comment: ``,
			Notes: `Warning table: gm_ips does not have a primary key defined, setting col position 1 name as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "account_id",
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
			GoFieldName:        "AccountID",
			GoFieldType:        "int32",
			JSONFieldName:      "account_id",
			ProtobufFieldName:  "account_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ip_address",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(15)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       15,
			GoFieldName:        "IPAddress",
			GoFieldType:        "string",
			JSONFieldName:      "ip_address",
			ProtobufFieldName:  "ip_address",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *GmIps) TableName() string {
	return "gm_ips"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GmIps) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GmIps) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GmIps) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GmIps) TableInfo() *TableInfo {
	return gm_ipsTableInfo
}
