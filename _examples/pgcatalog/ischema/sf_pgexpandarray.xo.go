package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"gitlab.com/rarify-protocol/xo/_examples/pgcatalog/pgtypes"
)

// PgExpandarray calls the stored function 'information_schema._pg_expandarray(anyarray) (anyelement, integer)' on db.
func PgExpandarray(ctx context.Context, db DB, p0 pgtypes.Anyarray) (pgtypes.Anyelement, int, error) {
	// call information_schema._pg_expandarray
	const sqlstr = `SELECT * FROM information_schema._pg_expandarray($1)`
	// run
	var x pgtypes.Anyelement
	var n int
	logf(sqlstr, p0)
	if err := db.QueryRowContext(ctx, sqlstr, p0).Scan(&x, &n); err != nil {
		return pgtypes.Anyelement{}, 0, logerror(err)
	}
	return x, n, nil
}
