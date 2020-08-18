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


CREATE TABLE `task_activities` (
  `taskid` int(11) unsigned NOT NULL DEFAULT 0,
  `activityid` int(11) unsigned NOT NULL DEFAULT 0,
  `step` int(11) unsigned NOT NULL DEFAULT 0,
  `activitytype` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `target_name` varchar(64) NOT NULL DEFAULT '',
  `item_list` varchar(128) NOT NULL DEFAULT '',
  `skill_list` varchar(64) NOT NULL DEFAULT '-1',
  `spell_list` varchar(64) NOT NULL DEFAULT '0',
  `description_override` varchar(128) NOT NULL DEFAULT '',
  `goalid` int(11) unsigned NOT NULL DEFAULT 0,
  `goalmethod` int(10) unsigned NOT NULL DEFAULT 0,
  `goalcount` int(11) DEFAULT 1,
  `delivertonpc` int(11) unsigned NOT NULL DEFAULT 0,
  `zones` varchar(64) NOT NULL DEFAULT '',
  `optional` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`taskid`,`activityid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "description_override": "aqNeGkcWXWVMyBuJXueyQKyRX",    "delivertonpc": 79,    "skill_list": "njrBRfETZPWsrFyKqKqACWCmb",    "goalcount": 17,    "taskid": 26,    "activityid": 92,    "target_name": "sVhJqENUfOesjhnemuwGlXnwZ",    "item_list": "KSbOmWbgyZjwnVZcKSdwZtbuY",    "goalmethod": 33,    "zones": "qxasvMmotEPcjiNNMViWGQLDm",    "optional": 42,    "step": 65,    "activitytype": 99,    "spell_list": "yUdtriByLytJlTHnZYfxVYCYt",    "goalid": 58}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned
[12] column is set for unsigned



*/

// TaskActivities struct is a row record of the task_activities table in the eqemu database
type TaskActivities struct {
	//[ 0] taskid                                         uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Taskid uint32 `gorm:"primary_key;column:taskid;type:uint;default:0;" json:"taskid"`
	//[ 1] activityid                                     uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: [0]
	Activityid uint32 `gorm:"primary_key;column:activityid;type:uint;default:0;" json:"activityid"`
	//[ 2] step                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Step uint32 `gorm:"column:step;type:uint;default:0;" json:"step"`
	//[ 3] activitytype                                   utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Activitytype uint32 `gorm:"column:activitytype;type:utinyint;default:0;" json:"activitytype"`
	//[ 4] target_name                                    varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	TargetName string `gorm:"column:target_name;type:varchar;size:64;default:'';" json:"target_name"`
	//[ 5] item_list                                      varchar(128)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 128     default: ['']
	ItemList string `gorm:"column:item_list;type:varchar;size:128;default:'';" json:"item_list"`
	//[ 6] skill_list                                     varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['-1']
	SkillList string `gorm:"column:skill_list;type:varchar;size:64;default:'-1';" json:"skill_list"`
	//[ 7] spell_list                                     varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['0']
	SpellList string `gorm:"column:spell_list;type:varchar;size:64;default:'0';" json:"spell_list"`
	//[ 8] description_override                           varchar(128)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 128     default: ['']
	DescriptionOverride string `gorm:"column:description_override;type:varchar;size:128;default:'';" json:"description_override"`
	//[ 9] goalid                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Goalid uint32 `gorm:"column:goalid;type:uint;default:0;" json:"goalid"`
	//[10] goalmethod                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Goalmethod uint32 `gorm:"column:goalmethod;type:uint;default:0;" json:"goalmethod"`
	//[11] goalcount                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	Goalcount null.Int `gorm:"column:goalcount;type:int;default:1;" json:"goalcount"`
	//[12] delivertonpc                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Delivertonpc uint32 `gorm:"column:delivertonpc;type:uint;default:0;" json:"delivertonpc"`
	//[13] zones                                          varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Zones string `gorm:"column:zones;type:varchar;size:64;default:'';" json:"zones"`
	//[14] optional                                       tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Optional int32 `gorm:"column:optional;type:tinyint;default:0;" json:"optional"`
}

var task_activitiesTableInfo = &TableInfo{
	Name: "task_activities",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "taskid",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Taskid",
			GoFieldType:        "uint32",
			JSONFieldName:      "taskid",
			ProtobufFieldName:  "taskid",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "activityid",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Activityid",
			GoFieldType:        "uint32",
			JSONFieldName:      "activityid",
			ProtobufFieldName:  "activityid",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "activitytype",
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
			GoFieldName:        "Activitytype",
			GoFieldType:        "uint32",
			JSONFieldName:      "activitytype",
			ProtobufFieldName:  "activitytype",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "target_name",
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
			GoFieldName:        "TargetName",
			GoFieldType:        "string",
			JSONFieldName:      "target_name",
			ProtobufFieldName:  "target_name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "item_list",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       128,
			GoFieldName:        "ItemList",
			GoFieldType:        "string",
			JSONFieldName:      "item_list",
			ProtobufFieldName:  "item_list",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "skill_list",
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
			GoFieldName:        "SkillList",
			GoFieldType:        "string",
			JSONFieldName:      "skill_list",
			ProtobufFieldName:  "skill_list",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "spell_list",
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
			GoFieldName:        "SpellList",
			GoFieldType:        "string",
			JSONFieldName:      "spell_list",
			ProtobufFieldName:  "spell_list",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "description_override",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       128,
			GoFieldName:        "DescriptionOverride",
			GoFieldType:        "string",
			JSONFieldName:      "description_override",
			ProtobufFieldName:  "description_override",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "goalid",
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
			GoFieldName:        "Goalid",
			GoFieldType:        "uint32",
			JSONFieldName:      "goalid",
			ProtobufFieldName:  "goalid",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "goalmethod",
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
			GoFieldName:        "Goalmethod",
			GoFieldType:        "uint32",
			JSONFieldName:      "goalmethod",
			ProtobufFieldName:  "goalmethod",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "goalcount",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Goalcount",
			GoFieldType:        "null.Int",
			JSONFieldName:      "goalcount",
			ProtobufFieldName:  "goalcount",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "delivertonpc",
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
			GoFieldName:        "Delivertonpc",
			GoFieldType:        "uint32",
			JSONFieldName:      "delivertonpc",
			ProtobufFieldName:  "delivertonpc",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "zones",
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
			GoFieldName:        "Zones",
			GoFieldType:        "string",
			JSONFieldName:      "zones",
			ProtobufFieldName:  "zones",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "optional",
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
			GoFieldName:        "Optional",
			GoFieldType:        "int32",
			JSONFieldName:      "optional",
			ProtobufFieldName:  "optional",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TaskActivities) TableName() string {
	return "task_activities"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TaskActivities) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TaskActivities) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TaskActivities) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TaskActivities) TableInfo() *TableInfo {
	return task_activitiesTableInfo
}
