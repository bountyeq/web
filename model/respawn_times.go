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


CREATE TABLE `respawn_times` (
  `id` int(11) NOT NULL DEFAULT 0,
  `start` int(11) NOT NULL DEFAULT 0,
  `duration` int(11) NOT NULL DEFAULT 0,
  `instance_id` smallint(6) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`instance_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 80,    "start": 28,    "duration": 42,    "instance_id": 12}



*/

// RespawnTimes struct is a row record of the respawn_times table in the eqemu database
type RespawnTimes struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	ID int32 `gorm:"primary_key;column:id;type:int;default:0;" json:"id"`
	//[ 1] start                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Start int32 `gorm:"column:start;type:int;default:0;" json:"start"`
	//[ 2] duration                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Duration int32 `gorm:"column:duration;type:int;default:0;" json:"duration"`
	//[ 3] instance_id                                    smallint             null: false  primary: true   isArray: false  auto: false  col: smallint        len: -1      default: [0]
	InstanceID int32 `gorm:"primary_key;column:instance_id;type:smallint;default:0;" json:"instance_id"`
}

var respawn_timesTableInfo = &TableInfo{
	Name: "respawn_times",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
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
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "start",
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
			GoFieldName:        "Start",
			GoFieldType:        "int32",
			JSONFieldName:      "start",
			ProtobufFieldName:  "start",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "duration",
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
			GoFieldName:        "Duration",
			GoFieldType:        "int32",
			JSONFieldName:      "duration",
			ProtobufFieldName:  "duration",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "instance_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "InstanceID",
			GoFieldType:        "int32",
			JSONFieldName:      "instance_id",
			ProtobufFieldName:  "instance_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *RespawnTimes) TableName() string {
	return "respawn_times"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RespawnTimes) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RespawnTimes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RespawnTimes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RespawnTimes) TableInfo() *TableInfo {
	return respawn_timesTableInfo
}
