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


CREATE TABLE `character_corpses` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `charid` int(11) unsigned NOT NULL DEFAULT 0,
  `charname` varchar(64) NOT NULL DEFAULT '',
  `zone_id` smallint(5) NOT NULL DEFAULT 0,
  `instance_id` smallint(5) unsigned NOT NULL DEFAULT 0,
  `x` float NOT NULL DEFAULT 0,
  `y` float NOT NULL DEFAULT 0,
  `z` float NOT NULL DEFAULT 0,
  `heading` float NOT NULL DEFAULT 0,
  `time_of_death` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `guild_consent_id` int(11) unsigned NOT NULL DEFAULT 0,
  `is_rezzed` tinyint(3) unsigned DEFAULT 0,
  `is_buried` tinyint(3) NOT NULL DEFAULT 0,
  `was_at_graveyard` tinyint(3) NOT NULL DEFAULT 0,
  `is_locked` tinyint(11) DEFAULT 0,
  `exp` int(11) unsigned DEFAULT 0,
  `size` int(11) unsigned DEFAULT 0,
  `level` int(11) unsigned DEFAULT 0,
  `race` int(11) unsigned DEFAULT 0,
  `gender` int(11) unsigned DEFAULT 0,
  `class` int(11) unsigned DEFAULT 0,
  `deity` int(11) unsigned DEFAULT 0,
  `texture` int(11) unsigned DEFAULT 0,
  `helm_texture` int(11) unsigned DEFAULT 0,
  `copper` int(11) unsigned DEFAULT 0,
  `silver` int(11) unsigned DEFAULT 0,
  `gold` int(11) unsigned DEFAULT 0,
  `platinum` int(11) unsigned DEFAULT 0,
  `hair_color` int(11) unsigned DEFAULT 0,
  `beard_color` int(11) unsigned DEFAULT 0,
  `eye_color_1` int(11) unsigned DEFAULT 0,
  `eye_color_2` int(11) unsigned DEFAULT 0,
  `hair_style` int(11) unsigned DEFAULT 0,
  `face` int(11) unsigned DEFAULT 0,
  `beard` int(11) unsigned DEFAULT 0,
  `drakkin_heritage` int(11) unsigned DEFAULT 0,
  `drakkin_tattoo` int(11) unsigned DEFAULT 0,
  `drakkin_details` int(11) unsigned DEFAULT 0,
  `wc_1` int(11) unsigned DEFAULT 0,
  `wc_2` int(11) unsigned DEFAULT 0,
  `wc_3` int(11) unsigned DEFAULT 0,
  `wc_4` int(11) unsigned DEFAULT 0,
  `wc_5` int(11) unsigned DEFAULT 0,
  `wc_6` int(11) unsigned DEFAULT 0,
  `wc_7` int(11) unsigned DEFAULT 0,
  `wc_8` int(11) unsigned DEFAULT 0,
  `wc_9` int(11) unsigned DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `zoneid` (`zone_id`),
  KEY `instanceid` (`instance_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "wc_2": 46,    "wc_8": 27,    "time_of_death": "2203-09-26T21:23:59.388111742-07:00",    "guild_consent_id": 51,    "exp": 88,    "gender": 17,    "texture": 76,    "wc_1": 94,    "level": 2,    "helm_texture": 29,    "hair_style": 74,    "drakkin_tattoo": 24,    "id": 49,    "x": 0.4529282,    "race": 55,    "beard_color": 13,    "wc_6": 29,    "instance_id": 23,    "class": 50,    "wc_4": 80,    "eye_color_1": 38,    "drakkin_details": 14,    "charname": "fOExhERgDjpbGJxsZMeJsPDMU",    "is_locked": 5,    "size": 35,    "deity": 35,    "copper": 20,    "gold": 62,    "wc_9": 36,    "z": 0.5645386,    "is_rezzed": 79,    "was_at_graveyard": 20,    "silver": 86,    "face": 66,    "wc_7": 55,    "wc_3": 25,    "wc_5": 35,    "zone_id": 70,    "heading": 0.52388763,    "is_buried": 4,    "hair_color": 58,    "eye_color_2": 84,    "beard": 22,    "charid": 87,    "y": 0.95367986,    "platinum": 96,    "drakkin_heritage": 80}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 4] column is set for unsigned
[10] column is set for unsigned
[11] column is set for unsigned
[15] column is set for unsigned
[16] column is set for unsigned
[17] column is set for unsigned
[18] column is set for unsigned
[19] column is set for unsigned
[20] column is set for unsigned
[21] column is set for unsigned
[22] column is set for unsigned
[23] column is set for unsigned
[24] column is set for unsigned
[25] column is set for unsigned
[26] column is set for unsigned
[27] column is set for unsigned
[28] column is set for unsigned
[29] column is set for unsigned
[30] column is set for unsigned
[31] column is set for unsigned
[32] column is set for unsigned
[33] column is set for unsigned
[34] column is set for unsigned
[35] column is set for unsigned
[36] column is set for unsigned
[37] column is set for unsigned
[38] column is set for unsigned
[39] column is set for unsigned
[40] column is set for unsigned
[41] column is set for unsigned
[42] column is set for unsigned
[43] column is set for unsigned
[44] column is set for unsigned
[45] column is set for unsigned
[46] column is set for unsigned



*/

// CharacterCorpses struct is a row record of the character_corpses table in the eqemu database
type CharacterCorpses struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] charid                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Charid uint32 `gorm:"column:charid;type:uint;default:0;" json:"charid"`
	//[ 2] charname                                       varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Charname string `gorm:"column:charname;type:varchar;size:64;default:'';" json:"charname"`
	//[ 3] zone_id                                        smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	ZoneID int32 `gorm:"column:zone_id;type:smallint;default:0;" json:"zone_id"`
	//[ 4] instance_id                                    usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	InstanceID uint32 `gorm:"column:instance_id;type:usmallint;default:0;" json:"instance_id"`
	//[ 5] x                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	X float32 `gorm:"column:x;type:float;default:0;" json:"x"`
	//[ 6] y                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Y float32 `gorm:"column:y;type:float;default:0;" json:"y"`
	//[ 7] z                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Z float32 `gorm:"column:z;type:float;default:0;" json:"z"`
	//[ 8] heading                                        float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Heading float32 `gorm:"column:heading;type:float;default:0;" json:"heading"`
	//[ 9] time_of_death                                  datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: ['0000-00-00 00:00:00']
	TimeOfDeath time.Time `gorm:"column:time_of_death;type:datetime;default:'0000-00-00 00:00:00';" json:"time_of_death"`
	//[10] guild_consent_id                               uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	GuildConsentID uint32 `gorm:"column:guild_consent_id;type:uint;default:0;" json:"guild_consent_id"`
	//[11] is_rezzed                                      utinyint             null: true   primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	IsRezzed null.Int `gorm:"column:is_rezzed;type:utinyint;default:0;" json:"is_rezzed"`
	//[12] is_buried                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsBuried int32 `gorm:"column:is_buried;type:tinyint;default:0;" json:"is_buried"`
	//[13] was_at_graveyard                               tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	WasAtGraveyard int32 `gorm:"column:was_at_graveyard;type:tinyint;default:0;" json:"was_at_graveyard"`
	//[14] is_locked                                      tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsLocked null.Int `gorm:"column:is_locked;type:tinyint;default:0;" json:"is_locked"`
	//[15] exp                                            uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Exp null.Int `gorm:"column:exp;type:uint;default:0;" json:"exp"`
	//[16] size                                           uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Size null.Int `gorm:"column:size;type:uint;default:0;" json:"size"`
	//[17] level                                          uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Level null.Int `gorm:"column:level;type:uint;default:0;" json:"level"`
	//[18] race                                           uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Race null.Int `gorm:"column:race;type:uint;default:0;" json:"race"`
	//[19] gender                                         uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Gender null.Int `gorm:"column:gender;type:uint;default:0;" json:"gender"`
	//[20] class                                          uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Class null.Int `gorm:"column:class;type:uint;default:0;" json:"class"`
	//[21] deity                                          uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Deity null.Int `gorm:"column:deity;type:uint;default:0;" json:"deity"`
	//[22] texture                                        uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Texture null.Int `gorm:"column:texture;type:uint;default:0;" json:"texture"`
	//[23] helm_texture                                   uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	HelmTexture null.Int `gorm:"column:helm_texture;type:uint;default:0;" json:"helm_texture"`
	//[24] copper                                         uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Copper null.Int `gorm:"column:copper;type:uint;default:0;" json:"copper"`
	//[25] silver                                         uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Silver null.Int `gorm:"column:silver;type:uint;default:0;" json:"silver"`
	//[26] gold                                           uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Gold null.Int `gorm:"column:gold;type:uint;default:0;" json:"gold"`
	//[27] platinum                                       uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Platinum null.Int `gorm:"column:platinum;type:uint;default:0;" json:"platinum"`
	//[28] hair_color                                     uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	HairColor null.Int `gorm:"column:hair_color;type:uint;default:0;" json:"hair_color"`
	//[29] beard_color                                    uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	BeardColor null.Int `gorm:"column:beard_color;type:uint;default:0;" json:"beard_color"`
	//[30] eye_color_1                                    uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EyeColor1 null.Int `gorm:"column:eye_color_1;type:uint;default:0;" json:"eye_color_1"`
	//[31] eye_color_2                                    uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EyeColor2 null.Int `gorm:"column:eye_color_2;type:uint;default:0;" json:"eye_color_2"`
	//[32] hair_style                                     uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	HairStyle null.Int `gorm:"column:hair_style;type:uint;default:0;" json:"hair_style"`
	//[33] face                                           uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Face null.Int `gorm:"column:face;type:uint;default:0;" json:"face"`
	//[34] beard                                          uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Beard null.Int `gorm:"column:beard;type:uint;default:0;" json:"beard"`
	//[35] drakkin_heritage                               uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	DrakkinHeritage null.Int `gorm:"column:drakkin_heritage;type:uint;default:0;" json:"drakkin_heritage"`
	//[36] drakkin_tattoo                                 uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	DrakkinTattoo null.Int `gorm:"column:drakkin_tattoo;type:uint;default:0;" json:"drakkin_tattoo"`
	//[37] drakkin_details                                uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	DrakkinDetails null.Int `gorm:"column:drakkin_details;type:uint;default:0;" json:"drakkin_details"`
	//[38] wc_1                                           uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Wc1 null.Int `gorm:"column:wc_1;type:uint;default:0;" json:"wc_1"`
	//[39] wc_2                                           uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Wc2 null.Int `gorm:"column:wc_2;type:uint;default:0;" json:"wc_2"`
	//[40] wc_3                                           uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Wc3 null.Int `gorm:"column:wc_3;type:uint;default:0;" json:"wc_3"`
	//[41] wc_4                                           uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Wc4 null.Int `gorm:"column:wc_4;type:uint;default:0;" json:"wc_4"`
	//[42] wc_5                                           uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Wc5 null.Int `gorm:"column:wc_5;type:uint;default:0;" json:"wc_5"`
	//[43] wc_6                                           uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Wc6 null.Int `gorm:"column:wc_6;type:uint;default:0;" json:"wc_6"`
	//[44] wc_7                                           uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Wc7 null.Int `gorm:"column:wc_7;type:uint;default:0;" json:"wc_7"`
	//[45] wc_8                                           uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Wc8 null.Int `gorm:"column:wc_8;type:uint;default:0;" json:"wc_8"`
	//[46] wc_9                                           uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Wc9 null.Int `gorm:"column:wc_9;type:uint;default:0;" json:"wc_9"`
}

var character_corpsesTableInfo = &TableInfo{
	Name: "character_corpses",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "charid",
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
			GoFieldName:        "Charid",
			GoFieldType:        "uint32",
			JSONFieldName:      "charid",
			ProtobufFieldName:  "charid",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "charname",
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
			GoFieldName:        "Charname",
			GoFieldType:        "string",
			JSONFieldName:      "charname",
			ProtobufFieldName:  "charname",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "zone_id",
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
			GoFieldName:        "ZoneID",
			GoFieldType:        "int32",
			JSONFieldName:      "zone_id",
			ProtobufFieldName:  "zone_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "instance_id",
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
			GoFieldName:        "InstanceID",
			GoFieldType:        "uint32",
			JSONFieldName:      "instance_id",
			ProtobufFieldName:  "instance_id",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			Name:               "time_of_death",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "TimeOfDeath",
			GoFieldType:        "time.Time",
			JSONFieldName:      "time_of_death",
			ProtobufFieldName:  "time_of_death",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "guild_consent_id",
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
			GoFieldName:        "GuildConsentID",
			GoFieldType:        "uint32",
			JSONFieldName:      "guild_consent_id",
			ProtobufFieldName:  "guild_consent_id",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "is_rezzed",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsRezzed",
			GoFieldType:        "null.Int",
			JSONFieldName:      "is_rezzed",
			ProtobufFieldName:  "is_rezzed",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "is_buried",
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
			GoFieldName:        "IsBuried",
			GoFieldType:        "int32",
			JSONFieldName:      "is_buried",
			ProtobufFieldName:  "is_buried",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "was_at_graveyard",
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
			GoFieldName:        "WasAtGraveyard",
			GoFieldType:        "int32",
			JSONFieldName:      "was_at_graveyard",
			ProtobufFieldName:  "was_at_graveyard",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "is_locked",
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
			GoFieldName:        "IsLocked",
			GoFieldType:        "null.Int",
			JSONFieldName:      "is_locked",
			ProtobufFieldName:  "is_locked",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "exp",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Exp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "exp",
			ProtobufFieldName:  "exp",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "size",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Size",
			GoFieldType:        "null.Int",
			JSONFieldName:      "size",
			ProtobufFieldName:  "size",
			ProtobufType:       "uint32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "level",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Level",
			GoFieldType:        "null.Int",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "uint32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "race",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Race",
			GoFieldType:        "null.Int",
			JSONFieldName:      "race",
			ProtobufFieldName:  "race",
			ProtobufType:       "uint32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "gender",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Gender",
			GoFieldType:        "null.Int",
			JSONFieldName:      "gender",
			ProtobufFieldName:  "gender",
			ProtobufType:       "uint32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "class",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Class",
			GoFieldType:        "null.Int",
			JSONFieldName:      "class",
			ProtobufFieldName:  "class",
			ProtobufType:       "uint32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "deity",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Deity",
			GoFieldType:        "null.Int",
			JSONFieldName:      "deity",
			ProtobufFieldName:  "deity",
			ProtobufType:       "uint32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "texture",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Texture",
			GoFieldType:        "null.Int",
			JSONFieldName:      "texture",
			ProtobufFieldName:  "texture",
			ProtobufType:       "uint32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "helm_texture",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "HelmTexture",
			GoFieldType:        "null.Int",
			JSONFieldName:      "helm_texture",
			ProtobufFieldName:  "helm_texture",
			ProtobufType:       "uint32",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "copper",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Copper",
			GoFieldType:        "null.Int",
			JSONFieldName:      "copper",
			ProtobufFieldName:  "copper",
			ProtobufType:       "uint32",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "silver",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Silver",
			GoFieldType:        "null.Int",
			JSONFieldName:      "silver",
			ProtobufFieldName:  "silver",
			ProtobufType:       "uint32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "gold",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Gold",
			GoFieldType:        "null.Int",
			JSONFieldName:      "gold",
			ProtobufFieldName:  "gold",
			ProtobufType:       "uint32",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "platinum",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Platinum",
			GoFieldType:        "null.Int",
			JSONFieldName:      "platinum",
			ProtobufFieldName:  "platinum",
			ProtobufType:       "uint32",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "hair_color",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "HairColor",
			GoFieldType:        "null.Int",
			JSONFieldName:      "hair_color",
			ProtobufFieldName:  "hair_color",
			ProtobufType:       "uint32",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "beard_color",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "BeardColor",
			GoFieldType:        "null.Int",
			JSONFieldName:      "beard_color",
			ProtobufFieldName:  "beard_color",
			ProtobufType:       "uint32",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "eye_color_1",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "EyeColor1",
			GoFieldType:        "null.Int",
			JSONFieldName:      "eye_color_1",
			ProtobufFieldName:  "eye_color_1",
			ProtobufType:       "uint32",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "eye_color_2",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "EyeColor2",
			GoFieldType:        "null.Int",
			JSONFieldName:      "eye_color_2",
			ProtobufFieldName:  "eye_color_2",
			ProtobufType:       "uint32",
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "hair_style",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "HairStyle",
			GoFieldType:        "null.Int",
			JSONFieldName:      "hair_style",
			ProtobufFieldName:  "hair_style",
			ProtobufType:       "uint32",
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
			Name:               "face",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Face",
			GoFieldType:        "null.Int",
			JSONFieldName:      "face",
			ProtobufFieldName:  "face",
			ProtobufType:       "uint32",
			ProtobufPos:        34,
		},

		&ColumnInfo{
			Index:              34,
			Name:               "beard",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Beard",
			GoFieldType:        "null.Int",
			JSONFieldName:      "beard",
			ProtobufFieldName:  "beard",
			ProtobufType:       "uint32",
			ProtobufPos:        35,
		},

		&ColumnInfo{
			Index:              35,
			Name:               "drakkin_heritage",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "DrakkinHeritage",
			GoFieldType:        "null.Int",
			JSONFieldName:      "drakkin_heritage",
			ProtobufFieldName:  "drakkin_heritage",
			ProtobufType:       "uint32",
			ProtobufPos:        36,
		},

		&ColumnInfo{
			Index:              36,
			Name:               "drakkin_tattoo",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "DrakkinTattoo",
			GoFieldType:        "null.Int",
			JSONFieldName:      "drakkin_tattoo",
			ProtobufFieldName:  "drakkin_tattoo",
			ProtobufType:       "uint32",
			ProtobufPos:        37,
		},

		&ColumnInfo{
			Index:              37,
			Name:               "drakkin_details",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "DrakkinDetails",
			GoFieldType:        "null.Int",
			JSONFieldName:      "drakkin_details",
			ProtobufFieldName:  "drakkin_details",
			ProtobufType:       "uint32",
			ProtobufPos:        38,
		},

		&ColumnInfo{
			Index:              38,
			Name:               "wc_1",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Wc1",
			GoFieldType:        "null.Int",
			JSONFieldName:      "wc_1",
			ProtobufFieldName:  "wc_1",
			ProtobufType:       "uint32",
			ProtobufPos:        39,
		},

		&ColumnInfo{
			Index:              39,
			Name:               "wc_2",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Wc2",
			GoFieldType:        "null.Int",
			JSONFieldName:      "wc_2",
			ProtobufFieldName:  "wc_2",
			ProtobufType:       "uint32",
			ProtobufPos:        40,
		},

		&ColumnInfo{
			Index:              40,
			Name:               "wc_3",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Wc3",
			GoFieldType:        "null.Int",
			JSONFieldName:      "wc_3",
			ProtobufFieldName:  "wc_3",
			ProtobufType:       "uint32",
			ProtobufPos:        41,
		},

		&ColumnInfo{
			Index:              41,
			Name:               "wc_4",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Wc4",
			GoFieldType:        "null.Int",
			JSONFieldName:      "wc_4",
			ProtobufFieldName:  "wc_4",
			ProtobufType:       "uint32",
			ProtobufPos:        42,
		},

		&ColumnInfo{
			Index:              42,
			Name:               "wc_5",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Wc5",
			GoFieldType:        "null.Int",
			JSONFieldName:      "wc_5",
			ProtobufFieldName:  "wc_5",
			ProtobufType:       "uint32",
			ProtobufPos:        43,
		},

		&ColumnInfo{
			Index:              43,
			Name:               "wc_6",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Wc6",
			GoFieldType:        "null.Int",
			JSONFieldName:      "wc_6",
			ProtobufFieldName:  "wc_6",
			ProtobufType:       "uint32",
			ProtobufPos:        44,
		},

		&ColumnInfo{
			Index:              44,
			Name:               "wc_7",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Wc7",
			GoFieldType:        "null.Int",
			JSONFieldName:      "wc_7",
			ProtobufFieldName:  "wc_7",
			ProtobufType:       "uint32",
			ProtobufPos:        45,
		},

		&ColumnInfo{
			Index:              45,
			Name:               "wc_8",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Wc8",
			GoFieldType:        "null.Int",
			JSONFieldName:      "wc_8",
			ProtobufFieldName:  "wc_8",
			ProtobufType:       "uint32",
			ProtobufPos:        46,
		},

		&ColumnInfo{
			Index:              46,
			Name:               "wc_9",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Wc9",
			GoFieldType:        "null.Int",
			JSONFieldName:      "wc_9",
			ProtobufFieldName:  "wc_9",
			ProtobufType:       "uint32",
			ProtobufPos:        47,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterCorpses) TableName() string {
	return "character_corpses"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterCorpses) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterCorpses) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterCorpses) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterCorpses) TableInfo() *TableInfo {
	return character_corpsesTableInfo
}
