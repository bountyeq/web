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


CREATE TABLE `ground_spawns` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `zoneid` int(10) unsigned NOT NULL DEFAULT 0,
  `version` smallint(5) NOT NULL DEFAULT 0,
  `max_x` float NOT NULL DEFAULT 2000,
  `max_y` float NOT NULL DEFAULT 2000,
  `max_z` float NOT NULL DEFAULT 10000,
  `min_x` float NOT NULL DEFAULT -2000,
  `min_y` float NOT NULL DEFAULT -2000,
  `heading` float NOT NULL DEFAULT 0,
  `name` varchar(16) NOT NULL DEFAULT '',
  `item` int(10) unsigned NOT NULL DEFAULT 0,
  `max_allowed` int(10) unsigned NOT NULL DEFAULT 1,
  `comment` varchar(255) NOT NULL DEFAULT '',
  `respawn_timer` int(11) unsigned NOT NULL DEFAULT 300,
  `min_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `max_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `content_flags` varchar(100) DEFAULT NULL,
  `content_flags_disabled` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `zone` (`zoneid`)
) ENGINE=InnoDB AUTO_INCREMENT=1748 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "max_allowed": 51,    "zoneid": 5,    "min_y": 0.13661472,    "name": "ZhnPAwsXcEUWRKRPnWluQwkmP",    "item": 2,    "content_flags_disabled": "MdpnUGEphQgMiqjAOeGXRgoBi",    "id": 61,    "max_y": 0.7615789,    "heading": 0.6927448,    "max_expansion": 70,    "min_expansion": 65,    "content_flags": "DBiTITsAIIPukhoAKCkPWqiZB",    "version": 97,    "min_x": 0.13411671,    "comment": "cVXwgpZLuNVwOEMbjWOhnXBqn",    "respawn_timer": 13,    "max_x": 0.853692,    "max_z": 0.20840307}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[10] column is set for unsigned
[11] column is set for unsigned
[13] column is set for unsigned
[14] column is set for unsigned
[15] column is set for unsigned



*/

// GroundSpawns struct is a row record of the ground_spawns table in the eqemu database
type GroundSpawns struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] zoneid                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Zoneid uint32 `gorm:"column:zoneid;type:uint;default:0;" json:"zoneid"`
	//[ 2] version                                        smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Version int32 `gorm:"column:version;type:smallint;default:0;" json:"version"`
	//[ 3] max_x                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [2000]
	MaxX float32 `gorm:"column:max_x;type:float;default:2000;" json:"max_x"`
	//[ 4] max_y                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [2000]
	MaxY float32 `gorm:"column:max_y;type:float;default:2000;" json:"max_y"`
	//[ 5] max_z                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [10000]
	MaxZ float32 `gorm:"column:max_z;type:float;default:10000;" json:"max_z"`
	//[ 6] min_x                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [-2000]
	MinX float32 `gorm:"column:min_x;type:float;default:-2000;" json:"min_x"`
	//[ 7] min_y                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [-2000]
	MinY float32 `gorm:"column:min_y;type:float;default:-2000;" json:"min_y"`
	//[ 8] heading                                        float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Heading float32 `gorm:"column:heading;type:float;default:0;" json:"heading"`
	//[ 9] name                                           varchar(16)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 16      default: ['']
	Name string `gorm:"column:name;type:varchar;size:16;default:'';" json:"name"`
	//[10] item                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Item uint32 `gorm:"column:item;type:uint;default:0;" json:"item"`
	//[11] max_allowed                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	MaxAllowed uint32 `gorm:"column:max_allowed;type:uint;default:1;" json:"max_allowed"`
	//[12] comment                                        varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: ['']
	Comment string `gorm:"column:comment;type:varchar;size:255;default:'';" json:"comment"`
	//[13] respawn_timer                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [300]
	RespawnTimer uint32 `gorm:"column:respawn_timer;type:uint;default:300;" json:"respawn_timer"`
	//[14] min_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MinExpansion uint32 `gorm:"column:min_expansion;type:utinyint;default:0;" json:"min_expansion"`
	//[15] max_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MaxExpansion uint32 `gorm:"column:max_expansion;type:utinyint;default:0;" json:"max_expansion"`
	//[16] content_flags                                  varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlags null.String `gorm:"column:content_flags;type:varchar;size:100;" json:"content_flags"`
	//[17] content_flags_disabled                         varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlagsDisabled null.String `gorm:"column:content_flags_disabled;type:varchar;size:100;" json:"content_flags_disabled"`
}

var ground_spawnsTableInfo = &TableInfo{
	Name: "ground_spawns",
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
			Name:               "zoneid",
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
			GoFieldName:        "Zoneid",
			GoFieldType:        "uint32",
			JSONFieldName:      "zoneid",
			ProtobufFieldName:  "zoneid",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "version",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "Version",
			GoFieldType:        "int32",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "max_x",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "MaxX",
			GoFieldType:        "float32",
			JSONFieldName:      "max_x",
			ProtobufFieldName:  "max_x",
			ProtobufType:       "float",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "max_y",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "MaxY",
			GoFieldType:        "float32",
			JSONFieldName:      "max_y",
			ProtobufFieldName:  "max_y",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "max_z",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "MaxZ",
			GoFieldType:        "float32",
			JSONFieldName:      "max_z",
			ProtobufFieldName:  "max_z",
			ProtobufType:       "float",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "min_x",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "MinX",
			GoFieldType:        "float32",
			JSONFieldName:      "min_x",
			ProtobufFieldName:  "min_x",
			ProtobufType:       "float",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "min_y",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "MinY",
			GoFieldType:        "float32",
			JSONFieldName:      "min_y",
			ProtobufFieldName:  "min_y",
			ProtobufType:       "float",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "heading",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "Heading",
			GoFieldType:        "float32",
			JSONFieldName:      "heading",
			ProtobufFieldName:  "heading",
			ProtobufType:       "float",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(16)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       16,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "item",
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
			GoFieldName:        "Item",
			GoFieldType:        "uint32",
			JSONFieldName:      "item",
			ProtobufFieldName:  "item",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "max_allowed",
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
			GoFieldName:        "MaxAllowed",
			GoFieldType:        "uint32",
			JSONFieldName:      "max_allowed",
			ProtobufFieldName:  "max_allowed",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "comment",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Comment",
			GoFieldType:        "string",
			JSONFieldName:      "comment",
			ProtobufFieldName:  "comment",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "respawn_timer",
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
			GoFieldName:        "RespawnTimer",
			GoFieldType:        "uint32",
			JSONFieldName:      "respawn_timer",
			ProtobufFieldName:  "respawn_timer",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "min_expansion",
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
			GoFieldName:        "MinExpansion",
			GoFieldType:        "uint32",
			JSONFieldName:      "min_expansion",
			ProtobufFieldName:  "min_expansion",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "max_expansion",
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
			GoFieldName:        "MaxExpansion",
			GoFieldType:        "uint32",
			JSONFieldName:      "max_expansion",
			ProtobufFieldName:  "max_expansion",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "content_flags",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "ContentFlags",
			GoFieldType:        "null.String",
			JSONFieldName:      "content_flags",
			ProtobufFieldName:  "content_flags",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "content_flags_disabled",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "ContentFlagsDisabled",
			GoFieldType:        "null.String",
			JSONFieldName:      "content_flags_disabled",
			ProtobufFieldName:  "content_flags_disabled",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *GroundSpawns) TableName() string {
	return "ground_spawns"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GroundSpawns) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GroundSpawns) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GroundSpawns) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GroundSpawns) TableInfo() *TableInfo {
	return ground_spawnsTableInfo
}
