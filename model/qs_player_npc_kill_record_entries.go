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


CREATE TABLE `qs_player_npc_kill_record_entries` (
  `event_id` int(11) DEFAULT 0,
  `char_id` int(11) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "event_id": 68,    "char_id": 21}


Comments
-------------------------------------
[ 0] Warning table: qs_player_npc_kill_record_entries does not have a primary key defined, setting col position 1 event_id as primary key
Warning table: qs_player_npc_kill_record_entries primary key column event_id is nullable column, setting it as NOT NULL




*/

// QsPlayerNpcKillRecordEntries struct is a row record of the qs_player_npc_kill_record_entries table in the eqemu database
type QsPlayerNpcKillRecordEntries struct {
	//[ 0] event_id                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	EventID int32 `gorm:"primary_key;column:event_id;type:int;default:0;" json:"event_id"`
	//[ 1] char_id                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CharID null.Int `gorm:"column:char_id;type:int;default:0;" json:"char_id"`
}

var qs_player_npc_kill_record_entriesTableInfo = &TableInfo{
	Name: "qs_player_npc_kill_record_entries",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "event_id",
			Comment: ``,
			Notes: `Warning table: qs_player_npc_kill_record_entries does not have a primary key defined, setting col position 1 event_id as primary key
Warning table: qs_player_npc_kill_record_entries primary key column event_id is nullable column, setting it as NOT NULL
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
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (q *QsPlayerNpcKillRecordEntries) TableName() string {
	return "qs_player_npc_kill_record_entries"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *QsPlayerNpcKillRecordEntries) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *QsPlayerNpcKillRecordEntries) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *QsPlayerNpcKillRecordEntries) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *QsPlayerNpcKillRecordEntries) TableInfo() *TableInfo {
	return qs_player_npc_kill_record_entriesTableInfo
}
