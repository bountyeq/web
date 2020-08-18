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


CREATE TABLE `launcher_zones` (
  `launcher` varchar(64) NOT NULL DEFAULT '',
  `zone` varchar(32) NOT NULL DEFAULT '',
  `port` mediumint(9) NOT NULL DEFAULT 0,
  PRIMARY KEY (`launcher`,`zone`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "launcher": "XBCqDBRlGqaXGKUQdKqlmdNwU",    "zone": "aCNlMqobrvIEplcwdTWwBRmNY",    "port": 21}



*/

// LauncherZones struct is a row record of the launcher_zones table in the eqemu database
type LauncherZones struct {
	//[ 0] launcher                                       varchar(64)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Launcher string `gorm:"primary_key;column:launcher;type:varchar;size:64;default:'';" json:"launcher"`
	//[ 1] zone                                           varchar(32)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 32      default: ['']
	Zone string `gorm:"primary_key;column:zone;type:varchar;size:32;default:'';" json:"zone"`
	//[ 2] port                                           mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	Port int32 `gorm:"column:port;type:mediumint;default:0;" json:"port"`
}

var launcher_zonesTableInfo = &TableInfo{
	Name: "launcher_zones",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "launcher",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Launcher",
			GoFieldType:        "string",
			JSONFieldName:      "launcher",
			ProtobufFieldName:  "launcher",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "zone",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Zone",
			GoFieldType:        "string",
			JSONFieldName:      "zone",
			ProtobufFieldName:  "zone",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "port",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "mediumint",
			DatabaseTypePretty: "mediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "mediumint",
			ColumnLength:       -1,
			GoFieldName:        "Port",
			GoFieldType:        "int32",
			JSONFieldName:      "port",
			ProtobufFieldName:  "port",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LauncherZones) TableName() string {
	return "launcher_zones"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LauncherZones) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LauncherZones) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LauncherZones) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LauncherZones) TableInfo() *TableInfo {
	return launcher_zonesTableInfo
}
