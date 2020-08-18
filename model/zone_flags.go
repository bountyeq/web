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


CREATE TABLE `zone_flags` (
  `charID` int(11) NOT NULL DEFAULT 0,
  `zoneID` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`charID`,`zoneID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "char_id": 56,    "zone_id": 19}



*/

// ZoneFlags struct is a row record of the zone_flags table in the eqemu database
type ZoneFlags struct {
	//[ 0] charID                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	CharID int32 `gorm:"primary_key;column:charID;type:int;default:0;" json:"char_id"`
	//[ 1] zoneID                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	ZoneID int32 `gorm:"primary_key;column:zoneID;type:int;default:0;" json:"zone_id"`
}

var zone_flagsTableInfo = &TableInfo{
	Name: "zone_flags",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "charID",
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
			GoFieldName:        "CharID",
			GoFieldType:        "int32",
			JSONFieldName:      "char_id",
			ProtobufFieldName:  "char_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "zoneID",
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
			GoFieldName:        "ZoneID",
			GoFieldType:        "int32",
			JSONFieldName:      "zone_id",
			ProtobufFieldName:  "zone_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (z *ZoneFlags) TableName() string {
	return "zone_flags"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (z *ZoneFlags) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (z *ZoneFlags) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (z *ZoneFlags) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (z *ZoneFlags) TableInfo() *TableInfo {
	return zone_flagsTableInfo
}
