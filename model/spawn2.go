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


CREATE TABLE `spawn2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `spawngroupID` int(11) NOT NULL DEFAULT 0,
  `zone` varchar(32) DEFAULT NULL,
  `version` smallint(5) NOT NULL DEFAULT 0,
  `x` float(14,6) NOT NULL DEFAULT 0.000000,
  `y` float(14,6) NOT NULL DEFAULT 0.000000,
  `z` float(14,6) NOT NULL DEFAULT 0.000000,
  `heading` float(14,6) NOT NULL DEFAULT 0.000000,
  `respawntime` int(11) NOT NULL DEFAULT 0,
  `variance` int(11) NOT NULL DEFAULT 0,
  `pathgrid` int(10) NOT NULL DEFAULT 0,
  `_condition` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `cond_value` mediumint(9) NOT NULL DEFAULT 1,
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT 1,
  `animation` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `min_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `max_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `content_flags` varchar(100) DEFAULT NULL,
  `content_flags_disabled` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ZoneGroup` (`zone`),
  KEY `spawn2_spawngroupid_idx` (`spawngroupID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=263866 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 34,    "z": 0.9249503,    "pathgrid": 90,    "content_flags": "PgUPIiFeADSSaguSxxIybXvyE",    "min_expansion": 62,    "max_expansion": 41,    "zone": "whVXJQhZpGbmTtUXdOrDoCIXG",    "respawntime": 10,    "_condition": 46,    "animation": 39,    "cond_value": 66,    "enabled": 45,    "content_flags_disabled": "XTblOxphNlhbdYbCqKRTDxlfV",    "heading": 0.6679745,    "variance": 37,    "spawngroup_id": 32,    "version": 70,    "x": 0.50775397,    "y": 0.5278718}


Comments
-------------------------------------
[11] column is set for unsigned
[13] column is set for unsigned
[14] column is set for unsigned
[15] column is set for unsigned
[16] column is set for unsigned



*/

// Spawn2 struct is a row record of the spawn2 table in the eqemu database
type Spawn2 struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] spawngroupID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SpawngroupID int32 `gorm:"column:spawngroupID;type:int;default:0;" json:"spawngroup_id"`
	//[ 2] zone                                           varchar(32)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 32      default: [NULL]
	Zone null.String `gorm:"column:zone;type:varchar;size:32;" json:"zone"`
	//[ 3] version                                        smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Version int32 `gorm:"column:version;type:smallint;default:0;" json:"version"`
	//[ 4] x                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0.000000]
	X float32 `gorm:"column:x;type:float;default:0.000000;" json:"x"`
	//[ 5] y                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0.000000]
	Y float32 `gorm:"column:y;type:float;default:0.000000;" json:"y"`
	//[ 6] z                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0.000000]
	Z float32 `gorm:"column:z;type:float;default:0.000000;" json:"z"`
	//[ 7] heading                                        float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0.000000]
	Heading float32 `gorm:"column:heading;type:float;default:0.000000;" json:"heading"`
	//[ 8] respawntime                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Respawntime int32 `gorm:"column:respawntime;type:int;default:0;" json:"respawntime"`
	//[ 9] variance                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Variance int32 `gorm:"column:variance;type:int;default:0;" json:"variance"`
	//[10] pathgrid                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Pathgrid int32 `gorm:"column:pathgrid;type:int;default:0;" json:"pathgrid"`
	//[11] _condition                                     umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Condition uint32 `gorm:"column:_condition;type:umediumint;default:0;" json:"_condition"`
	//[12] cond_value                                     mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [1]
	CondValue int32 `gorm:"column:cond_value;type:mediumint;default:1;" json:"cond_value"`
	//[13] enabled                                        utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	Enabled uint32 `gorm:"column:enabled;type:utinyint;default:1;" json:"enabled"`
	//[14] animation                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Animation uint32 `gorm:"column:animation;type:utinyint;default:0;" json:"animation"`
	//[15] min_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MinExpansion uint32 `gorm:"column:min_expansion;type:utinyint;default:0;" json:"min_expansion"`
	//[16] max_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MaxExpansion uint32 `gorm:"column:max_expansion;type:utinyint;default:0;" json:"max_expansion"`
	//[17] content_flags                                  varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlags null.String `gorm:"column:content_flags;type:varchar;size:100;" json:"content_flags"`
	//[18] content_flags_disabled                         varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlagsDisabled null.String `gorm:"column:content_flags_disabled;type:varchar;size:100;" json:"content_flags_disabled"`
}

var spawn2TableInfo = &TableInfo{
	Name: "spawn2",
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
			Name:               "spawngroupID",
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
			GoFieldName:        "SpawngroupID",
			GoFieldType:        "int32",
			JSONFieldName:      "spawngroup_id",
			ProtobufFieldName:  "spawngroup_id",
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			Name:               "respawntime",
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
			GoFieldName:        "Respawntime",
			GoFieldType:        "int32",
			JSONFieldName:      "respawntime",
			ProtobufFieldName:  "respawntime",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "variance",
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
			GoFieldName:        "Variance",
			GoFieldType:        "int32",
			JSONFieldName:      "variance",
			ProtobufFieldName:  "variance",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "pathgrid",
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
			GoFieldName:        "Pathgrid",
			GoFieldType:        "int32",
			JSONFieldName:      "pathgrid",
			ProtobufFieldName:  "pathgrid",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "_condition",
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
			GoFieldName:        "Condition",
			GoFieldType:        "uint32",
			JSONFieldName:      "_condition",
			ProtobufFieldName:  "_condition",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "cond_value",
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
			GoFieldName:        "CondValue",
			GoFieldType:        "int32",
			JSONFieldName:      "cond_value",
			ProtobufFieldName:  "cond_value",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "enabled",
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
			GoFieldName:        "Enabled",
			GoFieldType:        "uint32",
			JSONFieldName:      "enabled",
			ProtobufFieldName:  "enabled",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "animation",
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
			GoFieldName:        "Animation",
			GoFieldType:        "uint32",
			JSONFieldName:      "animation",
			ProtobufFieldName:  "animation",
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
func (s *Spawn2) TableName() string {
	return "spawn2"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *Spawn2) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *Spawn2) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *Spawn2) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *Spawn2) TableInfo() *TableInfo {
	return spawn2TableInfo
}
