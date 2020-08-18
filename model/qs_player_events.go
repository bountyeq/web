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


CREATE TABLE `qs_player_events` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `char_id` int(11) DEFAULT 0,
  `event` int(11) unsigned DEFAULT 0,
  `event_desc` varchar(255) DEFAULT NULL,
  `time` int(11) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 88,    "char_id": 15,    "event": 84,    "event_desc": "kLAJPsIuFtpsHhaileojbJrar",    "time": 16}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 4] column is set for unsigned



*/

// QsPlayerEvents struct is a row record of the qs_player_events table in the eqemu database
type QsPlayerEvents struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] char_id                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CharID null.Int `gorm:"column:char_id;type:int;default:0;" json:"char_id"`
	//[ 2] event                                          uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Event null.Int `gorm:"column:event;type:uint;default:0;" json:"event"`
	//[ 3] event_desc                                     varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: [NULL]
	EventDesc null.String `gorm:"column:event_desc;type:varchar;size:255;" json:"event_desc"`
	//[ 4] time                                           uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Time null.Int `gorm:"column:time;type:uint;default:0;" json:"time"`
}

var qs_player_eventsTableInfo = &TableInfo{
	Name: "qs_player_events",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
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

		&ColumnInfo{
			Index:              2,
			Name:               "event",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Event",
			GoFieldType:        "null.Int",
			JSONFieldName:      "event",
			ProtobufFieldName:  "event",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "event_desc",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "EventDesc",
			GoFieldType:        "null.String",
			JSONFieldName:      "event_desc",
			ProtobufFieldName:  "event_desc",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "time",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Time",
			GoFieldType:        "null.Int",
			JSONFieldName:      "time",
			ProtobufFieldName:  "time",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (q *QsPlayerEvents) TableName() string {
	return "qs_player_events"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *QsPlayerEvents) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *QsPlayerEvents) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *QsPlayerEvents) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *QsPlayerEvents) TableInfo() *TableInfo {
	return qs_player_eventsTableInfo
}
