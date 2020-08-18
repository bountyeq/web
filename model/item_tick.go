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


CREATE TABLE `item_tick` (
  `it_itemid` int(11) NOT NULL,
  `it_chance` int(11) NOT NULL,
  `it_level` int(11) NOT NULL,
  `it_id` int(11) NOT NULL AUTO_INCREMENT,
  `it_qglobal` varchar(50) NOT NULL,
  `it_bagslot` tinyint(4) NOT NULL,
  PRIMARY KEY (`it_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "it_itemid": 65,    "it_chance": 12,    "it_level": 97,    "it_id": 97,    "it_qglobal": "YDZZIvUmUuyTbDYgSMbXaJjIc",    "it_bagslot": 63}



*/

// ItemTick struct is a row record of the item_tick table in the eqemu database
type ItemTick struct {
	//[ 0] it_itemid                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ItItemid int32 `gorm:"column:it_itemid;type:int;" json:"it_itemid"`
	//[ 1] it_chance                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ItChance int32 `gorm:"column:it_chance;type:int;" json:"it_chance"`
	//[ 2] it_level                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ItLevel int32 `gorm:"column:it_level;type:int;" json:"it_level"`
	//[ 3] it_id                                          int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ItID int32 `gorm:"primary_key;AUTO_INCREMENT;column:it_id;type:int;" json:"it_id"`
	//[ 4] it_qglobal                                     varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	ItQglobal string `gorm:"column:it_qglobal;type:varchar;size:50;" json:"it_qglobal"`
	//[ 5] it_bagslot                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	ItBagslot int32 `gorm:"column:it_bagslot;type:tinyint;" json:"it_bagslot"`
}

var item_tickTableInfo = &TableInfo{
	Name: "item_tick",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "it_itemid",
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
			GoFieldName:        "ItItemid",
			GoFieldType:        "int32",
			JSONFieldName:      "it_itemid",
			ProtobufFieldName:  "it_itemid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "it_chance",
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
			GoFieldName:        "ItChance",
			GoFieldType:        "int32",
			JSONFieldName:      "it_chance",
			ProtobufFieldName:  "it_chance",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "it_level",
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
			GoFieldName:        "ItLevel",
			GoFieldType:        "int32",
			JSONFieldName:      "it_level",
			ProtobufFieldName:  "it_level",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "it_id",
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
			GoFieldName:        "ItID",
			GoFieldType:        "int32",
			JSONFieldName:      "it_id",
			ProtobufFieldName:  "it_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "it_qglobal",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "ItQglobal",
			GoFieldType:        "string",
			JSONFieldName:      "it_qglobal",
			ProtobufFieldName:  "it_qglobal",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "it_bagslot",
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
			GoFieldName:        "ItBagslot",
			GoFieldType:        "int32",
			JSONFieldName:      "it_bagslot",
			ProtobufFieldName:  "it_bagslot",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (i *ItemTick) TableName() string {
	return "item_tick"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *ItemTick) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *ItemTick) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *ItemTick) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *ItemTick) TableInfo() *TableInfo {
	return item_tickTableInfo
}
