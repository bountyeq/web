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


CREATE TABLE `proximities` (
  `zoneid` int(10) unsigned NOT NULL DEFAULT 0,
  `exploreid` int(10) unsigned NOT NULL DEFAULT 0,
  `minx` float(14,6) NOT NULL DEFAULT 0.000000,
  `maxx` float(14,6) NOT NULL DEFAULT 0.000000,
  `miny` float(14,6) NOT NULL DEFAULT 0.000000,
  `maxy` float(14,6) NOT NULL DEFAULT 0.000000,
  `minz` float(14,6) NOT NULL DEFAULT 0.000000,
  `maxz` float(14,6) NOT NULL DEFAULT 0.000000,
  PRIMARY KEY (`zoneid`,`exploreid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "exploreid": 68,    "minx": 0.25834373,    "maxx": 0.7272341,    "miny": 0.6599412,    "maxy": 0.64397997,    "minz": 0.51169044,    "maxz": 0.03906295,    "zoneid": 74}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// Proximities struct is a row record of the proximities table in the eqemu database
type Proximities struct {
	//[ 0] zoneid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Zoneid uint32 `gorm:"primary_key;column:zoneid;type:uint;default:0;" json:"zoneid"`
	//[ 1] exploreid                                      uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Exploreid uint32 `gorm:"primary_key;column:exploreid;type:uint;default:0;" json:"exploreid"`
	//[ 2] minx                                           float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0.000000]
	Minx float32 `gorm:"column:minx;type:float;default:0.000000;" json:"minx"`
	//[ 3] maxx                                           float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0.000000]
	Maxx float32 `gorm:"column:maxx;type:float;default:0.000000;" json:"maxx"`
	//[ 4] miny                                           float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0.000000]
	Miny float32 `gorm:"column:miny;type:float;default:0.000000;" json:"miny"`
	//[ 5] maxy                                           float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0.000000]
	Maxy float32 `gorm:"column:maxy;type:float;default:0.000000;" json:"maxy"`
	//[ 6] minz                                           float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0.000000]
	Minz float32 `gorm:"column:minz;type:float;default:0.000000;" json:"minz"`
	//[ 7] maxz                                           float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0.000000]
	Maxz float32 `gorm:"column:maxz;type:float;default:0.000000;" json:"maxz"`
}

var proximitiesTableInfo = &TableInfo{
	Name: "proximities",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "zoneid",
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
			GoFieldName:        "Zoneid",
			GoFieldType:        "uint32",
			JSONFieldName:      "zoneid",
			ProtobufFieldName:  "zoneid",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "exploreid",
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
			GoFieldName:        "Exploreid",
			GoFieldType:        "uint32",
			JSONFieldName:      "exploreid",
			ProtobufFieldName:  "exploreid",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "minx",
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
			GoFieldName:        "Minx",
			GoFieldType:        "float32",
			JSONFieldName:      "minx",
			ProtobufFieldName:  "minx",
			ProtobufType:       "float",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "maxx",
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
			GoFieldName:        "Maxx",
			GoFieldType:        "float32",
			JSONFieldName:      "maxx",
			ProtobufFieldName:  "maxx",
			ProtobufType:       "float",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "miny",
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
			GoFieldName:        "Miny",
			GoFieldType:        "float32",
			JSONFieldName:      "miny",
			ProtobufFieldName:  "miny",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "maxy",
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
			GoFieldName:        "Maxy",
			GoFieldType:        "float32",
			JSONFieldName:      "maxy",
			ProtobufFieldName:  "maxy",
			ProtobufType:       "float",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "minz",
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
			GoFieldName:        "Minz",
			GoFieldType:        "float32",
			JSONFieldName:      "minz",
			ProtobufFieldName:  "minz",
			ProtobufType:       "float",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "maxz",
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
			GoFieldName:        "Maxz",
			GoFieldType:        "float32",
			JSONFieldName:      "maxz",
			ProtobufFieldName:  "maxz",
			ProtobufType:       "float",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *Proximities) TableName() string {
	return "proximities"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Proximities) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Proximities) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Proximities) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Proximities) TableInfo() *TableInfo {
	return proximitiesTableInfo
}
