package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// AIndex represents a row from 'a_bit_of_everything.a_index'.
type AIndex struct {
	AKey sql.NullInt64 `json:"a_key"` // a_key

}

// AIndicesByAKey retrieves a row from 'a_bit_of_everything.a_index' as a AIndex.
//
// Generated from index 'a_index_idx'.
func AIndicesByAKey(ctx context.Context, db DB, aKey sql.NullInt64) ([]*AIndex, error) {
	// query
	const sqlstr = `SELECT ` +
		`a_key ` +
		`FROM a_bit_of_everything.a_index ` +
		`WHERE a_key = ?`
	// run
	logf(sqlstr, aKey)
	rows, err := db.QueryContext(ctx, sqlstr, aKey)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*AIndex
	for rows.Next() {
		ai := AIndex{}
		// scan
		if err := rows.Scan(&ai.AKey); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ai)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}