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


CREATE TABLE `auras` (
  `type` int(10) NOT NULL,
  `npc_type` int(10) NOT NULL,
  `name` varchar(64) NOT NULL,
  `spell_id` int(10) NOT NULL,
  `distance` int(10) NOT NULL DEFAULT 60,
  `aura_type` int(10) NOT NULL DEFAULT 1,
  `spawn_type` int(10) NOT NULL DEFAULT 0,
  `movement` int(10) NOT NULL DEFAULT 0,
  `duration` int(10) NOT NULL DEFAULT 5400,
  `icon` int(10) NOT NULL DEFAULT -1,
  `cast_time` int(10) NOT NULL DEFAULT 0,
  PRIMARY KEY (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "type": 84,    "npc_type": 25,    "name": "wBYvAwuqdXYsYMKopZAbSjSMV",    "spell_id": 82,    "distance": 47,    "duration": 17,    "aura_type": 16,    "spawn_type": 2,    "movement": 68,    "icon": 48,    "cast_time": 42}



*/

// Auras struct is a row record of the auras table in the eqemu database
type Auras struct {
	//[ 0] type                                           int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	Type int32 `gorm:"primary_key;column:type;type:int;" json:"type"`
	//[ 1] npc_type                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	NpcType int32 `gorm:"column:npc_type;type:int;" json:"npc_type"`
	//[ 2] name                                           varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Name string `gorm:"column:name;type:varchar;size:64;" json:"name"`
	//[ 3] spell_id                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SpellID int32 `gorm:"column:spell_id;type:int;" json:"spell_id"`
	//[ 4] distance                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [60]
	Distance int32 `gorm:"column:distance;type:int;default:60;" json:"distance"`
	//[ 5] aura_type                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	AuraType int32 `gorm:"column:aura_type;type:int;default:1;" json:"aura_type"`
	//[ 6] spawn_type                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SpawnType int32 `gorm:"column:spawn_type;type:int;default:0;" json:"spawn_type"`
	//[ 7] movement                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Movement int32 `gorm:"column:movement;type:int;default:0;" json:"movement"`
	//[ 8] duration                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [5400]
	Duration int32 `gorm:"column:duration;type:int;default:5400;" json:"duration"`
	//[ 9] icon                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	Icon int32 `gorm:"column:icon;type:int;default:-1;" json:"icon"`
	//[10] cast_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CastTime int32 `gorm:"column:cast_time;type:int;default:0;" json:"cast_time"`
}

var aurasTableInfo = &TableInfo{
	Name: "auras",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "npc_type",
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
			GoFieldName:        "NpcType",
			GoFieldType:        "int32",
			JSONFieldName:      "npc_type",
			ProtobufFieldName:  "npc_type",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "spell_id",
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
			GoFieldName:        "SpellID",
			GoFieldType:        "int32",
			JSONFieldName:      "spell_id",
			ProtobufFieldName:  "spell_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "distance",
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
			GoFieldName:        "Distance",
			GoFieldType:        "int32",
			JSONFieldName:      "distance",
			ProtobufFieldName:  "distance",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "aura_type",
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
			GoFieldName:        "AuraType",
			GoFieldType:        "int32",
			JSONFieldName:      "aura_type",
			ProtobufFieldName:  "aura_type",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "spawn_type",
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
			GoFieldName:        "SpawnType",
			GoFieldType:        "int32",
			JSONFieldName:      "spawn_type",
			ProtobufFieldName:  "spawn_type",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "movement",
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
			GoFieldName:        "Movement",
			GoFieldType:        "int32",
			JSONFieldName:      "movement",
			ProtobufFieldName:  "movement",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "duration",
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
			GoFieldName:        "Duration",
			GoFieldType:        "int32",
			JSONFieldName:      "duration",
			ProtobufFieldName:  "duration",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "cast_time",
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
			GoFieldName:        "CastTime",
			GoFieldType:        "int32",
			JSONFieldName:      "cast_time",
			ProtobufFieldName:  "cast_time",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *Auras) TableName() string {
	return "auras"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Auras) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Auras) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Auras) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Auras) TableInfo() *TableInfo {
	return aurasTableInfo
}
