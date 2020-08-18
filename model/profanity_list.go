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


CREATE TABLE `profanity_list` (
  `word` varchar(16) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "word": "mtseCDKWJSfpsVLkVRlLwFWue"}


Comments
-------------------------------------
[ 0] Warning table: profanity_list does not have a primary key defined, setting col position 1 word as primary key




*/

// ProfanityList struct is a row record of the profanity_list table in the eqemu database
type ProfanityList struct {
	//[ 0] word                                           varchar(16)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 16      default: []
	Word string `gorm:"primary_key;column:word;type:varchar;size:16;" json:"word"`
}

var profanity_listTableInfo = &TableInfo{
	Name: "profanity_list",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "word",
			Comment: ``,
			Notes: `Warning table: profanity_list does not have a primary key defined, setting col position 1 word as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(16)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       16,
			GoFieldName:        "Word",
			GoFieldType:        "string",
			JSONFieldName:      "word",
			ProtobufFieldName:  "word",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *ProfanityList) TableName() string {
	return "profanity_list"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *ProfanityList) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *ProfanityList) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *ProfanityList) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *ProfanityList) TableInfo() *TableInfo {
	return profanity_listTableInfo
}
