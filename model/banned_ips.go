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


CREATE TABLE `banned_ips` (
  `ip_address` varchar(32) NOT NULL,
  `notes` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`ip_address`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "ip_address": "FMQJQQVJJyojcZIcqBSQONmaK",    "notes": "NFRhHyqikYLDrqPWhWQMvaZdi"}



*/

// BannedIps struct is a row record of the banned_ips table in the eqemu database
type BannedIps struct {
	//[ 0] ip_address                                     varchar(32)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 32      default: []
	IPAddress string `gorm:"primary_key;column:ip_address;type:varchar;size:32;" json:"ip_address"`
	//[ 1] notes                                          varchar(32)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 32      default: [NULL]
	Notes null.String `gorm:"column:notes;type:varchar;size:32;" json:"notes"`
}

var banned_ipsTableInfo = &TableInfo{
	Name: "banned_ips",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ip_address",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "IPAddress",
			GoFieldType:        "string",
			JSONFieldName:      "ip_address",
			ProtobufFieldName:  "ip_address",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "notes",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Notes",
			GoFieldType:        "null.String",
			JSONFieldName:      "notes",
			ProtobufFieldName:  "notes",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BannedIps) TableName() string {
	return "banned_ips"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BannedIps) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BannedIps) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BannedIps) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BannedIps) TableInfo() *TableInfo {
	return banned_ipsTableInfo
}
