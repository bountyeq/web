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


CREATE TABLE `grid` (
  `id` int(10) NOT NULL DEFAULT 0,
  `zoneid` int(10) NOT NULL DEFAULT 0,
  `type` int(10) NOT NULL DEFAULT 0,
  `type2` int(10) NOT NULL DEFAULT 0,
  PRIMARY KEY (`zoneid`,`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "type": 87,    "type_2": 32,    "id": 23,    "zoneid": 60}



*/

// Grid struct is a row record of the grid table in the eqemu database
type Grid struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	ID int32 `gorm:"primary_key;column:id;type:int;default:0;" json:"id"`
	//[ 1] zoneid                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Zoneid int32 `gorm:"primary_key;column:zoneid;type:int;default:0;" json:"zoneid"`
	//[ 2] type                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Type int32 `gorm:"column:type;type:int;default:0;" json:"type"`
	//[ 3] type2                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Type2 int32 `gorm:"column:type2;type:int;default:0;" json:"type_2"`
}

var gridTableInfo = &TableInfo{
	Name: "grid",
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
			IsAutoIncrement:    false,
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
			IsPrimaryKey:       true,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "type2",
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
			GoFieldName:        "Type2",
			GoFieldType:        "int32",
			JSONFieldName:      "type_2",
			ProtobufFieldName:  "type_2",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *Grid) TableName() string {
	return "grid"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *Grid) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *Grid) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *Grid) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *Grid) TableInfo() *TableInfo {
	return gridTableInfo
}
