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


CREATE TABLE `forage` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `zoneid` int(4) NOT NULL DEFAULT 0,
  `Itemid` int(11) NOT NULL DEFAULT 0,
  `level` smallint(6) NOT NULL DEFAULT 0,
  `chance` smallint(6) NOT NULL DEFAULT 0,
  `min_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `max_expansion` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `content_flags` varchar(100) DEFAULT NULL,
  `content_flags_disabled` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=792 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "itemid": 43,    "level": 66,    "max_expansion": 0,    "content_flags": "yqlQmOcSOVgKHNoEkDgASLAaA",    "id": 78,    "zoneid": 38,    "chance": 0,    "min_expansion": 76,    "content_flags_disabled": "IFcEDDvcpYQZvsuDfBUnhwNXd"}


Comments
-------------------------------------
[ 5] column is set for unsigned
[ 6] column is set for unsigned



*/

// Forage struct is a row record of the forage table in the eqemu database
type Forage struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] zoneid                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Zoneid int32 `gorm:"column:zoneid;type:int;default:0;" json:"zoneid"`
	//[ 2] Itemid                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Itemid int32 `gorm:"column:Itemid;type:int;default:0;" json:"itemid"`
	//[ 3] level                                          smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Level int32 `gorm:"column:level;type:smallint;default:0;" json:"level"`
	//[ 4] chance                                         smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Chance int32 `gorm:"column:chance;type:smallint;default:0;" json:"chance"`
	//[ 5] min_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MinExpansion uint32 `gorm:"column:min_expansion;type:utinyint;default:0;" json:"min_expansion"`
	//[ 6] max_expansion                                  utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	MaxExpansion uint32 `gorm:"column:max_expansion;type:utinyint;default:0;" json:"max_expansion"`
	//[ 7] content_flags                                  varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlags null.String `gorm:"column:content_flags;type:varchar;size:100;" json:"content_flags"`
	//[ 8] content_flags_disabled                         varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [NULL]
	ContentFlagsDisabled null.String `gorm:"column:content_flags_disabled;type:varchar;size:100;" json:"content_flags_disabled"`
}

var forageTableInfo = &TableInfo{
	Name: "forage",
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
			Name:               "zoneid",
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
			GoFieldName:        "Zoneid",
			GoFieldType:        "int32",
			JSONFieldName:      "zoneid",
			ProtobufFieldName:  "zoneid",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "Itemid",
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
			GoFieldName:        "Itemid",
			GoFieldType:        "int32",
			JSONFieldName:      "itemid",
			ProtobufFieldName:  "itemid",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "level",
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
			GoFieldName:        "Level",
			GoFieldType:        "int32",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "chance",
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
			GoFieldName:        "Chance",
			GoFieldType:        "int32",
			JSONFieldName:      "chance",
			ProtobufFieldName:  "chance",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *Forage) TableName() string {
	return "forage"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *Forage) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *Forage) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *Forage) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *Forage) TableInfo() *TableInfo {
	return forageTableInfo
}
