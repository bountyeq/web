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


CREATE TABLE `login_server_admins` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `account_name` varchar(30) NOT NULL,
  `account_password` varchar(255) NOT NULL,
  `first_name` varchar(50) NOT NULL,
  `last_name` varchar(50) NOT NULL,
  `email` varchar(100) NOT NULL,
  `registration_date` datetime NOT NULL,
  `registration_ip_address` varchar(15) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "registration_ip_address": "oyTwFWHLPfEUZhMLwIbPTyMgX",    "id": 79,    "account_name": "ZUPQwNNsiJtcmFKguPOlwMTYC",    "account_password": "ufMfnuVhpOtnbJGlLrMxYeflF",    "first_name": "bgvnZhapPXZFCBvLqvteUVbQn",    "last_name": "ATRkoeJpWOTwTeEGsNlCLbvWf",    "email": "NulxrYeKFBdAfhgVxcQPovubO",    "registration_date": "2244-12-04T00:52:43.293586133-08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// LoginServerAdmins struct is a row record of the login_server_admins table in the eqemu database
type LoginServerAdmins struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] account_name                                   varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	AccountName string `gorm:"column:account_name;type:varchar;size:30;" json:"account_name"`
	//[ 2] account_password                               varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	AccountPassword string `gorm:"column:account_password;type:varchar;size:255;" json:"account_password"`
	//[ 3] first_name                                     varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	FirstName string `gorm:"column:first_name;type:varchar;size:50;" json:"first_name"`
	//[ 4] last_name                                      varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	LastName string `gorm:"column:last_name;type:varchar;size:50;" json:"last_name"`
	//[ 5] email                                          varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Email string `gorm:"column:email;type:varchar;size:100;" json:"email"`
	//[ 6] registration_date                              datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	RegistrationDate time.Time `gorm:"column:registration_date;type:datetime;" json:"registration_date"`
	//[ 7] registration_ip_address                        varchar(15)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 15      default: []
	RegistrationIPAddress string `gorm:"column:registration_ip_address;type:varchar;size:15;" json:"registration_ip_address"`
}

var login_server_adminsTableInfo = &TableInfo{
	Name: "login_server_admins",
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
			Name:               "account_name",
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
			GoFieldName:        "AccountName",
			GoFieldType:        "string",
			JSONFieldName:      "account_name",
			ProtobufFieldName:  "account_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "account_password",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "AccountPassword",
			GoFieldType:        "string",
			JSONFieldName:      "account_password",
			ProtobufFieldName:  "account_password",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "first_name",
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
			GoFieldName:        "FirstName",
			GoFieldType:        "string",
			JSONFieldName:      "first_name",
			ProtobufFieldName:  "first_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "last_name",
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
			GoFieldName:        "LastName",
			GoFieldType:        "string",
			JSONFieldName:      "last_name",
			ProtobufFieldName:  "last_name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "email",
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
			GoFieldName:        "Email",
			GoFieldType:        "string",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "registration_date",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "RegistrationDate",
			GoFieldType:        "time.Time",
			JSONFieldName:      "registration_date",
			ProtobufFieldName:  "registration_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "registration_ip_address",
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
			GoFieldName:        "RegistrationIPAddress",
			GoFieldType:        "string",
			JSONFieldName:      "registration_ip_address",
			ProtobufFieldName:  "registration_ip_address",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LoginServerAdmins) TableName() string {
	return "login_server_admins"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LoginServerAdmins) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LoginServerAdmins) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LoginServerAdmins) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LoginServerAdmins) TableInfo() *TableInfo {
	return login_server_adminsTableInfo
}
