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


CREATE TABLE `adventure_template` (
  `id` int(10) unsigned NOT NULL,
  `zone` varchar(64) NOT NULL,
  `zone_version` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `is_hard` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `is_raid` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `min_level` tinyint(3) unsigned NOT NULL DEFAULT 1,
  `max_level` tinyint(3) unsigned NOT NULL DEFAULT 65,
  `type` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `type_data` int(10) unsigned NOT NULL DEFAULT 0,
  `type_count` smallint(5) unsigned NOT NULL DEFAULT 0,
  `assa_x` float NOT NULL DEFAULT 0,
  `assa_y` float NOT NULL DEFAULT 0,
  `assa_z` float NOT NULL DEFAULT 0,
  `assa_h` float NOT NULL DEFAULT 0,
  `text` varchar(511) DEFAULT NULL,
  `duration` int(10) unsigned NOT NULL DEFAULT 7200,
  `zone_in_time` int(10) unsigned NOT NULL DEFAULT 1800,
  `win_points` smallint(5) unsigned NOT NULL DEFAULT 0,
  `lose_points` smallint(5) unsigned NOT NULL DEFAULT 0,
  `theme` tinyint(3) unsigned NOT NULL DEFAULT 1,
  `zone_in_zone_id` smallint(5) unsigned NOT NULL DEFAULT 0,
  `zone_in_x` float NOT NULL DEFAULT 0,
  `zone_in_y` float NOT NULL DEFAULT 0,
  `zone_in_object_id` smallint(4) NOT NULL DEFAULT 0,
  `dest_x` float NOT NULL DEFAULT 0,
  `dest_y` float NOT NULL DEFAULT 0,
  `dest_z` float NOT NULL DEFAULT 0,
  `dest_h` float NOT NULL DEFAULT 0,
  `graveyard_zone_id` int(10) unsigned NOT NULL DEFAULT 0,
  `graveyard_x` float NOT NULL DEFAULT 0,
  `graveyard_y` float NOT NULL DEFAULT 0,
  `graveyard_z` float NOT NULL DEFAULT 0,
  `graveyard_radius` float unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `id_2` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "dest_y": 0.8172902,    "graveyard_z": 0.07792117,    "id": 4,    "zone_version": 34,    "assa_z": 0.1611645,    "assa_h": 0.7732988,    "win_points": 22,    "zone_in_zone_id": 2,    "graveyard_radius": 0.7369534,    "graveyard_y": 0.15228078,    "type_count": 92,    "zone_in_object_id": 96,    "dest_x": 0.19116612,    "dest_h": 0.6758751,    "is_raid": 27,    "zone_in_time": 94,    "graveyard_x": 0.34704292,    "zone": "ZwlXdajPOjMVhTqbxmlIAnExr",    "type_data": 49,    "duration": 11,    "zone_in_y": 0.16814445,    "is_hard": 16,    "type": 68,    "text": "ExfEeWQDqCuaAPdnHmVmkaIHH",    "dest_z": 0.17834142,    "graveyard_zone_id": 93,    "max_level": 59,    "assa_x": 0.7371924,    "lose_points": 68,    "zone_in_x": 0.9585418,    "min_level": 42,    "assa_y": 0.39147863,    "theme": 91}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[15] column is set for unsigned
[16] column is set for unsigned
[17] column is set for unsigned
[18] column is set for unsigned
[19] column is set for unsigned
[20] column is set for unsigned
[28] column is set for unsigned
[32] column is set for unsigned



*/

// AdventureTemplate struct is a row record of the adventure_template table in the eqemu database
type AdventureTemplate struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;column:id;type:uint;" json:"id"`
	//[ 1] zone                                           varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Zone string `gorm:"column:zone;type:varchar;size:64;" json:"zone"`
	//[ 2] zone_version                                   utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	ZoneVersion uint32 `gorm:"column:zone_version;type:utinyint;default:0;" json:"zone_version"`
	//[ 3] is_hard                                        utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	IsHard uint32 `gorm:"column:is_hard;type:utinyint;default:0;" json:"is_hard"`
	//[ 4] is_raid                                        utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	IsRaid uint32 `gorm:"column:is_raid;type:utinyint;default:0;" json:"is_raid"`
	//[ 5] min_level                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	MinLevel uint32 `gorm:"column:min_level;type:utinyint;default:1;" json:"min_level"`
	//[ 6] max_level                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [65]
	MaxLevel uint32 `gorm:"column:max_level;type:utinyint;default:65;" json:"max_level"`
	//[ 7] type                                           utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Type uint32 `gorm:"column:type;type:utinyint;default:0;" json:"type"`
	//[ 8] type_data                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TypeData uint32 `gorm:"column:type_data;type:uint;default:0;" json:"type_data"`
	//[ 9] type_count                                     usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	TypeCount uint32 `gorm:"column:type_count;type:usmallint;default:0;" json:"type_count"`
	//[10] assa_x                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	AssaX float32 `gorm:"column:assa_x;type:float;default:0;" json:"assa_x"`
	//[11] assa_y                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	AssaY float32 `gorm:"column:assa_y;type:float;default:0;" json:"assa_y"`
	//[12] assa_z                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	AssaZ float32 `gorm:"column:assa_z;type:float;default:0;" json:"assa_z"`
	//[13] assa_h                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	AssaH float32 `gorm:"column:assa_h;type:float;default:0;" json:"assa_h"`
	//[14] text                                           varchar(511)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 511     default: [NULL]
	Text null.String `gorm:"column:text;type:varchar;size:511;" json:"text"`
	//[15] duration                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [7200]
	Duration uint32 `gorm:"column:duration;type:uint;default:7200;" json:"duration"`
	//[16] zone_in_time                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1800]
	ZoneInTime uint32 `gorm:"column:zone_in_time;type:uint;default:1800;" json:"zone_in_time"`
	//[17] win_points                                     usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	WinPoints uint32 `gorm:"column:win_points;type:usmallint;default:0;" json:"win_points"`
	//[18] lose_points                                    usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	LosePoints uint32 `gorm:"column:lose_points;type:usmallint;default:0;" json:"lose_points"`
	//[19] theme                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	Theme uint32 `gorm:"column:theme;type:utinyint;default:1;" json:"theme"`
	//[20] zone_in_zone_id                                usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	ZoneInZoneID uint32 `gorm:"column:zone_in_zone_id;type:usmallint;default:0;" json:"zone_in_zone_id"`
	//[21] zone_in_x                                      float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	ZoneInX float32 `gorm:"column:zone_in_x;type:float;default:0;" json:"zone_in_x"`
	//[22] zone_in_y                                      float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	ZoneInY float32 `gorm:"column:zone_in_y;type:float;default:0;" json:"zone_in_y"`
	//[23] zone_in_object_id                              smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	ZoneInObjectID int32 `gorm:"column:zone_in_object_id;type:smallint;default:0;" json:"zone_in_object_id"`
	//[24] dest_x                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	DestX float32 `gorm:"column:dest_x;type:float;default:0;" json:"dest_x"`
	//[25] dest_y                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	DestY float32 `gorm:"column:dest_y;type:float;default:0;" json:"dest_y"`
	//[26] dest_z                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	DestZ float32 `gorm:"column:dest_z;type:float;default:0;" json:"dest_z"`
	//[27] dest_h                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	DestH float32 `gorm:"column:dest_h;type:float;default:0;" json:"dest_h"`
	//[28] graveyard_zone_id                              uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	GraveyardZoneID uint32 `gorm:"column:graveyard_zone_id;type:uint;default:0;" json:"graveyard_zone_id"`
	//[29] graveyard_x                                    float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	GraveyardX float32 `gorm:"column:graveyard_x;type:float;default:0;" json:"graveyard_x"`
	//[30] graveyard_y                                    float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	GraveyardY float32 `gorm:"column:graveyard_y;type:float;default:0;" json:"graveyard_y"`
	//[31] graveyard_z                                    float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	GraveyardZ float32 `gorm:"column:graveyard_z;type:float;default:0;" json:"graveyard_z"`
	//[32] graveyard_radius                               ufloat               null: false  primary: false  isArray: false  auto: false  col: ufloat          len: -1      default: [0]
	GraveyardRadius float32 `gorm:"column:graveyard_radius;type:ufloat;default:0;" json:"graveyard_radius"`
}

var adventure_templateTableInfo = &TableInfo{
	Name: "adventure_template",
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
			IsAutoIncrement:    false,
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
			Name:               "zone",
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
			GoFieldName:        "Zone",
			GoFieldType:        "string",
			JSONFieldName:      "zone",
			ProtobufFieldName:  "zone",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "zone_version",
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
			GoFieldName:        "ZoneVersion",
			GoFieldType:        "uint32",
			JSONFieldName:      "zone_version",
			ProtobufFieldName:  "zone_version",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "is_hard",
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
			GoFieldName:        "IsHard",
			GoFieldType:        "uint32",
			JSONFieldName:      "is_hard",
			ProtobufFieldName:  "is_hard",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "is_raid",
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
			GoFieldName:        "IsRaid",
			GoFieldType:        "uint32",
			JSONFieldName:      "is_raid",
			ProtobufFieldName:  "is_raid",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "min_level",
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
			GoFieldName:        "MinLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "min_level",
			ProtobufFieldName:  "min_level",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "max_level",
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
			GoFieldName:        "MaxLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "max_level",
			ProtobufFieldName:  "max_level",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "uint32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "type_data",
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
			GoFieldName:        "TypeData",
			GoFieldType:        "uint32",
			JSONFieldName:      "type_data",
			ProtobufFieldName:  "type_data",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "type_count",
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
			GoFieldName:        "TypeCount",
			GoFieldType:        "uint32",
			JSONFieldName:      "type_count",
			ProtobufFieldName:  "type_count",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "assa_x",
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
			GoFieldName:        "AssaX",
			GoFieldType:        "float32",
			JSONFieldName:      "assa_x",
			ProtobufFieldName:  "assa_x",
			ProtobufType:       "float",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "assa_y",
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
			GoFieldName:        "AssaY",
			GoFieldType:        "float32",
			JSONFieldName:      "assa_y",
			ProtobufFieldName:  "assa_y",
			ProtobufType:       "float",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "assa_z",
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
			GoFieldName:        "AssaZ",
			GoFieldType:        "float32",
			JSONFieldName:      "assa_z",
			ProtobufFieldName:  "assa_z",
			ProtobufType:       "float",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "assa_h",
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
			GoFieldName:        "AssaH",
			GoFieldType:        "float32",
			JSONFieldName:      "assa_h",
			ProtobufFieldName:  "assa_h",
			ProtobufType:       "float",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "text",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(511)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       511,
			GoFieldName:        "Text",
			GoFieldType:        "null.String",
			JSONFieldName:      "text",
			ProtobufFieldName:  "text",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "zone_in_time",
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
			GoFieldName:        "ZoneInTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "zone_in_time",
			ProtobufFieldName:  "zone_in_time",
			ProtobufType:       "uint32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "win_points",
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
			GoFieldName:        "WinPoints",
			GoFieldType:        "uint32",
			JSONFieldName:      "win_points",
			ProtobufFieldName:  "win_points",
			ProtobufType:       "uint32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "lose_points",
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
			GoFieldName:        "LosePoints",
			GoFieldType:        "uint32",
			JSONFieldName:      "lose_points",
			ProtobufFieldName:  "lose_points",
			ProtobufType:       "uint32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "theme",
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
			GoFieldName:        "Theme",
			GoFieldType:        "uint32",
			JSONFieldName:      "theme",
			ProtobufFieldName:  "theme",
			ProtobufType:       "uint32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "zone_in_zone_id",
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
			GoFieldName:        "ZoneInZoneID",
			GoFieldType:        "uint32",
			JSONFieldName:      "zone_in_zone_id",
			ProtobufFieldName:  "zone_in_zone_id",
			ProtobufType:       "uint32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "zone_in_x",
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
			GoFieldName:        "ZoneInX",
			GoFieldType:        "float32",
			JSONFieldName:      "zone_in_x",
			ProtobufFieldName:  "zone_in_x",
			ProtobufType:       "float",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "zone_in_y",
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
			GoFieldName:        "ZoneInY",
			GoFieldType:        "float32",
			JSONFieldName:      "zone_in_y",
			ProtobufFieldName:  "zone_in_y",
			ProtobufType:       "float",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "zone_in_object_id",
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
			GoFieldName:        "ZoneInObjectID",
			GoFieldType:        "int32",
			JSONFieldName:      "zone_in_object_id",
			ProtobufFieldName:  "zone_in_object_id",
			ProtobufType:       "int32",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "dest_x",
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
			GoFieldName:        "DestX",
			GoFieldType:        "float32",
			JSONFieldName:      "dest_x",
			ProtobufFieldName:  "dest_x",
			ProtobufType:       "float",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "dest_y",
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
			GoFieldName:        "DestY",
			GoFieldType:        "float32",
			JSONFieldName:      "dest_y",
			ProtobufFieldName:  "dest_y",
			ProtobufType:       "float",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "dest_z",
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
			GoFieldName:        "DestZ",
			GoFieldType:        "float32",
			JSONFieldName:      "dest_z",
			ProtobufFieldName:  "dest_z",
			ProtobufType:       "float",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "dest_h",
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
			GoFieldName:        "DestH",
			GoFieldType:        "float32",
			JSONFieldName:      "dest_h",
			ProtobufFieldName:  "dest_h",
			ProtobufType:       "float",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "graveyard_zone_id",
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
			GoFieldName:        "GraveyardZoneID",
			GoFieldType:        "uint32",
			JSONFieldName:      "graveyard_zone_id",
			ProtobufFieldName:  "graveyard_zone_id",
			ProtobufType:       "uint32",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "graveyard_x",
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
			GoFieldName:        "GraveyardX",
			GoFieldType:        "float32",
			JSONFieldName:      "graveyard_x",
			ProtobufFieldName:  "graveyard_x",
			ProtobufType:       "float",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "graveyard_y",
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
			GoFieldName:        "GraveyardY",
			GoFieldType:        "float32",
			JSONFieldName:      "graveyard_y",
			ProtobufFieldName:  "graveyard_y",
			ProtobufType:       "float",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "graveyard_z",
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
			GoFieldName:        "GraveyardZ",
			GoFieldType:        "float32",
			JSONFieldName:      "graveyard_z",
			ProtobufFieldName:  "graveyard_z",
			ProtobufType:       "float",
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "graveyard_radius",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ufloat",
			DatabaseTypePretty: "ufloat",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ufloat",
			ColumnLength:       -1,
			GoFieldName:        "GraveyardRadius",
			GoFieldType:        "float32",
			JSONFieldName:      "graveyard_radius",
			ProtobufFieldName:  "graveyard_radius",
			ProtobufType:       "float",
			ProtobufPos:        33,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AdventureTemplate) TableName() string {
	return "adventure_template"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AdventureTemplate) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AdventureTemplate) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AdventureTemplate) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AdventureTemplate) TableInfo() *TableInfo {
	return adventure_templateTableInfo
}
