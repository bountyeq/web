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


CREATE TABLE `inventory_versions` (
  `version` int(11) unsigned NOT NULL DEFAULT 0,
  `step` int(11) unsigned NOT NULL DEFAULT 0,
  `bot_step` int(11) unsigned NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "step": 43,    "bot_step": 20,    "version": 84}


Comments
-------------------------------------
[ 0] column is set for unsignedWarning table: inventory_versions does not have a primary key defined, setting col position 1 version as primary key

[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// InventoryVersions struct is a row record of the inventory_versions table in the eqemu database
type InventoryVersions struct {
	//[ 0] version                                        uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Version uint32 `gorm:"primary_key;column:version;type:uint;default:0;" json:"version"`
	//[ 1] step                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Step uint32 `gorm:"column:step;type:uint;default:0;" json:"step"`
	//[ 2] bot_step                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	BotStep uint32 `gorm:"column:bot_step;type:uint;default:0;" json:"bot_step"`
}

var inventory_versionsTableInfo = &TableInfo{
	Name: "inventory_versions",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "version",
			Comment: ``,
			Notes: `column is set for unsignedWarning table: inventory_versions does not have a primary key defined, setting col position 1 version as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Version",
			GoFieldType:        "uint32",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "step",
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
			GoFieldName:        "Step",
			GoFieldType:        "uint32",
			JSONFieldName:      "step",
			ProtobufFieldName:  "step",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "bot_step",
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
			GoFieldName:        "BotStep",
			GoFieldType:        "uint32",
			JSONFieldName:      "bot_step",
			ProtobufFieldName:  "bot_step",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (i *InventoryVersions) TableName() string {
	return "inventory_versions"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *InventoryVersions) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *InventoryVersions) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *InventoryVersions) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *InventoryVersions) TableInfo() *TableInfo {
	return inventory_versionsTableInfo
}
