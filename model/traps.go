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


CREATE TABLE `traps` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `zone` varchar(16) NOT NULL DEFAULT '',
  `version` smallint(5) unsigned NOT NULL DEFAULT 0,
  `x` int(11) NOT NULL DEFAULT 0,
  `y` int(11) NOT NULL DEFAULT 0,
  `z` int(11) NOT NULL DEFAULT 0,
  `chance` tinyint(4) NOT NULL DEFAULT 0,
  `maxzdiff` float NOT NULL DEFAULT 0,
  `radius` float NOT NULL DEFAULT 0,
  `effect` int(11) NOT NULL DEFAULT 0,
  `effectvalue` int(11) NOT NULL DEFAULT 0,
  `effectvalue2` int(11) NOT NULL DEFAULT 0,
  `message` varchar(200) NOT NULL DEFAULT '',
  `skill` int(11) NOT NULL DEFAULT 0,
  `level` mediumint(4) unsigned NOT NULL DEFAULT 1,
  `respawn_time` int(11) unsigned NOT NULL DEFAULT 60,
  `respawn_var` int(11) unsigned NOT NULL DEFAULT 0,
  `triggered_number` tinyint(4) NOT NULL DEFAULT 0,
  `group` tinyint(4) NOT NULL DEFAULT 0,
  `despawn_when_triggered` tinyint(4) NOT NULL DEFAULT 0,
  `undetectable` tinyint(4) NOT NULL DEFAULT 0,
  `min_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `max_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `content_flags` varchar(100) DEFAULT NULL,
  `content_flags_disabled` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `zone` (`zone`)
) ENGINE=InnoDB AUTO_INCREMENT=855 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "level": 50,    "undetectable": 15,    "min_expansion": 64,    "max_expansion": 47,    "content_flags_disabled": "NKJgsuaiVckCrOAHKFvxKjDAA",    "maxzdiff": 0.2651147,    "message": "rMcRVnHEWZhcKHxTsdTLFGBUZ",    "skill": 35,    "respawn_time": 5,    "group": 82,    "content_flags": "JJMCEbrDYXxmKMiIOlnKnnQOv",    "z": 47,    "effectvalue": 91,    "effectvalue_2": 98,    "despawn_when_triggered": 22,    "version": 76,    "respawn_var": 6,    "triggered_number": 45,    "y": 60,    "chance": 23,    "radius": 0.6388013,    "effect": 91,    "id": 21,    "zone": "TMgmpyjsouHwmLImuGdlREZGT",    "x": 71}


Comments
-------------------------------------
[ 2] column is set for unsigned
[14] column is set for unsigned
[15] column is set for unsigned
[16] column is set for unsigned
[21] column is set for unsigned
[22] column is set for unsigned



*/

// Traps struct is a row record of the traps table in the eqemu database
type Traps struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] zone                                           varchar(16)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 16      default: ['']
	Zone string `gorm:"column:zone;type:varchar;size:16;default:'';" json:"zone"`
	//[ 2] version                                        usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	Version uint32 `gorm:"column:version;type:usmallint;default:0;" json:"version"`
	//[ 3] x                                              int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	X int32 `gorm:"column:x;type:int;default:0;" json:"x"`
	//[ 4] y                                              int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Y int32 `gorm:"column:y;type:int;default:0;" json:"y"`
	//[ 5] z                                              int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Z int32 `gorm:"column:z;type:int;default:0;" json:"z"`
	//[ 6] chance                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Chance int32 `gorm:"column:chance;type:tinyint;default:0;" json:"chance"`
	//[ 7] maxzdiff                                       float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Maxzdiff float32 `gorm:"column:maxzdiff;type:float;default:0;" json:"maxzdiff"`
	//[ 8] radius                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Radius float32 `gorm:"column:radius;type:float;default:0;" json:"radius"`
	//[ 9] effect                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Effect int32 `gorm:"column:effect;type:int;default:0;" json:"effect"`
	//[10] effectvalue                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Effectvalue int32 `gorm:"column:effectvalue;type:int;default:0;" json:"effectvalue"`
	//[11] effectvalue2                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Effectvalue2 int32 `gorm:"column:effectvalue2;type:int;default:0;" json:"effectvalue_2"`
	//[12] message                                        varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: ['']
	Message string `gorm:"column:message;type:varchar;size:200;default:'';" json:"message"`
	//[13] skill                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Skill int32 `gorm:"column:skill;type:int;default:0;" json:"skill"`
	//[14] level                                          umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [1]
	Level uint32 `gorm:"column:level;type:umediumint;default:1;" json:"level"`
	//[15] respawn_time                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [60]
	RespawnTime uint32 `gorm:"column:respawn_time;type:uint;default:60;" json:"respawn_time"`
	//[16] respawn_var                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	RespawnVar uint32 `gorm:"column:respawn_var;type:uint;default:0;" json:"respawn_var"`
	//[17] triggered_number                               tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	TriggeredNumber int32 `gorm:"column:triggered_number;type:tinyint;default:0;" json:"triggered_number"`
	//[18] group                                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Group int32 `gorm:"column:group;type:tinyint;default:0;" json:"group"`
	//[19] despawn_when_triggered                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	DespawnWhenTriggered int32 `gorm:"column:despawn_when_triggered;type:tinyint;default:0;" json:"despawn_when_triggered"`
	//[20] undetectable                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Undetectable int32 `gorm:"column:undetectable;type:tinyint;default:0;" json:"undetectable"`
	//[21] min_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MinExpansion uint32 `gorm:"column:min_expansion;type:utinyint;default:0;" json:"min_expansion"`
	//[22] max_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MaxExpansion uint32 `gorm:"column:max_expansion;type:utinyint;default:0;" json:"max_expansion"`
	//[23] content_flags                                  varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlags null.String `gorm:"column:content_flags;type:varchar;size:100;" json:"content_flags"`
	//[24] content_flags_disabled                         varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlagsDisabled null.String `gorm:"column:content_flags_disabled;type:varchar;size:100;" json:"content_flags_disabled"`
}

var trapsTableInfo = &TableInfo{
	Name: "traps",
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
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(16)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       16,
			GoFieldName:        "Zone",
			GoFieldType:        "string",
			JSONFieldName:      "zone",
			ProtobufFieldName:  "zone",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "version",
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
			GoFieldName:        "Version",
			GoFieldType:        "uint32",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "x",
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
			GoFieldName:        "X",
			GoFieldType:        "int32",
			JSONFieldName:      "x",
			ProtobufFieldName:  "x",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "y",
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
			GoFieldName:        "Y",
			GoFieldType:        "int32",
			JSONFieldName:      "y",
			ProtobufFieldName:  "y",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "z",
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
			GoFieldName:        "Z",
			GoFieldType:        "int32",
			JSONFieldName:      "z",
			ProtobufFieldName:  "z",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "chance",
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
			GoFieldName:        "Chance",
			GoFieldType:        "int32",
			JSONFieldName:      "chance",
			ProtobufFieldName:  "chance",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "maxzdiff",
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
			GoFieldName:        "Maxzdiff",
			GoFieldType:        "float32",
			JSONFieldName:      "maxzdiff",
			ProtobufFieldName:  "maxzdiff",
			ProtobufType:       "float",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "radius",
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
			GoFieldName:        "Radius",
			GoFieldType:        "float32",
			JSONFieldName:      "radius",
			ProtobufFieldName:  "radius",
			ProtobufType:       "float",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "effect",
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
			GoFieldName:        "Effect",
			GoFieldType:        "int32",
			JSONFieldName:      "effect",
			ProtobufFieldName:  "effect",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "effectvalue",
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
			GoFieldName:        "Effectvalue",
			GoFieldType:        "int32",
			JSONFieldName:      "effectvalue",
			ProtobufFieldName:  "effectvalue",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "effectvalue2",
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
			GoFieldName:        "Effectvalue2",
			GoFieldType:        "int32",
			JSONFieldName:      "effectvalue_2",
			ProtobufFieldName:  "effectvalue_2",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "message",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "Message",
			GoFieldType:        "string",
			JSONFieldName:      "message",
			ProtobufFieldName:  "message",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "skill",
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
			GoFieldName:        "Skill",
			GoFieldType:        "int32",
			JSONFieldName:      "skill",
			ProtobufFieldName:  "skill",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "level",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "umediumint",
			DatabaseTypePretty: "umediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "umediumint",
			ColumnLength:       -1,
			GoFieldName:        "Level",
			GoFieldType:        "uint32",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "respawn_time",
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
			GoFieldName:        "RespawnTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "respawn_time",
			ProtobufFieldName:  "respawn_time",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "respawn_var",
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
			GoFieldName:        "RespawnVar",
			GoFieldType:        "uint32",
			JSONFieldName:      "respawn_var",
			ProtobufFieldName:  "respawn_var",
			ProtobufType:       "uint32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "triggered_number",
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
			GoFieldName:        "TriggeredNumber",
			GoFieldType:        "int32",
			JSONFieldName:      "triggered_number",
			ProtobufFieldName:  "triggered_number",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "group",
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
			GoFieldName:        "Group",
			GoFieldType:        "int32",
			JSONFieldName:      "group",
			ProtobufFieldName:  "group",
			ProtobufType:       "int32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "despawn_when_triggered",
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
			GoFieldName:        "DespawnWhenTriggered",
			GoFieldType:        "int32",
			JSONFieldName:      "despawn_when_triggered",
			ProtobufFieldName:  "despawn_when_triggered",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "undetectable",
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
			GoFieldName:        "Undetectable",
			GoFieldType:        "int32",
			JSONFieldName:      "undetectable",
			ProtobufFieldName:  "undetectable",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
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
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
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
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
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
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
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
			ProtobufPos:        25,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *Traps) TableName() string {
	return "traps"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *Traps) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *Traps) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *Traps) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *Traps) TableInfo() *TableInfo {
	return trapsTableInfo
}
