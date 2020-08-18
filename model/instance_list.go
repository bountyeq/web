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


CREATE TABLE `instance_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `zone` int(11) unsigned NOT NULL DEFAULT 0,
  `version` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `is_global` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `start_time` int(11) unsigned NOT NULL DEFAULT 0,
  `duration` int(11) unsigned NOT NULL DEFAULT 0,
  `never_expires` tinyint(3) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `id_2` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "never_expires": 65,    "id": 15,    "zone": 77,    "version": 92,    "is_global": 29,    "start_time": 97,    "duration": 67}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned



*/

// InstanceList struct is a row record of the instance_list table in the eqemu database
type InstanceList struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] zone                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Zone uint32 `gorm:"column:zone;type:uint;default:0;" json:"zone"`
	//[ 2] version                                        utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Version uint32 `gorm:"column:version;type:utinyint;default:0;" json:"version"`
	//[ 3] is_global                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	IsGlobal uint32 `gorm:"column:is_global;type:utinyint;default:0;" json:"is_global"`
	//[ 4] start_time                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	StartTime uint32 `gorm:"column:start_time;type:uint;default:0;" json:"start_time"`
	//[ 5] duration                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Duration uint32 `gorm:"column:duration;type:uint;default:0;" json:"duration"`
	//[ 6] never_expires                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	NeverExpires uint32 `gorm:"column:never_expires;type:utinyint;default:0;" json:"never_expires"`
}

var instance_listTableInfo = &TableInfo{
	Name: "instance_list",
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
			IsAutoIncrement:    true,
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
			Name:               "zone",
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
			GoFieldName:        "Zone",
			GoFieldType:        "uint32",
			JSONFieldName:      "zone",
			ProtobufFieldName:  "zone",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "version",
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
			GoFieldName:        "Version",
			GoFieldType:        "uint32",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "is_global",
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
			GoFieldName:        "IsGlobal",
			GoFieldType:        "uint32",
			JSONFieldName:      "is_global",
			ProtobufFieldName:  "is_global",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "start_time",
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
			GoFieldName:        "StartTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "start_time",
			ProtobufFieldName:  "start_time",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "duration",
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
			GoFieldName:        "Duration",
			GoFieldType:        "uint32",
			JSONFieldName:      "duration",
			ProtobufFieldName:  "duration",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "never_expires",
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
			GoFieldName:        "NeverExpires",
			GoFieldType:        "uint32",
			JSONFieldName:      "never_expires",
			ProtobufFieldName:  "never_expires",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (i *InstanceList) TableName() string {
	return "instance_list"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *InstanceList) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *InstanceList) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *InstanceList) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *InstanceList) TableInfo() *TableInfo {
	return instance_listTableInfo
}
