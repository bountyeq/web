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


CREATE TABLE `spell_globals` (
  `spellid` int(11) NOT NULL,
  `spell_name` varchar(64) NOT NULL DEFAULT '',
  `qglobal` varchar(65) NOT NULL DEFAULT '',
  `value` varchar(65) NOT NULL DEFAULT '',
  PRIMARY KEY (`spellid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "spellid": 31,    "spell_name": "VdRIxqbXTawcwOekhmfBZicbr",    "qglobal": "NrrbhFQHwdYCqcGryDcEqVrPv",    "value": "agPoDyMAIAjWhZNHBPSRGDIRp"}



*/

// SpellGlobals struct is a row record of the spell_globals table in the eqemu database
type SpellGlobals struct {
	//[ 0] spellid                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	Spellid int32 `gorm:"primary_key;column:spellid;type:int;" json:"spellid"`
	//[ 1] spell_name                                     varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	SpellName string `gorm:"column:spell_name;type:varchar;size:64;default:'';" json:"spell_name"`
	//[ 2] qglobal                                        varchar(65)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 65      default: ['']
	Qglobal string `gorm:"column:qglobal;type:varchar;size:65;default:'';" json:"qglobal"`
	//[ 3] value                                          varchar(65)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 65      default: ['']
	Value string `gorm:"column:value;type:varchar;size:65;default:'';" json:"value"`
}

var spell_globalsTableInfo = &TableInfo{
	Name: "spell_globals",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "spellid",
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
			GoFieldName:        "Spellid",
			GoFieldType:        "int32",
			JSONFieldName:      "spellid",
			ProtobufFieldName:  "spellid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "spell_name",
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
			GoFieldName:        "SpellName",
			GoFieldType:        "string",
			JSONFieldName:      "spell_name",
			ProtobufFieldName:  "spell_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "qglobal",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(65)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       65,
			GoFieldName:        "Qglobal",
			GoFieldType:        "string",
			JSONFieldName:      "qglobal",
			ProtobufFieldName:  "qglobal",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(65)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       65,
			GoFieldName:        "Value",
			GoFieldType:        "string",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SpellGlobals) TableName() string {
	return "spell_globals"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SpellGlobals) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SpellGlobals) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SpellGlobals) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SpellGlobals) TableInfo() *TableInfo {
	return spell_globalsTableInfo
}
