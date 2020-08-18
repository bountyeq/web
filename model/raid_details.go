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


CREATE TABLE `raid_details` (
  `raidid` int(4) NOT NULL DEFAULT 0,
  `loottype` int(4) NOT NULL DEFAULT 0,
  `locked` tinyint(1) NOT NULL DEFAULT 0,
  `motd` varchar(1024) DEFAULT NULL,
  PRIMARY KEY (`raidid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "motd": "kKpWBjhUUUciJYPGkVMiGNsPm",    "raidid": 42,    "loottype": 89,    "locked": 47}



*/

// RaidDetails struct is a row record of the raid_details table in the eqemu database
type RaidDetails struct {
	//[ 0] raidid                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Raidid int32 `gorm:"primary_key;column:raidid;type:int;default:0;" json:"raidid"`
	//[ 1] loottype                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Loottype int32 `gorm:"column:loottype;type:int;default:0;" json:"loottype"`
	//[ 2] locked                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Locked int32 `gorm:"column:locked;type:tinyint;default:0;" json:"locked"`
	//[ 3] motd                                           varchar(1024)        null: true   primary: false  isArray: false  auto: false  col: varchar         len: 1024    default: [NULL]
	Motd null.String `gorm:"column:motd;type:varchar;size:1024;" json:"motd"`
}

var raid_detailsTableInfo = &TableInfo{
	Name: "raid_details",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "raidid",
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
			GoFieldName:        "Raidid",
			GoFieldType:        "int32",
			JSONFieldName:      "raidid",
			ProtobufFieldName:  "raidid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "loottype",
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
			GoFieldName:        "Loottype",
			GoFieldType:        "int32",
			JSONFieldName:      "loottype",
			ProtobufFieldName:  "loottype",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "locked",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Locked",
			GoFieldType:        "int32",
			JSONFieldName:      "locked",
			ProtobufFieldName:  "locked",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "motd",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       1024,
			GoFieldName:        "Motd",
			GoFieldType:        "null.String",
			JSONFieldName:      "motd",
			ProtobufFieldName:  "motd",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *RaidDetails) TableName() string {
	return "raid_details"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RaidDetails) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RaidDetails) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RaidDetails) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RaidDetails) TableInfo() *TableInfo {
	return raid_detailsTableInfo
}
