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


CREATE TABLE `qs_player_delete_record` (
  `delete_id` int(11) NOT NULL AUTO_INCREMENT,
  `time` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  `char_id` int(11) DEFAULT 0,
  `stack_size` mediumint(7) DEFAULT 0,
  `char_items` mediumint(7) DEFAULT 0,
  PRIMARY KEY (`delete_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "char_id": 94,    "stack_size": 30,    "char_items": 70,    "delete_id": 96,    "time": "2024-10-02T17:22:35.271947163-07:00"}



*/

// QsPlayerDeleteRecord struct is a row record of the qs_player_delete_record table in the eqemu database
type QsPlayerDeleteRecord struct {
	//[ 0] delete_id                                      int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	DeleteID int32 `gorm:"primary_key;AUTO_INCREMENT;column:delete_id;type:int;" json:"delete_id"`
	//[ 1] time                                           timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [NULL]
	Time null.Time `gorm:"column:time;type:timestamp;" json:"time"`
	//[ 2] char_id                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CharID null.Int `gorm:"column:char_id;type:int;default:0;" json:"char_id"`
	//[ 3] stack_size                                     mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	StackSize null.Int `gorm:"column:stack_size;type:mediumint;default:0;" json:"stack_size"`
	//[ 4] char_items                                     mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	CharItems null.Int `gorm:"column:char_items;type:mediumint;default:0;" json:"char_items"`
}

var qs_player_delete_recordTableInfo = &TableInfo{
	Name: "qs_player_delete_record",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "delete_id",
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
			GoFieldName:        "DeleteID",
			GoFieldType:        "int32",
			JSONFieldName:      "delete_id",
			ProtobufFieldName:  "delete_id",
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "stack_size",
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
			GoFieldName:        "StackSize",
			GoFieldType:        "null.Int",
			JSONFieldName:      "stack_size",
			ProtobufFieldName:  "stack_size",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (q *QsPlayerDeleteRecord) TableName() string {
	return "qs_player_delete_record"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *QsPlayerDeleteRecord) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *QsPlayerDeleteRecord) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *QsPlayerDeleteRecord) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *QsPlayerDeleteRecord) TableInfo() *TableInfo {
	return qs_player_delete_recordTableInfo
}
