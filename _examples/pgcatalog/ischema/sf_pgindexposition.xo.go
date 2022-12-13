package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"gitlab.com/rarimo/xo/_examples/pgcatalog/pgtypes"
)

// PgIndexPosition calls the stored function 'information_schema._pg_index_position(oid, smallint) integer' on db.
func PgIndexPosition(ctx context.Context, db DB, p0 pgtypes.Oid, p1 int16) (int, error) {
	// call information_schema._pg_index_position
	const sqlstr = `SELECT * FROM information_schema._pg_index_position($1, $2)`
	// run
	var r0 int
	logf(sqlstr, p0, p1)
	if err := db.QueryRowContext(ctx, sqlstr, p0, p1).Scan(&r0); err != nil {
		return 0, logerror(err)
	}
	return r0, nil
}
