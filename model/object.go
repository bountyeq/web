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


CREATE TABLE `object` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `zoneid` int(11) unsigned NOT NULL DEFAULT 0,
  `version` smallint(5) NOT NULL DEFAULT 0,
  `xpos` float NOT NULL DEFAULT 0,
  `ypos` float NOT NULL DEFAULT 0,
  `zpos` float NOT NULL DEFAULT 0,
  `heading` float NOT NULL DEFAULT 0,
  `itemid` int(11) NOT NULL DEFAULT 0,
  `charges` smallint(3) unsigned NOT NULL DEFAULT 0,
  `objectname` varchar(32) DEFAULT NULL,
  `type` int(11) NOT NULL DEFAULT 0,
  `icon` int(11) NOT NULL DEFAULT 0,
  `unknown08` mediumint(5) NOT NULL DEFAULT 0,
  `unknown10` mediumint(5) NOT NULL DEFAULT 0,
  `unknown20` int(11) NOT NULL DEFAULT 0,
  `unknown24` int(11) NOT NULL DEFAULT 0,
  `unknown60` int(11) NOT NULL DEFAULT 0,
  `unknown64` int(11) NOT NULL DEFAULT 0,
  `unknown68` int(11) NOT NULL DEFAULT 0,
  `unknown72` int(11) NOT NULL DEFAULT 0,
  `unknown76` int(11) NOT NULL DEFAULT 0,
  `unknown84` int(11) NOT NULL DEFAULT 0,
  `size` float NOT NULL DEFAULT 100,
  `tilt_x` float NOT NULL DEFAULT 0,
  `tilt_y` float NOT NULL DEFAULT 0,
  `display_name` varchar(64) DEFAULT NULL,
  `min_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `max_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `content_flags` varchar(100) DEFAULT NULL,
  `content_flags_disabled` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `zone` (`zoneid`)
) ENGINE=InnoDB AUTO_INCREMENT=228455 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "icon": 9,    "unknown_08": 7,    "charges": 15,    "tilt_x": 0.09635797,    "display_name": "PubVwRIdZODGVlRbLqsSfoioG",    "max_expansion": 30,    "version": 20,    "unknown_84": 68,    "id": 53,    "unknown_20": 0,    "unknown_60": 94,    "itemid": 20,    "type": 9,    "unknown_72": 46,    "size": 0.99690485,    "content_flags_disabled": "qLGSlrwPSDPCSyftScemcDtgr",    "objectname": "otfyxbTDhkrXCTeZRkAqbrjRA",    "unknown_76": 53,    "content_flags": "BGTKmbkfFdRBZkTopepQwoUyr",    "zoneid": 56,    "xpos": 0.7129688,    "ypos": 0.25346908,    "heading": 0.31268588,    "unknown_24": 66,    "unknown_64": 94,    "min_expansion": 85,    "zpos": 0.7439053,    "unknown_10": 22,    "unknown_68": 38,    "tilt_y": 0.40706787}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 8] column is set for unsigned
[26] column is set for unsigned
[27] column is set for unsigned



*/

// Object struct is a row record of the object table in the eqemu database
type Object struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] zoneid                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Zoneid uint32 `gorm:"column:zoneid;type:uint;default:0;" json:"zoneid"`
	//[ 2] version                                        smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Version int32 `gorm:"column:version;type:smallint;default:0;" json:"version"`
	//[ 3] xpos                                           float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Xpos float32 `gorm:"column:xpos;type:float;default:0;" json:"xpos"`
	//[ 4] ypos                                           float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Ypos float32 `gorm:"column:ypos;type:float;default:0;" json:"ypos"`
	//[ 5] zpos                                           float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Zpos float32 `gorm:"column:zpos;type:float;default:0;" json:"zpos"`
	//[ 6] heading                                        float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Heading float32 `gorm:"column:heading;type:float;default:0;" json:"heading"`
	//[ 7] itemid                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Itemid int32 `gorm:"column:itemid;type:int;default:0;" json:"itemid"`
	//[ 8] charges                                        usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	Charges uint32 `gorm:"column:charges;type:usmallint;default:0;" json:"charges"`
	//[ 9] objectname                                     varchar(32)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 32      default: [NULL]
	Objectname null.String `gorm:"column:objectname;type:varchar;size:32;" json:"objectname"`
	//[10] type                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Type int32 `gorm:"column:type;type:int;default:0;" json:"type"`
	//[11] icon                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Icon int32 `gorm:"column:icon;type:int;default:0;" json:"icon"`
	//[12] unknown08                                      mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	Unknown08 int32 `gorm:"column:unknown08;type:mediumint;default:0;" json:"unknown_08"`
	//[13] unknown10                                      mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	Unknown10 int32 `gorm:"column:unknown10;type:mediumint;default:0;" json:"unknown_10"`
	//[14] unknown20                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Unknown20 int32 `gorm:"column:unknown20;type:int;default:0;" json:"unknown_20"`
	//[15] unknown24                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Unknown24 int32 `gorm:"column:unknown24;type:int;default:0;" json:"unknown_24"`
	//[16] unknown60                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Unknown60 int32 `gorm:"column:unknown60;type:int;default:0;" json:"unknown_60"`
	//[17] unknown64                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Unknown64 int32 `gorm:"column:unknown64;type:int;default:0;" json:"unknown_64"`
	//[18] unknown68                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Unknown68 int32 `gorm:"column:unknown68;type:int;default:0;" json:"unknown_68"`
	//[19] unknown72                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Unknown72 int32 `gorm:"column:unknown72;type:int;default:0;" json:"unknown_72"`
	//[20] unknown76                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Unknown76 int32 `gorm:"column:unknown76;type:int;default:0;" json:"unknown_76"`
	//[21] unknown84                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Unknown84 int32 `gorm:"column:unknown84;type:int;default:0;" json:"unknown_84"`
	//[22] size                                           float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [100]
	Size float32 `gorm:"column:size;type:float;default:100;" json:"size"`
	//[23] tilt_x                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	TiltX float32 `gorm:"column:tilt_x;type:float;default:0;" json:"tilt_x"`
	//[24] tilt_y                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	TiltY float32 `gorm:"column:tilt_y;type:float;default:0;" json:"tilt_y"`
	//[25] display_name                                   varchar(64)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 64      default: [NULL]
	DisplayName null.String `gorm:"column:display_name;type:varchar;size:64;" json:"display_name"`
	//[26] min_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MinExpansion uint32 `gorm:"column:min_expansion;type:utinyint;default:0;" json:"min_expansion"`
	//[27] max_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MaxExpansion uint32 `gorm:"column:max_expansion;type:utinyint;default:0;" json:"max_expansion"`
	//[28] content_flags                                  varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlags null.String `gorm:"column:content_flags;type:varchar;size:100;" json:"content_flags"`
	//[29] content_flags_disabled                         varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlagsDisabled null.String `gorm:"column:content_flags_disabled;type:varchar;size:100;" json:"content_flags_disabled"`
}

var objectTableInfo = &TableInfo{
	Name: "object",
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
			Name:               "xpos",
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
			GoFieldName:        "Xpos",
			GoFieldType:        "float32",
			JSONFieldName:      "xpos",
			ProtobufFieldName:  "xpos",
			ProtobufType:       "float",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "ypos",
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
			GoFieldName:        "Ypos",
			GoFieldType:        "float32",
			JSONFieldName:      "ypos",
			ProtobufFieldName:  "ypos",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "zpos",
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
			GoFieldName:        "Zpos",
			GoFieldType:        "float32",
			JSONFieldName:      "zpos",
			ProtobufFieldName:  "zpos",
			ProtobufType:       "float",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "itemid",
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
			GoFieldName:        "Itemid",
			GoFieldType:        "int32",
			JSONFieldName:      "itemid",
			ProtobufFieldName:  "itemid",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "charges",
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
			GoFieldName:        "Charges",
			GoFieldType:        "uint32",
			JSONFieldName:      "charges",
			ProtobufFieldName:  "charges",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "objectname",
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
			GoFieldName:        "Objectname",
			GoFieldType:        "null.String",
			JSONFieldName:      "objectname",
			ProtobufFieldName:  "objectname",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "icon",
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
			GoFieldName:        "Icon",
			GoFieldType:        "int32",
			JSONFieldName:      "icon",
			ProtobufFieldName:  "icon",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "unknown08",
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
			GoFieldName:        "Unknown08",
			GoFieldType:        "int32",
			JSONFieldName:      "unknown_08",
			ProtobufFieldName:  "unknown_08",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "unknown10",
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
			GoFieldName:        "Unknown10",
			GoFieldType:        "int32",
			JSONFieldName:      "unknown_10",
			ProtobufFieldName:  "unknown_10",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "unknown20",
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
			GoFieldName:        "Unknown20",
			GoFieldType:        "int32",
			JSONFieldName:      "unknown_20",
			ProtobufFieldName:  "unknown_20",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "unknown24",
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
			GoFieldName:        "Unknown24",
			GoFieldType:        "int32",
			JSONFieldName:      "unknown_24",
			ProtobufFieldName:  "unknown_24",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "unknown60",
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
			GoFieldName:        "Unknown60",
			GoFieldType:        "int32",
			JSONFieldName:      "unknown_60",
			ProtobufFieldName:  "unknown_60",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "unknown64",
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
			GoFieldName:        "Unknown64",
			GoFieldType:        "int32",
			JSONFieldName:      "unknown_64",
			ProtobufFieldName:  "unknown_64",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "unknown68",
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
			GoFieldName:        "Unknown68",
			GoFieldType:        "int32",
			JSONFieldName:      "unknown_68",
			ProtobufFieldName:  "unknown_68",
			ProtobufType:       "int32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "unknown72",
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
			GoFieldName:        "Unknown72",
			GoFieldType:        "int32",
			JSONFieldName:      "unknown_72",
			ProtobufFieldName:  "unknown_72",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "unknown76",
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
			GoFieldName:        "Unknown76",
			GoFieldType:        "int32",
			JSONFieldName:      "unknown_76",
			ProtobufFieldName:  "unknown_76",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "unknown84",
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
			GoFieldName:        "Unknown84",
			GoFieldType:        "int32",
			JSONFieldName:      "unknown_84",
			ProtobufFieldName:  "unknown_84",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "size",
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
			GoFieldName:        "Size",
			GoFieldType:        "float32",
			JSONFieldName:      "size",
			ProtobufFieldName:  "size",
			ProtobufType:       "float",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "tilt_x",
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
			GoFieldName:        "TiltX",
			GoFieldType:        "float32",
			JSONFieldName:      "tilt_x",
			ProtobufFieldName:  "tilt_x",
			ProtobufType:       "float",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "tilt_y",
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
			GoFieldName:        "TiltY",
			GoFieldType:        "float32",
			JSONFieldName:      "tilt_y",
			ProtobufFieldName:  "tilt_y",
			ProtobufType:       "float",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "display_name",
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
			GoFieldName:        "DisplayName",
			GoFieldType:        "null.String",
			JSONFieldName:      "display_name",
			ProtobufFieldName:  "display_name",
			ProtobufType:       "string",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
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
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
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
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
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
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
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
			ProtobufPos:        30,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *Object) TableName() string {
	return "object"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *Object) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *Object) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *Object) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *Object) TableInfo() *TableInfo {
	return objectTableInfo
}
