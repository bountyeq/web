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


CREATE TABLE `start_zones` (
  `x` float NOT NULL DEFAULT 0,
  `y` float NOT NULL DEFAULT 0,
  `z` float NOT NULL DEFAULT 0,
  `heading` float NOT NULL DEFAULT 0,
  `zone_id` int(4) NOT NULL DEFAULT 0,
  `bind_id` int(4) NOT NULL DEFAULT 0,
  `player_choice` int(2) NOT NULL DEFAULT 0,
  `player_class` int(2) NOT NULL DEFAULT 0,
  `player_deity` int(4) NOT NULL DEFAULT 0,
  `player_race` int(4) NOT NULL DEFAULT 0,
  `start_zone` int(4) NOT NULL DEFAULT 0,
  `bind_x` float NOT NULL DEFAULT 0,
  `bind_y` float NOT NULL DEFAULT 0,
  `bind_z` float NOT NULL DEFAULT 0,
  `select_rank` tinyint(3) unsigned NOT NULL DEFAULT 50,
  `min_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `max_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `content_flags` varchar(100) DEFAULT NULL,
  `content_flags_disabled` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`player_choice`,`player_race`,`player_class`,`player_deity`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "heading": 0.8642705,    "zone_id": 71,    "min_expansion": 92,    "content_flags_disabled": "WyOJQXoBckaOtfHOWvCKemImL",    "bind_x": 0.6569122,    "select_rank": 77,    "max_expansion": 89,    "bind_id": 94,    "player_class": 44,    "player_deity": 19,    "content_flags": "ITumwtyukrEBLXYiNAJVwiCgn",    "start_zone": 41,    "bind_y": 0.36638525,    "bind_z": 0.4914426,    "x": 0.55867743,    "y": 0.23993534,    "z": 0.86166906,    "player_choice": 10,    "player_race": 53}


Comments
-------------------------------------
[14] column is set for unsigned
[15] column is set for unsigned
[16] column is set for unsigned



*/

// StartZones struct is a row record of the start_zones table in the eqemu database
type StartZones struct {
	//[ 0] x                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	X float32 `gorm:"column:x;type:float;default:0;" json:"x"`
	//[ 1] y                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Y float32 `gorm:"column:y;type:float;default:0;" json:"y"`
	//[ 2] z                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Z float32 `gorm:"column:z;type:float;default:0;" json:"z"`
	//[ 3] heading                                        float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Heading float32 `gorm:"column:heading;type:float;default:0;" json:"heading"`
	//[ 4] zone_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ZoneID int32 `gorm:"column:zone_id;type:int;default:0;" json:"zone_id"`
	//[ 5] bind_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	BindID int32 `gorm:"column:bind_id;type:int;default:0;" json:"bind_id"`
	//[ 6] player_choice                                  int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	PlayerChoice int32 `gorm:"primary_key;column:player_choice;type:int;default:0;" json:"player_choice"`
	//[ 7] player_class                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	PlayerClass int32 `gorm:"primary_key;column:player_class;type:int;default:0;" json:"player_class"`
	//[ 8] player_deity                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	PlayerDeity int32 `gorm:"primary_key;column:player_deity;type:int;default:0;" json:"player_deity"`
	//[ 9] player_race                                    int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	PlayerRace int32 `gorm:"primary_key;column:player_race;type:int;default:0;" json:"player_race"`
	//[10] start_zone                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	StartZone int32 `gorm:"column:start_zone;type:int;default:0;" json:"start_zone"`
	//[11] bind_x                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	BindX float32 `gorm:"column:bind_x;type:float;default:0;" json:"bind_x"`
	//[12] bind_y                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	BindY float32 `gorm:"column:bind_y;type:float;default:0;" json:"bind_y"`
	//[13] bind_z                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	BindZ float32 `gorm:"column:bind_z;type:float;default:0;" json:"bind_z"`
	//[14] select_rank                                    utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [50]
	SelectRank uint32 `gorm:"column:select_rank;type:utinyint;default:50;" json:"select_rank"`
	//[15] min_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MinExpansion uint32 `gorm:"column:min_expansion;type:utinyint;default:0;" json:"min_expansion"`
	//[16] max_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MaxExpansion uint32 `gorm:"column:max_expansion;type:utinyint;default:0;" json:"max_expansion"`
	//[17] content_flags                                  varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlags null.String `gorm:"column:content_flags;type:varchar;size:100;" json:"content_flags"`
	//[18] content_flags_disabled                         varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlagsDisabled null.String `gorm:"column:content_flags_disabled;type:varchar;size:100;" json:"content_flags_disabled"`
}

var start_zonesTableInfo = &TableInfo{
	Name: "start_zones",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
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
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "zone_id",
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
			GoFieldName:        "ZoneID",
			GoFieldType:        "int32",
			JSONFieldName:      "zone_id",
			ProtobufFieldName:  "zone_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "bind_id",
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
			GoFieldName:        "BindID",
			GoFieldType:        "int32",
			JSONFieldName:      "bind_id",
			ProtobufFieldName:  "bind_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "player_choice",
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
			GoFieldName:        "PlayerChoice",
			GoFieldType:        "int32",
			JSONFieldName:      "player_choice",
			ProtobufFieldName:  "player_choice",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "player_class",
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
			GoFieldName:        "PlayerClass",
			GoFieldType:        "int32",
			JSONFieldName:      "player_class",
			ProtobufFieldName:  "player_class",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "player_deity",
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
			GoFieldName:        "PlayerDeity",
			GoFieldType:        "int32",
			JSONFieldName:      "player_deity",
			ProtobufFieldName:  "player_deity",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "player_race",
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
			GoFieldName:        "PlayerRace",
			GoFieldType:        "int32",
			JSONFieldName:      "player_race",
			ProtobufFieldName:  "player_race",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "start_zone",
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
			GoFieldName:        "StartZone",
			GoFieldType:        "int32",
			JSONFieldName:      "start_zone",
			ProtobufFieldName:  "start_zone",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "bind_x",
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
			GoFieldName:        "BindX",
			GoFieldType:        "float32",
			JSONFieldName:      "bind_x",
			ProtobufFieldName:  "bind_x",
			ProtobufType:       "float",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "bind_y",
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
			GoFieldName:        "BindY",
			GoFieldType:        "float32",
			JSONFieldName:      "bind_y",
			ProtobufFieldName:  "bind_y",
			ProtobufType:       "float",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "bind_z",
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
			GoFieldName:        "BindZ",
			GoFieldType:        "float32",
			JSONFieldName:      "bind_z",
			ProtobufFieldName:  "bind_z",
			ProtobufType:       "float",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "select_rank",
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
			GoFieldName:        "SelectRank",
			GoFieldType:        "uint32",
			JSONFieldName:      "select_rank",
			ProtobufFieldName:  "select_rank",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
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
			ProtobufPos:        19,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *StartZones) TableName() string {
	return "start_zones"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *StartZones) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *StartZones) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *StartZones) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *StartZones) TableInfo() *TableInfo {
	return start_zonesTableInfo
}
