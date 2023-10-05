package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"github.com/rarimo/xo/_examples/pgcatalog/pgtypes"
)

// PgIntervalType calls the stored function 'information_schema._pg_interval_type(oid, integer) text' on db.
func PgIntervalType(ctx context.Context, db DB, typid pgtypes.Oid, mod int) (string, error) {
	// call information_schema._pg_interval_type
	const sqlstr = `SELECT * FROM information_schema._pg_interval_type($1, $2)`
	// run
	var r0 string
	logf(sqlstr, typid, mod)
	if err := db.QueryRowContext(ctx, sqlstr, typid, mod).Scan(&r0); err != nil {
		return "", logerror(err)
	}
	return r0, nil
}
