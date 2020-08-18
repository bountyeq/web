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


CREATE TABLE `faction_base_data` (
  `client_faction_id` smallint(6) NOT NULL,
  `min` smallint(6) DEFAULT -2000,
  `max` smallint(6) DEFAULT 2000,
  `unk_hero1` smallint(6) DEFAULT NULL,
  `unk_hero2` smallint(6) DEFAULT NULL,
  `unk_hero3` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`client_faction_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "client_faction_id": 29,    "min": 2,    "max": 69,    "unk_hero_1": 14,    "unk_hero_2": 17,    "unk_hero_3": 44}



*/

// FactionBaseData struct is a row record of the faction_base_data table in the eqemu database
type FactionBaseData struct {
	//[ 0] client_faction_id                              smallint             null: false  primary: true   isArray: false  auto: false  col: smallint        len: -1      default: []
	ClientFactionID int32 `gorm:"primary_key;column:client_faction_id;type:smallint;" json:"client_faction_id"`
	//[ 1] min                                            smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [-2000]
	Min null.Int `gorm:"column:min;type:smallint;default:-2000;" json:"min"`
	//[ 2] max                                            smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [2000]
	Max null.Int `gorm:"column:max;type:smallint;default:2000;" json:"max"`
	//[ 3] unk_hero1                                      smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [NULL]
	UnkHero1 null.Int `gorm:"column:unk_hero1;type:smallint;" json:"unk_hero_1"`
	//[ 4] unk_hero2                                      smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [NULL]
	UnkHero2 null.Int `gorm:"column:unk_hero2;type:smallint;" json:"unk_hero_2"`
	//[ 5] unk_hero3                                      smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [NULL]
	UnkHero3 null.Int `gorm:"column:unk_hero3;type:smallint;" json:"unk_hero_3"`
}

var faction_base_dataTableInfo = &TableInfo{
	Name: "faction_base_data",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "client_faction_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "ClientFactionID",
			GoFieldType:        "int32",
			JSONFieldName:      "client_faction_id",
			ProtobufFieldName:  "client_faction_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "min",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "Min",
			GoFieldType:        "null.Int",
			JSONFieldName:      "min",
			ProtobufFieldName:  "min",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "max",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "Max",
			GoFieldType:        "null.Int",
			JSONFieldName:      "max",
			ProtobufFieldName:  "max",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "unk_hero1",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "UnkHero1",
			GoFieldType:        "null.Int",
			JSONFieldName:      "unk_hero_1",
			ProtobufFieldName:  "unk_hero_1",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "unk_hero2",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "UnkHero2",
			GoFieldType:        "null.Int",
			JSONFieldName:      "unk_hero_2",
			ProtobufFieldName:  "unk_hero_2",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "unk_hero3",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "UnkHero3",
			GoFieldType:        "null.Int",
			JSONFieldName:      "unk_hero_3",
			ProtobufFieldName:  "unk_hero_3",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *FactionBaseData) TableName() string {
	return "faction_base_data"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FactionBaseData) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FactionBaseData) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FactionBaseData) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FactionBaseData) TableInfo() *TableInfo {
	return faction_base_dataTableInfo
}
