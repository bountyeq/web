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


CREATE TABLE `raid_members` (
  `raidid` int(4) NOT NULL DEFAULT 0,
  `charid` int(4) NOT NULL DEFAULT 0,
  `groupid` int(4) unsigned NOT NULL DEFAULT 0,
  `_class` tinyint(4) NOT NULL DEFAULT 0,
  `level` tinyint(4) NOT NULL DEFAULT 0,
  `name` varchar(64) NOT NULL DEFAULT '',
  `isgroupleader` tinyint(1) NOT NULL DEFAULT 0,
  `israidleader` tinyint(1) NOT NULL DEFAULT 0,
  `islooter` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`charid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "raidid": 62,    "level": 52,    "name": "UBxuUZnvIhSAlABucOBxtZWrm",    "charid": 51,    "groupid": 34,    "_class": 64,    "isgroupleader": 68,    "israidleader": 28,    "islooter": 29}


Comments
-------------------------------------
[ 2] column is set for unsigned



*/

// RaidMembers struct is a row record of the raid_members table in the eqemu database
type RaidMembers struct {
	//[ 0] raidid                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Raidid int32 `gorm:"column:raidid;type:int;default:0;" json:"raidid"`
	//[ 1] charid                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Charid int32 `gorm:"primary_key;column:charid;type:int;default:0;" json:"charid"`
	//[ 2] groupid                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Groupid uint32 `gorm:"column:groupid;type:uint;default:0;" json:"groupid"`
	//[ 3] _class                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Class int32 `gorm:"column:_class;type:tinyint;default:0;" json:"_class"`
	//[ 4] level                                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Level int32 `gorm:"column:level;type:tinyint;default:0;" json:"level"`
	//[ 5] name                                           varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Name string `gorm:"column:name;type:varchar;size:64;default:'';" json:"name"`
	//[ 6] isgroupleader                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Isgroupleader int32 `gorm:"column:isgroupleader;type:tinyint;default:0;" json:"isgroupleader"`
	//[ 7] israidleader                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Israidleader int32 `gorm:"column:israidleader;type:tinyint;default:0;" json:"israidleader"`
	//[ 8] islooter                                       tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Islooter int32 `gorm:"column:islooter;type:tinyint;default:0;" json:"islooter"`
}

var raid_membersTableInfo = &TableInfo{
	Name: "raid_members",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "raidid",
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
			GoFieldName:        "Raidid",
			GoFieldType:        "int32",
			JSONFieldName:      "raidid",
			ProtobufFieldName:  "raidid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "charid",
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
			GoFieldName:        "Charid",
			GoFieldType:        "int32",
			JSONFieldName:      "charid",
			ProtobufFieldName:  "charid",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "groupid",
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
			GoFieldName:        "Groupid",
			GoFieldType:        "uint32",
			JSONFieldName:      "groupid",
			ProtobufFieldName:  "groupid",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "_class",
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
			GoFieldName:        "Class",
			GoFieldType:        "int32",
			JSONFieldName:      "_class",
			ProtobufFieldName:  "_class",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "level",
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
			GoFieldName:        "Level",
			GoFieldType:        "int32",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "isgroupleader",
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
			GoFieldName:        "Isgroupleader",
			GoFieldType:        "int32",
			JSONFieldName:      "isgroupleader",
			ProtobufFieldName:  "isgroupleader",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "israidleader",
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
			GoFieldName:        "Israidleader",
			GoFieldType:        "int32",
			JSONFieldName:      "israidleader",
			ProtobufFieldName:  "israidleader",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "islooter",
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
			GoFieldName:        "Islooter",
			GoFieldType:        "int32",
			JSONFieldName:      "islooter",
			ProtobufFieldName:  "islooter",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *RaidMembers) TableName() string {
	return "raid_members"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RaidMembers) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RaidMembers) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RaidMembers) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RaidMembers) TableInfo() *TableInfo {
	return raid_membersTableInfo
}
