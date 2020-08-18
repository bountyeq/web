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


CREATE TABLE `qs_player_handin_record` (
  `handin_id` int(11) NOT NULL AUTO_INCREMENT,
  `time` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  `quest_id` int(11) DEFAULT 0,
  `char_id` int(11) DEFAULT 0,
  `char_pp` int(11) DEFAULT 0,
  `char_gp` int(11) DEFAULT 0,
  `char_sp` int(11) DEFAULT 0,
  `char_cp` int(11) DEFAULT 0,
  `char_items` mediumint(7) DEFAULT 0,
  `npc_id` int(11) DEFAULT 0,
  `npc_pp` int(11) DEFAULT 0,
  `npc_gp` int(11) DEFAULT 0,
  `npc_sp` int(11) DEFAULT 0,
  `npc_cp` int(11) DEFAULT 0,
  `npc_items` mediumint(7) DEFAULT 0,
  PRIMARY KEY (`handin_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "handin_id": 54,    "quest_id": 45,    "char_id": 34,    "npc_pp": 69,    "npc_sp": 48,    "char_pp": 71,    "char_sp": 34,    "char_items": 26,    "npc_cp": 66,    "npc_items": 73,    "time": "2103-09-02T00:56:55.162303686-07:00",    "char_gp": 1,    "char_cp": 19,    "npc_id": 50,    "npc_gp": 8}



*/

// QsPlayerHandinRecord struct is a row record of the qs_player_handin_record table in the eqemu database
type QsPlayerHandinRecord struct {
	//[ 0] handin_id                                      int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	HandinID int32 `gorm:"primary_key;AUTO_INCREMENT;column:handin_id;type:int;" json:"handin_id"`
	//[ 1] time                                           timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [NULL]
	Time null.Time `gorm:"column:time;type:timestamp;" json:"time"`
	//[ 2] quest_id                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	QuestID null.Int `gorm:"column:quest_id;type:int;default:0;" json:"quest_id"`
	//[ 3] char_id                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CharID null.Int `gorm:"column:char_id;type:int;default:0;" json:"char_id"`
	//[ 4] char_pp                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CharPp null.Int `gorm:"column:char_pp;type:int;default:0;" json:"char_pp"`
	//[ 5] char_gp                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CharGp null.Int `gorm:"column:char_gp;type:int;default:0;" json:"char_gp"`
	//[ 6] char_sp                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CharSp null.Int `gorm:"column:char_sp;type:int;default:0;" json:"char_sp"`
	//[ 7] char_cp                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CharCp null.Int `gorm:"column:char_cp;type:int;default:0;" json:"char_cp"`
	//[ 8] char_items                                     mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	CharItems null.Int `gorm:"column:char_items;type:mediumint;default:0;" json:"char_items"`
	//[ 9] npc_id                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	NpcID null.Int `gorm:"column:npc_id;type:int;default:0;" json:"npc_id"`
	//[10] npc_pp                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	NpcPp null.Int `gorm:"column:npc_pp;type:int;default:0;" json:"npc_pp"`
	//[11] npc_gp                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	NpcGp null.Int `gorm:"column:npc_gp;type:int;default:0;" json:"npc_gp"`
	//[12] npc_sp                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	NpcSp null.Int `gorm:"column:npc_sp;type:int;default:0;" json:"npc_sp"`
	//[13] npc_cp                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	NpcCp null.Int `gorm:"column:npc_cp;type:int;default:0;" json:"npc_cp"`
	//[14] npc_items                                      mediumint            null: true   primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	NpcItems null.Int `gorm:"column:npc_items;type:mediumint;default:0;" json:"npc_items"`
}

var qs_player_handin_recordTableInfo = &TableInfo{
	Name: "qs_player_handin_record",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "handin_id",
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
			GoFieldName:        "HandinID",
			GoFieldType:        "int32",
			JSONFieldName:      "handin_id",
			ProtobufFieldName:  "handin_id",
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
			Name:               "quest_id",
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
			GoFieldName:        "QuestID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "quest_id",
			ProtobufFieldName:  "quest_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "npc_pp",
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
			GoFieldName:        "NpcPp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "npc_pp",
			ProtobufFieldName:  "npc_pp",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "npc_gp",
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
			GoFieldName:        "NpcGp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "npc_gp",
			ProtobufFieldName:  "npc_gp",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "npc_sp",
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
			GoFieldName:        "NpcSp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "npc_sp",
			ProtobufFieldName:  "npc_sp",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "npc_cp",
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
			GoFieldName:        "NpcCp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "npc_cp",
			ProtobufFieldName:  "npc_cp",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "npc_items",
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
			GoFieldName:        "NpcItems",
			GoFieldType:        "null.Int",
			JSONFieldName:      "npc_items",
			ProtobufFieldName:  "npc_items",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},
	},
}

// TableName sets the insert table name for this struct type
func (q *QsPlayerHandinRecord) TableName() string {
	return "qs_player_handin_record"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *QsPlayerHandinRecord) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *QsPlayerHandinRecord) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *QsPlayerHandinRecord) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *QsPlayerHandinRecord) TableInfo() *TableInfo {
	return qs_player_handin_recordTableInfo
}
