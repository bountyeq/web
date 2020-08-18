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


CREATE TABLE `adventure_stats` (
  `player_id` int(10) unsigned NOT NULL,
  `guk_wins` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `mir_wins` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `mmc_wins` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `ruj_wins` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `tak_wins` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `guk_losses` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `mir_losses` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `mmc_losses` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `ruj_losses` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `tak_losses` mediumint(8) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`player_id`),
  UNIQUE KEY `player_id` (`player_id`),
  KEY `player_id_2` (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "player_id": 70,    "mir_wins": 66,    "tak_wins": 76,    "mir_losses": 35,    "mmc_losses": 53,    "tak_losses": 17,    "guk_wins": 77,    "mmc_wins": 71,    "ruj_wins": 72,    "guk_losses": 90,    "ruj_losses": 28}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned



*/

// AdventureStats struct is a row record of the adventure_stats table in the eqemu database
type AdventureStats struct {
	//[ 0] player_id                                      uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	PlayerID uint32 `gorm:"primary_key;column:player_id;type:uint;" json:"player_id"`
	//[ 1] guk_wins                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	GukWins uint32 `gorm:"column:guk_wins;type:umediumint;default:0;" json:"guk_wins"`
	//[ 2] mir_wins                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	MirWins uint32 `gorm:"column:mir_wins;type:umediumint;default:0;" json:"mir_wins"`
	//[ 3] mmc_wins                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	MmcWins uint32 `gorm:"column:mmc_wins;type:umediumint;default:0;" json:"mmc_wins"`
	//[ 4] ruj_wins                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	RujWins uint32 `gorm:"column:ruj_wins;type:umediumint;default:0;" json:"ruj_wins"`
	//[ 5] tak_wins                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	TakWins uint32 `gorm:"column:tak_wins;type:umediumint;default:0;" json:"tak_wins"`
	//[ 6] guk_losses                                     umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	GukLosses uint32 `gorm:"column:guk_losses;type:umediumint;default:0;" json:"guk_losses"`
	//[ 7] mir_losses                                     umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	MirLosses uint32 `gorm:"column:mir_losses;type:umediumint;default:0;" json:"mir_losses"`
	//[ 8] mmc_losses                                     umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	MmcLosses uint32 `gorm:"column:mmc_losses;type:umediumint;default:0;" json:"mmc_losses"`
	//[ 9] ruj_losses                                     umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	RujLosses uint32 `gorm:"column:ruj_losses;type:umediumint;default:0;" json:"ruj_losses"`
	//[10] tak_losses                                     umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	TakLosses uint32 `gorm:"column:tak_losses;type:umediumint;default:0;" json:"tak_losses"`
}

var adventure_statsTableInfo = &TableInfo{
	Name: "adventure_stats",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "player_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PlayerID",
			GoFieldType:        "uint32",
			JSONFieldName:      "player_id",
			ProtobufFieldName:  "player_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "guk_wins",
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
			GoFieldName:        "GukWins",
			GoFieldType:        "uint32",
			JSONFieldName:      "guk_wins",
			ProtobufFieldName:  "guk_wins",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "mir_wins",
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
			GoFieldName:        "MirWins",
			GoFieldType:        "uint32",
			JSONFieldName:      "mir_wins",
			ProtobufFieldName:  "mir_wins",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "mmc_wins",
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
			GoFieldName:        "MmcWins",
			GoFieldType:        "uint32",
			JSONFieldName:      "mmc_wins",
			ProtobufFieldName:  "mmc_wins",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "ruj_wins",
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
			GoFieldName:        "RujWins",
			GoFieldType:        "uint32",
			JSONFieldName:      "ruj_wins",
			ProtobufFieldName:  "ruj_wins",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "tak_wins",
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
			GoFieldName:        "TakWins",
			GoFieldType:        "uint32",
			JSONFieldName:      "tak_wins",
			ProtobufFieldName:  "tak_wins",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "guk_losses",
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
			GoFieldName:        "GukLosses",
			GoFieldType:        "uint32",
			JSONFieldName:      "guk_losses",
			ProtobufFieldName:  "guk_losses",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "mir_losses",
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
			GoFieldName:        "MirLosses",
			GoFieldType:        "uint32",
			JSONFieldName:      "mir_losses",
			ProtobufFieldName:  "mir_losses",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "mmc_losses",
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
			GoFieldName:        "MmcLosses",
			GoFieldType:        "uint32",
			JSONFieldName:      "mmc_losses",
			ProtobufFieldName:  "mmc_losses",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "ruj_losses",
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
			GoFieldName:        "RujLosses",
			GoFieldType:        "uint32",
			JSONFieldName:      "ruj_losses",
			ProtobufFieldName:  "ruj_losses",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "tak_losses",
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
			GoFieldName:        "TakLosses",
			GoFieldType:        "uint32",
			JSONFieldName:      "tak_losses",
			ProtobufFieldName:  "tak_losses",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AdventureStats) TableName() string {
	return "adventure_stats"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AdventureStats) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AdventureStats) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AdventureStats) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AdventureStats) TableInfo() *TableInfo {
	return adventure_statsTableInfo
}
