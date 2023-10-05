package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"

	"github.com/rarimo/xo/_examples/pgcatalog/pgtypes"
)

// SQLSizing represents a row from 'information_schema.sql_sizing'.
type SQLSizing struct {
	Tableoid       pgtypes.Oid    `json:"tableoid"`        // tableoid
	Cmax           pgtypes.Cid    `json:"cmax"`            // cmax
	Xmax           pgtypes.Xid    `json:"xmax"`            // xmax
	Cmin           pgtypes.Cid    `json:"cmin"`            // cmin
	Xmin           pgtypes.Xid    `json:"xmin"`            // xmin
	Ctid           pgtypes.Tid    `json:"ctid"`            // ctid
	SizingID       sql.NullInt64  `json:"sizing_id"`       // sizing_id
	SizingName     sql.NullString `json:"sizing_name"`     // sizing_name
	SupportedValue sql.NullInt64  `json:"supported_value"` // supported_value
	Comments       sql.NullString `json:"comments"`        // comments
}
