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


CREATE TABLE `login_accounts` (
  `id` int(11) unsigned NOT NULL,
  `account_name` varchar(50) NOT NULL,
  `account_password` text NOT NULL,
  `account_email` varchar(100) NOT NULL,
  `source_loginserver` varchar(64) DEFAULT NULL,
  `last_ip_address` varchar(15) NOT NULL,
  `last_login_date` datetime NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `source_loginserver_account_name` (`source_loginserver`,`account_name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "account_password": "REZuRKCtkiPifvZKNjFirGwwT",    "last_ip_address": "tXkcElaucRqRwXDWBPfpJLIRB",    "updated_at": "2094-06-21T21:43:48.885112973-07:00",    "account_name": "ZMZBMxeWqrwxgiBsCScVdhplE",    "account_email": "YHJQiikohmdndmHOgtEoUYlhp",    "source_loginserver": "qPMOxfDepIjgPpWJccItSCUTK",    "last_login_date": "2181-07-02T12:40:37.056746056-07:00",    "created_at": "2265-04-21T21:39:15.076553405-07:00",    "id": 91}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// LoginAccounts struct is a row record of the login_accounts table in the eqemu database
type LoginAccounts struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;column:id;type:uint;" json:"id"`
	//[ 1] account_name                                   varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	AccountName string `gorm:"column:account_name;type:varchar;size:50;" json:"account_name"`
	//[ 2] account_password                               text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	AccountPassword string `gorm:"column:account_password;type:text;size:65535;" json:"account_password"`
	//[ 3] account_email                                  varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	AccountEmail string `gorm:"column:account_email;type:varchar;size:100;" json:"account_email"`
	//[ 4] source_loginserver                             varchar(64)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 64      default: [NULL]
	SourceLoginserver null.String `gorm:"column:source_loginserver;type:varchar;size:64;" json:"source_loginserver"`
	//[ 5] last_ip_address                                varchar(15)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 15      default: []
	LastIPAddress string `gorm:"column:last_ip_address;type:varchar;size:15;" json:"last_ip_address"`
	//[ 6] last_login_date                                datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	LastLoginDate time.Time `gorm:"column:last_login_date;type:datetime;" json:"last_login_date"`
	//[ 7] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [NULL]
	CreatedAt null.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[ 8] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [current_timestamp()]
	UpdatedAt null.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var login_accountsTableInfo = &TableInfo{
	Name: "login_accounts",
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
			IsAutoIncrement:    false,
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
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
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
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "AccountPassword",
			GoFieldType:        "string",
			JSONFieldName:      "account_password",
			ProtobufFieldName:  "account_password",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "account_email",
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
			GoFieldName:        "AccountEmail",
			GoFieldType:        "string",
			JSONFieldName:      "account_email",
			ProtobufFieldName:  "account_email",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "source_loginserver",
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
			GoFieldName:        "SourceLoginserver",
			GoFieldType:        "null.String",
			JSONFieldName:      "source_loginserver",
			ProtobufFieldName:  "source_loginserver",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "last_ip_address",
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
			GoFieldName:        "LastIPAddress",
			GoFieldType:        "string",
			JSONFieldName:      "last_ip_address",
			ProtobufFieldName:  "last_ip_address",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "last_login_date",
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
			GoFieldName:        "LastLoginDate",
			GoFieldType:        "time.Time",
			JSONFieldName:      "last_login_date",
			ProtobufFieldName:  "last_login_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LoginAccounts) TableName() string {
	return "login_accounts"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LoginAccounts) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LoginAccounts) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LoginAccounts) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LoginAccounts) TableInfo() *TableInfo {
	return login_accountsTableInfo
}
