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


CREATE TABLE `tradeskill_recipe_entries` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `recipe_id` int(11) NOT NULL DEFAULT 0,
  `item_id` int(11) NOT NULL DEFAULT 0,
  `successcount` tinyint(2) NOT NULL DEFAULT 0,
  `failcount` tinyint(2) NOT NULL DEFAULT 0,
  `componentcount` tinyint(2) NOT NULL DEFAULT 1,
  `salvagecount` tinyint(2) NOT NULL DEFAULT 0,
  `iscontainer` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `recipe_id` (`recipe_id`),
  KEY `item_id` (`item_id`)
) ENGINE=InnoDB AUTO_INCREMENT=277755 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "recipe_id": 37,    "item_id": 72,    "successcount": 85,    "failcount": 8,    "componentcount": 63,    "salvagecount": 2,    "iscontainer": 95,    "id": 81}



*/

// TradeskillRecipeEntries struct is a row record of the tradeskill_recipe_entries table in the eqemu database
type TradeskillRecipeEntries struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] recipe_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RecipeID int32 `gorm:"column:recipe_id;type:int;default:0;" json:"recipe_id"`
	//[ 2] item_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ItemID int32 `gorm:"column:item_id;type:int;default:0;" json:"item_id"`
	//[ 3] successcount                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Successcount int32 `gorm:"column:successcount;type:tinyint;default:0;" json:"successcount"`
	//[ 4] failcount                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Failcount int32 `gorm:"column:failcount;type:tinyint;default:0;" json:"failcount"`
	//[ 5] componentcount                                 tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Componentcount int32 `gorm:"column:componentcount;type:tinyint;default:1;" json:"componentcount"`
	//[ 6] salvagecount                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Salvagecount int32 `gorm:"column:salvagecount;type:tinyint;default:0;" json:"salvagecount"`
	//[ 7] iscontainer                                    tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Iscontainer int32 `gorm:"column:iscontainer;type:tinyint;default:0;" json:"iscontainer"`
}

var tradeskill_recipe_entriesTableInfo = &TableInfo{
	Name: "tradeskill_recipe_entries",
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
			Name:               "recipe_id",
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
			GoFieldName:        "RecipeID",
			GoFieldType:        "int32",
			JSONFieldName:      "recipe_id",
			ProtobufFieldName:  "recipe_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "item_id",
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
			GoFieldName:        "ItemID",
			GoFieldType:        "int32",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "successcount",
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
			GoFieldName:        "Successcount",
			GoFieldType:        "int32",
			JSONFieldName:      "successcount",
			ProtobufFieldName:  "successcount",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "failcount",
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
			GoFieldName:        "Failcount",
			GoFieldType:        "int32",
			JSONFieldName:      "failcount",
			ProtobufFieldName:  "failcount",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "componentcount",
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
			GoFieldName:        "Componentcount",
			GoFieldType:        "int32",
			JSONFieldName:      "componentcount",
			ProtobufFieldName:  "componentcount",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "salvagecount",
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
			GoFieldName:        "Salvagecount",
			GoFieldType:        "int32",
			JSONFieldName:      "salvagecount",
			ProtobufFieldName:  "salvagecount",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "iscontainer",
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
			GoFieldName:        "Iscontainer",
			GoFieldType:        "int32",
			JSONFieldName:      "iscontainer",
			ProtobufFieldName:  "iscontainer",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TradeskillRecipeEntries) TableName() string {
	return "tradeskill_recipe_entries"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TradeskillRecipeEntries) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TradeskillRecipeEntries) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TradeskillRecipeEntries) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TradeskillRecipeEntries) TableInfo() *TableInfo {
	return tradeskill_recipe_entriesTableInfo
}
