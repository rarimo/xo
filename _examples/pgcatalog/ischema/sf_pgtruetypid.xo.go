package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"gitlab.com/rarimo/xo/_examples/pgcatalog/pgtypes"
)

// PgTruetypid calls the stored function 'information_schema._pg_truetypid(pg_attribute, pg_type) oid' on db.
func PgTruetypid(ctx context.Context, db DB, p0 pgtypes.PgAttribute, p1 pgtypes.PgType) (pgtypes.Oid, error) {
	// call information_schema._pg_truetypid
	const sqlstr = `SELECT * FROM information_schema._pg_truetypid($1, $2)`
	// run
	var r0 pgtypes.Oid
	logf(sqlstr, p0, p1)
	if err := db.QueryRowContext(ctx, sqlstr, p0, p1).Scan(&r0); err != nil {
		return pgtypes.Oid{}, logerror(err)
	}
	return r0, nil
}
