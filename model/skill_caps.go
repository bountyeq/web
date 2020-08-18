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


CREATE TABLE `skill_caps` (
  `skillID` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `class` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `level` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `cap` mediumint(8) unsigned NOT NULL DEFAULT 0,
  `class_` tinyint(3) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`skillID`,`class`,`level`,`class_`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "skill_id": 90,    "class": 9,    "level": 49,    "cap": 55,    "class_": 47}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned



*/

// SkillCaps struct is a row record of the skill_caps table in the eqemu database
type SkillCaps struct {
	//[ 0] skillID                                        utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	SkillID uint32 `gorm:"primary_key;column:skillID;type:utinyint;default:0;" json:"skill_id"`
	//[ 1] class                                          utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Class uint32 `gorm:"primary_key;column:class;type:utinyint;default:0;" json:"class"`
	//[ 2] level                                          utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Level uint32 `gorm:"primary_key;column:level;type:utinyint;default:0;" json:"level"`
	//[ 3] cap                                            umediumint           null: false  primary: false  isArray: false  auto: false  col: umediumint      len: -1      default: [0]
	Cap uint32 `gorm:"column:cap;type:umediumint;default:0;" json:"cap"`
	//[ 4] class_                                         utinyint             null: false  primary: true   isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Classalt1 uint32 `gorm:"primary_key;column:class_;type:utinyint;default:0;" json:"class_"`
}

var skill_capsTableInfo = &TableInfo{
	Name: "skill_caps",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "skillID",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "SkillID",
			GoFieldType:        "uint32",
			JSONFieldName:      "skill_id",
			ProtobufFieldName:  "skill_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "class",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Class",
			GoFieldType:        "uint32",
			JSONFieldName:      "class",
			ProtobufFieldName:  "class",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "level",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Level",
			GoFieldType:        "uint32",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "cap",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "umediumint",
			DatabaseTypePretty: "umediumint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "umediumint",
			ColumnLength:       -1,
			GoFieldName:        "Cap",
			GoFieldType:        "uint32",
			JSONFieldName:      "cap",
			ProtobufFieldName:  "cap",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "class_",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Classalt1",
			GoFieldType:        "uint32",
			JSONFieldName:      "class_",
			ProtobufFieldName:  "class_",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SkillCaps) TableName() string {
	return "skill_caps"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SkillCaps) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SkillCaps) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SkillCaps) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SkillCaps) TableInfo() *TableInfo {
	return skill_capsTableInfo
}
