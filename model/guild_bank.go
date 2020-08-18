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


CREATE TABLE `guild_bank` (
  `guildid` int(10) unsigned NOT NULL DEFAULT 0,
  `area` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `slot` int(4) unsigned NOT NULL DEFAULT 0,
  `itemid` int(10) unsigned NOT NULL DEFAULT 0,
  `qty` int(10) unsigned NOT NULL DEFAULT 0,
  `donator` varchar(64) DEFAULT NULL,
  `permissions` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `whofor` varchar(64) DEFAULT NULL,
  KEY `guildid` (`guildid`),
  KEY `area` (`area`),
  KEY `slot` (`slot`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "guildid": 60,    "area": 41,    "slot": 49,    "itemid": 83,    "qty": 47,    "donator": "YMyYSWbTCXmqyeIdVRfOJaISU",    "permissions": 75,    "whofor": "DNGaLdFnnvURcUlMMQEGICWtx"}


Comments
-------------------------------------
[ 0] column is set for unsignedWarning table: guild_bank does not have a primary key defined, setting col position 1 guildid as primary key

[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 6] column is set for unsigned



*/

// GuildBank struct is a row record of the guild_bank table in the eqemu database
type GuildBank struct {
	//[ 0] guildid                                        uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Guildid uint32 `gorm:"primary_key;column:guildid;type:uint;default:0;" json:"guildid"`
	//[ 1] area                                           utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Area uint32 `gorm:"column:area;type:utinyint;default:0;" json:"area"`
	//[ 2] slot                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Slot uint32 `gorm:"column:slot;type:uint;default:0;" json:"slot"`
	//[ 3] itemid                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Itemid uint32 `gorm:"column:itemid;type:uint;default:0;" json:"itemid"`
	//[ 4] qty                                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Qty uint32 `gorm:"column:qty;type:uint;default:0;" json:"qty"`
	//[ 5] donator                                        varchar(64)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 64      default: [NULL]
	Donator null.String `gorm:"column:donator;type:varchar;size:64;" json:"donator"`
	//[ 6] permissions                                    utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Permissions uint32 `gorm:"column:permissions;type:utinyint;default:0;" json:"permissions"`
	//[ 7] whofor                                         varchar(64)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 64      default: [NULL]
	Whofor null.String `gorm:"column:whofor;type:varchar;size:64;" json:"whofor"`
}

var guild_bankTableInfo = &TableInfo{
	Name: "guild_bank",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "guildid",
			Comment: ``,
			Notes: `column is set for unsignedWarning table: guild_bank does not have a primary key defined, setting col position 1 guildid as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Guildid",
			GoFieldType:        "uint32",
			JSONFieldName:      "guildid",
			ProtobufFieldName:  "guildid",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "area",
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
			GoFieldName:        "Area",
			GoFieldType:        "uint32",
			JSONFieldName:      "area",
			ProtobufFieldName:  "area",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "slot",
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
			GoFieldName:        "Slot",
			GoFieldType:        "uint32",
			JSONFieldName:      "slot",
			ProtobufFieldName:  "slot",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "itemid",
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
			GoFieldName:        "Itemid",
			GoFieldType:        "uint32",
			JSONFieldName:      "itemid",
			ProtobufFieldName:  "itemid",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "qty",
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
			GoFieldName:        "Qty",
			GoFieldType:        "uint32",
			JSONFieldName:      "qty",
			ProtobufFieldName:  "qty",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "donator",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Donator",
			GoFieldType:        "null.String",
			JSONFieldName:      "donator",
			ProtobufFieldName:  "donator",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "permissions",
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
			GoFieldName:        "Permissions",
			GoFieldType:        "uint32",
			JSONFieldName:      "permissions",
			ProtobufFieldName:  "permissions",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "whofor",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Whofor",
			GoFieldType:        "null.String",
			JSONFieldName:      "whofor",
			ProtobufFieldName:  "whofor",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *GuildBank) TableName() string {
	return "guild_bank"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GuildBank) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GuildBank) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GuildBank) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GuildBank) TableInfo() *TableInfo {
	return guild_bankTableInfo
}
