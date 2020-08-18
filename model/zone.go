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


CREATE TABLE `zone` (
  `short_name` varchar(32) DEFAULT NULL,
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `file_name` varchar(16) DEFAULT NULL,
  `long_name` text NOT NULL,
  `map_file_name` varchar(100) DEFAULT NULL,
  `safe_x` float NOT NULL DEFAULT 0,
  `safe_y` float NOT NULL DEFAULT 0,
  `safe_z` float NOT NULL DEFAULT 0,
  `graveyard_id` float NOT NULL DEFAULT 0,
  `min_level` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `min_status` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `zoneidnumber` int(4) NOT NULL DEFAULT 0,
  `version` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `timezone` int(5) NOT NULL DEFAULT 0,
  `maxclients` int(5) NOT NULL DEFAULT 0,
  `ruleset` int(10) unsigned NOT NULL DEFAULT 0,
  `note` varchar(80) DEFAULT NULL,
  `underworld` float NOT NULL DEFAULT 0,
  `minclip` float NOT NULL DEFAULT 450,
  `maxclip` float NOT NULL DEFAULT 450,
  `fog_minclip` float NOT NULL DEFAULT 450,
  `fog_maxclip` float NOT NULL DEFAULT 450,
  `fog_blue` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `fog_red` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `fog_green` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `sky` tinyint(3) unsigned NOT NULL DEFAULT 1,
  `ztype` tinyint(3) unsigned NOT NULL DEFAULT 1,
  `zone_exp_multiplier` decimal(6,2) NOT NULL DEFAULT 0.00,
  `walkspeed` float NOT NULL DEFAULT 0.4,
  `time_type` tinyint(3) unsigned NOT NULL DEFAULT 2,
  `fog_red1` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `fog_green1` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `fog_blue1` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `fog_minclip1` float NOT NULL DEFAULT 450,
  `fog_maxclip1` float NOT NULL DEFAULT 450,
  `fog_red2` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `fog_green2` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `fog_blue2` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `fog_minclip2` float NOT NULL DEFAULT 450,
  `fog_maxclip2` float NOT NULL DEFAULT 450,
  `fog_red3` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `fog_green3` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `fog_blue3` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `fog_minclip3` float NOT NULL DEFAULT 450,
  `fog_maxclip3` float NOT NULL DEFAULT 450,
  `fog_red4` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `fog_green4` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `fog_blue4` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `fog_minclip4` float NOT NULL DEFAULT 450,
  `fog_maxclip4` float NOT NULL DEFAULT 450,
  `fog_density` float NOT NULL DEFAULT 0,
  `flag_needed` varchar(128) NOT NULL DEFAULT '',
  `canbind` tinyint(4) NOT NULL DEFAULT 1,
  `cancombat` tinyint(4) NOT NULL DEFAULT 1,
  `canlevitate` tinyint(4) NOT NULL DEFAULT 1,
  `castoutdoor` tinyint(4) NOT NULL DEFAULT 1,
  `hotzone` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `insttype` tinyint(1) unsigned zerofill NOT NULL DEFAULT 0,
  `shutdowndelay` bigint(16) unsigned NOT NULL DEFAULT 5000,
  `peqzone` tinyint(4) NOT NULL DEFAULT 1,
  `expansion` tinyint(3) NOT NULL DEFAULT 0,
  `suspendbuffs` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `rain_chance1` int(4) NOT NULL DEFAULT 0,
  `rain_chance2` int(4) NOT NULL DEFAULT 0,
  `rain_chance3` int(4) NOT NULL DEFAULT 0,
  `rain_chance4` int(4) NOT NULL DEFAULT 0,
  `rain_duration1` int(4) NOT NULL DEFAULT 0,
  `rain_duration2` int(4) NOT NULL DEFAULT 0,
  `rain_duration3` int(4) NOT NULL DEFAULT 0,
  `rain_duration4` int(4) NOT NULL DEFAULT 0,
  `snow_chance1` int(4) NOT NULL DEFAULT 0,
  `snow_chance2` int(4) NOT NULL DEFAULT 0,
  `snow_chance3` int(4) NOT NULL DEFAULT 0,
  `snow_chance4` int(4) NOT NULL DEFAULT 0,
  `snow_duration1` int(4) NOT NULL DEFAULT 0,
  `snow_duration2` int(4) NOT NULL DEFAULT 0,
  `snow_duration3` int(4) NOT NULL DEFAULT 0,
  `snow_duration4` int(4) NOT NULL DEFAULT 0,
  `gravity` float NOT NULL DEFAULT 0.4,
  `type` int(3) NOT NULL DEFAULT 0,
  `skylock` tinyint(4) NOT NULL DEFAULT 0,
  `fast_regen_hp` int(11) NOT NULL DEFAULT 180,
  `fast_regen_mana` int(11) NOT NULL DEFAULT 180,
  `fast_regen_endurance` int(11) NOT NULL DEFAULT 180,
  `npc_max_aggro_dist` int(11) NOT NULL DEFAULT 600,
  `max_movement_update_range` int(11) unsigned NOT NULL DEFAULT 600,
  `min_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `max_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `content_flags` varchar(100) DEFAULT NULL,
  `content_flags_disabled` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `zoneidnumber` (`zoneidnumber`),
  KEY `zonename` (`short_name`)
) ENGINE=InnoDB AUTO_INCREMENT=6019 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "hotzone": 52,    "rain_chance_4": 10,    "rain_duration_3": 21,    "snow_chance_3": 6,    "fog_minclip": 0.8549852,    "fog_green": 14,    "fog_maxclip_3": 0.5159846,    "rain_duration_1": 79,    "snow_duration_3": 72,    "ztype": 71,    "fog_blue_3": 3,    "fog_minclip_3": 0.9724632,    "snow_chance_1": 87,    "id": 77,    "safe_y": 0.9503453,    "maxclients": 8,    "fog_red_1": 23,    "snow_duration_1": 41,    "shutdowndelay": 35,    "long_name": "dkACbAUHubYKqUYSXVWEZASUp",    "safe_x": 0.3764996,    "min_status": 40,    "ruleset": 30,    "sky": 0,    "fog_red_2": 18,    "fog_density": 0.6005665,    "min_level": 36,    "fog_green_3": 5,    "fast_regen_hp": 31,    "skylock": 65,    "file_name": "VnXYipbGIZoKVjwHWkXcRQRSL",    "version": 57,    "timezone": 34,    "fog_green_4": 81,    "canlevitate": 41,    "rain_duration_4": 58,    "gravity": 0.33532095,    "min_expansion": 74,    "short_name": "YSOuDIJpGskPpoFIXODLGFvmO",    "map_file_name": "FAvNbAVKtXybwtJEsAtcXQmFs",    "minclip": 0.14423566,    "fog_red": 15,    "rain_chance_3": 90,    "content_flags": "IrGkRLZNuhpfSFkjunVvGCNqe",    "time_type": 46,    "fog_green_1": 97,    "snow_chance_4": 73,    "snow_duration_2": 12,    "max_movement_update_range": 91,    "max_expansion": 57,    "graveyard_id": 0.9312882,    "note": "WmWWDjrcxeuRYgZjbMBVulxeg",    "underworld": 0.27708793,    "fog_maxclip_2": 0.8189912,    "fog_maxclip_4": 0.1250884,    "castoutdoor": 94,    "insttype": 77,    "walkspeed": 0.36936185,    "fog_minclip_4": 0.9892267,    "flag_needed": "EtSHsvKVcrvFgrQXDxwLqDKyM",    "rain_chance_1": 27,    "snow_duration_4": 13,    "fast_regen_mana": 9,    "fog_blue_2": 11,    "fog_minclip_2": 0.84144753,    "fog_red_3": 43,    "cancombat": 94,    "suspendbuffs": 24,    "rain_chance_2": 20,    "type": 3,    "fog_maxclip": 0.85339475,    "fog_blue": 30,    "expansion": 67,    "rain_duration_2": 89,    "snow_chance_2": 90,    "fog_minclip_1": 0.056973238,    "fog_maxclip_1": 0.7270401,    "canbind": 68,    "peqzone": 88,    "fast_regen_endurance": 90,    "safe_z": 0.7712859,    "zoneidnumber": 3,    "fog_green_2": 64,    "fog_red_4": 68,    "fog_blue_4": 78,    "npc_max_aggro_dist": 3,    "maxclip": 0.81998205,    "zone_exp_multiplier": 0.8655295689385535,    "fog_blue_1": 25,    "content_flags_disabled": "iRtfbAZgDNJTMnOPUOlkYlwcM"}


Comments
-------------------------------------
[ 9] column is set for unsigned
[10] column is set for unsigned
[12] column is set for unsigned
[15] column is set for unsigned
[22] column is set for unsigned
[23] column is set for unsigned
[24] column is set for unsigned
[25] column is set for unsigned
[26] column is set for unsigned
[29] column is set for unsigned
[30] column is set for unsigned
[31] column is set for unsigned
[32] column is set for unsigned
[35] column is set for unsigned
[36] column is set for unsigned
[37] column is set for unsigned
[40] column is set for unsigned
[41] column is set for unsigned
[42] column is set for unsigned
[45] column is set for unsigned
[46] column is set for unsigned
[47] column is set for unsigned
[56] column is set for unsigned
[57] column is set for unsigned
[58] column is set for unsigned
[61] column is set for unsigned
[85] column is set for unsigned
[86] column is set for unsigned
[87] column is set for unsigned



*/

// Zone struct is a row record of the zone table in the eqemu database
type Zone struct {
	//[ 0] short_name                                     varchar(32)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 32      default: [NULL]
	ShortName null.String `gorm:"column:short_name;type:varchar;size:32;" json:"short_name"`
	//[ 1] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 2] file_name                                      varchar(16)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 16      default: [NULL]
	FileName null.String `gorm:"column:file_name;type:varchar;size:16;" json:"file_name"`
	//[ 3] long_name                                      text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	LongName string `gorm:"column:long_name;type:text;size:65535;" json:"long_name"`
	//[ 4] map_file_name                                  varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	MapFileName null.String `gorm:"column:map_file_name;type:varchar;size:100;" json:"map_file_name"`
	//[ 5] safe_x                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	SafeX float32 `gorm:"column:safe_x;type:float;default:0;" json:"safe_x"`
	//[ 6] safe_y                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	SafeY float32 `gorm:"column:safe_y;type:float;default:0;" json:"safe_y"`
	//[ 7] safe_z                                         float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	SafeZ float32 `gorm:"column:safe_z;type:float;default:0;" json:"safe_z"`
	//[ 8] graveyard_id                                   float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	GraveyardID float32 `gorm:"column:graveyard_id;type:float;default:0;" json:"graveyard_id"`
	//[ 9] min_level                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MinLevel uint32 `gorm:"column:min_level;type:utinyint;default:0;" json:"min_level"`
	//[10] min_status                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MinStatus uint32 `gorm:"column:min_status;type:utinyint;default:0;" json:"min_status"`
	//[11] zoneidnumber                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Zoneidnumber int32 `gorm:"column:zoneidnumber;type:int;default:0;" json:"zoneidnumber"`
	//[12] version                                        utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Version uint32 `gorm:"column:version;type:utinyint;default:0;" json:"version"`
	//[13] timezone                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Timezone int32 `gorm:"column:timezone;type:int;default:0;" json:"timezone"`
	//[14] maxclients                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Maxclients int32 `gorm:"column:maxclients;type:int;default:0;" json:"maxclients"`
	//[15] ruleset                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Ruleset uint32 `gorm:"column:ruleset;type:uint;default:0;" json:"ruleset"`
	//[16] note                                           varchar(80)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 80      default: [NULL]
	Note null.String `gorm:"column:note;type:varchar;size:80;" json:"note"`
	//[17] underworld                                     float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Underworld float32 `gorm:"column:underworld;type:float;default:0;" json:"underworld"`
	//[18] minclip                                        float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [450]
	Minclip float32 `gorm:"column:minclip;type:float;default:450;" json:"minclip"`
	//[19] maxclip                                        float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [450]
	Maxclip float32 `gorm:"column:maxclip;type:float;default:450;" json:"maxclip"`
	//[20] fog_minclip                                    float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [450]
	FogMinclip float32 `gorm:"column:fog_minclip;type:float;default:450;" json:"fog_minclip"`
	//[21] fog_maxclip                                    float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [450]
	FogMaxclip float32 `gorm:"column:fog_maxclip;type:float;default:450;" json:"fog_maxclip"`
	//[22] fog_blue                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogBlue uint32 `gorm:"column:fog_blue;type:utinyint;default:0;" json:"fog_blue"`
	//[23] fog_red                                        utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogRed uint32 `gorm:"column:fog_red;type:utinyint;default:0;" json:"fog_red"`
	//[24] fog_green                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogGreen uint32 `gorm:"column:fog_green;type:utinyint;default:0;" json:"fog_green"`
	//[25] sky                                            utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	Sky uint32 `gorm:"column:sky;type:utinyint;default:1;" json:"sky"`
	//[26] ztype                                          utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	Ztype uint32 `gorm:"column:ztype;type:utinyint;default:1;" json:"ztype"`
	//[27] zone_exp_multiplier                            decimal              null: false  primary: false  isArray: false  auto: false  col: decimal         len: -1      default: [0.00]
	ZoneExpMultiplier float64 `gorm:"column:zone_exp_multiplier;type:decimal;default:0.00;" json:"zone_exp_multiplier"`
	//[28] walkspeed                                      float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0.4]
	Walkspeed float32 `gorm:"column:walkspeed;type:float;default:0.4;" json:"walkspeed"`
	//[29] time_type                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [2]
	TimeType uint32 `gorm:"column:time_type;type:utinyint;default:2;" json:"time_type"`
	//[30] fog_red1                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogRed1 uint32 `gorm:"column:fog_red1;type:utinyint;default:0;" json:"fog_red_1"`
	//[31] fog_green1                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogGreen1 uint32 `gorm:"column:fog_green1;type:utinyint;default:0;" json:"fog_green_1"`
	//[32] fog_blue1                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogBlue1 uint32 `gorm:"column:fog_blue1;type:utinyint;default:0;" json:"fog_blue_1"`
	//[33] fog_minclip1                                   float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [450]
	FogMinclip1 float32 `gorm:"column:fog_minclip1;type:float;default:450;" json:"fog_minclip_1"`
	//[34] fog_maxclip1                                   float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [450]
	FogMaxclip1 float32 `gorm:"column:fog_maxclip1;type:float;default:450;" json:"fog_maxclip_1"`
	//[35] fog_red2                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogRed2 uint32 `gorm:"column:fog_red2;type:utinyint;default:0;" json:"fog_red_2"`
	//[36] fog_green2                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogGreen2 uint32 `gorm:"column:fog_green2;type:utinyint;default:0;" json:"fog_green_2"`
	//[37] fog_blue2                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogBlue2 uint32 `gorm:"column:fog_blue2;type:utinyint;default:0;" json:"fog_blue_2"`
	//[38] fog_minclip2                                   float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [450]
	FogMinclip2 float32 `gorm:"column:fog_minclip2;type:float;default:450;" json:"fog_minclip_2"`
	//[39] fog_maxclip2                                   float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [450]
	FogMaxclip2 float32 `gorm:"column:fog_maxclip2;type:float;default:450;" json:"fog_maxclip_2"`
	//[40] fog_red3                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogRed3 uint32 `gorm:"column:fog_red3;type:utinyint;default:0;" json:"fog_red_3"`
	//[41] fog_green3                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogGreen3 uint32 `gorm:"column:fog_green3;type:utinyint;default:0;" json:"fog_green_3"`
	//[42] fog_blue3                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogBlue3 uint32 `gorm:"column:fog_blue3;type:utinyint;default:0;" json:"fog_blue_3"`
	//[43] fog_minclip3                                   float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [450]
	FogMinclip3 float32 `gorm:"column:fog_minclip3;type:float;default:450;" json:"fog_minclip_3"`
	//[44] fog_maxclip3                                   float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [450]
	FogMaxclip3 float32 `gorm:"column:fog_maxclip3;type:float;default:450;" json:"fog_maxclip_3"`
	//[45] fog_red4                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogRed4 uint32 `gorm:"column:fog_red4;type:utinyint;default:0;" json:"fog_red_4"`
	//[46] fog_green4                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogGreen4 uint32 `gorm:"column:fog_green4;type:utinyint;default:0;" json:"fog_green_4"`
	//[47] fog_blue4                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	FogBlue4 uint32 `gorm:"column:fog_blue4;type:utinyint;default:0;" json:"fog_blue_4"`
	//[48] fog_minclip4                                   float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [450]
	FogMinclip4 float32 `gorm:"column:fog_minclip4;type:float;default:450;" json:"fog_minclip_4"`
	//[49] fog_maxclip4                                   float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [450]
	FogMaxclip4 float32 `gorm:"column:fog_maxclip4;type:float;default:450;" json:"fog_maxclip_4"`
	//[50] fog_density                                    float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	FogDensity float32 `gorm:"column:fog_density;type:float;default:0;" json:"fog_density"`
	//[51] flag_needed                                    varchar(128)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 128     default: ['']
	FlagNeeded string `gorm:"column:flag_needed;type:varchar;size:128;default:'';" json:"flag_needed"`
	//[52] canbind                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Canbind int32 `gorm:"column:canbind;type:tinyint;default:1;" json:"canbind"`
	//[53] cancombat                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Cancombat int32 `gorm:"column:cancombat;type:tinyint;default:1;" json:"cancombat"`
	//[54] canlevitate                                    tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Canlevitate int32 `gorm:"column:canlevitate;type:tinyint;default:1;" json:"canlevitate"`
	//[55] castoutdoor                                    tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Castoutdoor int32 `gorm:"column:castoutdoor;type:tinyint;default:1;" json:"castoutdoor"`
	//[56] hotzone                                        utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Hotzone uint32 `gorm:"column:hotzone;type:utinyint;default:0;" json:"hotzone"`
	//[57] insttype                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Insttype uint32 `gorm:"column:insttype;type:utinyint;default:0;" json:"insttype"`
	//[58] shutdowndelay                                  ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [5000]
	Shutdowndelay uint64 `gorm:"column:shutdowndelay;type:ubigint;default:5000;" json:"shutdowndelay"`
	//[59] peqzone                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Peqzone int32 `gorm:"column:peqzone;type:tinyint;default:1;" json:"peqzone"`
	//[60] expansion                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Expansion int32 `gorm:"column:expansion;type:tinyint;default:0;" json:"expansion"`
	//[61] suspendbuffs                                   utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Suspendbuffs uint32 `gorm:"column:suspendbuffs;type:utinyint;default:0;" json:"suspendbuffs"`
	//[62] rain_chance1                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RainChance1 int32 `gorm:"column:rain_chance1;type:int;default:0;" json:"rain_chance_1"`
	//[63] rain_chance2                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RainChance2 int32 `gorm:"column:rain_chance2;type:int;default:0;" json:"rain_chance_2"`
	//[64] rain_chance3                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RainChance3 int32 `gorm:"column:rain_chance3;type:int;default:0;" json:"rain_chance_3"`
	//[65] rain_chance4                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RainChance4 int32 `gorm:"column:rain_chance4;type:int;default:0;" json:"rain_chance_4"`
	//[66] rain_duration1                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RainDuration1 int32 `gorm:"column:rain_duration1;type:int;default:0;" json:"rain_duration_1"`
	//[67] rain_duration2                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RainDuration2 int32 `gorm:"column:rain_duration2;type:int;default:0;" json:"rain_duration_2"`
	//[68] rain_duration3                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RainDuration3 int32 `gorm:"column:rain_duration3;type:int;default:0;" json:"rain_duration_3"`
	//[69] rain_duration4                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RainDuration4 int32 `gorm:"column:rain_duration4;type:int;default:0;" json:"rain_duration_4"`
	//[70] snow_chance1                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SnowChance1 int32 `gorm:"column:snow_chance1;type:int;default:0;" json:"snow_chance_1"`
	//[71] snow_chance2                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SnowChance2 int32 `gorm:"column:snow_chance2;type:int;default:0;" json:"snow_chance_2"`
	//[72] snow_chance3                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SnowChance3 int32 `gorm:"column:snow_chance3;type:int;default:0;" json:"snow_chance_3"`
	//[73] snow_chance4                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SnowChance4 int32 `gorm:"column:snow_chance4;type:int;default:0;" json:"snow_chance_4"`
	//[74] snow_duration1                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SnowDuration1 int32 `gorm:"column:snow_duration1;type:int;default:0;" json:"snow_duration_1"`
	//[75] snow_duration2                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SnowDuration2 int32 `gorm:"column:snow_duration2;type:int;default:0;" json:"snow_duration_2"`
	//[76] snow_duration3                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SnowDuration3 int32 `gorm:"column:snow_duration3;type:int;default:0;" json:"snow_duration_3"`
	//[77] snow_duration4                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SnowDuration4 int32 `gorm:"column:snow_duration4;type:int;default:0;" json:"snow_duration_4"`
	//[78] gravity                                        float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0.4]
	Gravity float32 `gorm:"column:gravity;type:float;default:0.4;" json:"gravity"`
	//[79] type                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Type int32 `gorm:"column:type;type:int;default:0;" json:"type"`
	//[80] skylock                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Skylock int32 `gorm:"column:skylock;type:tinyint;default:0;" json:"skylock"`
	//[81] fast_regen_hp                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [180]
	FastRegenHp int32 `gorm:"column:fast_regen_hp;type:int;default:180;" json:"fast_regen_hp"`
	//[82] fast_regen_mana                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [180]
	FastRegenMana int32 `gorm:"column:fast_regen_mana;type:int;default:180;" json:"fast_regen_mana"`
	//[83] fast_regen_endurance                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [180]
	FastRegenEndurance int32 `gorm:"column:fast_regen_endurance;type:int;default:180;" json:"fast_regen_endurance"`
	//[84] npc_max_aggro_dist                             int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [600]
	NpcMaxAggroDist int32 `gorm:"column:npc_max_aggro_dist;type:int;default:600;" json:"npc_max_aggro_dist"`
	//[85] max_movement_update_range                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [600]
	MaxMovementUpdateRange uint32 `gorm:"column:max_movement_update_range;type:uint;default:600;" json:"max_movement_update_range"`
	//[86] min_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MinExpansion uint32 `gorm:"column:min_expansion;type:utinyint;default:0;" json:"min_expansion"`
	//[87] max_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MaxExpansion uint32 `gorm:"column:max_expansion;type:utinyint;default:0;" json:"max_expansion"`
	//[88] content_flags                                  varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlags null.String `gorm:"column:content_flags;type:varchar;size:100;" json:"content_flags"`
	//[89] content_flags_disabled                         varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlagsDisabled null.String `gorm:"column:content_flags_disabled;type:varchar;size:100;" json:"content_flags_disabled"`
}

var zoneTableInfo = &TableInfo{
	Name: "zone",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "short_name",
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
			GoFieldName:        "ShortName",
			GoFieldType:        "null.String",
			JSONFieldName:      "short_name",
			ProtobufFieldName:  "short_name",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "file_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(16)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       16,
			GoFieldName:        "FileName",
			GoFieldType:        "null.String",
			JSONFieldName:      "file_name",
			ProtobufFieldName:  "file_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "long_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "LongName",
			GoFieldType:        "string",
			JSONFieldName:      "long_name",
			ProtobufFieldName:  "long_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "map_file_name",
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
			GoFieldName:        "MapFileName",
			GoFieldType:        "null.String",
			JSONFieldName:      "map_file_name",
			ProtobufFieldName:  "map_file_name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "safe_x",
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
			GoFieldName:        "SafeX",
			GoFieldType:        "float32",
			JSONFieldName:      "safe_x",
			ProtobufFieldName:  "safe_x",
			ProtobufType:       "float",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "safe_y",
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
			GoFieldName:        "SafeY",
			GoFieldType:        "float32",
			JSONFieldName:      "safe_y",
			ProtobufFieldName:  "safe_y",
			ProtobufType:       "float",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "safe_z",
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
			GoFieldName:        "SafeZ",
			GoFieldType:        "float32",
			JSONFieldName:      "safe_z",
			ProtobufFieldName:  "safe_z",
			ProtobufType:       "float",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "graveyard_id",
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
			GoFieldName:        "GraveyardID",
			GoFieldType:        "float32",
			JSONFieldName:      "graveyard_id",
			ProtobufFieldName:  "graveyard_id",
			ProtobufType:       "float",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "min_level",
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
			GoFieldName:        "MinLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "min_level",
			ProtobufFieldName:  "min_level",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "min_status",
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
			GoFieldName:        "MinStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "min_status",
			ProtobufFieldName:  "min_status",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "zoneidnumber",
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
			GoFieldName:        "Zoneidnumber",
			GoFieldType:        "int32",
			JSONFieldName:      "zoneidnumber",
			ProtobufFieldName:  "zoneidnumber",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "version",
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
			GoFieldName:        "Version",
			GoFieldType:        "uint32",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "timezone",
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
			GoFieldName:        "Timezone",
			GoFieldType:        "int32",
			JSONFieldName:      "timezone",
			ProtobufFieldName:  "timezone",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "maxclients",
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
			GoFieldName:        "Maxclients",
			GoFieldType:        "int32",
			JSONFieldName:      "maxclients",
			ProtobufFieldName:  "maxclients",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "ruleset",
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
			GoFieldName:        "Ruleset",
			GoFieldType:        "uint32",
			JSONFieldName:      "ruleset",
			ProtobufFieldName:  "ruleset",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "note",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(80)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       80,
			GoFieldName:        "Note",
			GoFieldType:        "null.String",
			JSONFieldName:      "note",
			ProtobufFieldName:  "note",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "underworld",
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
			GoFieldName:        "Underworld",
			GoFieldType:        "float32",
			JSONFieldName:      "underworld",
			ProtobufFieldName:  "underworld",
			ProtobufType:       "float",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "minclip",
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
			GoFieldName:        "Minclip",
			GoFieldType:        "float32",
			JSONFieldName:      "minclip",
			ProtobufFieldName:  "minclip",
			ProtobufType:       "float",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "maxclip",
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
			GoFieldName:        "Maxclip",
			GoFieldType:        "float32",
			JSONFieldName:      "maxclip",
			ProtobufFieldName:  "maxclip",
			ProtobufType:       "float",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "fog_minclip",
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
			GoFieldName:        "FogMinclip",
			GoFieldType:        "float32",
			JSONFieldName:      "fog_minclip",
			ProtobufFieldName:  "fog_minclip",
			ProtobufType:       "float",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "fog_maxclip",
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
			GoFieldName:        "FogMaxclip",
			GoFieldType:        "float32",
			JSONFieldName:      "fog_maxclip",
			ProtobufFieldName:  "fog_maxclip",
			ProtobufType:       "float",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "fog_blue",
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
			GoFieldName:        "FogBlue",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_blue",
			ProtobufFieldName:  "fog_blue",
			ProtobufType:       "uint32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "fog_red",
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
			GoFieldName:        "FogRed",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_red",
			ProtobufFieldName:  "fog_red",
			ProtobufType:       "uint32",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "fog_green",
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
			GoFieldName:        "FogGreen",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_green",
			ProtobufFieldName:  "fog_green",
			ProtobufType:       "uint32",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "sky",
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
			GoFieldName:        "Sky",
			GoFieldType:        "uint32",
			JSONFieldName:      "sky",
			ProtobufFieldName:  "sky",
			ProtobufType:       "uint32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "ztype",
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
			GoFieldName:        "Ztype",
			GoFieldType:        "uint32",
			JSONFieldName:      "ztype",
			ProtobufFieldName:  "ztype",
			ProtobufType:       "uint32",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "zone_exp_multiplier",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "ZoneExpMultiplier",
			GoFieldType:        "float64",
			JSONFieldName:      "zone_exp_multiplier",
			ProtobufFieldName:  "zone_exp_multiplier",
			ProtobufType:       "float",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "walkspeed",
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
			GoFieldName:        "Walkspeed",
			GoFieldType:        "float32",
			JSONFieldName:      "walkspeed",
			ProtobufFieldName:  "walkspeed",
			ProtobufType:       "float",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "time_type",
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
			GoFieldName:        "TimeType",
			GoFieldType:        "uint32",
			JSONFieldName:      "time_type",
			ProtobufFieldName:  "time_type",
			ProtobufType:       "uint32",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "fog_red1",
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
			GoFieldName:        "FogRed1",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_red_1",
			ProtobufFieldName:  "fog_red_1",
			ProtobufType:       "uint32",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "fog_green1",
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
			GoFieldName:        "FogGreen1",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_green_1",
			ProtobufFieldName:  "fog_green_1",
			ProtobufType:       "uint32",
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "fog_blue1",
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
			GoFieldName:        "FogBlue1",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_blue_1",
			ProtobufFieldName:  "fog_blue_1",
			ProtobufType:       "uint32",
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
			Name:               "fog_minclip1",
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
			GoFieldName:        "FogMinclip1",
			GoFieldType:        "float32",
			JSONFieldName:      "fog_minclip_1",
			ProtobufFieldName:  "fog_minclip_1",
			ProtobufType:       "float",
			ProtobufPos:        34,
		},

		&ColumnInfo{
			Index:              34,
			Name:               "fog_maxclip1",
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
			GoFieldName:        "FogMaxclip1",
			GoFieldType:        "float32",
			JSONFieldName:      "fog_maxclip_1",
			ProtobufFieldName:  "fog_maxclip_1",
			ProtobufType:       "float",
			ProtobufPos:        35,
		},

		&ColumnInfo{
			Index:              35,
			Name:               "fog_red2",
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
			GoFieldName:        "FogRed2",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_red_2",
			ProtobufFieldName:  "fog_red_2",
			ProtobufType:       "uint32",
			ProtobufPos:        36,
		},

		&ColumnInfo{
			Index:              36,
			Name:               "fog_green2",
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
			GoFieldName:        "FogGreen2",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_green_2",
			ProtobufFieldName:  "fog_green_2",
			ProtobufType:       "uint32",
			ProtobufPos:        37,
		},

		&ColumnInfo{
			Index:              37,
			Name:               "fog_blue2",
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
			GoFieldName:        "FogBlue2",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_blue_2",
			ProtobufFieldName:  "fog_blue_2",
			ProtobufType:       "uint32",
			ProtobufPos:        38,
		},

		&ColumnInfo{
			Index:              38,
			Name:               "fog_minclip2",
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
			GoFieldName:        "FogMinclip2",
			GoFieldType:        "float32",
			JSONFieldName:      "fog_minclip_2",
			ProtobufFieldName:  "fog_minclip_2",
			ProtobufType:       "float",
			ProtobufPos:        39,
		},

		&ColumnInfo{
			Index:              39,
			Name:               "fog_maxclip2",
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
			GoFieldName:        "FogMaxclip2",
			GoFieldType:        "float32",
			JSONFieldName:      "fog_maxclip_2",
			ProtobufFieldName:  "fog_maxclip_2",
			ProtobufType:       "float",
			ProtobufPos:        40,
		},

		&ColumnInfo{
			Index:              40,
			Name:               "fog_red3",
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
			GoFieldName:        "FogRed3",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_red_3",
			ProtobufFieldName:  "fog_red_3",
			ProtobufType:       "uint32",
			ProtobufPos:        41,
		},

		&ColumnInfo{
			Index:              41,
			Name:               "fog_green3",
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
			GoFieldName:        "FogGreen3",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_green_3",
			ProtobufFieldName:  "fog_green_3",
			ProtobufType:       "uint32",
			ProtobufPos:        42,
		},

		&ColumnInfo{
			Index:              42,
			Name:               "fog_blue3",
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
			GoFieldName:        "FogBlue3",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_blue_3",
			ProtobufFieldName:  "fog_blue_3",
			ProtobufType:       "uint32",
			ProtobufPos:        43,
		},

		&ColumnInfo{
			Index:              43,
			Name:               "fog_minclip3",
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
			GoFieldName:        "FogMinclip3",
			GoFieldType:        "float32",
			JSONFieldName:      "fog_minclip_3",
			ProtobufFieldName:  "fog_minclip_3",
			ProtobufType:       "float",
			ProtobufPos:        44,
		},

		&ColumnInfo{
			Index:              44,
			Name:               "fog_maxclip3",
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
			GoFieldName:        "FogMaxclip3",
			GoFieldType:        "float32",
			JSONFieldName:      "fog_maxclip_3",
			ProtobufFieldName:  "fog_maxclip_3",
			ProtobufType:       "float",
			ProtobufPos:        45,
		},

		&ColumnInfo{
			Index:              45,
			Name:               "fog_red4",
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
			GoFieldName:        "FogRed4",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_red_4",
			ProtobufFieldName:  "fog_red_4",
			ProtobufType:       "uint32",
			ProtobufPos:        46,
		},

		&ColumnInfo{
			Index:              46,
			Name:               "fog_green4",
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
			GoFieldName:        "FogGreen4",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_green_4",
			ProtobufFieldName:  "fog_green_4",
			ProtobufType:       "uint32",
			ProtobufPos:        47,
		},

		&ColumnInfo{
			Index:              47,
			Name:               "fog_blue4",
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
			GoFieldName:        "FogBlue4",
			GoFieldType:        "uint32",
			JSONFieldName:      "fog_blue_4",
			ProtobufFieldName:  "fog_blue_4",
			ProtobufType:       "uint32",
			ProtobufPos:        48,
		},

		&ColumnInfo{
			Index:              48,
			Name:               "fog_minclip4",
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
			GoFieldName:        "FogMinclip4",
			GoFieldType:        "float32",
			JSONFieldName:      "fog_minclip_4",
			ProtobufFieldName:  "fog_minclip_4",
			ProtobufType:       "float",
			ProtobufPos:        49,
		},

		&ColumnInfo{
			Index:              49,
			Name:               "fog_maxclip4",
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
			GoFieldName:        "FogMaxclip4",
			GoFieldType:        "float32",
			JSONFieldName:      "fog_maxclip_4",
			ProtobufFieldName:  "fog_maxclip_4",
			ProtobufType:       "float",
			ProtobufPos:        50,
		},

		&ColumnInfo{
			Index:              50,
			Name:               "fog_density",
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
			GoFieldName:        "FogDensity",
			GoFieldType:        "float32",
			JSONFieldName:      "fog_density",
			ProtobufFieldName:  "fog_density",
			ProtobufType:       "float",
			ProtobufPos:        51,
		},

		&ColumnInfo{
			Index:              51,
			Name:               "flag_needed",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       128,
			GoFieldName:        "FlagNeeded",
			GoFieldType:        "string",
			JSONFieldName:      "flag_needed",
			ProtobufFieldName:  "flag_needed",
			ProtobufType:       "string",
			ProtobufPos:        52,
		},

		&ColumnInfo{
			Index:              52,
			Name:               "canbind",
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
			GoFieldName:        "Canbind",
			GoFieldType:        "int32",
			JSONFieldName:      "canbind",
			ProtobufFieldName:  "canbind",
			ProtobufType:       "int32",
			ProtobufPos:        53,
		},

		&ColumnInfo{
			Index:              53,
			Name:               "cancombat",
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
			GoFieldName:        "Cancombat",
			GoFieldType:        "int32",
			JSONFieldName:      "cancombat",
			ProtobufFieldName:  "cancombat",
			ProtobufType:       "int32",
			ProtobufPos:        54,
		},

		&ColumnInfo{
			Index:              54,
			Name:               "canlevitate",
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
			GoFieldName:        "Canlevitate",
			GoFieldType:        "int32",
			JSONFieldName:      "canlevitate",
			ProtobufFieldName:  "canlevitate",
			ProtobufType:       "int32",
			ProtobufPos:        55,
		},

		&ColumnInfo{
			Index:              55,
			Name:               "castoutdoor",
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
			GoFieldName:        "Castoutdoor",
			GoFieldType:        "int32",
			JSONFieldName:      "castoutdoor",
			ProtobufFieldName:  "castoutdoor",
			ProtobufType:       "int32",
			ProtobufPos:        56,
		},

		&ColumnInfo{
			Index:              56,
			Name:               "hotzone",
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
			GoFieldName:        "Hotzone",
			GoFieldType:        "uint32",
			JSONFieldName:      "hotzone",
			ProtobufFieldName:  "hotzone",
			ProtobufType:       "uint32",
			ProtobufPos:        57,
		},

		&ColumnInfo{
			Index:              57,
			Name:               "insttype",
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
			GoFieldName:        "Insttype",
			GoFieldType:        "uint32",
			JSONFieldName:      "insttype",
			ProtobufFieldName:  "insttype",
			ProtobufType:       "uint32",
			ProtobufPos:        58,
		},

		&ColumnInfo{
			Index:              58,
			Name:               "shutdowndelay",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "Shutdowndelay",
			GoFieldType:        "uint64",
			JSONFieldName:      "shutdowndelay",
			ProtobufFieldName:  "shutdowndelay",
			ProtobufType:       "uint64",
			ProtobufPos:        59,
		},

		&ColumnInfo{
			Index:              59,
			Name:               "peqzone",
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
			GoFieldName:        "Peqzone",
			GoFieldType:        "int32",
			JSONFieldName:      "peqzone",
			ProtobufFieldName:  "peqzone",
			ProtobufType:       "int32",
			ProtobufPos:        60,
		},

		&ColumnInfo{
			Index:              60,
			Name:               "expansion",
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
			GoFieldName:        "Expansion",
			GoFieldType:        "int32",
			JSONFieldName:      "expansion",
			ProtobufFieldName:  "expansion",
			ProtobufType:       "int32",
			ProtobufPos:        61,
		},

		&ColumnInfo{
			Index:              61,
			Name:               "suspendbuffs",
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
			GoFieldName:        "Suspendbuffs",
			GoFieldType:        "uint32",
			JSONFieldName:      "suspendbuffs",
			ProtobufFieldName:  "suspendbuffs",
			ProtobufType:       "uint32",
			ProtobufPos:        62,
		},

		&ColumnInfo{
			Index:              62,
			Name:               "rain_chance1",
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
			GoFieldName:        "RainChance1",
			GoFieldType:        "int32",
			JSONFieldName:      "rain_chance_1",
			ProtobufFieldName:  "rain_chance_1",
			ProtobufType:       "int32",
			ProtobufPos:        63,
		},

		&ColumnInfo{
			Index:              63,
			Name:               "rain_chance2",
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
			GoFieldName:        "RainChance2",
			GoFieldType:        "int32",
			JSONFieldName:      "rain_chance_2",
			ProtobufFieldName:  "rain_chance_2",
			ProtobufType:       "int32",
			ProtobufPos:        64,
		},

		&ColumnInfo{
			Index:              64,
			Name:               "rain_chance3",
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
			GoFieldName:        "RainChance3",
			GoFieldType:        "int32",
			JSONFieldName:      "rain_chance_3",
			ProtobufFieldName:  "rain_chance_3",
			ProtobufType:       "int32",
			ProtobufPos:        65,
		},

		&ColumnInfo{
			Index:              65,
			Name:               "rain_chance4",
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
			GoFieldName:        "RainChance4",
			GoFieldType:        "int32",
			JSONFieldName:      "rain_chance_4",
			ProtobufFieldName:  "rain_chance_4",
			ProtobufType:       "int32",
			ProtobufPos:        66,
		},

		&ColumnInfo{
			Index:              66,
			Name:               "rain_duration1",
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
			GoFieldName:        "RainDuration1",
			GoFieldType:        "int32",
			JSONFieldName:      "rain_duration_1",
			ProtobufFieldName:  "rain_duration_1",
			ProtobufType:       "int32",
			ProtobufPos:        67,
		},

		&ColumnInfo{
			Index:              67,
			Name:               "rain_duration2",
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
			GoFieldName:        "RainDuration2",
			GoFieldType:        "int32",
			JSONFieldName:      "rain_duration_2",
			ProtobufFieldName:  "rain_duration_2",
			ProtobufType:       "int32",
			ProtobufPos:        68,
		},

		&ColumnInfo{
			Index:              68,
			Name:               "rain_duration3",
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
			GoFieldName:        "RainDuration3",
			GoFieldType:        "int32",
			JSONFieldName:      "rain_duration_3",
			ProtobufFieldName:  "rain_duration_3",
			ProtobufType:       "int32",
			ProtobufPos:        69,
		},

		&ColumnInfo{
			Index:              69,
			Name:               "rain_duration4",
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
			GoFieldName:        "RainDuration4",
			GoFieldType:        "int32",
			JSONFieldName:      "rain_duration_4",
			ProtobufFieldName:  "rain_duration_4",
			ProtobufType:       "int32",
			ProtobufPos:        70,
		},

		&ColumnInfo{
			Index:              70,
			Name:               "snow_chance1",
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
			GoFieldName:        "SnowChance1",
			GoFieldType:        "int32",
			JSONFieldName:      "snow_chance_1",
			ProtobufFieldName:  "snow_chance_1",
			ProtobufType:       "int32",
			ProtobufPos:        71,
		},

		&ColumnInfo{
			Index:              71,
			Name:               "snow_chance2",
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
			GoFieldName:        "SnowChance2",
			GoFieldType:        "int32",
			JSONFieldName:      "snow_chance_2",
			ProtobufFieldName:  "snow_chance_2",
			ProtobufType:       "int32",
			ProtobufPos:        72,
		},

		&ColumnInfo{
			Index:              72,
			Name:               "snow_chance3",
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
			GoFieldName:        "SnowChance3",
			GoFieldType:        "int32",
			JSONFieldName:      "snow_chance_3",
			ProtobufFieldName:  "snow_chance_3",
			ProtobufType:       "int32",
			ProtobufPos:        73,
		},

		&ColumnInfo{
			Index:              73,
			Name:               "snow_chance4",
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
			GoFieldName:        "SnowChance4",
			GoFieldType:        "int32",
			JSONFieldName:      "snow_chance_4",
			ProtobufFieldName:  "snow_chance_4",
			ProtobufType:       "int32",
			ProtobufPos:        74,
		},

		&ColumnInfo{
			Index:              74,
			Name:               "snow_duration1",
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
			GoFieldName:        "SnowDuration1",
			GoFieldType:        "int32",
			JSONFieldName:      "snow_duration_1",
			ProtobufFieldName:  "snow_duration_1",
			ProtobufType:       "int32",
			ProtobufPos:        75,
		},

		&ColumnInfo{
			Index:              75,
			Name:               "snow_duration2",
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
			GoFieldName:        "SnowDuration2",
			GoFieldType:        "int32",
			JSONFieldName:      "snow_duration_2",
			ProtobufFieldName:  "snow_duration_2",
			ProtobufType:       "int32",
			ProtobufPos:        76,
		},

		&ColumnInfo{
			Index:              76,
			Name:               "snow_duration3",
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
			GoFieldName:        "SnowDuration3",
			GoFieldType:        "int32",
			JSONFieldName:      "snow_duration_3",
			ProtobufFieldName:  "snow_duration_3",
			ProtobufType:       "int32",
			ProtobufPos:        77,
		},

		&ColumnInfo{
			Index:              77,
			Name:               "snow_duration4",
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
			GoFieldName:        "SnowDuration4",
			GoFieldType:        "int32",
			JSONFieldName:      "snow_duration_4",
			ProtobufFieldName:  "snow_duration_4",
			ProtobufType:       "int32",
			ProtobufPos:        78,
		},

		&ColumnInfo{
			Index:              78,
			Name:               "gravity",
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
			GoFieldName:        "Gravity",
			GoFieldType:        "float32",
			JSONFieldName:      "gravity",
			ProtobufFieldName:  "gravity",
			ProtobufType:       "float",
			ProtobufPos:        79,
		},

		&ColumnInfo{
			Index:              79,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        80,
		},

		&ColumnInfo{
			Index:              80,
			Name:               "skylock",
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
			GoFieldName:        "Skylock",
			GoFieldType:        "int32",
			JSONFieldName:      "skylock",
			ProtobufFieldName:  "skylock",
			ProtobufType:       "int32",
			ProtobufPos:        81,
		},

		&ColumnInfo{
			Index:              81,
			Name:               "fast_regen_hp",
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
			GoFieldName:        "FastRegenHp",
			GoFieldType:        "int32",
			JSONFieldName:      "fast_regen_hp",
			ProtobufFieldName:  "fast_regen_hp",
			ProtobufType:       "int32",
			ProtobufPos:        82,
		},

		&ColumnInfo{
			Index:              82,
			Name:               "fast_regen_mana",
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
			GoFieldName:        "FastRegenMana",
			GoFieldType:        "int32",
			JSONFieldName:      "fast_regen_mana",
			ProtobufFieldName:  "fast_regen_mana",
			ProtobufType:       "int32",
			ProtobufPos:        83,
		},

		&ColumnInfo{
			Index:              83,
			Name:               "fast_regen_endurance",
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
			GoFieldName:        "FastRegenEndurance",
			GoFieldType:        "int32",
			JSONFieldName:      "fast_regen_endurance",
			ProtobufFieldName:  "fast_regen_endurance",
			ProtobufType:       "int32",
			ProtobufPos:        84,
		},

		&ColumnInfo{
			Index:              84,
			Name:               "npc_max_aggro_dist",
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
			GoFieldName:        "NpcMaxAggroDist",
			GoFieldType:        "int32",
			JSONFieldName:      "npc_max_aggro_dist",
			ProtobufFieldName:  "npc_max_aggro_dist",
			ProtobufType:       "int32",
			ProtobufPos:        85,
		},

		&ColumnInfo{
			Index:              85,
			Name:               "max_movement_update_range",
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
			GoFieldName:        "MaxMovementUpdateRange",
			GoFieldType:        "uint32",
			JSONFieldName:      "max_movement_update_range",
			ProtobufFieldName:  "max_movement_update_range",
			ProtobufType:       "uint32",
			ProtobufPos:        86,
		},

		&ColumnInfo{
			Index:              86,
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
			ProtobufPos:        87,
		},

		&ColumnInfo{
			Index:              87,
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
			ProtobufPos:        88,
		},

		&ColumnInfo{
			Index:              88,
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
			ProtobufPos:        89,
		},

		&ColumnInfo{
			Index:              89,
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
			ProtobufPos:        90,
		},
	},
}

// TableName sets the insert table name for this struct type
func (z *Zone) TableName() string {
	return "zone"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (z *Zone) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (z *Zone) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (z *Zone) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (z *Zone) TableInfo() *TableInfo {
	return zoneTableInfo
}
