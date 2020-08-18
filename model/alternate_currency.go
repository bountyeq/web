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


CREATE TABLE `alternate_currency` (
  `id` int(10) NOT NULL,
  `item_id` int(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "item_id": 30,    "id": 22}



*/

// AlternateCurrency struct is a row record of the alternate_currency table in the eqemu database
type AlternateCurrency struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;column:id;type:int;" json:"id"`
	//[ 1] item_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ItemID int32 `gorm:"column:item_id;type:int;" json:"item_id"`
}

var alternate_currencyTableInfo = &TableInfo{
	Name: "alternate_currency",
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
			IsAutoIncrement:    false,
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
			Name:               "item_id",
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
			GoFieldName:        "ItemID",
			GoFieldType:        "int32",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AlternateCurrency) TableName() string {
	return "alternate_currency"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AlternateCurrency) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AlternateCurrency) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AlternateCurrency) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AlternateCurrency) TableInfo() *TableInfo {
	return alternate_currencyTableInfo
}
