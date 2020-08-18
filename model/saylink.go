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


CREATE TABLE `saylink` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `phrase` varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `phrase_index` (`phrase`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 35,    "phrase": "gLjtaorFAvLxbZRBNfWSgIhIC"}



*/

// Saylink struct is a row record of the saylink table in the eqemu database
type Saylink struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] phrase                                         varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Phrase string `gorm:"column:phrase;type:varchar;size:64;default:'';" json:"phrase"`
}

var saylinkTableInfo = &TableInfo{
	Name: "saylink",
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
			Name:               "phrase",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Phrase",
			GoFieldType:        "string",
			JSONFieldName:      "phrase",
			ProtobufFieldName:  "phrase",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *Saylink) TableName() string {
	return "saylink"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *Saylink) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *Saylink) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *Saylink) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *Saylink) TableInfo() *TableInfo {
	return saylinkTableInfo
}
