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


CREATE TABLE `faction_list_mod` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `faction_id` int(10) unsigned NOT NULL,
  `mod` smallint(6) NOT NULL,
  `mod_name` varchar(16) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `faction_id_mod_name` (`faction_id`,`mod_name`)
) ENGINE=InnoDB AUTO_INCREMENT=11304 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "mod_name": "HjacmuxDyIkiSFxViENpSNQGJ",    "id": 68,    "faction_id": 97,    "mod": 52}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// FactionListMod struct is a row record of the faction_list_mod table in the eqemu database
type FactionListMod struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] faction_id                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	FactionID uint32 `gorm:"column:faction_id;type:uint;" json:"faction_id"`
	//[ 2] mod                                            smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: []
	Mod int32 `gorm:"column:mod;type:smallint;" json:"mod"`
	//[ 3] mod_name                                       varchar(16)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 16      default: []
	ModName string `gorm:"column:mod_name;type:varchar;size:16;" json:"mod_name"`
}

var faction_list_modTableInfo = &TableInfo{
	Name: "faction_list_mod",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "faction_id",
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
			GoFieldName:        "FactionID",
			GoFieldType:        "uint32",
			JSONFieldName:      "faction_id",
			ProtobufFieldName:  "faction_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "mod",
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
			GoFieldName:        "Mod",
			GoFieldType:        "int32",
			JSONFieldName:      "mod",
			ProtobufFieldName:  "mod",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "mod_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(16)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       16,
			GoFieldName:        "ModName",
			GoFieldType:        "string",
			JSONFieldName:      "mod_name",
			ProtobufFieldName:  "mod_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *FactionListMod) TableName() string {
	return "faction_list_mod"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FactionListMod) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FactionListMod) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FactionListMod) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FactionListMod) TableInfo() *TableInfo {
	return faction_list_modTableInfo
}
