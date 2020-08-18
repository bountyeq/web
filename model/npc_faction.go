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


CREATE TABLE `npc_faction` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` tinytext DEFAULT NULL,
  `primaryfaction` int(11) NOT NULL DEFAULT 0,
  `ignore_primary_assist` tinyint(3) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1459 DEFAULT CHARSET=latin1 PACK_KEYS=0

JSON Sample
-------------------------------------
{    "id": 80,    "name": "MDquZOpAevFIRTqNNFsDdBevf",    "primaryfaction": 51,    "ignore_primary_assist": 46}



*/

// NpcFaction struct is a row record of the npc_faction table in the eqemu database
type NpcFaction struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] name                                           text(255)            null: true   primary: false  isArray: false  auto: false  col: text            len: 255     default: [NULL]
	Name null.String `gorm:"column:name;type:text;size:255;" json:"name"`
	//[ 2] primaryfaction                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Primaryfaction int32 `gorm:"column:primaryfaction;type:int;default:0;" json:"primaryfaction"`
	//[ 3] ignore_primary_assist                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IgnorePrimaryAssist int32 `gorm:"column:ignore_primary_assist;type:tinyint;default:0;" json:"ignore_primary_assist"`
}

var npc_factionTableInfo = &TableInfo{
	Name: "npc_faction",
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
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "null.String",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "primaryfaction",
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
			GoFieldName:        "Primaryfaction",
			GoFieldType:        "int32",
			JSONFieldName:      "primaryfaction",
			ProtobufFieldName:  "primaryfaction",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "ignore_primary_assist",
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
			GoFieldName:        "IgnorePrimaryAssist",
			GoFieldType:        "int32",
			JSONFieldName:      "ignore_primary_assist",
			ProtobufFieldName:  "ignore_primary_assist",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (n *NpcFaction) TableName() string {
	return "npc_faction"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (n *NpcFaction) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (n *NpcFaction) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (n *NpcFaction) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (n *NpcFaction) TableInfo() *TableInfo {
	return npc_factionTableInfo
}
