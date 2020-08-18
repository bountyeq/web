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


CREATE TABLE `level_exp_mods` (
  `level` int(11) NOT NULL DEFAULT 0,
  `exp_mod` float DEFAULT NULL,
  `aa_exp_mod` float DEFAULT NULL,
  PRIMARY KEY (`level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "level": 48,    "exp_mod": 0.38754717,    "aa_exp_mod": 0.8630825}



*/

// LevelExpMods struct is a row record of the level_exp_mods table in the eqemu database
type LevelExpMods struct {
	//[ 0] level                                          int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Level int32 `gorm:"primary_key;column:level;type:int;default:0;" json:"level"`
	//[ 1] exp_mod                                        float                null: true   primary: false  isArray: false  auto: false  col: float           len: -1      default: [NULL]
	ExpMod null.Float `gorm:"column:exp_mod;type:float;" json:"exp_mod"`
	//[ 2] aa_exp_mod                                     float                null: true   primary: false  isArray: false  auto: false  col: float           len: -1      default: [NULL]
	AaExpMod null.Float `gorm:"column:aa_exp_mod;type:float;" json:"aa_exp_mod"`
}

var level_exp_modsTableInfo = &TableInfo{
	Name: "level_exp_mods",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "level",
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
			GoFieldName:        "Level",
			GoFieldType:        "int32",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "exp_mod",
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
			GoFieldName:        "ExpMod",
			GoFieldType:        "null.Float",
			JSONFieldName:      "exp_mod",
			ProtobufFieldName:  "exp_mod",
			ProtobufType:       "float",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "aa_exp_mod",
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
			GoFieldName:        "AaExpMod",
			GoFieldType:        "null.Float",
			JSONFieldName:      "aa_exp_mod",
			ProtobufFieldName:  "aa_exp_mod",
			ProtobufType:       "float",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LevelExpMods) TableName() string {
	return "level_exp_mods"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LevelExpMods) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LevelExpMods) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LevelExpMods) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LevelExpMods) TableInfo() *TableInfo {
	return level_exp_modsTableInfo
}
