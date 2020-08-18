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


CREATE TABLE `qs_player_move_record` (
  `move_id` int(11) NOT NULL AUTO_INCREMENT,
  `time` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  `char_id` int(11) DEFAULT 0,
  `from_slot` mediumint(7) DEFAULT 0,
  `to_slot` mediumint(7) DEFAULT 0,
  `stack_size` mediumint(7) DEFAULT 0,
  `char_items` mediumint(7) DEFAULT 0,
  `postaction` tinyint(1) DEFAULT 0,
  PRIMARY KEY (`move_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "time": "2223-09-16T23:40:23.475037382-07:00",    "char_id": 73,    "from_slot": 19,    "to_slot": 1,    "stack_size": 83,    "char_items": 96,    "postaction": 18,    "move_id": 7}



*/

// QsPlayerMoveRecord struct is a row record of the qs_player_move_record table in the eqemu database
type QsPlayerMoveRecord struct {
	//[ 0] move_id                                        int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	MoveID int32 `gorm:"primary_key;AUTO_INCREMENT;column:move_id;type:int;" json:"move_id"`
	//[ 1] time                                           timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [NULL]
	Time null.Time `gorm:"column:time;type:timestamp;" json:"time"`
	//[ 2] char_id                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CharID null.Int `gorm:"column:char_id;type:int;default:0;" json:"char_id"`
	//[ 3] from_slot                                      mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	FromSlot null.Int `gorm:"column:from_slot;type:mediumint;default:0;" json:"from_slot"`
	//[ 4] to_slot                                        mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	ToSlot null.Int `gorm:"column:to_slot;type:mediumint;default:0;" json:"to_slot"`
	//[ 5] stack_size                                     mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	StackSize null.Int `gorm:"column:stack_size;type:mediumint;default:0;" json:"stack_size"`
	//[ 6] char_items                                     mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	CharItems null.Int `gorm:"column:char_items;type:mediumint;default:0;" json:"char_items"`
	//[ 7] postaction                                     tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Postaction null.Int `gorm:"column:postaction;type:tinyint;default:0;" json:"postaction"`
}

var qs_player_move_recordTableInfo = &TableInfo{
	Name: "qs_player_move_record",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "move_id",
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
			GoFieldName:        "MoveID",
			GoFieldType:        "int32",
			JSONFieldName:      "move_id",
			ProtobufFieldName:  "move_id",
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
			Name:               "from_slot",
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
			GoFieldName:        "FromSlot",
			GoFieldType:        "null.Int",
			JSONFieldName:      "from_slot",
			ProtobufFieldName:  "from_slot",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "to_slot",
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
			GoFieldName:        "ToSlot",
			GoFieldType:        "null.Int",
			JSONFieldName:      "to_slot",
			ProtobufFieldName:  "to_slot",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "postaction",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Postaction",
			GoFieldType:        "null.Int",
			JSONFieldName:      "postaction",
			ProtobufFieldName:  "postaction",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (q *QsPlayerMoveRecord) TableName() string {
	return "qs_player_move_record"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *QsPlayerMoveRecord) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *QsPlayerMoveRecord) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *QsPlayerMoveRecord) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *QsPlayerMoveRecord) TableInfo() *TableInfo {
	return qs_player_move_recordTableInfo
}
