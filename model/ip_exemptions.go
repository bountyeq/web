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


CREATE TABLE `ip_exemptions` (
  `exemption_id` int(11) NOT NULL AUTO_INCREMENT,
  `exemption_ip` varchar(255) DEFAULT NULL,
  `exemption_amount` int(11) DEFAULT NULL,
  PRIMARY KEY (`exemption_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "exemption_amount": 53,    "exemption_id": 46,    "exemption_ip": "sLouaYJwiBjDbREVpLmpqNpBe"}



*/

// IPExemptions struct is a row record of the ip_exemptions table in the eqemu database
type IPExemptions struct {
	//[ 0] exemption_id                                   int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ExemptionID int32 `gorm:"primary_key;AUTO_INCREMENT;column:exemption_id;type:int;" json:"exemption_id"`
	//[ 1] exemption_ip                                   varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: [NULL]
	ExemptionIP null.String `gorm:"column:exemption_ip;type:varchar;size:255;" json:"exemption_ip"`
	//[ 2] exemption_amount                               int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	ExemptionAmount null.Int `gorm:"column:exemption_amount;type:int;" json:"exemption_amount"`
}

var ip_exemptionsTableInfo = &TableInfo{
	Name: "ip_exemptions",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "exemption_id",
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
			GoFieldName:        "ExemptionID",
			GoFieldType:        "int32",
			JSONFieldName:      "exemption_id",
			ProtobufFieldName:  "exemption_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "exemption_ip",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ExemptionIP",
			GoFieldType:        "null.String",
			JSONFieldName:      "exemption_ip",
			ProtobufFieldName:  "exemption_ip",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "exemption_amount",
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
			GoFieldName:        "ExemptionAmount",
			GoFieldType:        "null.Int",
			JSONFieldName:      "exemption_amount",
			ProtobufFieldName:  "exemption_amount",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (i *IPExemptions) TableName() string {
	return "ip_exemptions"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *IPExemptions) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *IPExemptions) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *IPExemptions) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *IPExemptions) TableInfo() *TableInfo {
	return ip_exemptionsTableInfo
}
