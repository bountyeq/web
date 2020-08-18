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


CREATE TABLE `qs_player_move_record_entries` (
  `event_id` int(11) DEFAULT 0,
  `from_slot` mediumint(7) DEFAULT 0,
  `to_slot` mediumint(7) DEFAULT 0,
  `item_id` int(11) DEFAULT 0,
  `charges` mediumint(7) DEFAULT 0,
  `aug_1` int(11) DEFAULT 0,
  `aug_2` int(11) DEFAULT 0,
  `aug_3` int(11) DEFAULT 0,
  `aug_4` int(11) DEFAULT 0,
  `aug_5` int(11) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "from_slot": 24,    "to_slot": 92,    "charges": 18,    "aug_2": 6,    "aug_3": 0,    "aug_4": 81,    "aug_5": 61,    "event_id": 76,    "aug_1": 96,    "item_id": 13}


Comments
-------------------------------------
[ 0] Warning table: qs_player_move_record_entries does not have a primary key defined, setting col position 1 event_id as primary key
Warning table: qs_player_move_record_entries primary key column event_id is nullable column, setting it as NOT NULL




*/

// QsPlayerMoveRecordEntries struct is a row record of the qs_player_move_record_entries table in the eqemu database
type QsPlayerMoveRecordEntries struct {
	//[ 0] event_id                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	EventID int32 `gorm:"primary_key;column:event_id;type:int;default:0;" json:"event_id"`
	//[ 1] from_slot                                      mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	FromSlot null.Int `gorm:"column:from_slot;type:mediumint;default:0;" json:"from_slot"`
	//[ 2] to_slot                                        mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	ToSlot null.Int `gorm:"column:to_slot;type:mediumint;default:0;" json:"to_slot"`
	//[ 3] item_id                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ItemID null.Int `gorm:"column:item_id;type:int;default:0;" json:"item_id"`
	//[ 4] charges                                        mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	Charges null.Int `gorm:"column:charges;type:mediumint;default:0;" json:"charges"`
	//[ 5] aug_1                                          int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Aug1 null.Int `gorm:"column:aug_1;type:int;default:0;" json:"aug_1"`
	//[ 6] aug_2                                          int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Aug2 null.Int `gorm:"column:aug_2;type:int;default:0;" json:"aug_2"`
	//[ 7] aug_3                                          int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Aug3 null.Int `gorm:"column:aug_3;type:int;default:0;" json:"aug_3"`
	//[ 8] aug_4                                          int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Aug4 null.Int `gorm:"column:aug_4;type:int;default:0;" json:"aug_4"`
	//[ 9] aug_5                                          int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Aug5 null.Int `gorm:"column:aug_5;type:int;default:0;" json:"aug_5"`
}

var qs_player_move_record_entriesTableInfo = &TableInfo{
	Name: "qs_player_move_record_entries",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "event_id",
			Comment: ``,
			Notes: `Warning table: qs_player_move_record_entries does not have a primary key defined, setting col position 1 event_id as primary key
Warning table: qs_player_move_record_entries primary key column event_id is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "EventID",
			GoFieldType:        "int32",
			JSONFieldName:      "event_id",
			ProtobufFieldName:  "event_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "item_id",
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
			GoFieldName:        "ItemID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "charges",
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
			GoFieldName:        "Charges",
			GoFieldType:        "null.Int",
			JSONFieldName:      "charges",
			ProtobufFieldName:  "charges",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "aug_1",
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
			GoFieldName:        "Aug1",
			GoFieldType:        "null.Int",
			JSONFieldName:      "aug_1",
			ProtobufFieldName:  "aug_1",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "aug_2",
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
			GoFieldName:        "Aug2",
			GoFieldType:        "null.Int",
			JSONFieldName:      "aug_2",
			ProtobufFieldName:  "aug_2",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "aug_3",
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
			GoFieldName:        "Aug3",
			GoFieldType:        "null.Int",
			JSONFieldName:      "aug_3",
			ProtobufFieldName:  "aug_3",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "aug_4",
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
			GoFieldName:        "Aug4",
			GoFieldType:        "null.Int",
			JSONFieldName:      "aug_4",
			ProtobufFieldName:  "aug_4",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "aug_5",
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
			GoFieldName:        "Aug5",
			GoFieldType:        "null.Int",
			JSONFieldName:      "aug_5",
			ProtobufFieldName:  "aug_5",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (q *QsPlayerMoveRecordEntries) TableName() string {
	return "qs_player_move_record_entries"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *QsPlayerMoveRecordEntries) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *QsPlayerMoveRecordEntries) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *QsPlayerMoveRecordEntries) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *QsPlayerMoveRecordEntries) TableInfo() *TableInfo {
	return qs_player_move_record_entriesTableInfo
}
