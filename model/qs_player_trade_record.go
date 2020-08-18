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


CREATE TABLE `qs_player_trade_record` (
  `trade_id` int(11) NOT NULL AUTO_INCREMENT,
  `time` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  `char1_id` int(11) DEFAULT 0,
  `char1_pp` int(11) DEFAULT 0,
  `char1_gp` int(11) DEFAULT 0,
  `char1_sp` int(11) DEFAULT 0,
  `char1_cp` int(11) DEFAULT 0,
  `char1_items` mediumint(7) DEFAULT 0,
  `char2_id` int(11) DEFAULT 0,
  `char2_pp` int(11) DEFAULT 0,
  `char2_gp` int(11) DEFAULT 0,
  `char2_sp` int(11) DEFAULT 0,
  `char2_cp` int(11) DEFAULT 0,
  `char2_items` mediumint(7) DEFAULT 0,
  PRIMARY KEY (`trade_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "char_2_pp": 90,    "char_2_gp": 92,    "trade_id": 59,    "char_1_gp": 84,    "char_1_cp": 90,    "char_1_pp": 15,    "char_2_cp": 92,    "char_2_sp": 98,    "char_1_id": 9,    "char_1_items": 90,    "char_2_id": 86,    "time": "2027-07-06T03:51:24.376605792-07:00",    "char_1_sp": 12,    "char_2_items": 13}



*/

// QsPlayerTradeRecord struct is a row record of the qs_player_trade_record table in the eqemu database
type QsPlayerTradeRecord struct {
	//[ 0] trade_id                                       int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	TradeID int32 `gorm:"primary_key;AUTO_INCREMENT;column:trade_id;type:int;" json:"trade_id"`
	//[ 1] time                                           timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [NULL]
	Time null.Time `gorm:"column:time;type:timestamp;" json:"time"`
	//[ 2] char1_id                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Char1ID null.Int `gorm:"column:char1_id;type:int;default:0;" json:"char_1_id"`
	//[ 3] char1_pp                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Char1Pp null.Int `gorm:"column:char1_pp;type:int;default:0;" json:"char_1_pp"`
	//[ 4] char1_gp                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Char1Gp null.Int `gorm:"column:char1_gp;type:int;default:0;" json:"char_1_gp"`
	//[ 5] char1_sp                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Char1Sp null.Int `gorm:"column:char1_sp;type:int;default:0;" json:"char_1_sp"`
	//[ 6] char1_cp                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Char1Cp null.Int `gorm:"column:char1_cp;type:int;default:0;" json:"char_1_cp"`
	//[ 7] char1_items                                    mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	Char1Items null.Int `gorm:"column:char1_items;type:mediumint;default:0;" json:"char_1_items"`
	//[ 8] char2_id                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Char2ID null.Int `gorm:"column:char2_id;type:int;default:0;" json:"char_2_id"`
	//[ 9] char2_pp                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Char2Pp null.Int `gorm:"column:char2_pp;type:int;default:0;" json:"char_2_pp"`
	//[10] char2_gp                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Char2Gp null.Int `gorm:"column:char2_gp;type:int;default:0;" json:"char_2_gp"`
	//[11] char2_sp                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Char2Sp null.Int `gorm:"column:char2_sp;type:int;default:0;" json:"char_2_sp"`
	//[12] char2_cp                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Char2Cp null.Int `gorm:"column:char2_cp;type:int;default:0;" json:"char_2_cp"`
	//[13] char2_items                                    mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	Char2Items null.Int `gorm:"column:char2_items;type:mediumint;default:0;" json:"char_2_items"`
}

var qs_player_trade_recordTableInfo = &TableInfo{
	Name: "qs_player_trade_record",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "trade_id",
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
			GoFieldName:        "TradeID",
			GoFieldType:        "int32",
			JSONFieldName:      "trade_id",
			ProtobufFieldName:  "trade_id",
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
			Name:               "char1_id",
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
			GoFieldName:        "Char1ID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_1_id",
			ProtobufFieldName:  "char_1_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "char1_pp",
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
			GoFieldName:        "Char1Pp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_1_pp",
			ProtobufFieldName:  "char_1_pp",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "char1_gp",
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
			GoFieldName:        "Char1Gp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_1_gp",
			ProtobufFieldName:  "char_1_gp",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "char1_sp",
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
			GoFieldName:        "Char1Sp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_1_sp",
			ProtobufFieldName:  "char_1_sp",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "char1_cp",
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
			GoFieldName:        "Char1Cp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_1_cp",
			ProtobufFieldName:  "char_1_cp",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "char1_items",
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
			GoFieldName:        "Char1Items",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_1_items",
			ProtobufFieldName:  "char_1_items",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "char2_id",
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
			GoFieldName:        "Char2ID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_2_id",
			ProtobufFieldName:  "char_2_id",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "char2_pp",
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
			GoFieldName:        "Char2Pp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_2_pp",
			ProtobufFieldName:  "char_2_pp",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "char2_gp",
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
			GoFieldName:        "Char2Gp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_2_gp",
			ProtobufFieldName:  "char_2_gp",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "char2_sp",
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
			GoFieldName:        "Char2Sp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_2_sp",
			ProtobufFieldName:  "char_2_sp",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "char2_cp",
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
			GoFieldName:        "Char2Cp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_2_cp",
			ProtobufFieldName:  "char_2_cp",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "char2_items",
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
			GoFieldName:        "Char2Items",
			GoFieldType:        "null.Int",
			JSONFieldName:      "char_2_items",
			ProtobufFieldName:  "char_2_items",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (q *QsPlayerTradeRecord) TableName() string {
	return "qs_player_trade_record"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *QsPlayerTradeRecord) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *QsPlayerTradeRecord) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *QsPlayerTradeRecord) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *QsPlayerTradeRecord) TableInfo() *TableInfo {
	return qs_player_trade_recordTableInfo
}
