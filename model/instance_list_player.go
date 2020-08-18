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


CREATE TABLE `instance_list_player` (
  `id` int(11) unsigned NOT NULL DEFAULT 0,
  `charid` int(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`charid`,`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "id": 44,    "charid": 43}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// InstanceListPlayer struct is a row record of the instance_list_player table in the eqemu database
type InstanceListPlayer struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ID uint32 `gorm:"primary_key;column:id;type:uint;default:0;" json:"id"`
	//[ 1] charid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Charid uint32 `gorm:"primary_key;column:charid;type:uint;default:0;" json:"charid"`
}

var instance_list_playerTableInfo = &TableInfo{
	Name: "instance_list_player",
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
			Name:               "charid",
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
			GoFieldName:        "Charid",
			GoFieldType:        "uint32",
			JSONFieldName:      "charid",
			ProtobufFieldName:  "charid",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (i *InstanceListPlayer) TableName() string {
	return "instance_list_player"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *InstanceListPlayer) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *InstanceListPlayer) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *InstanceListPlayer) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *InstanceListPlayer) TableInfo() *TableInfo {
	return instance_list_playerTableInfo
}
