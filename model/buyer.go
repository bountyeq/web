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


CREATE TABLE `buyer` (
  `charid` int(11) NOT NULL DEFAULT 0,
  `buyslot` int(11) NOT NULL DEFAULT 0,
  `itemid` int(11) NOT NULL DEFAULT 0,
  `itemname` varchar(65) NOT NULL DEFAULT '',
  `quantity` int(11) NOT NULL DEFAULT 0,
  `price` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`charid`,`buyslot`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "buyslot": 91,    "itemid": 55,    "itemname": "fwCwNfPbyiasEDcKHGuuSKwFJ",    "quantity": 82,    "price": 29,    "charid": 85}



*/

// Buyer struct is a row record of the buyer table in the eqemu database
type Buyer struct {
	//[ 0] charid                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Charid int32 `gorm:"primary_key;column:charid;type:int;default:0;" json:"charid"`
	//[ 1] buyslot                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Buyslot int32 `gorm:"primary_key;column:buyslot;type:int;default:0;" json:"buyslot"`
	//[ 2] itemid                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Itemid int32 `gorm:"column:itemid;type:int;default:0;" json:"itemid"`
	//[ 3] itemname                                       varchar(65)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 65      default: ['']
	Itemname string `gorm:"column:itemname;type:varchar;size:65;default:'';" json:"itemname"`
	//[ 4] quantity                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Quantity int32 `gorm:"column:quantity;type:int;default:0;" json:"quantity"`
	//[ 5] price                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Price int32 `gorm:"column:price;type:int;default:0;" json:"price"`
}

var buyerTableInfo = &TableInfo{
	Name: "buyer",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "charid",
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
			GoFieldName:        "Charid",
			GoFieldType:        "int32",
			JSONFieldName:      "charid",
			ProtobufFieldName:  "charid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "buyslot",
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
			GoFieldName:        "Buyslot",
			GoFieldType:        "int32",
			JSONFieldName:      "buyslot",
			ProtobufFieldName:  "buyslot",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "itemid",
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
			GoFieldName:        "Itemid",
			GoFieldType:        "int32",
			JSONFieldName:      "itemid",
			ProtobufFieldName:  "itemid",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "itemname",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(65)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       65,
			GoFieldName:        "Itemname",
			GoFieldType:        "string",
			JSONFieldName:      "itemname",
			ProtobufFieldName:  "itemname",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "quantity",
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
			GoFieldName:        "Quantity",
			GoFieldType:        "int32",
			JSONFieldName:      "quantity",
			ProtobufFieldName:  "quantity",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "price",
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
			GoFieldName:        "Price",
			GoFieldType:        "int32",
			JSONFieldName:      "price",
			ProtobufFieldName:  "price",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *Buyer) TableName() string {
	return "buyer"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *Buyer) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *Buyer) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *Buyer) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *Buyer) TableInfo() *TableInfo {
	return buyerTableInfo
}
