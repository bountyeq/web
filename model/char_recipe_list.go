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


CREATE TABLE `char_recipe_list` (
  `char_id` int(11) NOT NULL,
  `recipe_id` int(11) NOT NULL,
  `madecount` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`char_id`,`recipe_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "recipe_id": 7,    "madecount": 47,    "char_id": 36}



*/

// CharRecipeList struct is a row record of the char_recipe_list table in the eqemu database
type CharRecipeList struct {
	//[ 0] char_id                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	CharID int32 `gorm:"primary_key;column:char_id;type:int;" json:"char_id"`
	//[ 1] recipe_id                                      int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	RecipeID int32 `gorm:"primary_key;column:recipe_id;type:int;" json:"recipe_id"`
	//[ 2] madecount                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Madecount int32 `gorm:"column:madecount;type:int;default:0;" json:"madecount"`
}

var char_recipe_listTableInfo = &TableInfo{
	Name: "char_recipe_list",
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
			Name:               "recipe_id",
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
			GoFieldName:        "RecipeID",
			GoFieldType:        "int32",
			JSONFieldName:      "recipe_id",
			ProtobufFieldName:  "recipe_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "madecount",
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
			GoFieldName:        "Madecount",
			GoFieldType:        "int32",
			JSONFieldName:      "madecount",
			ProtobufFieldName:  "madecount",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CharRecipeList) TableName() string {
	return "char_recipe_list"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CharRecipeList) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CharRecipeList) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CharRecipeList) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CharRecipeList) TableInfo() *TableInfo {
	return char_recipe_listTableInfo
}
