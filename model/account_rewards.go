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


CREATE TABLE `account_rewards` (
  `account_id` int(10) unsigned NOT NULL,
  `reward_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  UNIQUE KEY `account_reward` (`account_id`,`reward_id`),
  KEY `account_id` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "account_id": 30,    "reward_id": 66,    "amount": 44}


Comments
-------------------------------------
[ 0] column is set for unsignedWarning table: account_rewards does not have a primary key defined, setting col position 1 account_id as primary key

[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// AccountRewards struct is a row record of the account_rewards table in the eqemu database
type AccountRewards struct {
	//[ 0] account_id                                     uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	AccountID uint32 `gorm:"primary_key;column:account_id;type:uint;" json:"account_id"`
	//[ 1] reward_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	RewardID uint32 `gorm:"column:reward_id;type:uint;" json:"reward_id"`
	//[ 2] amount                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	Amount uint32 `gorm:"column:amount;type:uint;" json:"amount"`
}

var account_rewardsTableInfo = &TableInfo{
	Name: "account_rewards",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "account_id",
			Comment: ``,
			Notes: `column is set for unsignedWarning table: account_rewards does not have a primary key defined, setting col position 1 account_id as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "AccountID",
			GoFieldType:        "uint32",
			JSONFieldName:      "account_id",
			ProtobufFieldName:  "account_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "reward_id",
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
			GoFieldName:        "RewardID",
			GoFieldType:        "uint32",
			JSONFieldName:      "reward_id",
			ProtobufFieldName:  "reward_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "amount",
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
			GoFieldName:        "Amount",
			GoFieldType:        "uint32",
			JSONFieldName:      "amount",
			ProtobufFieldName:  "amount",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AccountRewards) TableName() string {
	return "account_rewards"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AccountRewards) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AccountRewards) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AccountRewards) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AccountRewards) TableInfo() *TableInfo {
	return account_rewardsTableInfo
}
