# xo

`xo` is a command-line tool to generate *idiomatic* code for different
languages code based on a database schema or a custom query.

#### Supported languages

At the moment, `xo` only supports [Go](https://golang.org). Support for other
languages will come soon.

#### How it works

In schema mode, `xo` connects to your database and generates code using Go
templates. `xo` works by using database metadata and SQL introspection queries to
discover the types  and relationships contained within a schema, and applying a
standard set of base (or customized) Go [templates](templates) against the
discovered relationships.

Currently, `xo` can generate types for tables, enums, stored procedures, and
custom SQL queries for PostgreSQL, MySQL, Oracle, Microsoft SQL Server, and
SQLite3 databases.

> **Note:** While the code generated by xo is production quality, it is not the
goal, nor the intention for xo to be a "silver bullet," nor to completely
eliminate the manual authoring of SQL / Go code.

In query mode, `xo` parses your query to generate code from Go templates.
It finds related tables in your database to ensure type safety.

## Database Feature Support

The following is a matrix of the feature support for each database:

|              | PostgreSQL       | MySQL            | Oracle           | Microsoft SQL Server| SQLite           |
| ------------ |:----------------:|:----------------:|:----------------:|:-------------------:|:----------------:|
| Models       |:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:   |:white_check_mark:|
| Primary Keys |:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:   |:white_check_mark:|
| Foreign Keys |:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:   |:white_check_mark:|
| Indexes      |:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:   |:white_check_mark:|
| Stored Procs |:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:   |:white_check_mark:|
| Functions    |:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:   |:white_check_mark:|
| ENUM types   |:white_check_mark:|:white_check_mark:|                  |                     |                  |
| Custom types |:white_check_mark:|                  |                  |                     |                  |

## Installation

For Go code generation, install the `goimports` dependency (if not already
installed):

```sh
$ go get -u golang.org/x/tools/cmd/goimports
```

Then, install `xo` in the usual Go way:

```sh
$ go get -u gitlab.com/rarify-protocol/xo
```

> **Note:** Go 1.16+ is needed for installing `xo` from source, as it makes use
of `go embed` to embed Go templates into the binaries, which is not compatible
with previous versions of Go. When compiling to Go, generated code can compile
with Go 1.3+ code, disabling context mode if necessary.

## Quickstart

The following is a quick overview of using `xo` on the command-line:

```sh
# Make an output directory for generated code.
$ mkdir -p models

# Generate code from your Postgres schema. (Default output folder is models)
$ xo schema pgsql://user:pass@host/dbname

# Generate code from a Microsoft SQL schema using a custom template directory (see notes below)
$ mkdir -p mssqlmodels
$ xo schema mssql://user:pass@host/dbname -o mssqlmodels --src custom/templates

# Generate code from a custom SQL query for Postgres
$ xo query pg://user:pass@host/dbname -M -B -T -2 AuthorResulto << ENDSQL
SELECT
  a.name::varchar AS name,
  b.type::integer AS my_type
FROM authors a
  INNER JOIN authortypes b ON a.id = b.author_id
WHERE
  a.id = %%authorID int%%
LIMIT %%limit int%%
ENDSQL

# Build generated code - verify it compiles
$ go build ./models/
$ go build ./mssqlmodels/
```

## Command Line Options

The following are `xo`'s command-line arguments and options:

```sh
$ xo --help-long
usage: xo [<flags>] <command> [<args> ...]

Flags:
      --help     Show context-sensitive help (also try --help-long and
                 --help-man).
  -v, --verbose  enable verbose output
      --version  display version and exit

Commands:
  help [<command>...]
    Show help.


  query [<flags>] <DSN>
    Generate code for a database custom query from a template.

    -s, --schema=<name>          database schema name
    -t, --template=go            template type (createdb, dot, go, json, yaml;
                                 default: go)
    -f, --suffix=<ext>           file extension suffix for generated files
                                 (otherwise set by template type)
    -o, --out=models             out path (default: models)
    -a, --append                 enable append mode
    -S, --single=<file>          enable single file output
    -D, --debug                  debug generated code (writes generated code to
                                 disk without post processing)
    -Q, --query=""               custom database query (uses stdin if not
                                 provided)
    -T, --type=<name>            type name
        --type-comment=""        type comment
    -F, --func=<name>            func name
        --func-comment=""        func comment
    -M, --trim                   enable trimming whitespace
    -B, --strip                  enable stripping type casts
    -1, --one                    enable returning single (only one) result
    -l, --flat                   enable returning unstructured values
    -X, --exec                   enable exec (no introspection performed)
    -I, --interpolate            enable interpolation of embedded params
    -L, --delimiter=%%           delimiter used for embedded params (default:
                                 %%)
    -Z, --fields=<field>         override field names for results
    -U, --allow-nulls            allow result fields with NULL values
    -d, --src=<path>             template source directory
    -2, --go-not-first           disable package comment (ie, not first
                                 generated file)
        --go-int32=int           int32 type (default: int)
        --go-uint32=uint         uint32 type (default: uint)
        --go-pkg=<name>          package name
        --go-tag="" ...          build tags
        --go-import="" ...       package imports
        --go-uuid=<pkg>          uuid type package
        --go-custom=<name>       package name for custom types
        --go-conflict=Val        name conflict suffix (default: Val)
        --go-esc=none ...        escape fields (none, schema, table, column,
                                 all; default: none)
    -g, --go-field-tag=<tag>     field tag
        --go-context=only        context mode (disable, both, only; default:
                                 only)
        --go-inject=""           insert code into generated file headers
        --go-inject-file=<file>  insert code into generated file headers from a
                                 file
        --json-indent="  "       indent spacing
        --json-ugly              disable indentation

  schema [<flags>] <DSN>
    Generate code for a database schema from a template.

    -s, --schema=<name>          database schema name
    -t, --template=go            template type (createdb, dot, go, json, yaml;
                                 default: go)
    -f, --suffix=<ext>           file extension suffix for generated files
                                 (otherwise set by template type)
    -o, --out=models             out path (default: models)
    -a, --append                 enable append mode
    -S, --single=<file>          enable single file output
    -D, --debug                  debug generated code (writes generated code to
                                 disk without post processing)
    -k, --fk-mode=smart          foreign key resolution mode (smart, parent,
                                 field, key; default: smart)
    -i, --include=<glob> ...     include types (<type>)
    -e, --exclude=<glob> ...     exclude types/fields (<type>[.<field>])
    -j, --use-index-names        use index names as defined in schema for
                                 generated code
    -d, --src=<path>             template source directory
        --createdb-fmt=<path>    fmt command (default: )
        --createdb-fmt-opts=<opts> ...
                                 fmt options (default: )
        --createdb-constraint    enable constraint name in output (postgres,
                                 mysql, sqlite3)
        --createdb-escape=none   escape mode (none, types, all; default: none)
        --createdb-engine=""     mysql table engine (default: InnoDB)
        --dot-defaults="" ...    default statements (default: node [shape=none,
                                 margin=0])
        --dot-bold               bold header row
        --dot-color=""           header color (default: lightblue)
        --dot-row=""             row value template (default: {{ .Name }}: {{
                                 .Datatype.Type }})
        --dot-direction          enable edge directions
    -2, --go-not-first           disable package comment (ie, not first
                                 generated file)
        --go-int32=int           int32 type (default: int)
        --go-uint32=uint         uint32 type (default: uint)
        --go-pkg=<name>          package name
        --go-tag="" ...          build tags
        --go-import="" ...       package imports
        --go-uuid=<pkg>          uuid type package
        --go-custom=<name>       package name for custom types
        --go-conflict=Val        name conflict suffix (default: Val)
        --go-esc=none ...        escape fields (none, schema, table, column,
                                 all; default: none)
    -g, --go-field-tag=<tag>     field tag
        --go-context=only        context mode (disable, both, only; default:
                                 only)
        --go-inject=""           insert code into generated file headers
        --go-inject-file=<file>  insert code into generated file headers from a
                                 file
        --json-indent="  "       indent spacing
        --json-ugly              disable indentation
        --postgres-oids          enable postgres OIDs

  dump [<flags>] <out>
    Dump internal templates to path.

    -t, --template=go   template type (createdb, dot, go, json, yaml; default:
                        go)
    -f, --suffix=<ext>  file extension suffix for generated files (otherwise set
                        by template type)
```

## About Base Templates

`xo` provides a set of generic "base" [templates](templates) for each of the
supported databases, but it is understood these templates are not suitable for
every organization or every schema out there. As such, you can author your own
custom templates, or modify the base templates available in the `xo` source
tree, and use those with `xo` by a passing a directory path via the `--src`
flag.

For non-trivial schemas, custom templates are the most practical, common, and
best way to use `xo` (see below quickstart and related example).

### Custom Template Quickstart

The following is a quick overview of copying the base templates contained in
the `xo` project's [`templates/`](templates) directory, editing to suit, and
using with `xo`:

```sh
# Create a template directory
$ mkdir -p templates

# Copy xo templates
$ xo dump templates

# edit base postgres templates
$ vi templates/*.go.tpl

# use with xo
$ xo pgsql://user:pass@host/db -o models --src templates
```

See the Custom Template example below for more information on adapting the base
templates in the `xo` source tree for use within your own project.

### Storing Project Templates

Ideally, the custom templates for your project/schema should be stored
within your project, and used in conjunction with a build pipeline such as
`go generate`:

```sh
# Add to custom xo command to go generate:
$ tee -a gen.go << ENDGO
package mypackage

//go:generate xo pgsql://user:pass@host/db -o models --src templates
ENDGO

# Run go generate
$ go generate

# Add custom templates and gen.go to project
$ git add templates gen.go && git commit -m 'Adding custom xo templates for models'
```

> **Note**: via the `--template` parameter of `xo dump` you can create
templates for other languages. The default is `go`.

### Template Language/Syntax

`xo` templates are standard Go text templates. Please see the [documentation
for Go's standard `text/template` package](https://pkg.go.dev/text/template)
for information concerning the syntax, logic, and variable use within Go
templates.

### Template Context and File Layout

The contexts (ie, the `.` identifier in templates) made available to custom
templates can be found in [templates/types.go](templates/types.go)
(see below table for more information on which file uses which type).

Each language, has its own set of templates for `$TYPE` and are
available in the [templates/](templates).

|        Template File       | [Type](templates/types.go) | Description                                                                            |
|:--------------------------:|:--------------------------:|----------------------------------------------------------------------------------------|
|        hdr.xo.*.tpl        |                            | Base template. Executed with content for a template.                                   |
|         db.xo.*.tpl        |                            | Package level template with base types and interface data. Generated once per package. |
|    schema/enum.xo.*.tpl    |            Enum            | Template for schema enum type definitions. Generates types and related methods.        |
| schema/foreignkey.xo.*.tpl |         ForeignKey         | Template for foreign key relationships. Generates  related method.                     |
|    schema/index.xo.*.tpl   |            Index           | Template for schema indexes. Generates related method.                                 |
|    schema/proc.xo.*.tpl    |            Proc            | Template to generate functions to call defined stored procedures in the db.            |
|   schema/typedef.xo.*.tpl  |            Type            | Template for schema table/views.                                                       |
|    query/custom.xo.*.tpl   |            Query           | Template for custom query execution.                                                   |
|   query/typedef.xo.*.tpl   |            Type            | Template for custom query's generated type.                                            |

For example, Go has
[`templates/gotpl/schema/foreignkey.xo.go.tpl`](templates/gotpl/schema/foreignkey.xo.go.tpl)
which defines the template used by `xo` for generating a function to get the
foreign key type in Go. The templates are designed to be Database agnostic, so
they are used for both PostgreSQL and Microsoft SQL the same, and all other
supported database types. The template is passed a different instance of
`templates.ForeignKey` instance (for each foreign key in a table). To get the
`Name` field in from `ForeignKey`, the template can use ` {{ .Data.Name }}`, or
any other field similarly.

#### Template Helpers

There is a set of well-defined template helpers in `funcs.go` for each supported
language that assist with writing templated Go code / SQL. Please review how the
base [`templates`](templates) make use of helpers, and the inline Go
documentation for the respective helper func definitions.

## Examples

### Example: End-to-End

Please see the [booktest example](_examples/booktest) for a full end-to-end
example for each supported database, showcasing how to use a database schema
with `xo`, and the resulting code generated by `xo`.

Additionally, please see the [northwind example](_examples/northwind) for a
demonstration of running `xo` against a large schema. Please note that this
example is a work in progress, and does not yet work properly with Microsoft
SQL Server and Oracle databases, and has no documentation (for now) -- however
it works very similarly to the booktest end-to-end example.

### Example: Ignoring Fields

Sometimes you may wish to have the database manage the values of columns
instead of having them managed by code generated by `xo`. As such, when you
need `xo` to ignore fields for a database schema, you can use the `-e` or
`--exclude` flag. For example, a common use case is to define a table with
`created_at` and/or `modified_at` timestamps fields, where the database is
responsible for setting column values on `INSERT` and `UPDATE`, respectively.

Consider the following PostgreSQL schema where a `users` table has a
`created_at` and `modified_at` field, where `created_at` has a default value of
`now()` and where `modified_at` is updated by a trigger on `UPDATE`:

```postgresql
CREATE TABLE users (
  id          SERIAL PRIMARY KEY,
  name        text NOT NULL DEFAULT '' UNIQUE,
  created_at  timestamptz   default now(),
  modified_at timestamptz   default now()
);

CREATEOR REPLACE FUNCTION update_modified_column() RETURNS TRIGGER AS $$
BEGIN
    NEW.modified_at= now();
RETURN NEW;
END;
$$language 'plpgsql';

CREATE TRIGGER update_users_modtime BEFORE UPDATE ON users
  FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
```

We can ensure that these columns are managed by PostgreSQL and not by the
application logic but by `xo` by passing the `--exclude` or `-e` flag:

```sh
# Ignore special fields
$ xo schema pgsql://user:pass@host/db -e users.created_at -e users.modified_at
# or, To ignore these fields in all tables
$ xo schema pgsql://user:pass@host/db -e *.created_at -e *.modified_at
```

### Example: Custom Template -- adding a `GetMostRecent` lookup for all tables (Go)

Often, a schema has a common layout/pattern, such as every table having a
`created_at` and `modified_at` field (as in the PostgreSQL schema in the
previous example). It is then a common use-case to have a `GetMostRecent`
lookup for each table type, retrieving the most recently modified rows for each
table (up to some limit, N).

To accomplish this with `xo`, we will need to create our own set of custom
templates, and then add a `GetMostRecent` lookup to the `.type.go.tpl`
template.

First, we create dump the base `xo` templates:

```sh
$ mkdir -p templates

$ xo dump templates
```

We can now modify the templates to suit our specific schema, adding lookups,
helpers, or anything else necessary for our schema.

To add a `GetMostRecent` lookup, we edit our copy of the `typedef.xo.go.tpl`
template:

```sh
$ vi templates/gotpl/schema/typedef.xo.go.tpl
```

And add the following templated `GetMostRecent` func at the end of the file:

```go
// GetMostRecent{{ $type.Name }} returns n most recent rows from '{{ $table }}',
// ordered by "created_at" in descending order.
func GetMostRecent{{ $type.Name }}(ctx context.Context, db DB, n int) ([]*{{ $type.Name }}, error) {
    const sqlstr = `SELECT ` +
        `{{ $type.Fields "created_at" "modified_at" }}` +
        `FROM {{ $table }} ` +
        `ORDER BY created_at DESC LIMIT $1`

    rows, err := db.QueryContext(ctx, sqlstr, n)
    if err != nil {
        return nil, logerror(err)
    }
    defer rows.Close()

    // load results
    var res []*{{ $type.Name }}
    for rows.Next() {
        {{ $short }} := {{ $type.Name }}{
        {{- if $type.PrimaryKey }}
            _exists: true,
        {{ end -}}
        }
        // scan
        if err := rows.Scan({{ fieldnames $type.Fields (print "&" $short) }}); err != nil {
            return nil, logerror(err)
        }
        res = append(res, &{{ $short }})
    }
    return res, nil
}
```

We can then use the templates in conjunction with `xo` to generate our "model"
code:

```sh
$ xo schema pgsql://user:pass@localhost/dbname --src templates/
```

There will now be a `GetMostRecentUsers` func defined in `models/user.xo.go`,
which can be used as follows:

```go
db, err := dburl.Open("pgsql://user:pass@localhost/dbname")
if err != nil { /* ... */ }

// retrieve 15 most recent items
mostRecentUsers, err := models.GetMostRecentUsers(context.Background(), db, 15)
if err != nil { /* ... */ }
for _, user := range users {
    log.Printf("got user: %+v", user)
}
```

## Using SQL Drivers

Please note that the base `xo` templates do not import any SQL drivers. It is
left for the user of `xo`'s generated code to import the actual drivers. For
reference, these are the expected drivers to use with the code generated by
`xo`:

| Database (driver)            | Package                                                                      |
|------------------------------|------------------------------------------------------------------------------|
| PostgreSQL (postgres)        | [github.com/lib/pq](https://github.com/lib/pq)                               |
| SQLite3 (sqlite3)            | [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)           |
| MySQL (mysql)                | [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)     |
| Microsoft SQL Server (mssql) | [github.com/denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb) |
| Oracle (godror)              | [github.com/godror/godror](https://github/godror/godror)                     |

Additionally, please see below for usage notes on specific SQL database
drivers.

### MySQL (mysql)

If your schema or custom query contains table or column names that need to be
escaped using any of the `--escape-*` options, you must pass the `sql_mode=ansi`
option to the MySQL driver:

```sh
$ xo --escape-all 'mysql://user:pass@host/?parseTime=true&sql_mode=ansi' -o models
```

And when opening a database connection:

```go
db, err := dburl.Open("mysql://user:pass@host/?parseTime=true&sql_mode=ansi")
```

Additionally, when working with date/time column types in MySQL, one should
pass the `parseTime=true` option to the MySQL driver:

```sh
$ xo 'mysql://user:pass@host/dbname?parseTime=true' -o models
```

And when opening a database connection:

```go
db, err := dburl.Open("mysql://user:pass@host/dbname?parseTime=true")
```

### SQLite3 (sqlite3)

While not required, one should specify the `loc=auto` option when using `xo`
with a SQLite3 database:

```sh
$ xo 'file:mydatabase.sqlite3?loc=auto' -o models
```

And when opening a database connection:

```go
db, err := dburl.Open("file:mydatabase.sqlite3?loc=auto")
```
#### Installing Oracle instantclient on Debian/Ubuntu

On Ubuntu/Debian, you may download the instantclient RPMs
[here](http://www.oracle.com/technetwork/topics/linuxx86-64soft-092277.html).

You should then be able to do the following:

```sh
# install alien, if not already installed
$ sudo aptitude install alien

# install the instantclient RPMs
$ sudo alien -i oracle-instantclient-12.1-basic-*.rpm
$ sudo alien -i oracle-instantclient-12.1-devel-*.rpm
$ sudo alien -i oracle-instantclient-12.1-sqlplus-*.rpm

# get xo
$ go get -u gitlab.com/rarify-protocol/xo

# copy oci8.pc from xo/contrib to system pkg-config directory
$ sudo cp $GOPATH/src/gitlab.com/rarify-protocol/xo/contrib/oci8.pc /usr/lib/pkgconfig/

# install rana's ora driver
$ go get -u gopkg.in/rana/ora.v4

# assuming the above succeeded, install xo with oracle support enabled
$ go install -tags oracle gitlab.com/rarify-protocol/xo
```

#### Contrib Scripts and Oracle Docker Image

It's of note that there are additional scripts available in the
[usql contrib](https://github.com/xo/usql/tree/master/contrib) directory that
can help when working with Oracle databases and `xo`.

For reference, the `xo` developers use the
[sath89/oracle-12c](https://hub.docker.com/r/sath89/oracle-12c/) Docker image
for testing `xo`'s Oracle database support.


## About Primary Keys
For row inserts `xo` determines whether the primary key is
automatically generated by the DB or must be provided by the application for the
table row being inserted. For example a table that has a primary key that is
also a foreign key to another table, or a table that has multiple primary keys
in a many-to-many link table, it is desired that the application provide the
primary key(s) for the insert rather than the DB.

`xo` will query the schema to determine if the database provides an automatic
primary key and if the table does not provide one then it will require that the
application provide the primary key for the object passed to the Insert method.
Below is information on how the logic works for each database type to determine
if the DB automatically provides the PK.

### PostgreSQL Auto PK Logic
* Checks for a sequence that is owned by the table in question.

### MySQL Auto PK Logic
* Checks for an autoincrement row in the information_schema for the table in
  question.

### SQLite Auto PK Logic
* Checks the SQL that is used to generate the table contains
the *AUTOINCREMENT* keyword.
* Checks that the table was created with the primary key type of *INTEGER*.

If either of the above conditions are satisfied then the PK is determined to be
automatically provided by the DB. For the case of integer PK's when you want to
override that the PK be manually provided then you can define the key type as
*INT* instead of *INTEGER*, for example as in the following many-to-many link
table:

```sql
  CREATE TABLE site_contacts (
  contact_id	INT NOT NULL,
  site_id	INT NOT NULL,
  PRIMARY KEY(contact_id,siteid),
  FOREIGN KEY(contact_id) REFERENCES contacts (contact_id),
  FOREIGN KEY(site_id) REFERENCES sites (site_id)
)
```

### SQL Server Auto PK Logic
* Checks for an identity associated with one of the columns for the table in
  question.

### Oracle Auto PK Logic
`ALWAYS GENERATED` types will be parsed as Auto PK types for Oracle.

## About xo: Design, Origin, Philosophy, and History

`xo` can likely get you 99% "of the way there" on medium or large database
schemas and 100% of the way there for small or trivial database schemas. In
short, xo is a great launching point for developing standardized packages for
standard database abstractions/relationships, and `xo`'s most common use-case is
indeed in a code generation pipeline, ala `stringer`.

### Design

`xo` is **NOT** designed to be an ORM or to generate an ORM. Instead, `xo` is
designed to vastly reduce the overhead/redundancy of (re-)writing types and
funcs for common database queries/relationships -- it is not meant to be
a "silver bullet".

### History

`xo` was originally developed while migrating a large application written in
PHP to Go. The schema in use in the original app, while well-designed, had
become inconsistent over multiple iterations/generations, mainly due to
different naming styles adopted by various developers/database admins over the
preceding years. Additionally, some components had been written in different
languages (Ruby, Java) and had also accumulated significant drift from the
original application and accompanying schema. Simultaneously, a large amount of
growth meant that the PHP/Ruby code could no longer efficiently serve the
traffic volumes.

In late 2014/early 2015, a decision was made to unify and strip out certain
backend services and to fully isolate the API from the original application,
allowing the various components to instead speak to a common API layer instead
of directly to the database, and to build that service layer in Go.

However, unraveling the old PHP/Ruby/Java code became a large headache, as the
code, the database, and the API, all had significant drift -- thus, underlying
function names, fields, and API methods no longer coincided with the actual
database schema, and were named differently in each language. As such, after a
round of standardizing names, dropping cruft, and adding a few relationship
changes to the schema, the various codebases were fixed to match the schema
changes. After that was determined to be a success, the next target was to
rewrite the backend services in Go.

In order to keep a similar and consistent workflow for the developers, the
previous code generator (written in PHP and Twig templates) was modified to
generate Go code. Additionally, at this time, but tangential to the story, the
API definitions were ported from JSON to Protobuf to make use of its code
generation abilities as well.

`xo` is the open source version of that code generation tool, and is the fruits
of those development efforts. It is hoped that others will be able to use and
expand `xo` to support other databases -- SQL or otherwise -- and that `xo` can
become a common tool in any Go developer's toolbox.

### Goals

Part of `xo`'s goals is to avoid writing an ORM, or an ORM-like in Go, and to
instead generate static, type-safe, fast, and idiomatic Go code across
languages and databases. Additionally, the `xo` developers are of the opinion
that relational databases should have proper, well-designed relationships and
all the related definitions should reside within the database schema itself:
ie, a "self-documenting" schema. `xo` is an end to that pursuit.

## Related Projects

* [dburl](https://github.com/xo/dburl) - a Go package providing a standard, URL
  style mechanism for parsing and opening database connection URLs
* [usql](https://github.com/xo/usql) - a universal command-line interface for
  SQL databases

### Other Projects

The following projects work with similar concepts as xo:

#### Go Generators
* [ModelQ](https://github.com/mijia/modelq)
* [sqlgen](https://github.com/drone/sqlgen)
* [squirrel](https://github.com/Masterminds/squirrel)
* [scaneo](https://github.com/variadico/scaneo)
* [acorn](https://github.com/willowtreeapps/acorn) and
  [rootx](https://github.com/willowtreeapps/rootx) ([read overview
  here](http://willowtreeapps.com/blog/go-generate-your-database-code/))

#### Go ORM-likes
* [sqlc](https://github.com/relops/sqlc)

## TODO
* Add (finish) stored proc support for Oracle + Microsoft SQL Server
* Unit tests / code coverage / continuous builds for binary package releases
* Move database introspection to separate package for reuse by other Go packages
* Overhaul/standardize type parsing
* Finish support for --{incl, excl}[ude] types
* Write/publish template set for protobuf
* Add support for generating models for other languages
* Finish many-to-many and link table support
* Finish example and code for generated *Slice types (also, only generate for the databases its needed for)
* Add example for many-to-many relationships and link tables
* Add support for supplying a file (ie, *.sql) for query generation
* Add support for full text types (tsvector, tsquery on PostgreSQL)
* Finish COMMENT support for PostgreSQL/MySQL and update templates accordingly.
* Add support for JSON types (json, jsonb on PostgreSQL, json on MySQL)
* Add support for GIN index queries (PostgreSQL)
