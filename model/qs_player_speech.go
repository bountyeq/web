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


CREATE TABLE `qs_player_speech` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `from` varchar(64) NOT NULL,
  `to` varchar(64) NOT NULL,
  `message` varchar(256) NOT NULL,
  `minstatus` smallint(5) NOT NULL,
  `guilddbid` int(11) NOT NULL,
  `type` tinyint(3) NOT NULL,
  `timerecorded` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "type": 39,    "timerecorded": "2116-04-03T17:44:33.791445287-07:00",    "id": 12,    "from": "IlkyYduYBSscbaSfSGumANbkV",    "to": "RjENtZwBStGOsdGnposHIGbhK",    "message": "eilLibjnsEbahBNxwURvInlpD",    "minstatus": 57,    "guilddbid": 62}



*/

// QsPlayerSpeech struct is a row record of the qs_player_speech table in the eqemu database
type QsPlayerSpeech struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] from                                           varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	From string `gorm:"column:from;type:varchar;size:64;" json:"from"`
	//[ 2] to                                             varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	To string `gorm:"column:to;type:varchar;size:64;" json:"to"`
	//[ 3] message                                        varchar(256)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	Message string `gorm:"column:message;type:varchar;size:256;" json:"message"`
	//[ 4] minstatus                                      smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: []
	Minstatus int32 `gorm:"column:minstatus;type:smallint;" json:"minstatus"`
	//[ 5] guilddbid                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Guilddbid int32 `gorm:"column:guilddbid;type:int;" json:"guilddbid"`
	//[ 6] type                                           tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	Type int32 `gorm:"column:type;type:tinyint;" json:"type"`
	//[ 7] timerecorded                                   timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [current_timestamp()]
	Timerecorded time.Time `gorm:"column:timerecorded;type:timestamp;" json:"timerecorded"`
}

var qs_player_speechTableInfo = &TableInfo{
	Name: "qs_player_speech",
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
			Name:               "from",
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
			GoFieldName:        "From",
			GoFieldType:        "string",
			JSONFieldName:      "from",
			ProtobufFieldName:  "from",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "to",
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
			GoFieldName:        "To",
			GoFieldType:        "string",
			JSONFieldName:      "to",
			ProtobufFieldName:  "to",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "message",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "Message",
			GoFieldType:        "string",
			JSONFieldName:      "message",
			ProtobufFieldName:  "message",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "minstatus",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "Minstatus",
			GoFieldType:        "int32",
			JSONFieldName:      "minstatus",
			ProtobufFieldName:  "minstatus",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "guilddbid",
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
			GoFieldName:        "Guilddbid",
			GoFieldType:        "int32",
			JSONFieldName:      "guilddbid",
			ProtobufFieldName:  "guilddbid",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "timerecorded",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "Timerecorded",
			GoFieldType:        "time.Time",
			JSONFieldName:      "timerecorded",
			ProtobufFieldName:  "timerecorded",
			ProtobufType:       "uint64",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (q *QsPlayerSpeech) TableName() string {
	return "qs_player_speech"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *QsPlayerSpeech) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *QsPlayerSpeech) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *QsPlayerSpeech) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *QsPlayerSpeech) TableInfo() *TableInfo {
	return qs_player_speechTableInfo
}
