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


CREATE TABLE `base_data` (
  `level` int(10) unsigned NOT NULL,
  `class` int(10) unsigned NOT NULL,
  `hp` double NOT NULL,
  `mana` double NOT NULL,
  `end` double NOT NULL,
  `unk1` double NOT NULL,
  `unk2` double NOT NULL,
  `hp_fac` double NOT NULL,
  `mana_fac` double NOT NULL,
  `end_fac` double NOT NULL,
  PRIMARY KEY (`level`,`class`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "level": 5,    "unk_2": 0.7219146363786503,    "hp_fac": 0.9883869456460982,    "class": 28,    "hp": 0.03585486834166689,    "mana": 0.2361797849536649,    "end": 0.902053365413524,    "unk_1": 0.4121373371068175,    "mana_fac": 0.8113534708864335,    "end_fac": 0.38938487224995244}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// BaseData struct is a row record of the base_data table in the eqemu database
type BaseData struct {
	//[ 0] level                                          uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	Level uint32 `gorm:"primary_key;column:level;type:uint;" json:"level"`
	//[ 1] class                                          uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	Class uint32 `gorm:"primary_key;column:class;type:uint;" json:"class"`
	//[ 2] hp                                             double               null: false  primary: false  isArray: false  auto: false  col: double          len: -1      default: []
	Hp float64 `gorm:"column:hp;type:double;" json:"hp"`
	//[ 3] mana                                           double               null: false  primary: false  isArray: false  auto: false  col: double          len: -1      default: []
	Mana float64 `gorm:"column:mana;type:double;" json:"mana"`
	//[ 4] end                                            double               null: false  primary: false  isArray: false  auto: false  col: double          len: -1      default: []
	End float64 `gorm:"column:end;type:double;" json:"end"`
	//[ 5] unk1                                           double               null: false  primary: false  isArray: false  auto: false  col: double          len: -1      default: []
	Unk1 float64 `gorm:"column:unk1;type:double;" json:"unk_1"`
	//[ 6] unk2                                           double               null: false  primary: false  isArray: false  auto: false  col: double          len: -1      default: []
	Unk2 float64 `gorm:"column:unk2;type:double;" json:"unk_2"`
	//[ 7] hp_fac                                         double               null: false  primary: false  isArray: false  auto: false  col: double          len: -1      default: []
	HpFac float64 `gorm:"column:hp_fac;type:double;" json:"hp_fac"`
	//[ 8] mana_fac                                       double               null: false  primary: false  isArray: false  auto: false  col: double          len: -1      default: []
	ManaFac float64 `gorm:"column:mana_fac;type:double;" json:"mana_fac"`
	//[ 9] end_fac                                        double               null: false  primary: false  isArray: false  auto: false  col: double          len: -1      default: []
	EndFac float64 `gorm:"column:end_fac;type:double;" json:"end_fac"`
}

var base_dataTableInfo = &TableInfo{
	Name: "base_data",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "level",
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
			GoFieldName:        "Level",
			GoFieldType:        "uint32",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "class",
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
			GoFieldName:        "Class",
			GoFieldType:        "uint32",
			JSONFieldName:      "class",
			ProtobufFieldName:  "class",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "hp",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "double",
			DatabaseTypePretty: "double",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "double",
			ColumnLength:       -1,
			GoFieldName:        "Hp",
			GoFieldType:        "float64",
			JSONFieldName:      "hp",
			ProtobufFieldName:  "hp",
			ProtobufType:       "float",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "mana",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "double",
			DatabaseTypePretty: "double",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "double",
			ColumnLength:       -1,
			GoFieldName:        "Mana",
			GoFieldType:        "float64",
			JSONFieldName:      "mana",
			ProtobufFieldName:  "mana",
			ProtobufType:       "float",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "end",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "double",
			DatabaseTypePretty: "double",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "double",
			ColumnLength:       -1,
			GoFieldName:        "End",
			GoFieldType:        "float64",
			JSONFieldName:      "end",
			ProtobufFieldName:  "end",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "unk1",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "double",
			DatabaseTypePretty: "double",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "double",
			ColumnLength:       -1,
			GoFieldName:        "Unk1",
			GoFieldType:        "float64",
			JSONFieldName:      "unk_1",
			ProtobufFieldName:  "unk_1",
			ProtobufType:       "float",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "unk2",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "double",
			DatabaseTypePretty: "double",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "double",
			ColumnLength:       -1,
			GoFieldName:        "Unk2",
			GoFieldType:        "float64",
			JSONFieldName:      "unk_2",
			ProtobufFieldName:  "unk_2",
			ProtobufType:       "float",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "hp_fac",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "double",
			DatabaseTypePretty: "double",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "double",
			ColumnLength:       -1,
			GoFieldName:        "HpFac",
			GoFieldType:        "float64",
			JSONFieldName:      "hp_fac",
			ProtobufFieldName:  "hp_fac",
			ProtobufType:       "float",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "mana_fac",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "double",
			DatabaseTypePretty: "double",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "double",
			ColumnLength:       -1,
			GoFieldName:        "ManaFac",
			GoFieldType:        "float64",
			JSONFieldName:      "mana_fac",
			ProtobufFieldName:  "mana_fac",
			ProtobufType:       "float",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "end_fac",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "double",
			DatabaseTypePretty: "double",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "double",
			ColumnLength:       -1,
			GoFieldName:        "EndFac",
			GoFieldType:        "float64",
			JSONFieldName:      "end_fac",
			ProtobufFieldName:  "end_fac",
			ProtobufType:       "float",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BaseData) TableName() string {
	return "base_data"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BaseData) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BaseData) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BaseData) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BaseData) TableInfo() *TableInfo {
	return base_dataTableInfo
}
