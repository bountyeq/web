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


CREATE TABLE `zone_points` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `zone` varchar(32) DEFAULT NULL,
  `version` int(11) NOT NULL DEFAULT 0,
  `number` smallint(4) unsigned NOT NULL DEFAULT 1,
  `y` float NOT NULL DEFAULT 0,
  `x` float NOT NULL DEFAULT 0,
  `z` float NOT NULL DEFAULT 0,
  `heading` float NOT NULL DEFAULT 0,
  `target_y` float NOT NULL DEFAULT 0,
  `target_x` float NOT NULL DEFAULT 0,
  `target_z` float NOT NULL DEFAULT 0,
  `target_heading` float NOT NULL DEFAULT 0,
  `zoneinst` smallint(5) unsigned DEFAULT 0,
  `target_zone_id` int(10) unsigned NOT NULL DEFAULT 0,
  `target_instance` int(10) unsigned NOT NULL DEFAULT 0,
  `buffer` float DEFAULT 0,
  `client_version_mask` int(10) unsigned NOT NULL DEFAULT 4294967295,
  `min_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `max_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `content_flags` varchar(100) DEFAULT NULL,
  `content_flags_disabled` varchar(100) DEFAULT NULL,
  `is_virtual` tinyint(4) NOT NULL DEFAULT 0,
  `height` int(11) NOT NULL DEFAULT 0,
  `width` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `NewIndex` (`number`,`zone`),
  KEY `zone_points_target_idx` (`target_zone_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2393 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "zone": "kgTaoErWbaXmrwNaGBJeNTrsp",    "version": 82,    "target_x": 0.2750955,    "id": 23,    "target_heading": 0.7583491,    "target_zone_id": 71,    "client_version_mask": 46,    "height": 55,    "z": 0.048796758,    "target_y": 0.43683007,    "target_z": 0.6280194,    "zoneinst": 24,    "min_expansion": 50,    "max_expansion": 10,    "content_flags_disabled": "DarMkKuBvghdpDmIeDDkgGYMr",    "width": 83,    "x": 0.4562789,    "y": 0.18928364,    "heading": 0.3953264,    "target_instance": 16,    "buffer": 0.85989577,    "content_flags": "mpwdtiIyucBKZnHIFuFYweVMJ",    "is_virtual": 83,    "number": 4}


Comments
-------------------------------------
[ 3] column is set for unsigned
[12] column is set for unsigned
[13] column is set for unsigned
[14] column is set for unsigned
[16] column is set for unsigned
[17] column is set for unsigned
[18] column is set for unsigned



*/

// ZonePoints struct is a row record of the zone_points table in the eqemu database
type ZonePoints struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] zone                                           varchar(32)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 32      default: [NULL]
	Zone null.String `gorm:"column:zone;type:varchar;size:32;" json:"zone"`
	//[ 2] version                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Version int32 `gorm:"column:version;type:int;default:0;" json:"version"`
	//[ 3] number                                         usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [1]
	Number uint32 `gorm:"column:number;type:usmallint;default:1;" json:"number"`
	//[ 4] y                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Y float32 `gorm:"column:y;type:float;default:0;" json:"y"`
	//[ 5] x                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	X float32 `gorm:"column:x;type:float;default:0;" json:"x"`
	//[ 6] z                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Z float32 `gorm:"column:z;type:float;default:0;" json:"z"`
	//[ 7] heading                                        float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Heading float32 `gorm:"column:heading;type:float;default:0;" json:"heading"`
	//[ 8] target_y                                       float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	TargetY float32 `gorm:"column:target_y;type:float;default:0;" json:"target_y"`
	//[ 9] target_x                                       float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	TargetX float32 `gorm:"column:target_x;type:float;default:0;" json:"target_x"`
	//[10] target_z                                       float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	TargetZ float32 `gorm:"column:target_z;type:float;default:0;" json:"target_z"`
	//[11] target_heading                                 float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	TargetHeading float32 `gorm:"column:target_heading;type:float;default:0;" json:"target_heading"`
	//[12] zoneinst                                       usmallint            null: true   primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	Zoneinst null.Int `gorm:"column:zoneinst;type:usmallint;default:0;" json:"zoneinst"`
	//[13] target_zone_id                                 uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TargetZoneID uint32 `gorm:"column:target_zone_id;type:uint;default:0;" json:"target_zone_id"`
	//[14] target_instance                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TargetInstance uint32 `gorm:"column:target_instance;type:uint;default:0;" json:"target_instance"`
	//[15] buffer                                         float                null: true   primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Buffer null.Float `gorm:"column:buffer;type:float;default:0;" json:"buffer"`
	//[16] client_version_mask                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [4294967295]
	ClientVersionMask uint32 `gorm:"column:client_version_mask;type:uint;default:4294967295;" json:"client_version_mask"`
	//[17] min_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MinExpansion uint32 `gorm:"column:min_expansion;type:utinyint;default:0;" json:"min_expansion"`
	//[18] max_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MaxExpansion uint32 `gorm:"column:max_expansion;type:utinyint;default:0;" json:"max_expansion"`
	//[19] content_flags                                  varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlags null.String `gorm:"column:content_flags;type:varchar;size:100;" json:"content_flags"`
	//[20] content_flags_disabled                         varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlagsDisabled null.String `gorm:"column:content_flags_disabled;type:varchar;size:100;" json:"content_flags_disabled"`
	//[21] is_virtual                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsVirtual int32 `gorm:"column:is_virtual;type:tinyint;default:0;" json:"is_virtual"`
	//[22] height                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Height int32 `gorm:"column:height;type:int;default:0;" json:"height"`
	//[23] width                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Width int32 `gorm:"column:width;type:int;default:0;" json:"width"`
}

var zone_pointsTableInfo = &TableInfo{
	Name: "zone_points",
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
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Zone",
			GoFieldType:        "null.String",
			JSONFieldName:      "zone",
			ProtobufFieldName:  "zone",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "version",
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
			GoFieldName:        "Version",
			GoFieldType:        "int32",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "number",
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
			GoFieldName:        "Number",
			GoFieldType:        "uint32",
			JSONFieldName:      "number",
			ProtobufFieldName:  "number",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "y",
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
			GoFieldName:        "Y",
			GoFieldType:        "float32",
			JSONFieldName:      "y",
			ProtobufFieldName:  "y",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "x",
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
			GoFieldName:        "X",
			GoFieldType:        "float32",
			JSONFieldName:      "x",
			ProtobufFieldName:  "x",
			ProtobufType:       "float",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "z",
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
			GoFieldName:        "Z",
			GoFieldType:        "float32",
			JSONFieldName:      "z",
			ProtobufFieldName:  "z",
			ProtobufType:       "float",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "target_y",
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
			GoFieldName:        "TargetY",
			GoFieldType:        "float32",
			JSONFieldName:      "target_y",
			ProtobufFieldName:  "target_y",
			ProtobufType:       "float",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "target_x",
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
			GoFieldName:        "TargetX",
			GoFieldType:        "float32",
			JSONFieldName:      "target_x",
			ProtobufFieldName:  "target_x",
			ProtobufType:       "float",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "target_z",
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
			GoFieldName:        "TargetZ",
			GoFieldType:        "float32",
			JSONFieldName:      "target_z",
			ProtobufFieldName:  "target_z",
			ProtobufType:       "float",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "target_heading",
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
			GoFieldName:        "TargetHeading",
			GoFieldType:        "float32",
			JSONFieldName:      "target_heading",
			ProtobufFieldName:  "target_heading",
			ProtobufType:       "float",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "zoneinst",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "Zoneinst",
			GoFieldType:        "null.Int",
			JSONFieldName:      "zoneinst",
			ProtobufFieldName:  "zoneinst",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "target_zone_id",
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
			GoFieldName:        "TargetZoneID",
			GoFieldType:        "uint32",
			JSONFieldName:      "target_zone_id",
			ProtobufFieldName:  "target_zone_id",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "target_instance",
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
			GoFieldName:        "TargetInstance",
			GoFieldType:        "uint32",
			JSONFieldName:      "target_instance",
			ProtobufFieldName:  "target_instance",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "buffer",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "Buffer",
			GoFieldType:        "null.Float",
			JSONFieldName:      "buffer",
			ProtobufFieldName:  "buffer",
			ProtobufType:       "float",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "client_version_mask",
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
			GoFieldName:        "ClientVersionMask",
			GoFieldType:        "uint32",
			JSONFieldName:      "client_version_mask",
			ProtobufFieldName:  "client_version_mask",
			ProtobufType:       "uint32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
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
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
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
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "is_virtual",
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
			GoFieldName:        "IsVirtual",
			GoFieldType:        "int32",
			JSONFieldName:      "is_virtual",
			ProtobufFieldName:  "is_virtual",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "height",
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
			GoFieldName:        "Height",
			GoFieldType:        "int32",
			JSONFieldName:      "height",
			ProtobufFieldName:  "height",
			ProtobufType:       "int32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "width",
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
			GoFieldName:        "Width",
			GoFieldType:        "int32",
			JSONFieldName:      "width",
			ProtobufFieldName:  "width",
			ProtobufType:       "int32",
			ProtobufPos:        24,
		},
	},
}

// TableName sets the insert table name for this struct type
func (z *ZonePoints) TableName() string {
	return "zone_points"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (z *ZonePoints) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (z *ZonePoints) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (z *ZonePoints) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (z *ZonePoints) TableInfo() *TableInfo {
	return zone_pointsTableInfo
}
