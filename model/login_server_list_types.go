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


CREATE TABLE `login_server_list_types` (
  `id` int(10) unsigned NOT NULL,
  `description` varchar(60) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 29,    "description": "baEkpGQmXOhPjtybpJHBtMYAa"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// LoginServerListTypes struct is a row record of the login_server_list_types table in the eqemu database
type LoginServerListTypes struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;column:id;type:uint;" json:"id"`
	//[ 1] description                                    varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	Description string `gorm:"column:description;type:varchar;size:60;" json:"description"`
}

var login_server_list_typesTableInfo = &TableInfo{
	Name: "login_server_list_types",
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
			Name:               "description",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "Description",
			GoFieldType:        "string",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LoginServerListTypes) TableName() string {
	return "login_server_list_types"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LoginServerListTypes) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LoginServerListTypes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LoginServerListTypes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LoginServerListTypes) TableInfo() *TableInfo {
	return login_server_list_typesTableInfo
}
