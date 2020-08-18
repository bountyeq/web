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


CREATE TABLE `character_pet_info` (
  `char_id` int(11) NOT NULL,
  `pet` int(11) NOT NULL,
  `petname` varchar(64) NOT NULL DEFAULT '',
  `petpower` int(11) NOT NULL DEFAULT 0,
  `spell_id` int(11) NOT NULL DEFAULT 0,
  `hp` int(11) NOT NULL DEFAULT 0,
  `mana` int(11) NOT NULL DEFAULT 0,
  `size` float NOT NULL DEFAULT 0,
  PRIMARY KEY (`char_id`,`pet`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "char_id": 23,    "pet": 54,    "petname": "WVVqwZsdYLREXeJVDBBCqkZNM",    "petpower": 10,    "spell_id": 72,    "hp": 28,    "mana": 56,    "size": 0.35944346}



*/

// CharacterPetInfo struct is a row record of the character_pet_info table in the eqemu database
type CharacterPetInfo struct {
	//[ 0] char_id                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	CharID int32 `gorm:"primary_key;column:char_id;type:int;" json:"char_id"`
	//[ 1] pet                                            int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	Pet int32 `gorm:"primary_key;column:pet;type:int;" json:"pet"`
	//[ 2] petname                                        varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Petname string `gorm:"column:petname;type:varchar;size:64;default:'';" json:"petname"`
	//[ 3] petpower                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Petpower int32 `gorm:"column:petpower;type:int;default:0;" json:"petpower"`
	//[ 4] spell_id                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SpellID int32 `gorm:"column:spell_id;type:int;default:0;" json:"spell_id"`
	//[ 5] hp                                             int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Hp int32 `gorm:"column:hp;type:int;default:0;" json:"hp"`
	//[ 6] mana                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Mana int32 `gorm:"column:mana;type:int;default:0;" json:"mana"`
	//[ 7] size                                           float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	Size float32 `gorm:"column:size;type:float;default:0;" json:"size"`
}

var character_pet_infoTableInfo = &TableInfo{
	Name: "character_pet_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "char_id",
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
			GoFieldName:        "CharID",
			GoFieldType:        "int32",
			JSONFieldName:      "char_id",
			ProtobufFieldName:  "char_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "pet",
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
			GoFieldName:        "Pet",
			GoFieldType:        "int32",
			JSONFieldName:      "pet",
			ProtobufFieldName:  "pet",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "petname",
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
			GoFieldName:        "Petname",
			GoFieldType:        "string",
			JSONFieldName:      "petname",
			ProtobufFieldName:  "petname",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "petpower",
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
			GoFieldName:        "Petpower",
			GoFieldType:        "int32",
			JSONFieldName:      "petpower",
			ProtobufFieldName:  "petpower",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "spell_id",
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
			GoFieldName:        "SpellID",
			GoFieldType:        "int32",
			JSONFieldName:      "spell_id",
			ProtobufFieldName:  "spell_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "hp",
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
			GoFieldName:        "Hp",
			GoFieldType:        "int32",
			JSONFieldName:      "hp",
			ProtobufFieldName:  "hp",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "mana",
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
			GoFieldName:        "Mana",
			GoFieldType:        "int32",
			JSONFieldName:      "mana",
			ProtobufFieldName:  "mana",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "size",
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
			GoFieldName:        "Size",
			GoFieldType:        "float32",
			JSONFieldName:      "size",
			ProtobufFieldName:  "size",
			ProtobufType:       "float",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharacterPetInfo) TableName() string {
	return "character_pet_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharacterPetInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharacterPetInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharacterPetInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharacterPetInfo) TableInfo() *TableInfo {
	return character_pet_infoTableInfo
}
