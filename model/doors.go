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


CREATE TABLE `doors` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `doorid` smallint(4) NOT NULL DEFAULT 0,
  `zone` varchar(32) DEFAULT NULL,
  `version` smallint(5) NOT NULL DEFAULT 0,
  `name` varchar(32) NOT NULL DEFAULT '',
  `pos_y` float NOT NULL DEFAULT 0,
  `pos_x` float NOT NULL DEFAULT 0,
  `pos_z` float NOT NULL DEFAULT 0,
  `heading` float NOT NULL DEFAULT 0,
  `opentype` smallint(4) NOT NULL DEFAULT 0,
  `guild` smallint(4) NOT NULL DEFAULT 0,
  `lockpick` smallint(4) NOT NULL DEFAULT 0,
  `keyitem` int(11) NOT NULL DEFAULT 0,
  `nokeyring` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `triggerdoor` smallint(4) NOT NULL DEFAULT 0,
  `triggertype` smallint(4) NOT NULL DEFAULT 0,
  `disable_timer` tinyint(2) NOT NULL DEFAULT 0,
  `doorisopen` smallint(4) NOT NULL DEFAULT 0,
  `door_param` int(4) NOT NULL DEFAULT 0,
  `dest_zone` varchar(32) DEFAULT 'NONE',
  `dest_instance` int(10) unsigned NOT NULL DEFAULT 0,
  `dest_x` float DEFAULT 0,
  `dest_y` float DEFAULT 0,
  `dest_z` float DEFAULT 0,
  `dest_heading` float DEFAULT 0,
  `invert_state` int(11) DEFAULT 0,
  `incline` int(11) DEFAULT 0,
  `size` smallint(5) unsigned NOT NULL DEFAULT 100,
  `buffer` float DEFAULT 0,
  `client_version_mask` int(10) unsigned NOT NULL DEFAULT 4294967295,
  `is_ldon_door` smallint(6) NOT NULL DEFAULT 0,
  `min_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `max_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `content_flags` varchar(100) DEFAULT NULL,
  `content_flags_disabled` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `DoorIndex` (`zone`,`doorid`,`version`)
) ENGINE=InnoDB AUTO_INCREMENT=40216 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "name": "HyrZjBiVtXIBQQNioOZKCqiiC",    "version": 89,    "pos_x": 0.23732185,    "heading": 0.67723805,    "lockpick": 56,    "incline": 8,    "door_param": 0,    "dest_x": 0.09187742,    "dest_heading": 0.8566821,    "size": 36,    "id": 6,    "pos_z": 0.47029352,    "guild": 78,    "nokeyring": 36,    "content_flags": "LijwqybtNjDBcRoEFqMaPABbX",    "keyitem": 26,    "triggerdoor": 51,    "dest_zone": "NpNCFoaoTyhimVPfRXcfMSePL",    "dest_z": 0.20060411,    "max_expansion": 38,    "pos_y": 0.204131,    "opentype": 1,    "triggertype": 58,    "doorisopen": 37,    "dest_instance": 39,    "is_ldon_door": 61,    "min_expansion": 13,    "doorid": 95,    "zone": "AvlbQkwPHkDiqjYgqcTVeYGYx",    "disable_timer": 29,    "dest_y": 0.61544424,    "client_version_mask": 30,    "invert_state": 70,    "buffer": 0.57940185,    "content_flags_disabled": "IJWTdQLKwbOCfTTdqdDAoGVpO"}


Comments
-------------------------------------
[13] column is set for unsigned
[20] column is set for unsigned
[27] column is set for unsigned
[29] column is set for unsigned
[31] column is set for unsigned
[32] column is set for unsigned



*/

// Doors struct is a row record of the doors table in the eqemu database
type Doors struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] doorid                                         smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Doorid int32 `gorm:"column:doorid;type:smallint;default:0;" json:"doorid"`
	//[ 2] zone                                           varchar(32)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 32      default: [NULL]
	Zone null.String `gorm:"column:zone;type:varchar;size:32;" json:"zone"`
	//[ 3] version                                        smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Version int32 `gorm:"column:version;type:smallint;default:0;" json:"version"`
	//[ 4] name                                           varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['']
	Name string `gorm:"column:name;type:varchar;size:32;default:'';" json:"name"`
	//[ 5] pos_y                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	PosY float32 `gorm:"column:pos_y;type:float;default:0;" json:"pos_y"`
	//[ 6] pos_x                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	PosX float32 `gorm:"column:pos_x;type:float;default:0;" json:"pos_x"`
	//[ 7] pos_z                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	PosZ float32 `gorm:"column:pos_z;type:float;default:0;" json:"pos_z"`
	//[ 8] heading                                        float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Heading float32 `gorm:"column:heading;type:float;default:0;" json:"heading"`
	//[ 9] opentype                                       smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Opentype int32 `gorm:"column:opentype;type:smallint;default:0;" json:"opentype"`
	//[10] guild                                          smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Guild int32 `gorm:"column:guild;type:smallint;default:0;" json:"guild"`
	//[11] lockpick                                       smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Lockpick int32 `gorm:"column:lockpick;type:smallint;default:0;" json:"lockpick"`
	//[12] keyitem                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Keyitem int32 `gorm:"column:keyitem;type:int;default:0;" json:"keyitem"`
	//[13] nokeyring                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Nokeyring uint32 `gorm:"column:nokeyring;type:utinyint;default:0;" json:"nokeyring"`
	//[14] triggerdoor                                    smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Triggerdoor int32 `gorm:"column:triggerdoor;type:smallint;default:0;" json:"triggerdoor"`
	//[15] triggertype                                    smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Triggertype int32 `gorm:"column:triggertype;type:smallint;default:0;" json:"triggertype"`
	//[16] disable_timer                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	DisableTimer int32 `gorm:"column:disable_timer;type:tinyint;default:0;" json:"disable_timer"`
	//[17] doorisopen                                     smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Doorisopen int32 `gorm:"column:doorisopen;type:smallint;default:0;" json:"doorisopen"`
	//[18] door_param                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DoorParam int32 `gorm:"column:door_param;type:int;default:0;" json:"door_param"`
	//[19] dest_zone                                      varchar(32)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['NONE']
	DestZone null.String `gorm:"column:dest_zone;type:varchar;size:32;default:'NONE';" json:"dest_zone"`
	//[20] dest_instance                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	DestInstance uint32 `gorm:"column:dest_instance;type:uint;default:0;" json:"dest_instance"`
	//[21] dest_x                                         float                null: true   primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	DestX null.Float `gorm:"column:dest_x;type:float;default:0;" json:"dest_x"`
	//[22] dest_y                                         float                null: true   primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	DestY null.Float `gorm:"column:dest_y;type:float;default:0;" json:"dest_y"`
	//[23] dest_z                                         float                null: true   primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	DestZ null.Float `gorm:"column:dest_z;type:float;default:0;" json:"dest_z"`
	//[24] dest_heading                                   float                null: true   primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	DestHeading null.Float `gorm:"column:dest_heading;type:float;default:0;" json:"dest_heading"`
	//[25] invert_state                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	InvertState null.Int `gorm:"column:invert_state;type:int;default:0;" json:"invert_state"`
	//[26] incline                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Incline null.Int `gorm:"column:incline;type:int;default:0;" json:"incline"`
	//[27] size                                           usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [100]
	Size uint32 `gorm:"column:size;type:usmallint;default:100;" json:"size"`
	//[28] buffer                                         float                null: true   primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Buffer null.Float `gorm:"column:buffer;type:float;default:0;" json:"buffer"`
	//[29] client_version_mask                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [4294967295]
	ClientVersionMask uint32 `gorm:"column:client_version_mask;type:uint;default:4294967295;" json:"client_version_mask"`
	//[30] is_ldon_door                                   smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	IsLdonDoor int32 `gorm:"column:is_ldon_door;type:smallint;default:0;" json:"is_ldon_door"`
	//[31] min_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MinExpansion uint32 `gorm:"column:min_expansion;type:utinyint;default:0;" json:"min_expansion"`
	//[32] max_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MaxExpansion uint32 `gorm:"column:max_expansion;type:utinyint;default:0;" json:"max_expansion"`
	//[33] content_flags                                  varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlags null.String `gorm:"column:content_flags;type:varchar;size:100;" json:"content_flags"`
	//[34] content_flags_disabled                         varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlagsDisabled null.String `gorm:"column:content_flags_disabled;type:varchar;size:100;" json:"content_flags_disabled"`
}

var doorsTableInfo = &TableInfo{
	Name: "doors",
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
			Name:               "doorid",
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
			GoFieldName:        "Doorid",
			GoFieldType:        "int32",
			JSONFieldName:      "doorid",
			ProtobufFieldName:  "doorid",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "pos_y",
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
			GoFieldName:        "PosY",
			GoFieldType:        "float32",
			JSONFieldName:      "pos_y",
			ProtobufFieldName:  "pos_y",
			ProtobufType:       "float",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "pos_x",
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
			GoFieldName:        "PosX",
			GoFieldType:        "float32",
			JSONFieldName:      "pos_x",
			ProtobufFieldName:  "pos_x",
			ProtobufType:       "float",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "pos_z",
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
			GoFieldName:        "PosZ",
			GoFieldType:        "float32",
			JSONFieldName:      "pos_z",
			ProtobufFieldName:  "pos_z",
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
			Name:               "opentype",
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
			GoFieldName:        "Opentype",
			GoFieldType:        "int32",
			JSONFieldName:      "opentype",
			ProtobufFieldName:  "opentype",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "guild",
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
			GoFieldName:        "Guild",
			GoFieldType:        "int32",
			JSONFieldName:      "guild",
			ProtobufFieldName:  "guild",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "lockpick",
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
			GoFieldName:        "Lockpick",
			GoFieldType:        "int32",
			JSONFieldName:      "lockpick",
			ProtobufFieldName:  "lockpick",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "keyitem",
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
			GoFieldName:        "Keyitem",
			GoFieldType:        "int32",
			JSONFieldName:      "keyitem",
			ProtobufFieldName:  "keyitem",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "nokeyring",
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
			GoFieldName:        "Nokeyring",
			GoFieldType:        "uint32",
			JSONFieldName:      "nokeyring",
			ProtobufFieldName:  "nokeyring",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "triggerdoor",
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
			GoFieldName:        "Triggerdoor",
			GoFieldType:        "int32",
			JSONFieldName:      "triggerdoor",
			ProtobufFieldName:  "triggerdoor",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "triggertype",
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
			GoFieldName:        "Triggertype",
			GoFieldType:        "int32",
			JSONFieldName:      "triggertype",
			ProtobufFieldName:  "triggertype",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "disable_timer",
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
			GoFieldName:        "DisableTimer",
			GoFieldType:        "int32",
			JSONFieldName:      "disable_timer",
			ProtobufFieldName:  "disable_timer",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "doorisopen",
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
			GoFieldName:        "Doorisopen",
			GoFieldType:        "int32",
			JSONFieldName:      "doorisopen",
			ProtobufFieldName:  "doorisopen",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "door_param",
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
			GoFieldName:        "DoorParam",
			GoFieldType:        "int32",
			JSONFieldName:      "door_param",
			ProtobufFieldName:  "door_param",
			ProtobufType:       "int32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "dest_zone",
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
			GoFieldName:        "DestZone",
			GoFieldType:        "null.String",
			JSONFieldName:      "dest_zone",
			ProtobufFieldName:  "dest_zone",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "dest_instance",
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
			GoFieldName:        "DestInstance",
			GoFieldType:        "uint32",
			JSONFieldName:      "dest_instance",
			ProtobufFieldName:  "dest_instance",
			ProtobufType:       "uint32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "dest_x",
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
			GoFieldName:        "DestX",
			GoFieldType:        "null.Float",
			JSONFieldName:      "dest_x",
			ProtobufFieldName:  "dest_x",
			ProtobufType:       "float",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "dest_y",
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
			GoFieldName:        "DestY",
			GoFieldType:        "null.Float",
			JSONFieldName:      "dest_y",
			ProtobufFieldName:  "dest_y",
			ProtobufType:       "float",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "dest_z",
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
			GoFieldName:        "DestZ",
			GoFieldType:        "null.Float",
			JSONFieldName:      "dest_z",
			ProtobufFieldName:  "dest_z",
			ProtobufType:       "float",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "dest_heading",
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
			GoFieldName:        "DestHeading",
			GoFieldType:        "null.Float",
			JSONFieldName:      "dest_heading",
			ProtobufFieldName:  "dest_heading",
			ProtobufType:       "float",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "invert_state",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "InvertState",
			GoFieldType:        "null.Int",
			JSONFieldName:      "invert_state",
			ProtobufFieldName:  "invert_state",
			ProtobufType:       "int32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "incline",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Incline",
			GoFieldType:        "null.Int",
			JSONFieldName:      "incline",
			ProtobufFieldName:  "incline",
			ProtobufType:       "int32",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "size",
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
			GoFieldName:        "Size",
			GoFieldType:        "uint32",
			JSONFieldName:      "size",
			ProtobufFieldName:  "size",
			ProtobufType:       "uint32",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
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
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
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
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "is_ldon_door",
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
			GoFieldName:        "IsLdonDoor",
			GoFieldType:        "int32",
			JSONFieldName:      "is_ldon_door",
			ProtobufFieldName:  "is_ldon_door",
			ProtobufType:       "int32",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
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
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
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
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
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
			ProtobufPos:        34,
		},

		&ColumnInfo{
			Index:              34,
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
			ProtobufPos:        35,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *Doors) TableName() string {
	return "doors"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *Doors) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *Doors) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *Doors) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *Doors) TableInfo() *TableInfo {
	return doorsTableInfo
}
