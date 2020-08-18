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


CREATE TABLE `login_api_tokens` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `token` varchar(200) DEFAULT NULL,
  `can_write` int(11) DEFAULT 0,
  `can_read` int(11) DEFAULT 0,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "created_at": "2163-12-23T16:21:20.586971076-08:00",    "updated_at": "2068-07-07T02:12:07.144394175-07:00",    "id": 6,    "token": "QAwDVFjNuYaACXWQlxkgSCKnH",    "can_write": 65,    "can_read": 21}



*/

// LoginAPITokens struct is a row record of the login_api_tokens table in the eqemu database
type LoginAPITokens struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] token                                          varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: [NULL]
	Token null.String `gorm:"column:token;type:varchar;size:200;" json:"token"`
	//[ 2] can_write                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CanWrite null.Int `gorm:"column:can_write;type:int;default:0;" json:"can_write"`
	//[ 3] can_read                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CanRead null.Int `gorm:"column:can_read;type:int;default:0;" json:"can_read"`
	//[ 4] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [NULL]
	CreatedAt null.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[ 5] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [current_timestamp()]
	UpdatedAt null.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var login_api_tokensTableInfo = &TableInfo{
	Name: "login_api_tokens",
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
			Name:               "token",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "Token",
			GoFieldType:        "null.String",
			JSONFieldName:      "token",
			ProtobufFieldName:  "token",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "can_write",
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
			GoFieldName:        "CanWrite",
			GoFieldType:        "null.Int",
			JSONFieldName:      "can_write",
			ProtobufFieldName:  "can_write",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "can_read",
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
			GoFieldName:        "CanRead",
			GoFieldType:        "null.Int",
			JSONFieldName:      "can_read",
			ProtobufFieldName:  "can_read",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LoginAPITokens) TableName() string {
	return "login_api_tokens"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LoginAPITokens) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LoginAPITokens) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LoginAPITokens) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LoginAPITokens) TableInfo() *TableInfo {
	return login_api_tokensTableInfo
}
