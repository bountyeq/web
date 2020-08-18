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


CREATE TABLE `petitions` (
  `dib` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `petid` int(10) unsigned NOT NULL DEFAULT 0,
  `charname` varchar(32) NOT NULL DEFAULT '',
  `accountname` varchar(32) NOT NULL DEFAULT '',
  `lastgm` varchar(32) NOT NULL DEFAULT '',
  `petitiontext` text NOT NULL,
  `gmtext` text DEFAULT NULL,
  `zone` varchar(32) NOT NULL DEFAULT '',
  `urgency` int(11) NOT NULL DEFAULT 0,
  `charclass` int(11) NOT NULL DEFAULT 0,
  `charrace` int(11) NOT NULL DEFAULT 0,
  `charlevel` int(11) NOT NULL DEFAULT 0,
  `checkouts` int(11) NOT NULL DEFAULT 0,
  `unavailables` int(11) NOT NULL DEFAULT 0,
  `ischeckedout` tinyint(4) NOT NULL DEFAULT 0,
  `senttime` bigint(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`dib`),
  KEY `petid` (`petid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "charlevel": 81,    "unavailables": 47,    "petid": 12,    "lastgm": "mugNcbudUcdFIQtvOGUraZZQW",    "senttime": 47,    "urgency": 85,    "charrace": 74,    "zone": "tBLeePxfBEeARhZTNtfHjBESs",    "checkouts": 96,    "ischeckedout": 51,    "dib": 45,    "charname": "aKTgaTETLemPXRLimhPGNEIHe",    "gmtext": "ytmujNjbqmrOUlPAxWkZWfMtG",    "charclass": 36,    "accountname": "FBmqwGRRAuEvJvRsDtvMExKGZ",    "petitiontext": "rZyswAKTSnVfbihANOFFjslrb"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// Petitions struct is a row record of the petitions table in the eqemu database
type Petitions struct {
	//[ 0] dib                                            uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	Dib uint32 `gorm:"primary_key;AUTO_INCREMENT;column:dib;type:uint;" json:"dib"`
	//[ 1] petid                                          uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Petid uint32 `gorm:"column:petid;type:uint;default:0;" json:"petid"`
	//[ 2] charname                                       varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['']
	Charname string `gorm:"column:charname;type:varchar;size:32;default:'';" json:"charname"`
	//[ 3] accountname                                    varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['']
	Accountname string `gorm:"column:accountname;type:varchar;size:32;default:'';" json:"accountname"`
	//[ 4] lastgm                                         varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['']
	Lastgm string `gorm:"column:lastgm;type:varchar;size:32;default:'';" json:"lastgm"`
	//[ 5] petitiontext                                   text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Petitiontext string `gorm:"column:petitiontext;type:text;size:65535;" json:"petitiontext"`
	//[ 6] gmtext                                         text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: [NULL]
	Gmtext null.String `gorm:"column:gmtext;type:text;size:65535;" json:"gmtext"`
	//[ 7] zone                                           varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['']
	Zone string `gorm:"column:zone;type:varchar;size:32;default:'';" json:"zone"`
	//[ 8] urgency                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Urgency int32 `gorm:"column:urgency;type:int;default:0;" json:"urgency"`
	//[ 9] charclass                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Charclass int32 `gorm:"column:charclass;type:int;default:0;" json:"charclass"`
	//[10] charrace                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Charrace int32 `gorm:"column:charrace;type:int;default:0;" json:"charrace"`
	//[11] charlevel                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Charlevel int32 `gorm:"column:charlevel;type:int;default:0;" json:"charlevel"`
	//[12] checkouts                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Checkouts int32 `gorm:"column:checkouts;type:int;default:0;" json:"checkouts"`
	//[13] unavailables                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Unavailables int32 `gorm:"column:unavailables;type:int;default:0;" json:"unavailables"`
	//[14] ischeckedout                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Ischeckedout int32 `gorm:"column:ischeckedout;type:tinyint;default:0;" json:"ischeckedout"`
	//[15] senttime                                       bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	Senttime int64 `gorm:"column:senttime;type:bigint;default:0;" json:"senttime"`
}

var petitionsTableInfo = &TableInfo{
	Name: "petitions",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "dib",
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
			GoFieldName:        "Dib",
			GoFieldType:        "uint32",
			JSONFieldName:      "dib",
			ProtobufFieldName:  "dib",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "petid",
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
			GoFieldName:        "Petid",
			GoFieldType:        "uint32",
			JSONFieldName:      "petid",
			ProtobufFieldName:  "petid",
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
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Charname",
			GoFieldType:        "string",
			JSONFieldName:      "charname",
			ProtobufFieldName:  "charname",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "accountname",
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
			GoFieldName:        "Accountname",
			GoFieldType:        "string",
			JSONFieldName:      "accountname",
			ProtobufFieldName:  "accountname",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "lastgm",
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
			GoFieldName:        "Lastgm",
			GoFieldType:        "string",
			JSONFieldName:      "lastgm",
			ProtobufFieldName:  "lastgm",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "petitiontext",
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
			GoFieldName:        "Petitiontext",
			GoFieldType:        "string",
			JSONFieldName:      "petitiontext",
			ProtobufFieldName:  "petitiontext",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "gmtext",
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
			GoFieldName:        "Gmtext",
			GoFieldType:        "null.String",
			JSONFieldName:      "gmtext",
			ProtobufFieldName:  "gmtext",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "zone",
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
			GoFieldName:        "Zone",
			GoFieldType:        "string",
			JSONFieldName:      "zone",
			ProtobufFieldName:  "zone",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "urgency",
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
			GoFieldName:        "Urgency",
			GoFieldType:        "int32",
			JSONFieldName:      "urgency",
			ProtobufFieldName:  "urgency",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "charclass",
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
			GoFieldName:        "Charclass",
			GoFieldType:        "int32",
			JSONFieldName:      "charclass",
			ProtobufFieldName:  "charclass",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "charrace",
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
			GoFieldName:        "Charrace",
			GoFieldType:        "int32",
			JSONFieldName:      "charrace",
			ProtobufFieldName:  "charrace",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "charlevel",
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
			GoFieldName:        "Charlevel",
			GoFieldType:        "int32",
			JSONFieldName:      "charlevel",
			ProtobufFieldName:  "charlevel",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "checkouts",
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
			GoFieldName:        "Checkouts",
			GoFieldType:        "int32",
			JSONFieldName:      "checkouts",
			ProtobufFieldName:  "checkouts",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "unavailables",
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
			GoFieldName:        "Unavailables",
			GoFieldType:        "int32",
			JSONFieldName:      "unavailables",
			ProtobufFieldName:  "unavailables",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "ischeckedout",
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
			GoFieldName:        "Ischeckedout",
			GoFieldType:        "int32",
			JSONFieldName:      "ischeckedout",
			ProtobufFieldName:  "ischeckedout",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "senttime",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "Senttime",
			GoFieldType:        "int64",
			JSONFieldName:      "senttime",
			ProtobufFieldName:  "senttime",
			ProtobufType:       "int64",
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *Petitions) TableName() string {
	return "petitions"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Petitions) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Petitions) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Petitions) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Petitions) TableInfo() *TableInfo {
	return petitionsTableInfo
}
