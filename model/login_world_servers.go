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


CREATE TABLE `login_world_servers` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `long_name` varchar(100) NOT NULL,
  `short_name` varchar(100) NOT NULL,
  `tag_description` varchar(50) NOT NULL DEFAULT '',
  `login_server_list_type_id` int(11) NOT NULL,
  `last_login_date` datetime DEFAULT NULL,
  `last_ip_address` varchar(15) DEFAULT NULL,
  `login_server_admin_id` int(11) NOT NULL,
  `is_server_trusted` int(11) NOT NULL,
  `note` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "short_name": "HXYGVKnFdLiQuEfGkEbvcaVyl",    "tag_description": "YpDhSJwqGcTHqVXRKbdApCNKO",    "last_login_date": "2045-09-04T02:25:07.276733022-07:00",    "login_server_admin_id": 51,    "id": 29,    "long_name": "LHPNyyxxQgDZxinqHdAkbmtoI",    "login_server_list_type_id": 21,    "last_ip_address": "VJdXhoOxJNSdBukEeHYVvSWZP",    "is_server_trusted": 73,    "note": "GTvCAofiqGyVjhOEoXJPPlfIh"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// LoginWorldServers struct is a row record of the login_world_servers table in the eqemu database
type LoginWorldServers struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] long_name                                      varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	LongName string `gorm:"column:long_name;type:varchar;size:100;" json:"long_name"`
	//[ 2] short_name                                     varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	ShortName string `gorm:"column:short_name;type:varchar;size:100;" json:"short_name"`
	//[ 3] tag_description                                varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: ['']
	TagDescription string `gorm:"column:tag_description;type:varchar;size:50;default:'';" json:"tag_description"`
	//[ 4] login_server_list_type_id                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LoginServerListTypeID int32 `gorm:"column:login_server_list_type_id;type:int;" json:"login_server_list_type_id"`
	//[ 5] last_login_date                                datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [NULL]
	LastLoginDate null.Time `gorm:"column:last_login_date;type:datetime;" json:"last_login_date"`
	//[ 6] last_ip_address                                varchar(15)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 15      default: [NULL]
	LastIPAddress null.String `gorm:"column:last_ip_address;type:varchar;size:15;" json:"last_ip_address"`
	//[ 7] login_server_admin_id                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LoginServerAdminID int32 `gorm:"column:login_server_admin_id;type:int;" json:"login_server_admin_id"`
	//[ 8] is_server_trusted                              int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	IsServerTrusted int32 `gorm:"column:is_server_trusted;type:int;" json:"is_server_trusted"`
	//[ 9] note                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: [NULL]
	Note null.String `gorm:"column:note;type:varchar;size:255;" json:"note"`
}

var login_world_serversTableInfo = &TableInfo{
	Name: "login_world_servers",
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
			Name:               "long_name",
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
			GoFieldName:        "LongName",
			GoFieldType:        "string",
			JSONFieldName:      "long_name",
			ProtobufFieldName:  "long_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "short_name",
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
			GoFieldName:        "ShortName",
			GoFieldType:        "string",
			JSONFieldName:      "short_name",
			ProtobufFieldName:  "short_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "tag_description",
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
			GoFieldName:        "TagDescription",
			GoFieldType:        "string",
			JSONFieldName:      "tag_description",
			ProtobufFieldName:  "tag_description",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "login_server_list_type_id",
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
			GoFieldName:        "LoginServerListTypeID",
			GoFieldType:        "int32",
			JSONFieldName:      "login_server_list_type_id",
			ProtobufFieldName:  "login_server_list_type_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "last_login_date",
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
			GoFieldName:        "LastLoginDate",
			GoFieldType:        "null.Time",
			JSONFieldName:      "last_login_date",
			ProtobufFieldName:  "last_login_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "last_ip_address",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(15)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       15,
			GoFieldName:        "LastIPAddress",
			GoFieldType:        "null.String",
			JSONFieldName:      "last_ip_address",
			ProtobufFieldName:  "last_ip_address",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "login_server_admin_id",
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
			GoFieldName:        "LoginServerAdminID",
			GoFieldType:        "int32",
			JSONFieldName:      "login_server_admin_id",
			ProtobufFieldName:  "login_server_admin_id",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "is_server_trusted",
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
			GoFieldName:        "IsServerTrusted",
			GoFieldType:        "int32",
			JSONFieldName:      "is_server_trusted",
			ProtobufFieldName:  "is_server_trusted",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "note",
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
			GoFieldName:        "Note",
			GoFieldType:        "null.String",
			JSONFieldName:      "note",
			ProtobufFieldName:  "note",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LoginWorldServers) TableName() string {
	return "login_world_servers"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LoginWorldServers) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LoginWorldServers) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LoginWorldServers) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LoginWorldServers) TableInfo() *TableInfo {
	return login_world_serversTableInfo
}
