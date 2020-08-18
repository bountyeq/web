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


CREATE TABLE `npc_scale_global_base` (
  `type` int(11) NOT NULL DEFAULT 0,
  `level` int(11) NOT NULL,
  `ac` int(11) DEFAULT NULL,
  `hp` int(11) DEFAULT NULL,
  `accuracy` int(11) DEFAULT NULL,
  `slow_mitigation` int(11) DEFAULT NULL,
  `attack` int(11) DEFAULT NULL,
  `strength` int(11) DEFAULT NULL,
  `stamina` int(11) DEFAULT NULL,
  `dexterity` int(11) DEFAULT NULL,
  `agility` int(11) DEFAULT NULL,
  `intelligence` int(11) DEFAULT NULL,
  `wisdom` int(11) DEFAULT NULL,
  `charisma` int(11) DEFAULT NULL,
  `magic_resist` int(11) DEFAULT NULL,
  `cold_resist` int(11) DEFAULT NULL,
  `fire_resist` int(11) DEFAULT NULL,
  `poison_resist` int(11) DEFAULT NULL,
  `disease_resist` int(11) DEFAULT NULL,
  `corruption_resist` int(11) DEFAULT NULL,
  `physical_resist` int(11) DEFAULT NULL,
  `min_dmg` int(11) DEFAULT NULL,
  `max_dmg` int(11) DEFAULT NULL,
  `hp_regen_rate` int(11) DEFAULT NULL,
  `attack_delay` int(11) DEFAULT NULL,
  `spell_scale` int(11) DEFAULT 100,
  `heal_scale` int(11) DEFAULT 100,
  `special_abilities` text DEFAULT NULL,
  PRIMARY KEY (`type`,`level`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

JSON Sample
-------------------------------------
{    "wisdom": 34,    "ac": 5,    "charisma": 8,    "hp_regen_rate": 17,    "stamina": 79,    "dexterity": 34,    "corruption_resist": 69,    "spell_scale": 34,    "fire_resist": 6,    "physical_resist": 4,    "max_dmg": 41,    "poison_resist": 74,    "min_dmg": 5,    "attack_delay": 62,    "hp": 99,    "attack": 38,    "agility": 7,    "special_abilities": "imofpwpggpehuQewKphYnVcRc",    "accuracy": 8,    "strength": 40,    "intelligence": 55,    "slow_mitigation": 84,    "cold_resist": 42,    "disease_resist": 18,    "heal_scale": 80,    "type": 62,    "level": 44,    "magic_resist": 30}



*/

// NpcScaleGlobalBase struct is a row record of the npc_scale_global_base table in the eqemu database
type NpcScaleGlobalBase struct {
	//[ 0] type                                           int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Type int32 `gorm:"primary_key;column:type;type:int;default:0;" json:"type"`
	//[ 1] level                                          int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	Level int32 `gorm:"primary_key;column:level;type:int;" json:"level"`
	//[ 2] ac                                             int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	Ac null.Int `gorm:"column:ac;type:int;" json:"ac"`
	//[ 3] hp                                             int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	Hp null.Int `gorm:"column:hp;type:int;" json:"hp"`
	//[ 4] accuracy                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	Accuracy null.Int `gorm:"column:accuracy;type:int;" json:"accuracy"`
	//[ 5] slow_mitigation                                int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	SlowMitigation null.Int `gorm:"column:slow_mitigation;type:int;" json:"slow_mitigation"`
	//[ 6] attack                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	Attack null.Int `gorm:"column:attack;type:int;" json:"attack"`
	//[ 7] strength                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	Strength null.Int `gorm:"column:strength;type:int;" json:"strength"`
	//[ 8] stamina                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	Stamina null.Int `gorm:"column:stamina;type:int;" json:"stamina"`
	//[ 9] dexterity                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	Dexterity null.Int `gorm:"column:dexterity;type:int;" json:"dexterity"`
	//[10] agility                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	Agility null.Int `gorm:"column:agility;type:int;" json:"agility"`
	//[11] intelligence                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	Intelligence null.Int `gorm:"column:intelligence;type:int;" json:"intelligence"`
	//[12] wisdom                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	Wisdom null.Int `gorm:"column:wisdom;type:int;" json:"wisdom"`
	//[13] charisma                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	Charisma null.Int `gorm:"column:charisma;type:int;" json:"charisma"`
	//[14] magic_resist                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	MagicResist null.Int `gorm:"column:magic_resist;type:int;" json:"magic_resist"`
	//[15] cold_resist                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	ColdResist null.Int `gorm:"column:cold_resist;type:int;" json:"cold_resist"`
	//[16] fire_resist                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	FireResist null.Int `gorm:"column:fire_resist;type:int;" json:"fire_resist"`
	//[17] poison_resist                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	PoisonResist null.Int `gorm:"column:poison_resist;type:int;" json:"poison_resist"`
	//[18] disease_resist                                 int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	DiseaseResist null.Int `gorm:"column:disease_resist;type:int;" json:"disease_resist"`
	//[19] corruption_resist                              int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	CorruptionResist null.Int `gorm:"column:corruption_resist;type:int;" json:"corruption_resist"`
	//[20] physical_resist                                int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	PhysicalResist null.Int `gorm:"column:physical_resist;type:int;" json:"physical_resist"`
	//[21] min_dmg                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	MinDmg null.Int `gorm:"column:min_dmg;type:int;" json:"min_dmg"`
	//[22] max_dmg                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	MaxDmg null.Int `gorm:"column:max_dmg;type:int;" json:"max_dmg"`
	//[23] hp_regen_rate                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	HpRegenRate null.Int `gorm:"column:hp_regen_rate;type:int;" json:"hp_regen_rate"`
	//[24] attack_delay                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [NULL]
	AttackDelay null.Int `gorm:"column:attack_delay;type:int;" json:"attack_delay"`
	//[25] spell_scale                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [100]
	SpellScale null.Int `gorm:"column:spell_scale;type:int;default:100;" json:"spell_scale"`
	//[26] heal_scale                                     int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [100]
	HealScale null.Int `gorm:"column:heal_scale;type:int;default:100;" json:"heal_scale"`
	//[27] special_abilities                              text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: [NULL]
	SpecialAbilities null.String `gorm:"column:special_abilities;type:text;size:65535;" json:"special_abilities"`
}

var npc_scale_global_baseTableInfo = &TableInfo{
	Name: "npc_scale_global_base",
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ac",
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
			GoFieldName:        "Ac",
			GoFieldType:        "null.Int",
			JSONFieldName:      "ac",
			ProtobufFieldName:  "ac",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "hp",
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
			GoFieldName:        "Hp",
			GoFieldType:        "null.Int",
			JSONFieldName:      "hp",
			ProtobufFieldName:  "hp",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "accuracy",
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
			GoFieldName:        "Accuracy",
			GoFieldType:        "null.Int",
			JSONFieldName:      "accuracy",
			ProtobufFieldName:  "accuracy",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "slow_mitigation",
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
			GoFieldName:        "SlowMitigation",
			GoFieldType:        "null.Int",
			JSONFieldName:      "slow_mitigation",
			ProtobufFieldName:  "slow_mitigation",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "attack",
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
			GoFieldName:        "Attack",
			GoFieldType:        "null.Int",
			JSONFieldName:      "attack",
			ProtobufFieldName:  "attack",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "strength",
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
			GoFieldName:        "Strength",
			GoFieldType:        "null.Int",
			JSONFieldName:      "strength",
			ProtobufFieldName:  "strength",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "stamina",
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
			GoFieldName:        "Stamina",
			GoFieldType:        "null.Int",
			JSONFieldName:      "stamina",
			ProtobufFieldName:  "stamina",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "dexterity",
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
			GoFieldName:        "Dexterity",
			GoFieldType:        "null.Int",
			JSONFieldName:      "dexterity",
			ProtobufFieldName:  "dexterity",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "agility",
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
			GoFieldName:        "Agility",
			GoFieldType:        "null.Int",
			JSONFieldName:      "agility",
			ProtobufFieldName:  "agility",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "intelligence",
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
			GoFieldName:        "Intelligence",
			GoFieldType:        "null.Int",
			JSONFieldName:      "intelligence",
			ProtobufFieldName:  "intelligence",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "wisdom",
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
			GoFieldName:        "Wisdom",
			GoFieldType:        "null.Int",
			JSONFieldName:      "wisdom",
			ProtobufFieldName:  "wisdom",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "charisma",
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
			GoFieldName:        "Charisma",
			GoFieldType:        "null.Int",
			JSONFieldName:      "charisma",
			ProtobufFieldName:  "charisma",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "magic_resist",
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
			GoFieldName:        "MagicResist",
			GoFieldType:        "null.Int",
			JSONFieldName:      "magic_resist",
			ProtobufFieldName:  "magic_resist",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "cold_resist",
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
			GoFieldName:        "ColdResist",
			GoFieldType:        "null.Int",
			JSONFieldName:      "cold_resist",
			ProtobufFieldName:  "cold_resist",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "fire_resist",
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
			GoFieldName:        "FireResist",
			GoFieldType:        "null.Int",
			JSONFieldName:      "fire_resist",
			ProtobufFieldName:  "fire_resist",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "poison_resist",
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
			GoFieldName:        "PoisonResist",
			GoFieldType:        "null.Int",
			JSONFieldName:      "poison_resist",
			ProtobufFieldName:  "poison_resist",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "disease_resist",
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
			GoFieldName:        "DiseaseResist",
			GoFieldType:        "null.Int",
			JSONFieldName:      "disease_resist",
			ProtobufFieldName:  "disease_resist",
			ProtobufType:       "int32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "corruption_resist",
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
			GoFieldName:        "CorruptionResist",
			GoFieldType:        "null.Int",
			JSONFieldName:      "corruption_resist",
			ProtobufFieldName:  "corruption_resist",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "physical_resist",
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
			GoFieldName:        "PhysicalResist",
			GoFieldType:        "null.Int",
			JSONFieldName:      "physical_resist",
			ProtobufFieldName:  "physical_resist",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "min_dmg",
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
			GoFieldName:        "MinDmg",
			GoFieldType:        "null.Int",
			JSONFieldName:      "min_dmg",
			ProtobufFieldName:  "min_dmg",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "max_dmg",
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
			GoFieldName:        "MaxDmg",
			GoFieldType:        "null.Int",
			JSONFieldName:      "max_dmg",
			ProtobufFieldName:  "max_dmg",
			ProtobufType:       "int32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "hp_regen_rate",
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
			GoFieldName:        "HpRegenRate",
			GoFieldType:        "null.Int",
			JSONFieldName:      "hp_regen_rate",
			ProtobufFieldName:  "hp_regen_rate",
			ProtobufType:       "int32",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "attack_delay",
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
			GoFieldName:        "AttackDelay",
			GoFieldType:        "null.Int",
			JSONFieldName:      "attack_delay",
			ProtobufFieldName:  "attack_delay",
			ProtobufType:       "int32",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "spell_scale",
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
			GoFieldName:        "SpellScale",
			GoFieldType:        "null.Int",
			JSONFieldName:      "spell_scale",
			ProtobufFieldName:  "spell_scale",
			ProtobufType:       "int32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "heal_scale",
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
			GoFieldName:        "HealScale",
			GoFieldType:        "null.Int",
			JSONFieldName:      "heal_scale",
			ProtobufFieldName:  "heal_scale",
			ProtobufType:       "int32",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "special_abilities",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "SpecialAbilities",
			GoFieldType:        "null.String",
			JSONFieldName:      "special_abilities",
			ProtobufFieldName:  "special_abilities",
			ProtobufType:       "string",
			ProtobufPos:        28,
		},
	},
}

// TableName sets the insert table name for this struct type
func (n *NpcScaleGlobalBase) TableName() string {
	return "npc_scale_global_base"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (n *NpcScaleGlobalBase) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (n *NpcScaleGlobalBase) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (n *NpcScaleGlobalBase) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (n *NpcScaleGlobalBase) TableInfo() *TableInfo {
	return npc_scale_global_baseTableInfo
}
