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


CREATE TABLE `object_contents` (
  `zoneid` int(11) unsigned NOT NULL DEFAULT 0,
  `parentid` int(11) unsigned NOT NULL DEFAULT 0,
  `bagidx` int(11) unsigned NOT NULL DEFAULT 0,
  `itemid` int(11) unsigned NOT NULL DEFAULT 0,
  `charges` smallint(3) NOT NULL DEFAULT 0,
  `droptime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `augslot1` mediumint(7) unsigned DEFAULT 0,
  `augslot2` mediumint(7) unsigned DEFAULT 0,
  `augslot3` mediumint(7) unsigned DEFAULT 0,
  `augslot4` mediumint(7) unsigned DEFAULT 0,
  `augslot5` mediumint(7) unsigned DEFAULT 0,
  `augslot6` mediumint(7) NOT NULL DEFAULT 0,
  PRIMARY KEY (`parentid`,`bagidx`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "augslot_2": 37,    "augslot_3": 31,    "augslot_4": 99,    "augslot_5": 6,    "bagidx": 46,    "augslot_1": 32,    "itemid": 93,    "charges": 40,    "droptime": "2206-04-18T15:01:32.152080571-07:00",    "augslot_6": 77,    "zoneid": 42,    "parentid": 89}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned



*/

// ObjectContents struct is a row record of the object_contents table in the eqemu database
type ObjectContents struct {
	//[ 0] zoneid                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Zoneid uint32 `gorm:"column:zoneid;type:uint;default:0;" json:"zoneid"`
	//[ 1] parentid                                       uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Parentid uint32 `gorm:"primary_key;column:parentid;type:uint;default:0;" json:"parentid"`
	//[ 2] bagidx                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Bagidx uint32 `gorm:"primary_key;column:bagidx;type:uint;default:0;" json:"bagidx"`
	//[ 3] itemid                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Itemid uint32 `gorm:"column:itemid;type:uint;default:0;" json:"itemid"`
	//[ 4] charges                                        smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Charges int32 `gorm:"column:charges;type:smallint;default:0;" json:"charges"`
	//[ 5] droptime                                       datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: ['0000-00-00 00:00:00']
	Droptime time.Time `gorm:"column:droptime;type:datetime;default:'0000-00-00 00:00:00';" json:"droptime"`
	//[ 6] augslot1                                       umediumint           null: true   primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot1 null.Int `gorm:"column:augslot1;type:umediumint;default:0;" json:"augslot_1"`
	//[ 7] augslot2                                       umediumint           null: true   primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot2 null.Int `gorm:"column:augslot2;type:umediumint;default:0;" json:"augslot_2"`
	//[ 8] augslot3                                       umediumint           null: true   primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot3 null.Int `gorm:"column:augslot3;type:umediumint;default:0;" json:"augslot_3"`
	//[ 9] augslot4                                       umediumint           null: true   primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot4 null.Int `gorm:"column:augslot4;type:umediumint;default:0;" json:"augslot_4"`
	//[10] augslot5                                       umediumint           null: true   primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Augslot5 null.Int `gorm:"column:augslot5;type:umediumint;default:0;" json:"augslot_5"`
	//[11] augslot6                                       mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [0]
	Augslot6 int32 `gorm:"column:augslot6;type:mediumint;default:0;" json:"augslot_6"`
}

var object_contentsTableInfo = &TableInfo{
	Name: "object_contents",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "zoneid",
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
			GoFieldName:        "Zoneid",
			GoFieldType:        "uint32",
			JSONFieldName:      "zoneid",
			ProtobufFieldName:  "zoneid",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "parentid",
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
			GoFieldName:        "Parentid",
			GoFieldType:        "uint32",
			JSONFieldName:      "parentid",
			ProtobufFieldName:  "parentid",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "bagidx",
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
			GoFieldName:        "Bagidx",
			GoFieldType:        "uint32",
			JSONFieldName:      "bagidx",
			ProtobufFieldName:  "bagidx",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "itemid",
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
			GoFieldName:        "Itemid",
			GoFieldType:        "uint32",
			JSONFieldName:      "itemid",
			ProtobufFieldName:  "itemid",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "charges",
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
			GoFieldName:        "Charges",
			GoFieldType:        "int32",
			JSONFieldName:      "charges",
			ProtobufFieldName:  "charges",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "droptime",
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
			GoFieldName:        "Droptime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "droptime",
			ProtobufFieldName:  "droptime",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "augslot1",
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
			GoFieldName:        "Augslot1",
			GoFieldType:        "null.Int",
			JSONFieldName:      "augslot_1",
			ProtobufFieldName:  "augslot_1",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "augslot2",
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
			GoFieldName:        "Augslot2",
			GoFieldType:        "null.Int",
			JSONFieldName:      "augslot_2",
			ProtobufFieldName:  "augslot_2",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "augslot3",
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
			GoFieldName:        "Augslot3",
			GoFieldType:        "null.Int",
			JSONFieldName:      "augslot_3",
			ProtobufFieldName:  "augslot_3",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "augslot4",
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
			GoFieldName:        "Augslot4",
			GoFieldType:        "null.Int",
			JSONFieldName:      "augslot_4",
			ProtobufFieldName:  "augslot_4",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "augslot5",
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
			GoFieldName:        "Augslot5",
			GoFieldType:        "null.Int",
			JSONFieldName:      "augslot_5",
			ProtobufFieldName:  "augslot_5",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *ObjectContents) TableName() string {
	return "object_contents"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *ObjectContents) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *ObjectContents) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *ObjectContents) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *ObjectContents) TableInfo() *TableInfo {
	return object_contentsTableInfo
}
