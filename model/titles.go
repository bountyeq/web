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


CREATE TABLE `titles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `skill_id` tinyint(3) NOT NULL DEFAULT -1,
  `min_skill_value` mediumint(8) NOT NULL DEFAULT -1,
  `max_skill_value` mediumint(8) NOT NULL DEFAULT -1,
  `min_aa_points` mediumint(8) NOT NULL DEFAULT -1,
  `max_aa_points` mediumint(8) NOT NULL DEFAULT -1,
  `class` tinyint(4) NOT NULL DEFAULT -1,
  `gender` tinyint(1) NOT NULL DEFAULT -1 COMMENT '-1 = either, 0 = male, 1 = female',
  `char_id` int(11) NOT NULL DEFAULT -1,
  `status` int(11) NOT NULL DEFAULT -1,
  `item_id` int(11) NOT NULL DEFAULT -1,
  `prefix` varchar(32) NOT NULL DEFAULT '',
  `suffix` varchar(32) NOT NULL DEFAULT '',
  `title_set` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "skill_id": 30,    "min_skill_value": 13,    "max_aa_points": 51,    "char_id": 7,    "suffix": "XujcaqdEhkGlFccUScCrISDCm",    "title_set": 44,    "class": 31,    "item_id": 66,    "id": 50,    "gender": 18,    "status": 69,    "prefix": "dttERlPPVEmYCAJZyIPMiZneG",    "max_skill_value": 99,    "min_aa_points": 62}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// Titles struct is a row record of the titles table in the eqemu database
type Titles struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] skill_id                                       tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [-1]
	SkillID int32 `gorm:"column:skill_id;type:tinyint;default:-1;" json:"skill_id"`
	//[ 2] min_skill_value                                mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [-1]
	MinSkillValue int32 `gorm:"column:min_skill_value;type:mediumint;default:-1;" json:"min_skill_value"`
	//[ 3] max_skill_value                                mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [-1]
	MaxSkillValue int32 `gorm:"column:max_skill_value;type:mediumint;default:-1;" json:"max_skill_value"`
	//[ 4] min_aa_points                                  mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [-1]
	MinAaPoints int32 `gorm:"column:min_aa_points;type:mediumint;default:-1;" json:"min_aa_points"`
	//[ 5] max_aa_points                                  mediumint            null: false  primary: false  isArray: false  auto: false  col: mediumint       len: -1      default: [-1]
	MaxAaPoints int32 `gorm:"column:max_aa_points;type:mediumint;default:-1;" json:"max_aa_points"`
	//[ 6] class                                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [-1]
	Class int32 `gorm:"column:class;type:tinyint;default:-1;" json:"class"`
	//[ 7] gender                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [-1]
	Gender int32 `gorm:"column:gender;type:tinyint;default:-1;" json:"gender"` // -1 = either, 0 = male, 1 = female
	//[ 8] char_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	CharID int32 `gorm:"column:char_id;type:int;default:-1;" json:"char_id"`
	//[ 9] status                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	Status int32 `gorm:"column:status;type:int;default:-1;" json:"status"`
	//[10] item_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	ItemID int32 `gorm:"column:item_id;type:int;default:-1;" json:"item_id"`
	//[11] prefix                                         varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['']
	Prefix string `gorm:"column:prefix;type:varchar;size:32;default:'';" json:"prefix"`
	//[12] suffix                                         varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: ['']
	Suffix string `gorm:"column:suffix;type:varchar;size:32;default:'';" json:"suffix"`
	//[13] title_set                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TitleSet int32 `gorm:"column:title_set;type:int;default:0;" json:"title_set"`
}

var titlesTableInfo = &TableInfo{
	Name: "titles",
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
			Name:               "skill_id",
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
			GoFieldName:        "SkillID",
			GoFieldType:        "int32",
			JSONFieldName:      "skill_id",
			ProtobufFieldName:  "skill_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "min_skill_value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "mediumint",
			DatabaseTypePretty: "mediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "mediumint",
			ColumnLength:       -1,
			GoFieldName:        "MinSkillValue",
			GoFieldType:        "int32",
			JSONFieldName:      "min_skill_value",
			ProtobufFieldName:  "min_skill_value",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "max_skill_value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "mediumint",
			DatabaseTypePretty: "mediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "mediumint",
			ColumnLength:       -1,
			GoFieldName:        "MaxSkillValue",
			GoFieldType:        "int32",
			JSONFieldName:      "max_skill_value",
			ProtobufFieldName:  "max_skill_value",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "min_aa_points",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "mediumint",
			DatabaseTypePretty: "mediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "mediumint",
			ColumnLength:       -1,
			GoFieldName:        "MinAaPoints",
			GoFieldType:        "int32",
			JSONFieldName:      "min_aa_points",
			ProtobufFieldName:  "min_aa_points",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "max_aa_points",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "mediumint",
			DatabaseTypePretty: "mediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "mediumint",
			ColumnLength:       -1,
			GoFieldName:        "MaxAaPoints",
			GoFieldType:        "int32",
			JSONFieldName:      "max_aa_points",
			ProtobufFieldName:  "max_aa_points",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "class",
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
			GoFieldName:        "Class",
			GoFieldType:        "int32",
			JSONFieldName:      "class",
			ProtobufFieldName:  "class",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "gender",
			Comment:            `-1 = either, 0 = male, 1 = female`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "char_id",
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
			GoFieldName:        "CharID",
			GoFieldType:        "int32",
			JSONFieldName:      "char_id",
			ProtobufFieldName:  "char_id",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "status",
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
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "prefix",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Prefix",
			GoFieldType:        "string",
			JSONFieldName:      "prefix",
			ProtobufFieldName:  "prefix",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "suffix",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Suffix",
			GoFieldType:        "string",
			JSONFieldName:      "suffix",
			ProtobufFieldName:  "suffix",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "title_set",
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
			GoFieldName:        "TitleSet",
			GoFieldType:        "int32",
			JSONFieldName:      "title_set",
			ProtobufFieldName:  "title_set",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *Titles) TableName() string {
	return "titles"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *Titles) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *Titles) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *Titles) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *Titles) TableInfo() *TableInfo {
	return titlesTableInfo
}
