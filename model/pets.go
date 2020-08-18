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


CREATE TABLE `pets` (
  `type` varchar(64) NOT NULL DEFAULT '',
  `petpower` int(11) NOT NULL DEFAULT 0,
  `npcID` int(11) NOT NULL DEFAULT 0,
  `temp` tinyint(4) NOT NULL DEFAULT 0,
  `petcontrol` tinyint(4) NOT NULL DEFAULT 0,
  `petnaming` tinyint(4) NOT NULL DEFAULT 0,
  `monsterflag` tinyint(4) NOT NULL DEFAULT 0,
  `equipmentset` int(11) NOT NULL DEFAULT -1,
  PRIMARY KEY (`type`,`petpower`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "equipmentset": 24,    "type": "MMhyvRCRMmjEpPZYUnAAcCDbn",    "petpower": 92,    "npc_id": 65,    "temp": 52,    "petcontrol": 50,    "petnaming": 17,    "monsterflag": 37}



*/

// Pets struct is a row record of the pets table in the eqemu database
type Pets struct {
	//[ 0] type                                           varchar(64)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Type string `gorm:"primary_key;column:type;type:varchar;size:64;default:'';" json:"type"`
	//[ 1] petpower                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Petpower int32 `gorm:"primary_key;column:petpower;type:int;default:0;" json:"petpower"`
	//[ 2] npcID                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	NpcID int32 `gorm:"column:npcID;type:int;default:0;" json:"npc_id"`
	//[ 3] temp                                           tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Temp int32 `gorm:"column:temp;type:tinyint;default:0;" json:"temp"`
	//[ 4] petcontrol                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Petcontrol int32 `gorm:"column:petcontrol;type:tinyint;default:0;" json:"petcontrol"`
	//[ 5] petnaming                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Petnaming int32 `gorm:"column:petnaming;type:tinyint;default:0;" json:"petnaming"`
	//[ 6] monsterflag                                    tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Monsterflag int32 `gorm:"column:monsterflag;type:tinyint;default:0;" json:"monsterflag"`
	//[ 7] equipmentset                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	Equipmentset int32 `gorm:"column:equipmentset;type:int;default:-1;" json:"equipmentset"`
}

var petsTableInfo = &TableInfo{
	Name: "pets",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Type",
			GoFieldType:        "string",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "petpower",
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
			GoFieldName:        "Petpower",
			GoFieldType:        "int32",
			JSONFieldName:      "petpower",
			ProtobufFieldName:  "petpower",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "npcID",
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
			GoFieldName:        "NpcID",
			GoFieldType:        "int32",
			JSONFieldName:      "npc_id",
			ProtobufFieldName:  "npc_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "temp",
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
			GoFieldName:        "Temp",
			GoFieldType:        "int32",
			JSONFieldName:      "temp",
			ProtobufFieldName:  "temp",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "petcontrol",
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
			GoFieldName:        "Petcontrol",
			GoFieldType:        "int32",
			JSONFieldName:      "petcontrol",
			ProtobufFieldName:  "petcontrol",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "petnaming",
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
			GoFieldName:        "Petnaming",
			GoFieldType:        "int32",
			JSONFieldName:      "petnaming",
			ProtobufFieldName:  "petnaming",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "monsterflag",
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
			GoFieldName:        "Monsterflag",
			GoFieldType:        "int32",
			JSONFieldName:      "monsterflag",
			ProtobufFieldName:  "monsterflag",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "equipmentset",
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
			GoFieldName:        "Equipmentset",
			GoFieldType:        "int32",
			JSONFieldName:      "equipmentset",
			ProtobufFieldName:  "equipmentset",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *Pets) TableName() string {
	return "pets"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Pets) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Pets) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Pets) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Pets) TableInfo() *TableInfo {
	return petsTableInfo
}
