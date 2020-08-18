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


CREATE TABLE `chatchannels` (
  `name` varchar(64) NOT NULL DEFAULT '',
  `owner` varchar(64) NOT NULL DEFAULT '',
  `password` varchar(64) NOT NULL DEFAULT '',
  `minstatus` int(5) NOT NULL DEFAULT 0,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "name": "LkYHcAFQmEuuHsRcblUDoWYAY",    "owner": "WsaayMaiFjSrWChjaoCOgaLxi",    "password": "iVRlVwHmuKGXaOkqHbZIsnDPI",    "minstatus": 85}



*/

// Chatchannels struct is a row record of the chatchannels table in the eqemu database
type Chatchannels struct {
	//[ 0] name                                           varchar(64)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Name string `gorm:"primary_key;column:name;type:varchar;size:64;default:'';" json:"name"`
	//[ 1] owner                                          varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Owner string `gorm:"column:owner;type:varchar;size:64;default:'';" json:"owner"`
	//[ 2] password                                       varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: ['']
	Password string `gorm:"column:password;type:varchar;size:64;default:'';" json:"password"`
	//[ 3] minstatus                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Minstatus int32 `gorm:"column:minstatus;type:int;default:0;" json:"minstatus"`
}

var chatchannelsTableInfo = &TableInfo{
	Name: "chatchannels",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "owner",
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
			GoFieldName:        "Owner",
			GoFieldType:        "string",
			JSONFieldName:      "owner",
			ProtobufFieldName:  "owner",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "password",
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
			GoFieldName:        "Password",
			GoFieldType:        "string",
			JSONFieldName:      "password",
			ProtobufFieldName:  "password",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "minstatus",
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
			GoFieldName:        "Minstatus",
			GoFieldType:        "int32",
			JSONFieldName:      "minstatus",
			ProtobufFieldName:  "minstatus",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *Chatchannels) TableName() string {
	return "chatchannels"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Chatchannels) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Chatchannels) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Chatchannels) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *Chatchannels) TableInfo() *TableInfo {
	return chatchannelsTableInfo
}
