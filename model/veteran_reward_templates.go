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


CREATE TABLE `veteran_reward_templates` (
  `claim_id` int(10) unsigned NOT NULL,
  `name` varchar(64) NOT NULL,
  `item_id` int(10) unsigned NOT NULL,
  `charges` smallint(5) unsigned NOT NULL,
  `reward_slot` tinyint(3) unsigned NOT NULL,
  UNIQUE KEY `claim_reward` (`claim_id`,`reward_slot`),
  KEY `claim_id` (`claim_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "claim_id": 61,    "name": "NkvPoNIjQlgRnBbnoHIjNZNNN",    "item_id": 10,    "charges": 52,    "reward_slot": 34}


Comments
-------------------------------------
[ 0] column is set for unsignedWarning table: veteran_reward_templates does not have a primary key defined, setting col position 1 claim_id as primary key

[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned



*/

// VeteranRewardTemplates struct is a row record of the veteran_reward_templates table in the eqemu database
type VeteranRewardTemplates struct {
	//[ 0] claim_id                                       uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	ClaimID uint32 `gorm:"primary_key;column:claim_id;type:uint;" json:"claim_id"`
	//[ 1] name                                           varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Name string `gorm:"column:name;type:varchar;size:64;" json:"name"`
	//[ 2] item_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	ItemID uint32 `gorm:"column:item_id;type:uint;" json:"item_id"`
	//[ 3] charges                                        usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: []
	Charges uint32 `gorm:"column:charges;type:usmallint;" json:"charges"`
	//[ 4] reward_slot                                    utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: []
	RewardSlot uint32 `gorm:"column:reward_slot;type:utinyint;" json:"reward_slot"`
}

var veteran_reward_templatesTableInfo = &TableInfo{
	Name: "veteran_reward_templates",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "claim_id",
			Comment: ``,
			Notes: `column is set for unsignedWarning table: veteran_reward_templates does not have a primary key defined, setting col position 1 claim_id as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ClaimID",
			GoFieldType:        "uint32",
			JSONFieldName:      "claim_id",
			ProtobufFieldName:  "claim_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "item_id",
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
			GoFieldName:        "ItemID",
			GoFieldType:        "uint32",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "charges",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "Charges",
			GoFieldType:        "uint32",
			JSONFieldName:      "charges",
			ProtobufFieldName:  "charges",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "reward_slot",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "RewardSlot",
			GoFieldType:        "uint32",
			JSONFieldName:      "reward_slot",
			ProtobufFieldName:  "reward_slot",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (v *VeteranRewardTemplates) TableName() string {
	return "veteran_reward_templates"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (v *VeteranRewardTemplates) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (v *VeteranRewardTemplates) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (v *VeteranRewardTemplates) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (v *VeteranRewardTemplates) TableInfo() *TableInfo {
	return veteran_reward_templatesTableInfo
}
