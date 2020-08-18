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


CREATE TABLE `keyring` (
  `char_id` int(11) NOT NULL DEFAULT 0,
  `item_id` int(11) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "char_id": 70,    "item_id": 15}


Comments
-------------------------------------
[ 0] Warning table: keyring does not have a primary key defined, setting col position 1 char_id as primary key




*/

// Keyring struct is a row record of the keyring table in the eqemu database
type Keyring struct {
	//[ 0] char_id                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	CharID int32 `gorm:"primary_key;column:char_id;type:int;default:0;" json:"char_id"`
	//[ 1] item_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ItemID int32 `gorm:"column:item_id;type:int;default:0;" json:"item_id"`
}

var keyringTableInfo = &TableInfo{
	Name: "keyring",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "char_id",
			Comment: ``,
			Notes: `Warning table: keyring does not have a primary key defined, setting col position 1 char_id as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CharID",
			GoFieldType:        "int32",
			JSONFieldName:      "char_id",
			ProtobufFieldName:  "char_id",
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
func (k *Keyring) TableName() string {
	return "keyring"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (k *Keyring) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (k *Keyring) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (k *Keyring) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (k *Keyring) TableInfo() *TableInfo {
	return keyringTableInfo
}
