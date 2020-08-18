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


CREATE TABLE `sharedbank` (
  `acctid` int(11) unsigned DEFAULT 0,
  `slotid` mediumint(7) unsigned DEFAULT 0,
  `itemid` int(11) unsigned DEFAULT 0,
  `charges` smallint(3) unsigned DEFAULT 0,
  `augslot1` mediumint(7) unsigned NOT NULL DEFAULT 0,
  `augslot2` mediumint(7) unsigned NOT NULL DEFAULT 0,
  `augslot3` mediumint(7) unsigned NOT NULL DEFAULT 0,
  `augslot4` mediumint(7) unsigned NOT NULL DEFAULT 0,
  `augslot5` mediumint(7) unsigned NOT NULL DEFAULT 0,
  `augslot6` mediumint(7) NOT NULL DEFAULT 0,
  `custom_data` text DEFAULT NULL,
  UNIQUE KEY `account` (`acctid`,`slotid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "acctid": 65,    "augslot_2": 60,    "augslot_4": 82,    "custom_data": "KmSWarrrbyOliQsSvyUZNIWEF",    "slotid": 65,    "itemid": 26,    "charges": 25,    "augslot_1": 34,    "augslot_3": 13,    "augslot_5": 29,    "augslot_6": 27}


Comments
-------------------------------------
[ 0] column is set for unsignedWarning table: sharedbank does not have a primary key defined, setting col position 1 acctid as primary key
Warning table: sharedbank primary key column acctid is nullable column, setting it as NOT NULL

[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned



*/

// Sharedbank struct is a row record of the sharedbank table in the eqemu database
type Sharedbank struct {
	//[ 0] acctid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Acctid uint32 `gorm:"primary_key;column:acctid;type:uint;default:0;" json:"acctid"`
	//[ 1] slotid                                         umediumint           null: true   primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Slotid null.Int `gorm:"column:slotid;type:umediumint;default:0;" json:"slotid"`
	//[ 2] itemid                                         uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Itemid null.Int `gorm:"column:itemid;type:uint;default:0;" json:"itemid"`
	//[ 3] charges                                        usmallint            null: true   primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	Charges null.Int `gorm:"column:charges;type:usmallint;default:0;" json:"charges"`
	//[ 4] augslot1                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot1 uint32 `gorm:"column:augslot1;type:umediumint;default:0;" json:"augslot_1"`
	//[ 5] augslot2                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot2 uint32 `gorm:"column:augslot2;type:umediumint;default:0;" json:"augslot_2"`
	//[ 6] augslot3                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot3 uint32 `gorm:"column:augslot3;type:umediumint;default:0;" json:"augslot_3"`
	//[ 7] augslot4                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot4 uint32 `gorm:"column:augslot4;type:umediumint;default:0;" json:"augslot_4"`
	//[ 8] augslot5                                       umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot5 uint32 `gorm:"column:augslot5;type:umediumint;default:0;" json:"augslot_5"`
	//[ 9] augslot6                                       mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	Augslot6 int32 `gorm:"column:augslot6;type:mediumint;default:0;" json:"augslot_6"`
	//[10] custom_data                                    text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: [NULL]
	CustomData null.String `gorm:"column:custom_data;type:text;size:65535;" json:"custom_data"`
}

var sharedbankTableInfo = &TableInfo{
	Name: "sharedbank",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "acctid",
			Comment: ``,
			Notes: `column is set for unsignedWarning table: sharedbank does not have a primary key defined, setting col position 1 acctid as primary key
Warning table: sharedbank primary key column acctid is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Acctid",
			GoFieldType:        "uint32",
			JSONFieldName:      "acctid",
			ProtobufFieldName:  "acctid",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "slotid",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "umediumint",
			DatabaseTypePretty: "umediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "umediumint",
			ColumnLength:       -1,
			GoFieldName:        "Slotid",
			GoFieldType:        "null.Int",
			JSONFieldName:      "slotid",
			ProtobufFieldName:  "slotid",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "itemid",
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
			GoFieldName:        "Itemid",
			GoFieldType:        "null.Int",
			JSONFieldName:      "itemid",
			ProtobufFieldName:  "itemid",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "charges",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "Charges",
			GoFieldType:        "null.Int",
			JSONFieldName:      "charges",
			ProtobufFieldName:  "charges",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "augslot1",
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
			GoFieldName:        "Augslot1",
			GoFieldType:        "uint32",
			JSONFieldName:      "augslot_1",
			ProtobufFieldName:  "augslot_1",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "augslot2",
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
			GoFieldName:        "Augslot2",
			GoFieldType:        "uint32",
			JSONFieldName:      "augslot_2",
			ProtobufFieldName:  "augslot_2",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "augslot3",
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
			GoFieldName:        "Augslot3",
			GoFieldType:        "uint32",
			JSONFieldName:      "augslot_3",
			ProtobufFieldName:  "augslot_3",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "augslot4",
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
			GoFieldName:        "Augslot4",
			GoFieldType:        "uint32",
			JSONFieldName:      "augslot_4",
			ProtobufFieldName:  "augslot_4",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "augslot5",
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
			GoFieldName:        "Augslot5",
			GoFieldType:        "uint32",
			JSONFieldName:      "augslot_5",
			ProtobufFieldName:  "augslot_5",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "augslot6",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "mediumint",
			DatabaseTypePretty: "mediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "mediumint",
			ColumnLength:       -1,
			GoFieldName:        "Augslot6",
			GoFieldType:        "int32",
			JSONFieldName:      "augslot_6",
			ProtobufFieldName:  "augslot_6",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "custom_data",
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
			GoFieldName:        "CustomData",
			GoFieldType:        "null.String",
			JSONFieldName:      "custom_data",
			ProtobufFieldName:  "custom_data",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *Sharedbank) TableName() string {
	return "sharedbank"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *Sharedbank) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *Sharedbank) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *Sharedbank) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *Sharedbank) TableInfo() *TableInfo {
	return sharedbankTableInfo
}
