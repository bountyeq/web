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


CREATE TABLE `logsys_categories` (
  `log_category_id` int(11) NOT NULL,
  `log_category_description` varchar(150) DEFAULT NULL,
  `log_to_console` smallint(11) DEFAULT 0,
  `log_to_file` smallint(11) DEFAULT 0,
  `log_to_gmsay` smallint(11) DEFAULT 0,
  PRIMARY KEY (`log_category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "log_category_id": 35,    "log_category_description": "MWtAEpiGCBWNQIhSqNlJcoZmA",    "log_to_console": 11,    "log_to_file": 41,    "log_to_gmsay": 5}



*/

// LogsysCategories struct is a row record of the logsys_categories table in the eqemu database
type LogsysCategories struct {
	//[ 0] log_category_id                                int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	LogCategoryID int32 `gorm:"primary_key;column:log_category_id;type:int;" json:"log_category_id"`
	//[ 1] log_category_description                       varchar(150)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 150     default: [NULL]
	LogCategoryDescription null.String `gorm:"column:log_category_description;type:varchar;size:150;" json:"log_category_description"`
	//[ 2] log_to_console                                 smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	LogToConsole null.Int `gorm:"column:log_to_console;type:smallint;default:0;" json:"log_to_console"`
	//[ 3] log_to_file                                    smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	LogToFile null.Int `gorm:"column:log_to_file;type:smallint;default:0;" json:"log_to_file"`
	//[ 4] log_to_gmsay                                   smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	LogToGmsay null.Int `gorm:"column:log_to_gmsay;type:smallint;default:0;" json:"log_to_gmsay"`
}

var logsys_categoriesTableInfo = &TableInfo{
	Name: "logsys_categories",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "log_category_id",
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
			GoFieldName:        "LogCategoryID",
			GoFieldType:        "int32",
			JSONFieldName:      "log_category_id",
			ProtobufFieldName:  "log_category_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "log_category_description",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(150)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       150,
			GoFieldName:        "LogCategoryDescription",
			GoFieldType:        "null.String",
			JSONFieldName:      "log_category_description",
			ProtobufFieldName:  "log_category_description",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "log_to_console",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "LogToConsole",
			GoFieldType:        "null.Int",
			JSONFieldName:      "log_to_console",
			ProtobufFieldName:  "log_to_console",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "log_to_file",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "LogToFile",
			GoFieldType:        "null.Int",
			JSONFieldName:      "log_to_file",
			ProtobufFieldName:  "log_to_file",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "log_to_gmsay",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "LogToGmsay",
			GoFieldType:        "null.Int",
			JSONFieldName:      "log_to_gmsay",
			ProtobufFieldName:  "log_to_gmsay",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LogsysCategories) TableName() string {
	return "logsys_categories"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LogsysCategories) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LogsysCategories) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LogsysCategories) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LogsysCategories) TableInfo() *TableInfo {
	return logsys_categoriesTableInfo
}
