// Command xo generates code from database schemas and custom queries. Works
// with PostgreSQL, MySQL, Microsoft SQL Server, Oracle Database, and SQLite3.
package main

//go:generate ./gen.sh models
//go:generate go generate ./internal

import (
	"context"
	"fmt"
	"os"

	// drivers
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/sijms/go-ora/v2"

	// templates
	_ "gitlab.com/rarify-protocol/templates/createdbtpl"
	_ "gitlab.com/rarify-protocol/templates/dottpl"
	_ "gitlab.com/rarify-protocol/templates/gotpl"
	_ "gitlab.com/rarify-protocol/templates/jsontpl"
	_ "gitlab.com/rarify-protocol/templates/yamltpl"

	"gitlab.com/rarify-protocol/cmd"
	"gitlab.com/rarify-protocol/internal"
	"gitlab.com/rarify-protocol/templates"
)

// version is the app version.
var version = "dhaidashenko_fork-1.0.1"

func main() {
	ctx := context.WithValue(context.Background(), templates.SymbolsKey, internal.Symbols)
	if err := cmd.Run(ctx, "xo", version); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
