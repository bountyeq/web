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


CREATE TABLE `character_data` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL DEFAULT 0,
  `name` varchar(64) NOT NULL DEFAULT '',
  `last_name` varchar(64) NOT NULL DEFAULT '',
  `title` varchar(32) NOT NULL DEFAULT '',
  `suffix` varchar(32) NOT NULL DEFAULT '',
  `zone_id` int(11) unsigned NOT NULL DEFAULT 0,
  `zone_instance` int(11) unsigned NOT NULL DEFAULT 0,
  `y` float NOT NULL DEFAULT 0,
  `x` float NOT NULL DEFAULT 0,
  `z` float NOT NULL DEFAULT 0,
  `heading` float NOT NULL DEFAULT 0,
  `gender` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `race` smallint(11) unsigned NOT NULL DEFAULT 0,
  `class` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `level` int(11) unsigned NOT NULL DEFAULT 0,
  `deity` int(11) unsigned NOT NULL DEFAULT 0,
  `birthday` int(11) unsigned NOT NULL DEFAULT 0,
  `last_login` int(11) unsigned NOT NULL DEFAULT 0,
  `time_played` int(11) unsigned NOT NULL DEFAULT 0,
  `level2` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `anon` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `gm` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `face` int(11) unsigned NOT NULL DEFAULT 0,
  `hair_color` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `hair_style` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `beard` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `beard_color` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `eye_color_1` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `eye_color_2` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `drakkin_heritage` int(11) unsigned NOT NULL DEFAULT 0,
  `drakkin_tattoo` int(11) unsigned NOT NULL DEFAULT 0,
  `drakkin_details` int(11) unsigned NOT NULL DEFAULT 0,
  `ability_time_seconds` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `ability_number` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `ability_time_minutes` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `ability_time_hours` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `exp` int(11) unsigned NOT NULL DEFAULT 0,
  `aa_points_spent` int(11) unsigned NOT NULL DEFAULT 0,
  `aa_exp` int(11) unsigned NOT NULL DEFAULT 0,
  `aa_points` int(11) unsigned NOT NULL DEFAULT 0,
  `group_leadership_exp` int(11) unsigned NOT NULL DEFAULT 0,
  `raid_leadership_exp` int(11) unsigned NOT NULL DEFAULT 0,
  `group_leadership_points` int(11) unsigned NOT NULL DEFAULT 0,
  `raid_leadership_points` int(11) unsigned NOT NULL DEFAULT 0,
  `points` int(11) unsigned NOT NULL DEFAULT 0,
  `cur_hp` int(11) unsigned NOT NULL DEFAULT 0,
  `mana` int(11) unsigned NOT NULL DEFAULT 0,
  `endurance` int(11) unsigned NOT NULL DEFAULT 0,
  `intoxication` int(11) unsigned NOT NULL DEFAULT 0,
  `str` int(11) unsigned NOT NULL DEFAULT 0,
  `sta` int(11) unsigned NOT NULL DEFAULT 0,
  `cha` int(11) unsigned NOT NULL DEFAULT 0,
  `dex` int(11) unsigned NOT NULL DEFAULT 0,
  `int` int(11) unsigned NOT NULL DEFAULT 0,
  `agi` int(11) unsigned NOT NULL DEFAULT 0,
  `wis` int(11) unsigned NOT NULL DEFAULT 0,
  `zone_change_count` int(11) unsigned NOT NULL DEFAULT 0,
  `toxicity` int(11) unsigned NOT NULL DEFAULT 0,
  `hunger_level` int(11) unsigned NOT NULL DEFAULT 0,
  `thirst_level` int(11) unsigned NOT NULL DEFAULT 0,
  `ability_up` int(11) unsigned NOT NULL DEFAULT 0,
  `ldon_points_guk` int(11) unsigned NOT NULL DEFAULT 0,
  `ldon_points_mir` int(11) unsigned NOT NULL DEFAULT 0,
  `ldon_points_mmc` int(11) unsigned NOT NULL DEFAULT 0,
  `ldon_points_ruj` int(11) unsigned NOT NULL DEFAULT 0,
  `ldon_points_tak` int(11) unsigned NOT NULL DEFAULT 0,
  `ldon_points_available` int(11) unsigned NOT NULL DEFAULT 0,
  `tribute_time_remaining` int(11) unsigned NOT NULL DEFAULT 0,
  `career_tribute_points` int(11) unsigned NOT NULL DEFAULT 0,
  `tribute_points` int(11) unsigned NOT NULL DEFAULT 0,
  `tribute_active` int(11) unsigned NOT NULL DEFAULT 0,
  `pvp_status` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `pvp_kills` int(11) unsigned NOT NULL DEFAULT 0,
  `pvp_deaths` int(11) unsigned NOT NULL DEFAULT 0,
  `pvp_current_points` int(11) unsigned NOT NULL DEFAULT 0,
  `pvp_career_points` int(11) unsigned NOT NULL DEFAULT 0,
  `pvp_best_kill_streak` int(11) unsigned NOT NULL DEFAULT 0,
  `pvp_worst_death_streak` int(11) unsigned NOT NULL DEFAULT 0,
  `pvp_current_kill_streak` int(11) unsigned NOT NULL DEFAULT 0,
  `pvp2` int(11) unsigned NOT NULL DEFAULT 0,
  `pvp_type` int(11) unsigned NOT NULL DEFAULT 0,
  `show_helm` int(11) unsigned NOT NULL DEFAULT 0,
  `group_auto_consent` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `raid_auto_consent` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `guild_auto_consent` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `leadership_exp_on` tinyint(11) unsigned NOT NULL DEFAULT 0,
  `RestTimer` int(11) unsigned NOT NULL DEFAULT 0,
  `air_remaining` int(11) unsigned NOT NULL DEFAULT 0,
  `autosplit_enabled` int(11) unsigned NOT NULL DEFAULT 0,
  `lfp` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `lfg` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `mailkey` char(16) NOT NULL DEFAULT '',
  `xtargets` tinyint(3) unsigned NOT NULL DEFAULT 5,
  `firstlogon` tinyint(3) NOT NULL DEFAULT 0,
  `e_aa_effects` int(11) unsigned NOT NULL DEFAULT 0,
  `e_percent_to_aa` int(11) unsigned NOT NULL DEFAULT 0,
  `e_expended_aa_spent` int(11) unsigned NOT NULL DEFAULT 0,
  `aa_points_spent_old` int(11) unsigned NOT NULL DEFAULT 0,
  `aa_points_old` int(11) unsigned NOT NULL DEFAULT 0,
  `e_last_invsnapshot` int(11) unsigned NOT NULL DEFAULT 0,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `account_id` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "level": 82,    "career_tribute_points": 48,    "level_2": 65,    "drakkin_heritage": 43,    "ability_time_seconds": 63,    "tribute_active": 30,    "pvp_deaths": 90,    "lfp": 15,    "id": 88,    "eye_color_2": 1,    "ability_time_minutes": 23,    "aa_points": 13,    "group_leadership_points": 24,    "e_expended_aa_spent": 72,    "aa_points_old": 25,    "deity": 54,    "ability_number": 97,    "endurance": 22,    "agi": 3,    "hunger_level": 90,    "ldon_points_ruj": 8,    "group_auto_consent": 91,    "heading": 0.084231935,    "exp": 61,    "pvp_kills": 27,    "pvp_career_points": 67,    "show_helm": 42,    "autosplit_enabled": 43,    "x": 0.6636293,    "beard": 55,    "beard_color": 55,    "ability_time_hours": 13,    "cha": 26,    "pvp_worst_death_streak": 38,    "last_login": 85,    "zone_change_count": 12,    "toxicity": 87,    "ldon_points_available": 59,    "e_percent_to_aa": 14,    "gender": 20,    "anon": 5,    "pvp_status": 44,    "pvp_type": 36,    "lfg": 6,    "aa_points_spent_old": 21,    "suffix": "OyXBKDZpDKykffeQDHtskIigU",    "eye_color_1": 97,    "aa_points_spent": 4,    "group_leadership_exp": 84,    "cur_hp": 2,    "str": 99,    "hair_style": 70,    "wis": 19,    "tribute_points": 48,    "guild_auto_consent": 29,    "leadership_exp_on": 59,    "account_id": 14,    "y": 0.5008992,    "face": 57,    "drakkin_details": 50,    "raid_leadership_points": 68,    "sta": 64,    "pvp_current_kill_streak": 85,    "pvp_current_points": 52,    "title": "gfThYtjtjbCxgskaiVXCaydSE",    "z": 0.13876428,    "class": 54,    "gm": 81,    "hair_color": 97,    "mana": 60,    "ldon_points_guk": 48,    "xtargets": 91,    "name": "cxbeuAHBQmgWTsDodkYkFssok",    "zone_instance": 38,    "intoxication": 78,    "int": 2,    "pvp_2": 11,    "rest_timer": 26,    "air_remaining": 2,    "last_name": "mrJpGRBeSCmDLFFBshosCCGuO",    "raid_leadership_exp": 26,    "points": 22,    "ability_up": 4,    "ldon_points_mmc": 69,    "tribute_time_remaining": 36,    "pvp_best_kill_streak": 53,    "mailkey": "KXUEKoLoEdRIxUpOadgrlMrfN",    "firstlogon": 0,    "deleted_at": "2300-04-03T12:03:24.978063165-07:00",    "zone_id": 10,    "race": 45,    "birthday": 85,    "aa_exp": 46,    "thirst_level": 22,    "ldon_points_mir": 32,    "e_aa_effects": 28,    "time_played": 35,    "drakkin_tattoo": 27,    "dex": 4,    "ldon_points_tak": 62,    "raid_auto_consent": 32,    "e_last_invsnapshot": 1}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[12] column is set for unsigned
[13] column is set for unsigned
[14] column is set for unsigned
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
[47] column is set for unsigned
[48] column is set for unsigned
[49] column is set for unsigned
[50] column is set for unsigned
[51] column is set for unsigned
[52] column is set for unsigned
[53] column is set for unsigned
[54] column is set for unsigned
[55] column is set for unsigned
[56] column is set for unsigned
[57] column is set for unsigned
[58] column is set for unsigned
[59] column is set for unsigned
[60] column is set for unsigned
[61] column is set for unsigned
[62] column is set for unsigned
[63] column is set for unsigned
[64] column is set for unsigned
[65] column is set for unsigned
[66] column is set for unsigned
[67] column is set for unsigned
[68] column is set for unsigned
[69] column is set for unsigned
[70] column is set for unsigned
[71] column is set for unsigned
[72] column is set for unsigned
[73] column is set for unsigned
[74] column is set for unsigned
[75] column is set for unsigned
[76] column is set for unsigned
[77] column is set for unsigned
[78] column is set for unsigned
[79] column is set for unsigned
[80] column is set for unsigned
[81] column is set for unsigned
[82] column is set for unsigned
[83] column is set for unsigned
[84] column is set for unsigned
[85] column is set for unsigned
[86] column is set for unsigned
[87] column is set for unsigned
[88] column is set for unsigned
[89] column is set for unsigned
[90] column is set for unsigned
[91] column is set for unsigned
[93] column is set for unsigned
[95] column is set for unsigned
[96] column is set for unsigned
[97] column is set for unsigned
[98] column is set for unsigned
[99] column is set for unsigned
[100] column is set for unsigned



*/

// CharacterData struct is a row record of the character_data table in the eqemu database
type CharacterData struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] account_id                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AccountID int32 `gorm:"column:account_id;type:int;default:0;" json:"account_id"`
	//[ 2] name                                           varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Name string `gorm:"column:name;type:varchar;size:64;default:'';" json:"name"`
	//[ 3] last_name                                      varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	LastName string `gorm:"column:last_name;type:varchar;size:64;default:'';" json:"last_name"`
	//[ 4] title                                          varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['']
	Title string `gorm:"column:title;type:varchar;size:32;default:'';" json:"title"`
	//[ 5] suffix                                         varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['']
	Suffix string `gorm:"column:suffix;type:varchar;size:32;default:'';" json:"suffix"`
	//[ 6] zone_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ZoneID uint32 `gorm:"column:zone_id;type:uint;default:0;" json:"zone_id"`
	//[ 7] zone_instance                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ZoneInstance uint32 `gorm:"column:zone_instance;type:uint;default:0;" json:"zone_instance"`
	//[ 8] y                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Y float32 `gorm:"column:y;type:float;default:0;" json:"y"`
	//[ 9] x                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	X float32 `gorm:"column:x;type:float;default:0;" json:"x"`
	//[10] z                                              float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Z float32 `gorm:"column:z;type:float;default:0;" json:"z"`
	//[11] heading                                        float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Heading float32 `gorm:"column:heading;type:float;default:0;" json:"heading"`
	//[12] gender                                         utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Gender uint32 `gorm:"column:gender;type:utinyint;default:0;" json:"gender"`
	//[13] race                                           usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	Race uint32 `gorm:"column:race;type:usmallint;default:0;" json:"race"`
	//[14] class                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Class uint32 `gorm:"column:class;type:utinyint;default:0;" json:"class"`
	//[15] level                                          uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Level uint32 `gorm:"column:level;type:uint;default:0;" json:"level"`
	//[16] deity                                          uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Deity uint32 `gorm:"column:deity;type:uint;default:0;" json:"deity"`
	//[17] birthday                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Birthday uint32 `gorm:"column:birthday;type:uint;default:0;" json:"birthday"`
	//[18] last_login                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastLogin uint32 `gorm:"column:last_login;type:uint;default:0;" json:"last_login"`
	//[19] time_played                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TimePlayed uint32 `gorm:"column:time_played;type:uint;default:0;" json:"time_played"`
	//[20] level2                                         utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Level2 uint32 `gorm:"column:level2;type:utinyint;default:0;" json:"level_2"`
	//[21] anon                                           utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Anon uint32 `gorm:"column:anon;type:utinyint;default:0;" json:"anon"`
	//[22] gm                                             utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Gm uint32 `gorm:"column:gm;type:utinyint;default:0;" json:"gm"`
	//[23] face                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Face uint32 `gorm:"column:face;type:uint;default:0;" json:"face"`
	//[24] hair_color                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	HairColor uint32 `gorm:"column:hair_color;type:utinyint;default:0;" json:"hair_color"`
	//[25] hair_style                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	HairStyle uint32 `gorm:"column:hair_style;type:utinyint;default:0;" json:"hair_style"`
	//[26] beard                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Beard uint32 `gorm:"column:beard;type:utinyint;default:0;" json:"beard"`
	//[27] beard_color                                    utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	BeardColor uint32 `gorm:"column:beard_color;type:utinyint;default:0;" json:"beard_color"`
	//[28] eye_color_1                                    utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	EyeColor1 uint32 `gorm:"column:eye_color_1;type:utinyint;default:0;" json:"eye_color_1"`
	//[29] eye_color_2                                    utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	EyeColor2 uint32 `gorm:"column:eye_color_2;type:utinyint;default:0;" json:"eye_color_2"`
	//[30] drakkin_heritage                               uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	DrakkinHeritage uint32 `gorm:"column:drakkin_heritage;type:uint;default:0;" json:"drakkin_heritage"`
	//[31] drakkin_tattoo                                 uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	DrakkinTattoo uint32 `gorm:"column:drakkin_tattoo;type:uint;default:0;" json:"drakkin_tattoo"`
	//[32] drakkin_details                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	DrakkinDetails uint32 `gorm:"column:drakkin_details;type:uint;default:0;" json:"drakkin_details"`
	//[33] ability_time_seconds                           utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	AbilityTimeSeconds uint32 `gorm:"column:ability_time_seconds;type:utinyint;default:0;" json:"ability_time_seconds"`
	//[34] ability_number                                 utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	AbilityNumber uint32 `gorm:"column:ability_number;type:utinyint;default:0;" json:"ability_number"`
	//[35] ability_time_minutes                           utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	AbilityTimeMinutes uint32 `gorm:"column:ability_time_minutes;type:utinyint;default:0;" json:"ability_time_minutes"`
	//[36] ability_time_hours                             utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	AbilityTimeHours uint32 `gorm:"column:ability_time_hours;type:utinyint;default:0;" json:"ability_time_hours"`
	//[37] exp                                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Exp uint32 `gorm:"column:exp;type:uint;default:0;" json:"exp"`
	//[38] aa_points_spent                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AaPointsSpent uint32 `gorm:"column:aa_points_spent;type:uint;default:0;" json:"aa_points_spent"`
	//[39] aa_exp                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AaExp uint32 `gorm:"column:aa_exp;type:uint;default:0;" json:"aa_exp"`
	//[40] aa_points                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AaPoints uint32 `gorm:"column:aa_points;type:uint;default:0;" json:"aa_points"`
	//[41] group_leadership_exp                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	GroupLeadershipExp uint32 `gorm:"column:group_leadership_exp;type:uint;default:0;" json:"group_leadership_exp"`
	//[42] raid_leadership_exp                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	RaidLeadershipExp uint32 `gorm:"column:raid_leadership_exp;type:uint;default:0;" json:"raid_leadership_exp"`
	//[43] group_leadership_points                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	GroupLeadershipPoints uint32 `gorm:"column:group_leadership_points;type:uint;default:0;" json:"group_leadership_points"`
	//[44] raid_leadership_points                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	RaidLeadershipPoints uint32 `gorm:"column:raid_leadership_points;type:uint;default:0;" json:"raid_leadership_points"`
	//[45] points                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Points uint32 `gorm:"column:points;type:uint;default:0;" json:"points"`
	//[46] cur_hp                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CurHp uint32 `gorm:"column:cur_hp;type:uint;default:0;" json:"cur_hp"`
	//[47] mana                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Mana uint32 `gorm:"column:mana;type:uint;default:0;" json:"mana"`
	//[48] endurance                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Endurance uint32 `gorm:"column:endurance;type:uint;default:0;" json:"endurance"`
	//[49] intoxication                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Intoxication uint32 `gorm:"column:intoxication;type:uint;default:0;" json:"intoxication"`
	//[50] str                                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Str uint32 `gorm:"column:str;type:uint;default:0;" json:"str"`
	//[51] sta                                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Sta uint32 `gorm:"column:sta;type:uint;default:0;" json:"sta"`
	//[52] cha                                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Cha uint32 `gorm:"column:cha;type:uint;default:0;" json:"cha"`
	//[53] dex                                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Dex uint32 `gorm:"column:dex;type:uint;default:0;" json:"dex"`
	//[54] int                                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Int uint32 `gorm:"column:int;type:uint;default:0;" json:"int"`
	//[55] agi                                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Agi uint32 `gorm:"column:agi;type:uint;default:0;" json:"agi"`
	//[56] wis                                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Wis uint32 `gorm:"column:wis;type:uint;default:0;" json:"wis"`
	//[57] zone_change_count                              uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ZoneChangeCount uint32 `gorm:"column:zone_change_count;type:uint;default:0;" json:"zone_change_count"`
	//[58] toxicity                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Toxicity uint32 `gorm:"column:toxicity;type:uint;default:0;" json:"toxicity"`
	//[59] hunger_level                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	HungerLevel uint32 `gorm:"column:hunger_level;type:uint;default:0;" json:"hunger_level"`
	//[60] thirst_level                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ThirstLevel uint32 `gorm:"column:thirst_level;type:uint;default:0;" json:"thirst_level"`
	//[61] ability_up                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AbilityUp uint32 `gorm:"column:ability_up;type:uint;default:0;" json:"ability_up"`
	//[62] ldon_points_guk                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LdonPointsGuk uint32 `gorm:"column:ldon_points_guk;type:uint;default:0;" json:"ldon_points_guk"`
	//[63] ldon_points_mir                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LdonPointsMir uint32 `gorm:"column:ldon_points_mir;type:uint;default:0;" json:"ldon_points_mir"`
	//[64] ldon_points_mmc                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LdonPointsMmc uint32 `gorm:"column:ldon_points_mmc;type:uint;default:0;" json:"ldon_points_mmc"`
	//[65] ldon_points_ruj                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LdonPointsRuj uint32 `gorm:"column:ldon_points_ruj;type:uint;default:0;" json:"ldon_points_ruj"`
	//[66] ldon_points_tak                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LdonPointsTak uint32 `gorm:"column:ldon_points_tak;type:uint;default:0;" json:"ldon_points_tak"`
	//[67] ldon_points_available                          uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LdonPointsAvailable uint32 `gorm:"column:ldon_points_available;type:uint;default:0;" json:"ldon_points_available"`
	//[68] tribute_time_remaining                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TributeTimeRemaining uint32 `gorm:"column:tribute_time_remaining;type:uint;default:0;" json:"tribute_time_remaining"`
	//[69] career_tribute_points                          uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CareerTributePoints uint32 `gorm:"column:career_tribute_points;type:uint;default:0;" json:"career_tribute_points"`
	//[70] tribute_points                                 uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TributePoints uint32 `gorm:"column:tribute_points;type:uint;default:0;" json:"tribute_points"`
	//[71] tribute_active                                 uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TributeActive uint32 `gorm:"column:tribute_active;type:uint;default:0;" json:"tribute_active"`
	//[72] pvp_status                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	PvpStatus uint32 `gorm:"column:pvp_status;type:utinyint;default:0;" json:"pvp_status"`
	//[73] pvp_kills                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PvpKills uint32 `gorm:"column:pvp_kills;type:uint;default:0;" json:"pvp_kills"`
	//[74] pvp_deaths                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PvpDeaths uint32 `gorm:"column:pvp_deaths;type:uint;default:0;" json:"pvp_deaths"`
	//[75] pvp_current_points                             uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PvpCurrentPoints uint32 `gorm:"column:pvp_current_points;type:uint;default:0;" json:"pvp_current_points"`
	//[76] pvp_career_points                              uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PvpCareerPoints uint32 `gorm:"column:pvp_career_points;type:uint;default:0;" json:"pvp_career_points"`
	//[77] pvp_best_kill_streak                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PvpBestKillStreak uint32 `gorm:"column:pvp_best_kill_streak;type:uint;default:0;" json:"pvp_best_kill_streak"`
	//[78] pvp_worst_death_streak                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PvpWorstDeathStreak uint32 `gorm:"column:pvp_worst_death_streak;type:uint;default:0;" json:"pvp_worst_death_streak"`
	//[79] pvp_current_kill_streak                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PvpCurrentKillStreak uint32 `gorm:"column:pvp_current_kill_streak;type:uint;default:0;" json:"pvp_current_kill_streak"`
	//[80] pvp2                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Pvp2 uint32 `gorm:"column:pvp2;type:uint;default:0;" json:"pvp_2"`
	//[81] pvp_type                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PvpType uint32 `gorm:"column:pvp_type;type:uint;default:0;" json:"pvp_type"`
	//[82] show_helm                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ShowHelm uint32 `gorm:"column:show_helm;type:uint;default:0;" json:"show_helm"`
	//[83] group_auto_consent                             utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	GroupAutoConsent uint32 `gorm:"column:group_auto_consent;type:utinyint;default:0;" json:"group_auto_consent"`
	//[84] raid_auto_consent                              utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	RaidAutoConsent uint32 `gorm:"column:raid_auto_consent;type:utinyint;default:0;" json:"raid_auto_consent"`
	//[85] guild_auto_consent                             utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	GuildAutoConsent uint32 `gorm:"column:guild_auto_consent;type:utinyint;default:0;" json:"guild_auto_consent"`
	//[86] leadership_exp_on                              utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	LeadershipExpOn uint32 `gorm:"column:leadership_exp_on;type:utinyint;default:0;" json:"leadership_exp_on"`
	//[87] RestTimer                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	RestTimer uint32 `gorm:"column:RestTimer;type:uint;default:0;" json:"rest_timer"`
	//[88] air_remaining                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AirRemaining uint32 `gorm:"column:air_remaining;type:uint;default:0;" json:"air_remaining"`
	//[89] autosplit_enabled                              uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AutosplitEnabled uint32 `gorm:"column:autosplit_enabled;type:uint;default:0;" json:"autosplit_enabled"`
	//[90] lfp                                            utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Lfp uint32 `gorm:"column:lfp;type:utinyint;default:0;" json:"lfp"`
	//[91] lfg                                            utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Lfg uint32 `gorm:"column:lfg;type:utinyint;default:0;" json:"lfg"`
	//[92] mailkey                                        char(16)             null: false  primary: false  isArray: false  auto: false  col: char            len: 16      default: ['']
	Mailkey string `gorm:"column:mailkey;type:char;size:16;default:'';" json:"mailkey"`
	//[93] xtargets                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [5]
	Xtargets uint32 `gorm:"column:xtargets;type:utinyint;default:5;" json:"xtargets"`
	//[94] firstlogon                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Firstlogon int32 `gorm:"column:firstlogon;type:tinyint;default:0;" json:"firstlogon"`
	//[95] e_aa_effects                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EAaEffects uint32 `gorm:"column:e_aa_effects;type:uint;default:0;" json:"e_aa_effects"`
	//[96] e_percent_to_aa                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EPercentToAa uint32 `gorm:"column:e_percent_to_aa;type:uint;default:0;" json:"e_percent_to_aa"`
	//[97] e_expended_aa_spent                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EExpendedAaSpent uint32 `gorm:"column:e_expended_aa_spent;type:uint;default:0;" json:"e_expended_aa_spent"`
	//[98] aa_points_spent_old                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AaPointsSpentOld uint32 `gorm:"column:aa_points_spent_old;type:uint;default:0;" json:"aa_points_spent_old"`
	//[99] aa_points_old                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AaPointsOld uint32 `gorm:"column:aa_points_old;type:uint;default:0;" json:"aa_points_old"`
	//[100] e_last_invsnapshot                             uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ELastInvsnapshot uint32 `gorm:"column:e_last_invsnapshot;type:uint;default:0;" json:"e_last_invsnapshot"`
	//[101] deleted_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [NULL]
	DeletedAt null.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
}

var character_dataTableInfo = &TableInfo{
	Name: "character_data",
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
			Name:               "account_id",
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
			GoFieldName:        "AccountID",
			GoFieldType:        "int32",
			JSONFieldName:      "account_id",
			ProtobufFieldName:  "account_id",
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
			Name:               "last_name",
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
			GoFieldName:        "LastName",
			GoFieldType:        "string",
			JSONFieldName:      "last_name",
			ProtobufFieldName:  "last_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "title",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Title",
			GoFieldType:        "string",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "suffix",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Suffix",
			GoFieldType:        "string",
			JSONFieldName:      "suffix",
			ProtobufFieldName:  "suffix",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "zone_id",
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
			GoFieldName:        "ZoneID",
			GoFieldType:        "uint32",
			JSONFieldName:      "zone_id",
			ProtobufFieldName:  "zone_id",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "zone_instance",
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
			GoFieldName:        "ZoneInstance",
			GoFieldType:        "uint32",
			JSONFieldName:      "zone_instance",
			ProtobufFieldName:  "zone_instance",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "gender",
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
			GoFieldName:        "Gender",
			GoFieldType:        "uint32",
			JSONFieldName:      "gender",
			ProtobufFieldName:  "gender",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "race",
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
			GoFieldName:        "Race",
			GoFieldType:        "uint32",
			JSONFieldName:      "race",
			ProtobufFieldName:  "race",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "class",
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
			GoFieldName:        "Class",
			GoFieldType:        "uint32",
			JSONFieldName:      "class",
			ProtobufFieldName:  "class",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "level",
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
			GoFieldName:        "Level",
			GoFieldType:        "uint32",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "deity",
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
			GoFieldName:        "Deity",
			GoFieldType:        "uint32",
			JSONFieldName:      "deity",
			ProtobufFieldName:  "deity",
			ProtobufType:       "uint32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "birthday",
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
			GoFieldName:        "Birthday",
			GoFieldType:        "uint32",
			JSONFieldName:      "birthday",
			ProtobufFieldName:  "birthday",
			ProtobufType:       "uint32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "last_login",
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
			GoFieldName:        "LastLogin",
			GoFieldType:        "uint32",
			JSONFieldName:      "last_login",
			ProtobufFieldName:  "last_login",
			ProtobufType:       "uint32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "time_played",
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
			GoFieldName:        "TimePlayed",
			GoFieldType:        "uint32",
			JSONFieldName:      "time_played",
			ProtobufFieldName:  "time_played",
			ProtobufType:       "uint32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "level2",
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
			GoFieldName:        "Level2",
			GoFieldType:        "uint32",
			JSONFieldName:      "level_2",
			ProtobufFieldName:  "level_2",
			ProtobufType:       "uint32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "anon",
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
			GoFieldName:        "Anon",
			GoFieldType:        "uint32",
			JSONFieldName:      "anon",
			ProtobufFieldName:  "anon",
			ProtobufType:       "uint32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "gm",
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
			GoFieldName:        "Gm",
			GoFieldType:        "uint32",
			JSONFieldName:      "gm",
			ProtobufFieldName:  "gm",
			ProtobufType:       "uint32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "face",
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
			GoFieldName:        "Face",
			GoFieldType:        "uint32",
			JSONFieldName:      "face",
			ProtobufFieldName:  "face",
			ProtobufType:       "uint32",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "hair_color",
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
			GoFieldName:        "HairColor",
			GoFieldType:        "uint32",
			JSONFieldName:      "hair_color",
			ProtobufFieldName:  "hair_color",
			ProtobufType:       "uint32",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "hair_style",
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
			GoFieldName:        "HairStyle",
			GoFieldType:        "uint32",
			JSONFieldName:      "hair_style",
			ProtobufFieldName:  "hair_style",
			ProtobufType:       "uint32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "beard",
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
			GoFieldName:        "Beard",
			GoFieldType:        "uint32",
			JSONFieldName:      "beard",
			ProtobufFieldName:  "beard",
			ProtobufType:       "uint32",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "beard_color",
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
			GoFieldName:        "BeardColor",
			GoFieldType:        "uint32",
			JSONFieldName:      "beard_color",
			ProtobufFieldName:  "beard_color",
			ProtobufType:       "uint32",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "eye_color_1",
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
			GoFieldName:        "EyeColor1",
			GoFieldType:        "uint32",
			JSONFieldName:      "eye_color_1",
			ProtobufFieldName:  "eye_color_1",
			ProtobufType:       "uint32",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "eye_color_2",
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
			GoFieldName:        "EyeColor2",
			GoFieldType:        "uint32",
			JSONFieldName:      "eye_color_2",
			ProtobufFieldName:  "eye_color_2",
			ProtobufType:       "uint32",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "drakkin_heritage",
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
			GoFieldName:        "DrakkinHeritage",
			GoFieldType:        "uint32",
			JSONFieldName:      "drakkin_heritage",
			ProtobufFieldName:  "drakkin_heritage",
			ProtobufType:       "uint32",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "drakkin_tattoo",
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
			GoFieldName:        "DrakkinTattoo",
			GoFieldType:        "uint32",
			JSONFieldName:      "drakkin_tattoo",
			ProtobufFieldName:  "drakkin_tattoo",
			ProtobufType:       "uint32",
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "drakkin_details",
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
			GoFieldName:        "DrakkinDetails",
			GoFieldType:        "uint32",
			JSONFieldName:      "drakkin_details",
			ProtobufFieldName:  "drakkin_details",
			ProtobufType:       "uint32",
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
			Name:               "ability_time_seconds",
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
			GoFieldName:        "AbilityTimeSeconds",
			GoFieldType:        "uint32",
			JSONFieldName:      "ability_time_seconds",
			ProtobufFieldName:  "ability_time_seconds",
			ProtobufType:       "uint32",
			ProtobufPos:        34,
		},

		&ColumnInfo{
			Index:              34,
			Name:               "ability_number",
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
			GoFieldName:        "AbilityNumber",
			GoFieldType:        "uint32",
			JSONFieldName:      "ability_number",
			ProtobufFieldName:  "ability_number",
			ProtobufType:       "uint32",
			ProtobufPos:        35,
		},

		&ColumnInfo{
			Index:              35,
			Name:               "ability_time_minutes",
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
			GoFieldName:        "AbilityTimeMinutes",
			GoFieldType:        "uint32",
			JSONFieldName:      "ability_time_minutes",
			ProtobufFieldName:  "ability_time_minutes",
			ProtobufType:       "uint32",
			ProtobufPos:        36,
		},

		&ColumnInfo{
			Index:              36,
			Name:               "ability_time_hours",
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
			GoFieldName:        "AbilityTimeHours",
			GoFieldType:        "uint32",
			JSONFieldName:      "ability_time_hours",
			ProtobufFieldName:  "ability_time_hours",
			ProtobufType:       "uint32",
			ProtobufPos:        37,
		},

		&ColumnInfo{
			Index:              37,
			Name:               "exp",
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
			GoFieldName:        "Exp",
			GoFieldType:        "uint32",
			JSONFieldName:      "exp",
			ProtobufFieldName:  "exp",
			ProtobufType:       "uint32",
			ProtobufPos:        38,
		},

		&ColumnInfo{
			Index:              38,
			Name:               "aa_points_spent",
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
			GoFieldName:        "AaPointsSpent",
			GoFieldType:        "uint32",
			JSONFieldName:      "aa_points_spent",
			ProtobufFieldName:  "aa_points_spent",
			ProtobufType:       "uint32",
			ProtobufPos:        39,
		},

		&ColumnInfo{
			Index:              39,
			Name:               "aa_exp",
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
			GoFieldName:        "AaExp",
			GoFieldType:        "uint32",
			JSONFieldName:      "aa_exp",
			ProtobufFieldName:  "aa_exp",
			ProtobufType:       "uint32",
			ProtobufPos:        40,
		},

		&ColumnInfo{
			Index:              40,
			Name:               "aa_points",
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
			GoFieldName:        "AaPoints",
			GoFieldType:        "uint32",
			JSONFieldName:      "aa_points",
			ProtobufFieldName:  "aa_points",
			ProtobufType:       "uint32",
			ProtobufPos:        41,
		},

		&ColumnInfo{
			Index:              41,
			Name:               "group_leadership_exp",
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
			GoFieldName:        "GroupLeadershipExp",
			GoFieldType:        "uint32",
			JSONFieldName:      "group_leadership_exp",
			ProtobufFieldName:  "group_leadership_exp",
			ProtobufType:       "uint32",
			ProtobufPos:        42,
		},

		&ColumnInfo{
			Index:              42,
			Name:               "raid_leadership_exp",
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
			GoFieldName:        "RaidLeadershipExp",
			GoFieldType:        "uint32",
			JSONFieldName:      "raid_leadership_exp",
			ProtobufFieldName:  "raid_leadership_exp",
			ProtobufType:       "uint32",
			ProtobufPos:        43,
		},

		&ColumnInfo{
			Index:              43,
			Name:               "group_leadership_points",
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
			GoFieldName:        "GroupLeadershipPoints",
			GoFieldType:        "uint32",
			JSONFieldName:      "group_leadership_points",
			ProtobufFieldName:  "group_leadership_points",
			ProtobufType:       "uint32",
			ProtobufPos:        44,
		},

		&ColumnInfo{
			Index:              44,
			Name:               "raid_leadership_points",
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
			GoFieldName:        "RaidLeadershipPoints",
			GoFieldType:        "uint32",
			JSONFieldName:      "raid_leadership_points",
			ProtobufFieldName:  "raid_leadership_points",
			ProtobufType:       "uint32",
			ProtobufPos:        45,
		},

		&ColumnInfo{
			Index:              45,
			Name:               "points",
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
			GoFieldName:        "Points",
			GoFieldType:        "uint32",
			JSONFieldName:      "points",
			ProtobufFieldName:  "points",
			ProtobufType:       "uint32",
			ProtobufPos:        46,
		},

		&ColumnInfo{
			Index:              46,
			Name:               "cur_hp",
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
			GoFieldName:        "CurHp",
			GoFieldType:        "uint32",
			JSONFieldName:      "cur_hp",
			ProtobufFieldName:  "cur_hp",
			ProtobufType:       "uint32",
			ProtobufPos:        47,
		},

		&ColumnInfo{
			Index:              47,
			Name:               "mana",
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
			GoFieldName:        "Mana",
			GoFieldType:        "uint32",
			JSONFieldName:      "mana",
			ProtobufFieldName:  "mana",
			ProtobufType:       "uint32",
			ProtobufPos:        48,
		},

		&ColumnInfo{
			Index:              48,
			Name:               "endurance",
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
			GoFieldName:        "Endurance",
			GoFieldType:        "uint32",
			JSONFieldName:      "endurance",
			ProtobufFieldName:  "endurance",
			ProtobufType:       "uint32",
			ProtobufPos:        49,
		},

		&ColumnInfo{
			Index:              49,
			Name:               "intoxication",
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
			GoFieldName:        "Intoxication",
			GoFieldType:        "uint32",
			JSONFieldName:      "intoxication",
			ProtobufFieldName:  "intoxication",
			ProtobufType:       "uint32",
			ProtobufPos:        50,
		},

		&ColumnInfo{
			Index:              50,
			Name:               "str",
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
			GoFieldName:        "Str",
			GoFieldType:        "uint32",
			JSONFieldName:      "str",
			ProtobufFieldName:  "str",
			ProtobufType:       "uint32",
			ProtobufPos:        51,
		},

		&ColumnInfo{
			Index:              51,
			Name:               "sta",
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
			GoFieldName:        "Sta",
			GoFieldType:        "uint32",
			JSONFieldName:      "sta",
			ProtobufFieldName:  "sta",
			ProtobufType:       "uint32",
			ProtobufPos:        52,
		},

		&ColumnInfo{
			Index:              52,
			Name:               "cha",
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
			GoFieldName:        "Cha",
			GoFieldType:        "uint32",
			JSONFieldName:      "cha",
			ProtobufFieldName:  "cha",
			ProtobufType:       "uint32",
			ProtobufPos:        53,
		},

		&ColumnInfo{
			Index:              53,
			Name:               "dex",
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
			GoFieldName:        "Dex",
			GoFieldType:        "uint32",
			JSONFieldName:      "dex",
			ProtobufFieldName:  "dex",
			ProtobufType:       "uint32",
			ProtobufPos:        54,
		},

		&ColumnInfo{
			Index:              54,
			Name:               "int",
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
			GoFieldName:        "Int",
			GoFieldType:        "uint32",
			JSONFieldName:      "int",
			ProtobufFieldName:  "int",
			ProtobufType:       "uint32",
			ProtobufPos:        55,
		},

		&ColumnInfo{
			Index:              55,
			Name:               "agi",
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
			GoFieldName:        "Agi",
			GoFieldType:        "uint32",
			JSONFieldName:      "agi",
			ProtobufFieldName:  "agi",
			ProtobufType:       "uint32",
			ProtobufPos:        56,
		},

		&ColumnInfo{
			Index:              56,
			Name:               "wis",
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
			GoFieldName:        "Wis",
			GoFieldType:        "uint32",
			JSONFieldName:      "wis",
			ProtobufFieldName:  "wis",
			ProtobufType:       "uint32",
			ProtobufPos:        57,
		},

		&ColumnInfo{
			Index:              57,
			Name:               "zone_change_count",
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
			GoFieldName:        "ZoneChangeCount",
			GoFieldType:        "uint32",
			JSONFieldName:      "zone_change_count",
			ProtobufFieldName:  "zone_change_count",
			ProtobufType:       "uint32",
			ProtobufPos:        58,
		},

		&ColumnInfo{
			Index:              58,
			Name:               "toxicity",
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
			GoFieldName:        "Toxicity",
			GoFieldType:        "uint32",
			JSONFieldName:      "toxicity",
			ProtobufFieldName:  "toxicity",
			ProtobufType:       "uint32",
			ProtobufPos:        59,
		},

		&ColumnInfo{
			Index:              59,
			Name:               "hunger_level",
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
			GoFieldName:        "HungerLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "hunger_level",
			ProtobufFieldName:  "hunger_level",
			ProtobufType:       "uint32",
			ProtobufPos:        60,
		},

		&ColumnInfo{
			Index:              60,
			Name:               "thirst_level",
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
			GoFieldName:        "ThirstLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "thirst_level",
			ProtobufFieldName:  "thirst_level",
			ProtobufType:       "uint32",
			ProtobufPos:        61,
		},

		&ColumnInfo{
			Index:              61,
			Name:               "ability_up",
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
			GoFieldName:        "AbilityUp",
			GoFieldType:        "uint32",
			JSONFieldName:      "ability_up",
			ProtobufFieldName:  "ability_up",
			ProtobufType:       "uint32",
			ProtobufPos:        62,
		},

		&ColumnInfo{
			Index:              62,
			Name:               "ldon_points_guk",
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
			GoFieldName:        "LdonPointsGuk",
			GoFieldType:        "uint32",
			JSONFieldName:      "ldon_points_guk",
			ProtobufFieldName:  "ldon_points_guk",
			ProtobufType:       "uint32",
			ProtobufPos:        63,
		},

		&ColumnInfo{
			Index:              63,
			Name:               "ldon_points_mir",
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
			GoFieldName:        "LdonPointsMir",
			GoFieldType:        "uint32",
			JSONFieldName:      "ldon_points_mir",
			ProtobufFieldName:  "ldon_points_mir",
			ProtobufType:       "uint32",
			ProtobufPos:        64,
		},

		&ColumnInfo{
			Index:              64,
			Name:               "ldon_points_mmc",
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
			GoFieldName:        "LdonPointsMmc",
			GoFieldType:        "uint32",
			JSONFieldName:      "ldon_points_mmc",
			ProtobufFieldName:  "ldon_points_mmc",
			ProtobufType:       "uint32",
			ProtobufPos:        65,
		},

		&ColumnInfo{
			Index:              65,
			Name:               "ldon_points_ruj",
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
			GoFieldName:        "LdonPointsRuj",
			GoFieldType:        "uint32",
			JSONFieldName:      "ldon_points_ruj",
			ProtobufFieldName:  "ldon_points_ruj",
			ProtobufType:       "uint32",
			ProtobufPos:        66,
		},

		&ColumnInfo{
			Index:              66,
			Name:               "ldon_points_tak",
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
			GoFieldName:        "LdonPointsTak",
			GoFieldType:        "uint32",
			JSONFieldName:      "ldon_points_tak",
			ProtobufFieldName:  "ldon_points_tak",
			ProtobufType:       "uint32",
			ProtobufPos:        67,
		},

		&ColumnInfo{
			Index:              67,
			Name:               "ldon_points_available",
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
			GoFieldName:        "LdonPointsAvailable",
			GoFieldType:        "uint32",
			JSONFieldName:      "ldon_points_available",
			ProtobufFieldName:  "ldon_points_available",
			ProtobufType:       "uint32",
			ProtobufPos:        68,
		},

		&ColumnInfo{
			Index:              68,
			Name:               "tribute_time_remaining",
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
			GoFieldName:        "TributeTimeRemaining",
			GoFieldType:        "uint32",
			JSONFieldName:      "tribute_time_remaining",
			ProtobufFieldName:  "tribute_time_remaining",
			ProtobufType:       "uint32",
			ProtobufPos:        69,
		},

		&ColumnInfo{
			Index:              69,
			Name:               "career_tribute_points",
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
			GoFieldName:        "CareerTributePoints",
			GoFieldType:        "uint32",
			JSONFieldName:      "career_tribute_points",
			ProtobufFieldName:  "career_tribute_points",
			ProtobufType:       "uint32",
			ProtobufPos:        70,
		},

		&ColumnInfo{
			Index:              70,
			Name:               "tribute_points",
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
			GoFieldName:        "TributePoints",
			GoFieldType:        "uint32",
			JSONFieldName:      "tribute_points",
			ProtobufFieldName:  "tribute_points",
			ProtobufType:       "uint32",
			ProtobufPos:        71,
		},

		&ColumnInfo{
			Index:              71,
			Name:               "tribute_active",
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
			GoFieldName:        "TributeActive",
			GoFieldType:        "uint32",
			JSONFieldName:      "tribute_active",
			ProtobufFieldName:  "tribute_active",
			ProtobufType:       "uint32",
			ProtobufPos:        72,
		},

		&ColumnInfo{
			Index:              72,
			Name:               "pvp_status",
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
			GoFieldName:        "PvpStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "pvp_status",
			ProtobufFieldName:  "pvp_status",
			ProtobufType:       "uint32",
			ProtobufPos:        73,
		},

		&ColumnInfo{
			Index:              73,
			Name:               "pvp_kills",
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
			GoFieldName:        "PvpKills",
			GoFieldType:        "uint32",
			JSONFieldName:      "pvp_kills",
			ProtobufFieldName:  "pvp_kills",
			ProtobufType:       "uint32",
			ProtobufPos:        74,
		},

		&ColumnInfo{
			Index:              74,
			Name:               "pvp_deaths",
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
			GoFieldName:        "PvpDeaths",
			GoFieldType:        "uint32",
			JSONFieldName:      "pvp_deaths",
			ProtobufFieldName:  "pvp_deaths",
			ProtobufType:       "uint32",
			ProtobufPos:        75,
		},

		&ColumnInfo{
			Index:              75,
			Name:               "pvp_current_points",
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
			GoFieldName:        "PvpCurrentPoints",
			GoFieldType:        "uint32",
			JSONFieldName:      "pvp_current_points",
			ProtobufFieldName:  "pvp_current_points",
			ProtobufType:       "uint32",
			ProtobufPos:        76,
		},

		&ColumnInfo{
			Index:              76,
			Name:               "pvp_career_points",
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
			GoFieldName:        "PvpCareerPoints",
			GoFieldType:        "uint32",
			JSONFieldName:      "pvp_career_points",
			ProtobufFieldName:  "pvp_career_points",
			ProtobufType:       "uint32",
			ProtobufPos:        77,
		},

		&ColumnInfo{
			Index:              77,
			Name:               "pvp_best_kill_streak",
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
			GoFieldName:        "PvpBestKillStreak",
			GoFieldType:        "uint32",
			JSONFieldName:      "pvp_best_kill_streak",
			ProtobufFieldName:  "pvp_best_kill_streak",
			ProtobufType:       "uint32",
			ProtobufPos:        78,
		},

		&ColumnInfo{
			Index:              78,
			Name:               "pvp_worst_death_streak",
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
			GoFieldName:        "PvpWorstDeathStreak",
			GoFieldType:        "uint32",
			JSONFieldName:      "pvp_worst_death_streak",
			ProtobufFieldName:  "pvp_worst_death_streak",
			ProtobufType:       "uint32",
			ProtobufPos:        79,
		},

		&ColumnInfo{
			Index:              79,
			Name:               "pvp_current_kill_streak",
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
			GoFieldName:        "PvpCurrentKillStreak",
			GoFieldType:        "uint32",
			JSONFieldName:      "pvp_current_kill_streak",
			ProtobufFieldName:  "pvp_current_kill_streak",
			ProtobufType:       "uint32",
			ProtobufPos:        80,
		},

		&ColumnInfo{
			Index:              80,
			Name:               "pvp2",
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
			GoFieldName:        "Pvp2",
			GoFieldType:        "uint32",
			JSONFieldName:      "pvp_2",
			ProtobufFieldName:  "pvp_2",
			ProtobufType:       "uint32",
			ProtobufPos:        81,
		},

		&ColumnInfo{
			Index:              81,
			Name:               "pvp_type",
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
			GoFieldName:        "PvpType",
			GoFieldType:        "uint32",
			JSONFieldName:      "pvp_type",
			ProtobufFieldName:  "pvp_type",
			ProtobufType:       "uint32",
			ProtobufPos:        82,
		},

		&ColumnInfo{
			Index:              82,
			Name:               "show_helm",
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
			GoFieldName:        "ShowHelm",
			GoFieldType:        "uint32",
			JSONFieldName:      "show_helm",
			ProtobufFieldName:  "show_helm",
			ProtobufType:       "uint32",
			ProtobufPos:        83,
		},

		&ColumnInfo{
			Index:              83,
			Name:               "group_auto_consent",
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
			GoFieldName:        "GroupAutoConsent",
			GoFieldType:        "uint32",
			JSONFieldName:      "group_auto_consent",
			ProtobufFieldName:  "group_auto_consent",
			ProtobufType:       "uint32",
			ProtobufPos:        84,
		},

		&ColumnInfo{
			Index:              84,
			Name:               "raid_auto_consent",
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
			GoFieldName:        "RaidAutoConsent",
			GoFieldType:        "uint32",
			JSONFieldName:      "raid_auto_consent",
			ProtobufFieldName:  "raid_auto_consent",
			ProtobufType:       "uint32",
			ProtobufPos:        85,
		},

		&ColumnInfo{
			Index:              85,
			Name:               "guild_auto_consent",
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
			GoFieldName:        "GuildAutoConsent",
			GoFieldType:        "uint32",
			JSONFieldName:      "guild_auto_consent",
			ProtobufFieldName:  "guild_auto_consent",
			ProtobufType:       "uint32",
			ProtobufPos:        86,
		},

		&ColumnInfo{
			Index:              86,
			Name:               "leadership_exp_on",
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
			GoFieldName:        "LeadershipExpOn",
			GoFieldType:        "uint32",
			JSONFieldName:      "leadership_exp_on",
			ProtobufFieldName:  "leadership_exp_on",
			ProtobufType:       "uint32",
			ProtobufPos:        87,
		},

		&ColumnInfo{
			Index:              87,
			Name:               "RestTimer",
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
			GoFieldName:        "RestTimer",
			GoFieldType:        "uint32",
			JSONFieldName:      "rest_timer",
			ProtobufFieldName:  "rest_timer",
			ProtobufType:       "uint32",
			ProtobufPos:        88,
		},

		&ColumnInfo{
			Index:              88,
			Name:               "air_remaining",
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
			GoFieldName:        "AirRemaining",
			GoFieldType:        "uint32",
			JSONFieldName:      "air_remaining",
			ProtobufFieldName:  "air_remaining",
			ProtobufType:       "uint32",
			ProtobufPos:        89,
		},

		&ColumnInfo{
			Index:              89,
			Name:               "autosplit_enabled",
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
			GoFieldName:        "AutosplitEnabled",
			GoFieldType:        "uint32",
			JSONFieldName:      "autosplit_enabled",
			ProtobufFieldName:  "autosplit_enabled",
			ProtobufType:       "uint32",
			ProtobufPos:        90,
		},

		&ColumnInfo{
			Index:              90,
			Name:               "lfp",
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
			GoFieldName:        "Lfp",
			GoFieldType:        "uint32",
			JSONFieldName:      "lfp",
			ProtobufFieldName:  "lfp",
			ProtobufType:       "uint32",
			ProtobufPos:        91,
		},

		&ColumnInfo{
			Index:              91,
			Name:               "lfg",
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
			GoFieldName:        "Lfg",
			GoFieldType:        "uint32",
			JSONFieldName:      "lfg",
			ProtobufFieldName:  "lfg",
			ProtobufType:       "uint32",
			ProtobufPos:        92,
		},

		&ColumnInfo{
			Index:              92,
			Name:               "mailkey",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(16)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       16,
			GoFieldName:        "Mailkey",
			GoFieldType:        "string",
			JSONFieldName:      "mailkey",
			ProtobufFieldName:  "mailkey",
			ProtobufType:       "string",
			ProtobufPos:        93,
		},

		&ColumnInfo{
			Index:              93,
			Name:               "xtargets",
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
			GoFieldName:        "Xtargets",
			GoFieldType:        "uint32",
			JSONFieldName:      "xtargets",
			ProtobufFieldName:  "xtargets",
			ProtobufType:       "uint32",
			ProtobufPos:        94,
		},

		&ColumnInfo{
			Index:              94,
			Name:               "firstlogon",
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
			GoFieldName:        "Firstlogon",
			GoFieldType:        "int32",
			JSONFieldName:      "firstlogon",
			ProtobufFieldName:  "firstlogon",
			ProtobufType:       "int32",
			ProtobufPos:        95,
		},

		&ColumnInfo{
			Index:              95,
			Name:               "e_aa_effects",
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
			GoFieldName:        "EAaEffects",
			GoFieldType:        "uint32",
			JSONFieldName:      "e_aa_effects",
			ProtobufFieldName:  "e_aa_effects",
			ProtobufType:       "uint32",
			ProtobufPos:        96,
		},

		&ColumnInfo{
			Index:              96,
			Name:               "e_percent_to_aa",
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
			GoFieldName:        "EPercentToAa",
			GoFieldType:        "uint32",
			JSONFieldName:      "e_percent_to_aa",
			ProtobufFieldName:  "e_percent_to_aa",
			ProtobufType:       "uint32",
			ProtobufPos:        97,
		},

		&ColumnInfo{
			Index:              97,
			Name:               "e_expended_aa_spent",
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
			GoFieldName:        "EExpendedAaSpent",
			GoFieldType:        "uint32",
			JSONFieldName:      "e_expended_aa_spent",
			ProtobufFieldName:  "e_expended_aa_spent",
			ProtobufType:       "uint32",
			ProtobufPos:        98,
		},

		&ColumnInfo{
			Index:              98,
			Name:               "aa_points_spent_old",
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
			GoFieldName:        "AaPointsSpentOld",
			GoFieldType:        "uint32",
			JSONFieldName:      "aa_points_spent_old",
			ProtobufFieldName:  "aa_points_spent_old",
			ProtobufType:       "uint32",
			ProtobufPos:        99,
		},

		&ColumnInfo{
			Index:              99,
			Name:               "aa_points_old",
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
			GoFieldName:        "AaPointsOld",
			GoFieldType:        "uint32",
			JSONFieldName:      "aa_points_old",
			ProtobufFieldName:  "aa_points_old",
			ProtobufType:       "uint32",
			ProtobufPos:        100,
		},

		&ColumnInfo{
			Index:              100,
			Name:               "e_last_invsnapshot",
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
			GoFieldName:        "ELastInvsnapshot",
			GoFieldType:        "uint32",
			JSONFieldName:      "e_last_invsnapshot",
			ProtobufFieldName:  "e_last_invsnapshot",
			ProtobufType:       "uint32",
			ProtobufPos:        101,
		},

		&ColumnInfo{
			Index:              101,
			Name:               "deleted_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "DeletedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "deleted_at",
			ProtobufFieldName:  "deleted_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        102,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterData) TableName() string {
	return "character_data"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterData) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterData) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterData) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterData) TableInfo() *TableInfo {
	return character_dataTableInfo
}
