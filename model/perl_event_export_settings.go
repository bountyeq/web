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


CREATE TABLE `perl_event_export_settings` (
  `event_id` int(11) NOT NULL,
  `event_description` varchar(150) DEFAULT NULL,
  `export_qglobals` smallint(11) DEFAULT 0,
  `export_mob` smallint(11) DEFAULT 0,
  `export_zone` smallint(11) DEFAULT 0,
  `export_item` smallint(11) DEFAULT 0,
  `export_event` smallint(11) DEFAULT 0,
  PRIMARY KEY (`event_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "event_description": "guUPFZmfnVTEbxhFUymyIPJfR",    "export_qglobals": 35,    "export_mob": 56,    "export_zone": 54,    "export_item": 83,    "export_event": 24,    "event_id": 66}



*/

// PerlEventExportSettings struct is a row record of the perl_event_export_settings table in the eqemu database
type PerlEventExportSettings struct {
	//[ 0] event_id                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	EventID int32 `gorm:"primary_key;column:event_id;type:int;" json:"event_id"`
	//[ 1] event_description                              varchar(150)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 150     default: [NULL]
	EventDescription null.String `gorm:"column:event_description;type:varchar;size:150;" json:"event_description"`
	//[ 2] export_qglobals                                smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	ExportQglobals null.Int `gorm:"column:export_qglobals;type:smallint;default:0;" json:"export_qglobals"`
	//[ 3] export_mob                                     smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	ExportMob null.Int `gorm:"column:export_mob;type:smallint;default:0;" json:"export_mob"`
	//[ 4] export_zone                                    smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	ExportZone null.Int `gorm:"column:export_zone;type:smallint;default:0;" json:"export_zone"`
	//[ 5] export_item                                    smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	ExportItem null.Int `gorm:"column:export_item;type:smallint;default:0;" json:"export_item"`
	//[ 6] export_event                                   smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	ExportEvent null.Int `gorm:"column:export_event;type:smallint;default:0;" json:"export_event"`
}

var perl_event_export_settingsTableInfo = &TableInfo{
	Name: "perl_event_export_settings",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "event_id",
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
			GoFieldName:        "EventID",
			GoFieldType:        "int32",
			JSONFieldName:      "event_id",
			ProtobufFieldName:  "event_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "event_description",
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
			GoFieldName:        "EventDescription",
			GoFieldType:        "null.String",
			JSONFieldName:      "event_description",
			ProtobufFieldName:  "event_description",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "export_qglobals",
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
			GoFieldName:        "ExportQglobals",
			GoFieldType:        "null.Int",
			JSONFieldName:      "export_qglobals",
			ProtobufFieldName:  "export_qglobals",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "export_mob",
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
			GoFieldName:        "ExportMob",
			GoFieldType:        "null.Int",
			JSONFieldName:      "export_mob",
			ProtobufFieldName:  "export_mob",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "export_zone",
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
			GoFieldName:        "ExportZone",
			GoFieldType:        "null.Int",
			JSONFieldName:      "export_zone",
			ProtobufFieldName:  "export_zone",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "export_item",
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
			GoFieldName:        "ExportItem",
			GoFieldType:        "null.Int",
			JSONFieldName:      "export_item",
			ProtobufFieldName:  "export_item",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "export_event",
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
			GoFieldName:        "ExportEvent",
			GoFieldType:        "null.Int",
			JSONFieldName:      "export_event",
			ProtobufFieldName:  "export_event",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PerlEventExportSettings) TableName() string {
	return "perl_event_export_settings"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PerlEventExportSettings) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PerlEventExportSettings) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PerlEventExportSettings) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PerlEventExportSettings) TableInfo() *TableInfo {
	return perl_event_export_settingsTableInfo
}
