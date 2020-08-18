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


CREATE TABLE `character_alt_currency` (
  `char_id` int(10) unsigned NOT NULL,
  `currency_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  PRIMARY KEY (`char_id`,`currency_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "char_id": 56,    "currency_id": 13,    "amount": 77}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// CharacterAltCurrency struct is a row record of the character_alt_currency table in the eqemu database
type CharacterAltCurrency struct {
	//[ 0] char_id                                        uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	CharID uint32 `gorm:"primary_key;column:char_id;type:uint;" json:"char_id"`
	//[ 1] currency_id                                    uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	CurrencyID uint32 `gorm:"primary_key;column:currency_id;type:uint;" json:"currency_id"`
	//[ 2] amount                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	Amount uint32 `gorm:"column:amount;type:uint;" json:"amount"`
}

var character_alt_currencyTableInfo = &TableInfo{
	Name: "character_alt_currency",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "char_id",
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
			GoFieldName:        "CharID",
			GoFieldType:        "uint32",
			JSONFieldName:      "char_id",
			ProtobufFieldName:  "char_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "currency_id",
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
			GoFieldName:        "CurrencyID",
			GoFieldType:        "uint32",
			JSONFieldName:      "currency_id",
			ProtobufFieldName:  "currency_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "amount",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Amount",
			GoFieldType:        "uint32",
			JSONFieldName:      "amount",
			ProtobufFieldName:  "amount",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterAltCurrency) TableName() string {
	return "character_alt_currency"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterAltCurrency) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterAltCurrency) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterAltCurrency) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterAltCurrency) TableInfo() *TableInfo {
	return character_alt_currencyTableInfo
}
