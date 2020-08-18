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


CREATE TABLE `adventure_details` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `adventure_id` smallint(5) unsigned NOT NULL,
  `instance_id` int(11) NOT NULL DEFAULT -1,
  `count` smallint(5) unsigned NOT NULL DEFAULT 0,
  `assassinate_count` smallint(5) unsigned NOT NULL DEFAULT 0,
  `status` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `time_created` int(10) unsigned NOT NULL DEFAULT 0,
  `time_zoned` int(10) unsigned NOT NULL DEFAULT 0,
  `time_completed` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "adventure_id": 50,    "instance_id": 16,    "count": 11,    "assassinate_count": 89,    "status": 50,    "time_created": 67,    "id": 62,    "time_zoned": 80,    "time_completed": 44}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned



*/

// AdventureDetails struct is a row record of the adventure_details table in the eqemu database
type AdventureDetails struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] adventure_id                                   usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: []
	AdventureID uint32 `gorm:"column:adventure_id;type:usmallint;" json:"adventure_id"`
	//[ 2] instance_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	InstanceID int32 `gorm:"column:instance_id;type:int;default:-1;" json:"instance_id"`
	//[ 3] count                                          usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	Count uint32 `gorm:"column:count;type:usmallint;default:0;" json:"count"`
	//[ 4] assassinate_count                              usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	AssassinateCount uint32 `gorm:"column:assassinate_count;type:usmallint;default:0;" json:"assassinate_count"`
	//[ 5] status                                         utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Status uint32 `gorm:"column:status;type:utinyint;default:0;" json:"status"`
	//[ 6] time_created                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TimeCreated uint32 `gorm:"column:time_created;type:uint;default:0;" json:"time_created"`
	//[ 7] time_zoned                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TimeZoned uint32 `gorm:"column:time_zoned;type:uint;default:0;" json:"time_zoned"`
	//[ 8] time_completed                                 uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TimeCompleted uint32 `gorm:"column:time_completed;type:uint;default:0;" json:"time_completed"`
}

var adventure_detailsTableInfo = &TableInfo{
	Name: "adventure_details",
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
			Name:               "adventure_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "AdventureID",
			GoFieldType:        "uint32",
			JSONFieldName:      "adventure_id",
			ProtobufFieldName:  "adventure_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "instance_id",
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
			GoFieldName:        "InstanceID",
			GoFieldType:        "int32",
			JSONFieldName:      "instance_id",
			ProtobufFieldName:  "instance_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "count",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "Count",
			GoFieldType:        "uint32",
			JSONFieldName:      "count",
			ProtobufFieldName:  "count",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "assassinate_count",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "AssassinateCount",
			GoFieldType:        "uint32",
			JSONFieldName:      "assassinate_count",
			ProtobufFieldName:  "assassinate_count",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "status",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "uint32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "time_created",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "TimeCreated",
			GoFieldType:        "uint32",
			JSONFieldName:      "time_created",
			ProtobufFieldName:  "time_created",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "time_zoned",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "TimeZoned",
			GoFieldType:        "uint32",
			JSONFieldName:      "time_zoned",
			ProtobufFieldName:  "time_zoned",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "time_completed",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "TimeCompleted",
			GoFieldType:        "uint32",
			JSONFieldName:      "time_completed",
			ProtobufFieldName:  "time_completed",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AdventureDetails) TableName() string {
	return "adventure_details"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AdventureDetails) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AdventureDetails) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AdventureDetails) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AdventureDetails) TableInfo() *TableInfo {
	return adventure_detailsTableInfo
}
