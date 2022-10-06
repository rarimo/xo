package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"gitlab.com/rarify-protocol/xo/_examples/pgcatalog/pgtypes"
)

// PgNumericPrecisionRadix calls the stored function 'information_schema._pg_numeric_precision_radix(oid, integer) integer' on db.
func PgNumericPrecisionRadix(ctx context.Context, db DB, typid pgtypes.Oid, typmod int) (int, error) {
	// call information_schema._pg_numeric_precision_radix
	const sqlstr = `SELECT * FROM information_schema._pg_numeric_precision_radix($1, $2)`
	// run
	var r0 int
	logf(sqlstr, typid, typmod)
	if err := db.QueryRowContext(ctx, sqlstr, typid, typmod).Scan(&r0); err != nil {
		return 0, logerror(err)
	}
	return r0, nil
}
