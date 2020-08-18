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


CREATE TABLE `global_loot` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `description` varchar(255) DEFAULT NULL,
  `loottable_id` int(11) NOT NULL,
  `enabled` tinyint(4) NOT NULL DEFAULT 1,
  `min_level` int(11) NOT NULL DEFAULT 0,
  `max_level` int(11) NOT NULL DEFAULT 0,
  `rare` tinyint(4) DEFAULT NULL,
  `raid` tinyint(4) DEFAULT NULL,
  `race` mediumtext DEFAULT NULL,
  `class` mediumtext DEFAULT NULL,
  `bodytype` mediumtext DEFAULT NULL,
  `zone` mediumtext DEFAULT NULL,
  `hot_zone` tinyint(4) DEFAULT NULL,
  `min_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `max_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `content_flags` varchar(100) DEFAULT NULL,
  `content_flags_disabled` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "content_flags": "mMNKUUVXoZBrDMNcWHYymnrtT",    "content_flags_disabled": "CNOZyddyogoSvILEAksHyUftN",    "bodytype": "xTaPUpEIRpbAxHfPqOgMsrYNo",    "zone": "QNgNbRaOkkHfKaNFHsjoQqFtn",    "max_expansion": 59,    "loottable_id": 39,    "min_level": 50,    "max_level": 3,    "raid": 61,    "class": "yJxxUWYmwGtecgkJIcxFTZcac",    "enabled": 39,    "hot_zone": 41,    "min_expansion": 67,    "id": 33,    "description": "CpsAmHTwfLcUqMHMOnkilROoR",    "rare": 95,    "race": "OxJmbqAqirmWVVhjthaIVAYfi"}


Comments
-------------------------------------
[13] column is set for unsigned
[14] column is set for unsigned



*/

// GlobalLoot struct is a row record of the global_loot table in the eqemu database
type GlobalLoot struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] description                                    varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: [NULL]
	Description null.String `gorm:"column:description;type:varchar;size:255;" json:"description"`
	//[ 2] loottable_id                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LoottableID int32 `gorm:"column:loottable_id;type:int;" json:"loottable_id"`
	//[ 3] enabled                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Enabled int32 `gorm:"column:enabled;type:tinyint;default:1;" json:"enabled"`
	//[ 4] min_level                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MinLevel int32 `gorm:"column:min_level;type:int;default:0;" json:"min_level"`
	//[ 5] max_level                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MaxLevel int32 `gorm:"column:max_level;type:int;default:0;" json:"max_level"`
	//[ 6] rare                                           tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [NULL]
	Rare null.Int `gorm:"column:rare;type:tinyint;" json:"rare"`
	//[ 7] raid                                           tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [NULL]
	Raid null.Int `gorm:"column:raid;type:tinyint;" json:"raid"`
	//[ 8] race                                           text(16777215)       null: true   primary: false  isArray: false  auto: false  col: text            len: 16777215 default: [NULL]
	Race null.String `gorm:"column:race;type:text;size:16777215;" json:"race"`
	//[ 9] class                                          text(16777215)       null: true   primary: false  isArray: false  auto: false  col: text            len: 16777215 default: [NULL]
	Class null.String `gorm:"column:class;type:text;size:16777215;" json:"class"`
	//[10] bodytype                                       text(16777215)       null: true   primary: false  isArray: false  auto: false  col: text            len: 16777215 default: [NULL]
	Bodytype null.String `gorm:"column:bodytype;type:text;size:16777215;" json:"bodytype"`
	//[11] zone                                           text(16777215)       null: true   primary: false  isArray: false  auto: false  col: text            len: 16777215 default: [NULL]
	Zone null.String `gorm:"column:zone;type:text;size:16777215;" json:"zone"`
	//[12] hot_zone                                       tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [NULL]
	HotZone null.Int `gorm:"column:hot_zone;type:tinyint;" json:"hot_zone"`
	//[13] min_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MinExpansion uint32 `gorm:"column:min_expansion;type:utinyint;default:0;" json:"min_expansion"`
	//[14] max_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MaxExpansion uint32 `gorm:"column:max_expansion;type:utinyint;default:0;" json:"max_expansion"`
	//[15] content_flags                                  varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlags null.String `gorm:"column:content_flags;type:varchar;size:100;" json:"content_flags"`
	//[16] content_flags_disabled                         varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlagsDisabled null.String `gorm:"column:content_flags_disabled;type:varchar;size:100;" json:"content_flags_disabled"`
}

var global_lootTableInfo = &TableInfo{
	Name: "global_loot",
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
			Name:               "description",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Description",
			GoFieldType:        "null.String",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "loottable_id",
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
			GoFieldName:        "LoottableID",
			GoFieldType:        "int32",
			JSONFieldName:      "loottable_id",
			ProtobufFieldName:  "loottable_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "enabled",
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
			GoFieldName:        "Enabled",
			GoFieldType:        "int32",
			JSONFieldName:      "enabled",
			ProtobufFieldName:  "enabled",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "min_level",
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
			GoFieldName:        "MinLevel",
			GoFieldType:        "int32",
			JSONFieldName:      "min_level",
			ProtobufFieldName:  "min_level",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "max_level",
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
			GoFieldName:        "MaxLevel",
			GoFieldType:        "int32",
			JSONFieldName:      "max_level",
			ProtobufFieldName:  "max_level",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "rare",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Rare",
			GoFieldType:        "null.Int",
			JSONFieldName:      "rare",
			ProtobufFieldName:  "rare",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "raid",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Raid",
			GoFieldType:        "null.Int",
			JSONFieldName:      "raid",
			ProtobufFieldName:  "raid",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "race",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(16777215)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       16777215,
			GoFieldName:        "Race",
			GoFieldType:        "null.String",
			JSONFieldName:      "race",
			ProtobufFieldName:  "race",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "class",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(16777215)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       16777215,
			GoFieldName:        "Class",
			GoFieldType:        "null.String",
			JSONFieldName:      "class",
			ProtobufFieldName:  "class",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "bodytype",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(16777215)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       16777215,
			GoFieldName:        "Bodytype",
			GoFieldType:        "null.String",
			JSONFieldName:      "bodytype",
			ProtobufFieldName:  "bodytype",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "zone",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(16777215)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       16777215,
			GoFieldName:        "Zone",
			GoFieldType:        "null.String",
			JSONFieldName:      "zone",
			ProtobufFieldName:  "zone",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "hot_zone",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "HotZone",
			GoFieldType:        "null.Int",
			JSONFieldName:      "hot_zone",
			ProtobufFieldName:  "hot_zone",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *GlobalLoot) TableName() string {
	return "global_loot"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GlobalLoot) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GlobalLoot) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GlobalLoot) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GlobalLoot) TableInfo() *TableInfo {
	return global_lootTableInfo
}
