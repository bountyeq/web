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


CREATE TABLE `qs_player_aa_rate_hourly` (
  `char_id` int(11) NOT NULL DEFAULT 0,
  `hour_time` int(11) NOT NULL,
  `aa_count` varchar(11) DEFAULT NULL,
  PRIMARY KEY (`char_id`,`hour_time`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "char_id": 38,    "hour_time": 61,    "aa_count": "AqkQohybxJMCgoWnsuSZxoDkN"}



*/

// QsPlayerAaRateHourly struct is a row record of the qs_player_aa_rate_hourly table in the eqemu database
type QsPlayerAaRateHourly struct {
	//[ 0] char_id                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	CharID int32 `gorm:"primary_key;column:char_id;type:int;default:0;" json:"char_id"`
	//[ 1] hour_time                                      int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	HourTime int32 `gorm:"primary_key;column:hour_time;type:int;" json:"hour_time"`
	//[ 2] aa_count                                       varchar(11)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 11      default: [NULL]
	AaCount null.String `gorm:"column:aa_count;type:varchar;size:11;" json:"aa_count"`
}

var qs_player_aa_rate_hourlyTableInfo = &TableInfo{
	Name: "qs_player_aa_rate_hourly",
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
			Name:               "hour_time",
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
			GoFieldName:        "HourTime",
			GoFieldType:        "int32",
			JSONFieldName:      "hour_time",
			ProtobufFieldName:  "hour_time",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "aa_count",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(11)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       11,
			GoFieldName:        "AaCount",
			GoFieldType:        "null.String",
			JSONFieldName:      "aa_count",
			ProtobufFieldName:  "aa_count",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (q *QsPlayerAaRateHourly) TableName() string {
	return "qs_player_aa_rate_hourly"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *QsPlayerAaRateHourly) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *QsPlayerAaRateHourly) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *QsPlayerAaRateHourly) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *QsPlayerAaRateHourly) TableInfo() *TableInfo {
	return qs_player_aa_rate_hourlyTableInfo
}
