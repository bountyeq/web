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


CREATE TABLE `qs_player_npc_kill_record` (
  `fight_id` int(11) NOT NULL AUTO_INCREMENT,
  `npc_id` int(11) DEFAULT NULL,
  `type` int(11) DEFAULT NULL,
  `zone_id` int(11) DEFAULT NULL,
  `time` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  PRIMARY KEY (`fight_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "fight_id": 44,    "npc_id": 54,    "type": 99,    "zone_id": 25,    "time": "2084-07-24T12:19:28.781816693-07:00"}



*/

// QsPlayerNpcKillRecord struct is a row record of the qs_player_npc_kill_record table in the eqemu database
type QsPlayerNpcKillRecord struct {
	//[ 0] fight_id                                       int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	FightID int32 `gorm:"primary_key;AUTO_INCREMENT;column:fight_id;type:int;" json:"fight_id"`
	//[ 1] npc_id                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	NpcID null.Int `gorm:"column:npc_id;type:int;" json:"npc_id"`
	//[ 2] type                                           int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	Type null.Int `gorm:"column:type;type:int;" json:"type"`
	//[ 3] zone_id                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	ZoneID null.Int `gorm:"column:zone_id;type:int;" json:"zone_id"`
	//[ 4] time                                           timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [NULL]
	Time null.Time `gorm:"column:time;type:timestamp;" json:"time"`
}

var qs_player_npc_kill_recordTableInfo = &TableInfo{
	Name: "qs_player_npc_kill_record",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "fight_id",
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
			GoFieldName:        "FightID",
			GoFieldType:        "int32",
			JSONFieldName:      "fight_id",
			ProtobufFieldName:  "fight_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "npc_id",
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
			GoFieldName:        "NpcID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "npc_id",
			ProtobufFieldName:  "npc_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "null.Int",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (q *QsPlayerNpcKillRecord) TableName() string {
	return "qs_player_npc_kill_record"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *QsPlayerNpcKillRecord) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *QsPlayerNpcKillRecord) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *QsPlayerNpcKillRecord) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *QsPlayerNpcKillRecord) TableInfo() *TableInfo {
	return qs_player_npc_kill_recordTableInfo
}
