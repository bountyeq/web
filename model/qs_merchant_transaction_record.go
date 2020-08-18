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


CREATE TABLE `qs_merchant_transaction_record` (
  `transaction_id` int(11) NOT NULL AUTO_INCREMENT,
  `time` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  `zone_id` int(11) DEFAULT 0,
  `merchant_id` int(11) DEFAULT 0,
  `merchant_pp` int(11) DEFAULT 0,
  `merchant_gp` int(11) DEFAULT 0,
  `merchant_sp` int(11) DEFAULT 0,
  `merchant_cp` int(11) DEFAULT 0,
  `merchant_items` mediumint(7) DEFAULT 0,
  `char_id` int(11) DEFAULT 0,
  `char_pp` int(11) DEFAULT 0,
  `char_gp` int(11) DEFAULT 0,
  `char_sp` int(11) DEFAULT 0,
  `char_cp` int(11) DEFAULT 0,
  `char_items` mediumint(7) DEFAULT 0,
  PRIMARY KEY (`transaction_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "merchant_gp": 37,    "char_cp": 33,    "merchant_sp": 42,    "char_id": 20,    "char_pp": 47,    "char_sp": 93,    "zone_id": 39,    "merchant_id": 33,    "merchant_pp": 64,    "char_gp": 37,    "transaction_id": 39,    "time": "2168-10-04T12:50:15.20532635-07:00",    "merchant_cp": 52,    "merchant_items": 15,    "char_items": 52}



*/

// QsMerchantTransactionRecord struct is a row record of the qs_merchant_transaction_record table in the eqemu database
type QsMerchantTransactionRecord struct {
	//[ 0] transaction_id                                 int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	TransactionID int32 `gorm:"primary_key;AUTO_INCREMENT;column:transaction_id;type:int;" json:"transaction_id"`
	//[ 1] time                                           timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [NULL]
	Time null.Time `gorm:"column:time;type:timestamp;" json:"time"`
	//[ 2] zone_id                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ZoneID null.Int `gorm:"column:zone_id;type:int;default:0;" json:"zone_id"`
	//[ 3] merchant_id                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MerchantID null.Int `gorm:"column:merchant_id;type:int;default:0;" json:"merchant_id"`
	//[ 4] merchant_pp                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MerchantPp null.Int `gorm:"column:merchant_pp;type:int;default:0;" json:"merchant_pp"`
	//[ 5] merchant_gp                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MerchantGp null.Int `gorm:"column:merchant_gp;type:int;default:0;" json:"merchant_gp"`
	//[ 6] merchant_sp                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MerchantSp null.Int `gorm:"column:merchant_sp;type:int;default:0;" json:"merchant_sp"`
	//[ 7] merchant_cp                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MerchantCp null.Int `gorm:"column:merchant_cp;type:int;default:0;" json:"merchant_cp"`
	//[ 8] merchant_items                                 mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	MerchantItems null.Int `gorm:"column:merchant_items;type:mediumint;default:0;" json:"merchant_items"`
	//[ 9] char_id                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CharID null.Int `gorm:"column:char_id;type:int;default:0;" json:"char_id"`
	//[10] char_pp                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CharPp null.Int `gorm:"column:char_pp;type:int;default:0;" json:"char_pp"`
	//[11] char_gp                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CharGp null.Int `gorm:"column:char_gp;type:int;default:0;" json:"char_gp"`
	//[12] char_sp                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CharSp null.Int `gorm:"column:char_sp;type:int;default:0;" json:"char_sp"`
	//[13] char_cp                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CharCp null.Int `gorm:"column:char_cp;type:int;default:0;" json:"char_cp"`
	//[14] char_items                                     mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	CharItems null.Int `gorm:"column:char_items;type:mediumint;default:0;" json:"char_items"`
}

var qs_merchant_transaction_recordTableInfo = &TableInfo{
	Name: "qs_merchant_transaction_record",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "transaction_id",
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
			GoFieldName:        "TransactionID",
			GoFieldType:        "int32",
			JSONFieldName:      "transaction_id",
			ProtobufFieldName:  "transaction_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "Time",
			GoFieldType:        "null.Time",
			JSONFieldName:      "time",
			ProtobufFieldName:  "time",
			ProtobufType:       "uint64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "zone_id",
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
			GoFieldName:        "ZoneID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "zone_id",
			ProtobufFieldName:  "zone_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "merchant_id",
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
			GoFieldName:        "MerchantID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "merchant_pp",
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
			GoFieldName:        "MerchantPp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "merchant_pp",
			ProtobufFieldName:  "merchant_pp",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "merchant_gp",
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
			GoFieldName:        "MerchantGp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "merchant_gp",
			ProtobufFieldName:  "merchant_gp",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "merchant_sp",
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
			GoFieldName:        "MerchantSp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "merchant_sp",
			ProtobufFieldName:  "merchant_sp",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "merchant_cp",
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
			GoFieldName:        "MerchantCp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "merchant_cp",
			ProtobufFieldName:  "merchant_cp",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "merchant_items",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "mediumint",
			DatabaseTypePretty: "mediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "mediumint",
			ColumnLength:       -1,
			GoFieldName:        "MerchantItems",
			GoFieldType:        "null.Int",
			JSONFieldName:      "merchant_items",
			ProtobufFieldName:  "merchant_items",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "char_id",
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
			GoFieldName:        "CharID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_id",
			ProtobufFieldName:  "char_id",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "char_pp",
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
			GoFieldName:        "CharPp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_pp",
			ProtobufFieldName:  "char_pp",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "char_gp",
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
			GoFieldName:        "CharGp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_gp",
			ProtobufFieldName:  "char_gp",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "char_sp",
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
			GoFieldName:        "CharSp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_sp",
			ProtobufFieldName:  "char_sp",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "char_cp",
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
			GoFieldName:        "CharCp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_cp",
			ProtobufFieldName:  "char_cp",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "char_items",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "mediumint",
			DatabaseTypePretty: "mediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "mediumint",
			ColumnLength:       -1,
			GoFieldName:        "CharItems",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_items",
			ProtobufFieldName:  "char_items",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},
	},
}

// TableName sets the insert table name for this struct type
func (q *QsMerchantTransactionRecord) TableName() string {
	return "qs_merchant_transaction_record"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *QsMerchantTransactionRecord) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *QsMerchantTransactionRecord) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *QsMerchantTransactionRecord) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *QsMerchantTransactionRecord) TableInfo() *TableInfo {
	return qs_merchant_transaction_recordTableInfo
}
