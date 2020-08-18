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


CREATE TABLE `horses` (
  `filename` varchar(32) NOT NULL DEFAULT '',
  `race` smallint(3) NOT NULL DEFAULT 216,
  `gender` tinyint(1) NOT NULL DEFAULT 0,
  `texture` tinyint(2) NOT NULL DEFAULT 0,
  `mountspeed` float(4,2) NOT NULL DEFAULT 0.75,
  `notes` varchar(64) DEFAULT 'Notes',
  PRIMARY KEY (`filename`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "mountspeed": 0.93963003,    "notes": "cFENagaucwPQhuHQrPlJAZMsv",    "filename": "oLCNYbZJqYSqmYhRibEDdSlKr",    "race": 46,    "gender": 71,    "texture": 58}



*/

// Horses struct is a row record of the horses table in the eqemu database
type Horses struct {
	//[ 0] filename                                       varchar(32)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 32      default: ['']
	Filename string `gorm:"primary_key;column:filename;type:varchar;size:32;default:'';" json:"filename"`
	//[ 1] race                                           smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [216]
	Race int32 `gorm:"column:race;type:smallint;default:216;" json:"race"`
	//[ 2] gender                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Gender int32 `gorm:"column:gender;type:tinyint;default:0;" json:"gender"`
	//[ 3] texture                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Texture int32 `gorm:"column:texture;type:tinyint;default:0;" json:"texture"`
	//[ 4] mountspeed                                     float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0.75]
	Mountspeed float32 `gorm:"column:mountspeed;type:float;default:0.75;" json:"mountspeed"`
	//[ 5] notes                                          varchar(64)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['Notes']
	Notes null.String `gorm:"column:notes;type:varchar;size:64;default:'Notes';" json:"notes"`
}

var horsesTableInfo = &TableInfo{
	Name: "horses",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "filename",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Filename",
			GoFieldType:        "string",
			JSONFieldName:      "filename",
			ProtobufFieldName:  "filename",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "race",
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
			GoFieldName:        "Race",
			GoFieldType:        "int32",
			JSONFieldName:      "race",
			ProtobufFieldName:  "race",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "gender",
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
			GoFieldName:        "Gender",
			GoFieldType:        "int32",
			JSONFieldName:      "gender",
			ProtobufFieldName:  "gender",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "texture",
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
			GoFieldName:        "Texture",
			GoFieldType:        "int32",
			JSONFieldName:      "texture",
			ProtobufFieldName:  "texture",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "mountspeed",
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
			GoFieldName:        "Mountspeed",
			GoFieldType:        "float32",
			JSONFieldName:      "mountspeed",
			ProtobufFieldName:  "mountspeed",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "notes",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Notes",
			GoFieldType:        "null.String",
			JSONFieldName:      "notes",
			ProtobufFieldName:  "notes",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (h *Horses) TableName() string {
	return "horses"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (h *Horses) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (h *Horses) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (h *Horses) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (h *Horses) TableInfo() *TableInfo {
	return horsesTableInfo
}
