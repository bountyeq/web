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


CREATE TABLE `account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `charname` varchar(64) NOT NULL DEFAULT '',
  `sharedplat` int(11) unsigned NOT NULL DEFAULT 0,
  `password` varchar(50) NOT NULL DEFAULT '',
  `status` int(5) NOT NULL DEFAULT 0,
  `ls_id` varchar(64) DEFAULT 'eqemu',
  `lsaccount_id` int(11) unsigned DEFAULT NULL,
  `gmspeed` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `revoked` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `karma` int(5) unsigned NOT NULL DEFAULT 0,
  `minilogin_ip` varchar(32) NOT NULL DEFAULT '',
  `hideme` tinyint(4) NOT NULL DEFAULT 0,
  `rulesflag` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `suspendeduntil` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `time_creation` int(10) unsigned NOT NULL DEFAULT 0,
  `expansion` tinyint(4) NOT NULL DEFAULT 0,
  `ban_reason` text DEFAULT NULL,
  `suspend_reason` text DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_ls_id` (`name`,`ls_id`),
  UNIQUE KEY `ls_id_lsaccount_id` (`ls_id`,`lsaccount_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 71,    "status": 46,    "ls_id": "YnMepSuTMtFEITtBxushwHHbT",    "karma": 25,    "time_creation": 29,    "expansion": 92,    "name": "klugYeFoDdouSWFsSMqnvWyTN",    "gmspeed": 67,    "sharedplat": 8,    "lsaccount_id": 20,    "revoked": 46,    "rulesflag": 39,    "ban_reason": "wBbwvbnkGMjNcVFSQnHYcRuLw",    "suspend_reason": "sjSOJKgLnVaLXYdZcCXIiPjrG",    "charname": "WLyjFgGLakGtiTDmATHEARJps",    "password": "lydAicYbjXbbGnWdolXbrRXSk",    "minilogin_ip": "tdBFvxbGEEqItfPIUBiMJHxcL",    "hideme": 79,    "suspendeduntil": "2255-09-03T19:56:10.027156703-07:00"}


Comments
-------------------------------------
[ 3] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned
[13] column is set for unsigned
[15] column is set for unsigned



*/

// Account struct is a row record of the account table in the eqemu database
type Account struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] name                                           varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: ['']
	Name string `gorm:"column:name;type:varchar;size:30;default:'';" json:"name"`
	//[ 2] charname                                       varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Charname string `gorm:"column:charname;type:varchar;size:64;default:'';" json:"charname"`
	//[ 3] sharedplat                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Sharedplat uint32 `gorm:"column:sharedplat;type:uint;default:0;" json:"sharedplat"`
	//[ 4] password                                       varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: ['']
	Password string `gorm:"column:password;type:varchar;size:50;default:'';" json:"password"`
	//[ 5] status                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Status int32 `gorm:"column:status;type:int;default:0;" json:"status"`
	//[ 6] ls_id                                          varchar(64)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['eqemu']
	LsID null.String `gorm:"column:ls_id;type:varchar;size:64;default:'eqemu';" json:"ls_id"`
	//[ 7] lsaccount_id                                   uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [NULL]
	LsaccountID null.Int `gorm:"column:lsaccount_id;type:uint;" json:"lsaccount_id"`
	//[ 8] gmspeed                                        utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Gmspeed uint32 `gorm:"column:gmspeed;type:utinyint;default:0;" json:"gmspeed"`
	//[ 9] revoked                                        utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Revoked uint32 `gorm:"column:revoked;type:utinyint;default:0;" json:"revoked"`
	//[10] karma                                          uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Karma uint32 `gorm:"column:karma;type:uint;default:0;" json:"karma"`
	//[11] minilogin_ip                                   varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['']
	MiniloginIP string `gorm:"column:minilogin_ip;type:varchar;size:32;default:'';" json:"minilogin_ip"`
	//[12] hideme                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Hideme int32 `gorm:"column:hideme;type:tinyint;default:0;" json:"hideme"`
	//[13] rulesflag                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Rulesflag uint32 `gorm:"column:rulesflag;type:utinyint;default:0;" json:"rulesflag"`
	//[14] suspendeduntil                                 datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: ['0000-00-00 00:00:00']
	Suspendeduntil time.Time `gorm:"column:suspendeduntil;type:datetime;default:'0000-00-00 00:00:00';" json:"suspendeduntil"`
	//[15] time_creation                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TimeCreation uint32 `gorm:"column:time_creation;type:uint;default:0;" json:"time_creation"`
	//[16] expansion                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Expansion int32 `gorm:"column:expansion;type:tinyint;default:0;" json:"expansion"`
	//[17] ban_reason                                     text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: [NULL]
	BanReason null.String `gorm:"column:ban_reason;type:text;size:65535;" json:"ban_reason"`
	//[18] suspend_reason                                 text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: [NULL]
	SuspendReason null.String `gorm:"column:suspend_reason;type:text;size:65535;" json:"suspend_reason"`
}

var accountTableInfo = &TableInfo{
	Name: "account",
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
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
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
			Name:               "sharedplat",
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
			GoFieldName:        "Sharedplat",
			GoFieldType:        "uint32",
			JSONFieldName:      "sharedplat",
			ProtobufFieldName:  "sharedplat",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "password",
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
			GoFieldName:        "Password",
			GoFieldType:        "string",
			JSONFieldName:      "password",
			ProtobufFieldName:  "password",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "status",
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
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "ls_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "LsID",
			GoFieldType:        "null.String",
			JSONFieldName:      "ls_id",
			ProtobufFieldName:  "ls_id",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "lsaccount_id",
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
			GoFieldName:        "LsaccountID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "lsaccount_id",
			ProtobufFieldName:  "lsaccount_id",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "gmspeed",
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
			GoFieldName:        "Gmspeed",
			GoFieldType:        "uint32",
			JSONFieldName:      "gmspeed",
			ProtobufFieldName:  "gmspeed",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "revoked",
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
			GoFieldName:        "Revoked",
			GoFieldType:        "uint32",
			JSONFieldName:      "revoked",
			ProtobufFieldName:  "revoked",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "karma",
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
			GoFieldName:        "Karma",
			GoFieldType:        "uint32",
			JSONFieldName:      "karma",
			ProtobufFieldName:  "karma",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "minilogin_ip",
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
			GoFieldName:        "MiniloginIP",
			GoFieldType:        "string",
			JSONFieldName:      "minilogin_ip",
			ProtobufFieldName:  "minilogin_ip",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "hideme",
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
			GoFieldName:        "Hideme",
			GoFieldType:        "int32",
			JSONFieldName:      "hideme",
			ProtobufFieldName:  "hideme",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "rulesflag",
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
			GoFieldName:        "Rulesflag",
			GoFieldType:        "uint32",
			JSONFieldName:      "rulesflag",
			ProtobufFieldName:  "rulesflag",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "suspendeduntil",
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
			GoFieldName:        "Suspendeduntil",
			GoFieldType:        "time.Time",
			JSONFieldName:      "suspendeduntil",
			ProtobufFieldName:  "suspendeduntil",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "time_creation",
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
			GoFieldName:        "TimeCreation",
			GoFieldType:        "uint32",
			JSONFieldName:      "time_creation",
			ProtobufFieldName:  "time_creation",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "ban_reason",
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
			GoFieldName:        "BanReason",
			GoFieldType:        "null.String",
			JSONFieldName:      "ban_reason",
			ProtobufFieldName:  "ban_reason",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "suspend_reason",
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
			GoFieldName:        "SuspendReason",
			GoFieldType:        "null.String",
			JSONFieldName:      "suspend_reason",
			ProtobufFieldName:  "suspend_reason",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *Account) TableName() string {
	return "account"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Account) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Account) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Account) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Account) TableInfo() *TableInfo {
	return accountTableInfo
}
