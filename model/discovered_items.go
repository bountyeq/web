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


CREATE TABLE `discovered_items` (
  `item_id` int(11) unsigned NOT NULL DEFAULT 0,
  `char_name` varchar(64) NOT NULL DEFAULT '',
  `discovered_date` int(11) unsigned NOT NULL DEFAULT 0,
  `account_status` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "item_id": 51,    "char_name": "PJFabZVQWGdDBiqZlExBIYlQo",    "discovered_date": 58,    "account_status": 63}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned



*/

// DiscoveredItems struct is a row record of the discovered_items table in the eqemu database
type DiscoveredItems struct {
	//[ 0] item_id                                        uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	ItemID uint32 `gorm:"primary_key;column:item_id;type:uint;default:0;" json:"item_id"`
	//[ 1] char_name                                      varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	CharName string `gorm:"column:char_name;type:varchar;size:64;default:'';" json:"char_name"`
	//[ 2] discovered_date                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	DiscoveredDate uint32 `gorm:"column:discovered_date;type:uint;default:0;" json:"discovered_date"`
	//[ 3] account_status                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AccountStatus int32 `gorm:"column:account_status;type:int;default:0;" json:"account_status"`
}

var discovered_itemsTableInfo = &TableInfo{
	Name: "discovered_items",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "item_id",
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
			GoFieldName:        "ItemID",
			GoFieldType:        "uint32",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "char_name",
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
			GoFieldName:        "CharName",
			GoFieldType:        "string",
			JSONFieldName:      "char_name",
			ProtobufFieldName:  "char_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "discovered_date",
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
			GoFieldName:        "DiscoveredDate",
			GoFieldType:        "uint32",
			JSONFieldName:      "discovered_date",
			ProtobufFieldName:  "discovered_date",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "account_status",
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
			GoFieldName:        "AccountStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "account_status",
			ProtobufFieldName:  "account_status",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *DiscoveredItems) TableName() string {
	return "discovered_items"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DiscoveredItems) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DiscoveredItems) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DiscoveredItems) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *DiscoveredItems) TableInfo() *TableInfo {
	return discovered_itemsTableInfo
}
