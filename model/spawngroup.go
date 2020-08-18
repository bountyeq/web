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


CREATE TABLE `spawngroup` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `spawn_limit` tinyint(4) NOT NULL DEFAULT 0,
  `dist` float NOT NULL DEFAULT 0,
  `max_x` float NOT NULL DEFAULT 0,
  `min_x` float NOT NULL DEFAULT 0,
  `max_y` float NOT NULL DEFAULT 0,
  `min_y` float NOT NULL DEFAULT 0,
  `delay` int(11) NOT NULL DEFAULT 45000,
  `mindelay` int(11) NOT NULL DEFAULT 15000,
  `despawn` tinyint(3) NOT NULL DEFAULT 0,
  `despawn_timer` int(11) NOT NULL DEFAULT 100,
  `wp_spawns` tinyint(1) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=287745 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 81,    "spawn_limit": 56,    "dist": 0.043497857,    "max_y": 0.6763093,    "mindelay": 42,    "despawn": 83,    "name": "qZbAUerurcwKxmFOOHucYakkt",    "max_x": 0.33362326,    "min_x": 0.3694564,    "min_y": 0.048789714,    "delay": 41,    "despawn_timer": 94,    "wp_spawns": 82}


Comments
-------------------------------------
[12] column is set for unsigned



*/

// Spawngroup struct is a row record of the spawngroup table in the eqemu database
type Spawngroup struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] name                                           varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: ['']
	Name string `gorm:"column:name;type:varchar;size:50;default:'';" json:"name"`
	//[ 2] spawn_limit                                    tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	SpawnLimit int32 `gorm:"column:spawn_limit;type:tinyint;default:0;" json:"spawn_limit"`
	//[ 3] dist                                           float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Dist float32 `gorm:"column:dist;type:float;default:0;" json:"dist"`
	//[ 4] max_x                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	MaxX float32 `gorm:"column:max_x;type:float;default:0;" json:"max_x"`
	//[ 5] min_x                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	MinX float32 `gorm:"column:min_x;type:float;default:0;" json:"min_x"`
	//[ 6] max_y                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	MaxY float32 `gorm:"column:max_y;type:float;default:0;" json:"max_y"`
	//[ 7] min_y                                          float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	MinY float32 `gorm:"column:min_y;type:float;default:0;" json:"min_y"`
	//[ 8] delay                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [45000]
	Delay int32 `gorm:"column:delay;type:int;default:45000;" json:"delay"`
	//[ 9] mindelay                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [15000]
	Mindelay int32 `gorm:"column:mindelay;type:int;default:15000;" json:"mindelay"`
	//[10] despawn                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Despawn int32 `gorm:"column:despawn;type:tinyint;default:0;" json:"despawn"`
	//[11] despawn_timer                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [100]
	DespawnTimer int32 `gorm:"column:despawn_timer;type:int;default:100;" json:"despawn_timer"`
	//[12] wp_spawns                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	WpSpawns uint32 `gorm:"column:wp_spawns;type:utinyint;default:0;" json:"wp_spawns"`
}

var spawngroupTableInfo = &TableInfo{
	Name: "spawngroup",
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
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "spawn_limit",
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
			GoFieldName:        "SpawnLimit",
			GoFieldType:        "int32",
			JSONFieldName:      "spawn_limit",
			ProtobufFieldName:  "spawn_limit",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "dist",
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
			GoFieldName:        "Dist",
			GoFieldType:        "float32",
			JSONFieldName:      "dist",
			ProtobufFieldName:  "dist",
			ProtobufType:       "float",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			Name:               "delay",
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
			GoFieldName:        "Delay",
			GoFieldType:        "int32",
			JSONFieldName:      "delay",
			ProtobufFieldName:  "delay",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "mindelay",
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
			GoFieldName:        "Mindelay",
			GoFieldType:        "int32",
			JSONFieldName:      "mindelay",
			ProtobufFieldName:  "mindelay",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "despawn",
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
			GoFieldName:        "Despawn",
			GoFieldType:        "int32",
			JSONFieldName:      "despawn",
			ProtobufFieldName:  "despawn",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "despawn_timer",
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
			GoFieldName:        "DespawnTimer",
			GoFieldType:        "int32",
			JSONFieldName:      "despawn_timer",
			ProtobufFieldName:  "despawn_timer",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "wp_spawns",
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
			GoFieldName:        "WpSpawns",
			GoFieldType:        "uint32",
			JSONFieldName:      "wp_spawns",
			ProtobufFieldName:  "wp_spawns",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *Spawngroup) TableName() string {
	return "spawngroup"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *Spawngroup) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *Spawngroup) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *Spawngroup) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *Spawngroup) TableInfo() *TableInfo {
	return spawngroupTableInfo
}
