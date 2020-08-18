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


CREATE TABLE `db_version` (
  `version` int(11) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "version": 9}


Comments
-------------------------------------
[ 0] Warning table: db_version does not have a primary key defined, setting col position 1 version as primary key
Warning table: db_version primary key column version is nullable column, setting it as NOT NULL




*/

// DbVersion struct is a row record of the db_version table in the eqemu database
type DbVersion struct {
	//[ 0] version                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	Version int32 `gorm:"primary_key;column:version;type:int;default:0;" json:"version"`
}

var db_versionTableInfo = &TableInfo{
	Name: "db_version",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "version",
			Comment: ``,
			Notes: `Warning table: db_version does not have a primary key defined, setting col position 1 version as primary key
Warning table: db_version primary key column version is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Version",
			GoFieldType:        "int32",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *DbVersion) TableName() string {
	return "db_version"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DbVersion) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DbVersion) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DbVersion) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *DbVersion) TableInfo() *TableInfo {
	return db_versionTableInfo
}
