package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"

	"gitlab.com/rarify-protocol/xo/_examples/pgcatalog/pgtypes"
)

// SQLPart represents a row from 'information_schema.sql_parts'.
type SQLPart struct {
	Tableoid     pgtypes.Oid    `json:"tableoid"`       // tableoid
	Cmax         pgtypes.Cid    `json:"cmax"`           // cmax
	Xmax         pgtypes.Xid    `json:"xmax"`           // xmax
	Cmin         pgtypes.Cid    `json:"cmin"`           // cmin
	Xmin         pgtypes.Xid    `json:"xmin"`           // xmin
	Ctid         pgtypes.Tid    `json:"ctid"`           // ctid
	FeatureID    sql.NullString `json:"feature_id"`     // feature_id
	FeatureName  sql.NullString `json:"feature_name"`   // feature_name
	IsSupported  sql.NullString `json:"is_supported"`   // is_supported
	IsVerifiedBy sql.NullString `json:"is_verified_by"` // is_verified_by
	Comments     sql.NullString `json:"comments"`       // comments
}
