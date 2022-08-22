package query

import (
	"bytes"
	"database/sql"
	"io"

	"github.com/myceliums/gdb/dialect"
)

// Querier is an interface that accepts both *sql.DB and *sql.Tx
type Querier interface {
	Exec(string, ...interface{}) (sql.Result, error)
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)

	Query(string, ...interface{}) (*sql.Rows, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)

	QueryRow(string, ...interface{}) *sql.Row
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row

	Prepare(string) (*sql.Stmt, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
}

// Field contains the field identifiers
type Field interface {
	SQL() string
	Args() []interface{}
}

// Condition
type Condition interface {
	SQL(io.Writer, dialect.Dialect)
}

// Column describes a table column
type Column struct {
	Table string
	Name  string
}

// SQL is an implementation of Field
func (x Column) SQL() string {
	return x.Table + `.` + x.Name
}

// Args is an implementation of Field
func (x Column) Args() []interface{} {
	return []interface{}{}
}

func newContext(source string, targets []Column) *context {
	x := &context{}
	x.source = source
	x.targets = targets

	return x
}

type context struct {
	dialect dialect.Dialect
	source  string
	targets []Column
	joins   []join
	where   Where
	values  []interface{}
}

func (x *context) join(col, ref Column, direction string) {
	var j join
	j.column = col
	j.ref = ref
	j.direction = direction

	x.joins = append(x.joins, j)
}

// InnerJoin adds a join statement to the query
func (x *context) InnerJoin(col, ref Column) {
	x.join(col, ref, `INNER`)
}

// OuterJoin adds a join statement to the query
func (x *context) OuterJoin(col, ref Column) {
	x.join(col, ref, `OUTER`)
}

// LeftJoin adds a join statement to the query
func (x *context) LeftJoin(col, ref Column) {
	x.join(col, ref, `LEFT`)
}

// RightJoin adds a join statement to the query
func (x *context) RightJoin(col, ref Column) {
	x.join(col, ref, `RIGHT`)
}

// Where adds a where statement to the query
func (x *context) Where(condition Condition) {
	x.where = condition
}

// Write is an implementation of io.Writer
func (x *context) Write(in []byte) (int, error) {
	wr := &bytes.Buffer{}

	// TODO

	return wr.Write(in)
}

type join struct {
	column    Column
	ref       Column
	direction string
}

type where struct {
	operator string
	column   Column
	ref      Field
}

type UpdateQuery interface {
	Set(col Column, val interface{}) UpdateQuery
	Where(qc ...Condition) UpdateQuery
	Join(col, ref Column) UpdateQuery
	InnerJoin(col, ref Column) UpdateQuery
	OuterJoin(col, ref Column) UpdateQuery
	LeftJoin(col, ref Column) UpdateQuery
	RightJoin(col, ref Column) UpdateQuery

	Exec(qr Querier) (sql.Result, error)
	ExecContext(ctx context.Context, qr Querier) (sql.Result, error)
}

type DeleteQuery interface {
	Where(qc ...Condition) UpdateQuery
	Join(col, ref Column) UpdateQuery
	InnerJoin(col, ref Column) UpdateQuery
	OuterJoin(col, ref Column) UpdateQuery
	LeftJoin(col, ref Column) UpdateQuery
	RightJoin(col, ref Column) UpdateQuery

	Exec(qr Querier) (sql.Result, error)
	ExecContext(ctx context.Context, qr Querier) (sql.Result, error)
}

//func NewInsert(dialect dialect.Dialect, table string, columns ...string) InsertQuery {
//	x := newQuery(dialect)
//	x.dia.Insert(x.wr, table, columns)
//
//	return x
//}
//
//func NewUpdate(dialect dialect.Dialect, table string) UpdateQuery {
//	x := newQuery(dialect)
//	x.dia.Update(x.wr, table)
//
//	return x
//}
//
//func NewDelete(dialect dialect.Dialect, table string) DeleteQuery {
//	x := newQuery(dialect)
//	x.dia.Delete(x.wr, table)
//
//	return x
//}
//
//func newQuery(dialect dialect.Dialect) *Query {
//	x := &Query{}
//	x.dia = dialect
//	x.wr = &bytes.Buffer{}
//
//	return x
//}

//// Query contains all the metadata to compile an sql query
//type Query struct {
//	wr   *bytes.Buffer
//	dia  dialect.Dialect
//	args []interface{}
//}
//
//// Read is an implementation of io.Reader
//func (x Query) Read(in []byte) (int, error) {
//	return x.wr.Read(in)
//}
//
//func (x Query) String() string {
//	return x.wr.String()
//}
//
//// Join adds a JOIN statement to the query
//func (x *Query) Join(col, ref Column) *Query {
//	return x.InnerJoin(col, ref)
//}
//
//// InnerJoin adds a INNER JOIN statement to the query
//func (x *Query) InnerJoin(col, ref Column) *Query {
//	x.dia.InnerJoin(x.wr, col.Table, col.Name, ref.Table, ref.Name)
//
//	return x
//}
//
////// Where adds a WHERE statement to the query
////func (x *Query) Where(qc Condition) *Query {
////	x.dia.Where(x.wr, qc)
////
////	return x
////}
//
//// Add adds your curstom piece of query to the query
//func (x *Query) Add(q string, args ...interface{}) *Query {
//	x.wr.WriteString(q)
//	x.args = append(x.args, args...)
//
//	return x
//}
//
//// ExecContext execute queries without returning rows.
//func (x Query) ExecContext(ctx context.Context, qr Querier) (sql.Result, error) {
//	return qr.ExecContext(ctx, fmt.Sprint(x), x.args...)
//}
//
//// Exec execute queries without returning rows.
//func (x Query) Exec(qr Querier) (sql.Result, error) {
//	return x.ExecContext(context.Background(), qr)
//}
//
//// QueryContext executes a query that returns rows, typically a SELECT.
//func (x Query) QueryContext(ctx context.Context, qr Querier) (*sql.Rows, error) {
//	return qr.QueryContext(ctx, fmt.Sprint(x), x.args...)
//}
//
//// Query executes a query that returns rows, typically a SELECT.
//func (x Query) Query(qr Querier) (*sql.Rows, error) {
//	return x.QueryContext(context.Background(), qr)
//}
//
//// QueryRowContext executes a query that returns one row. typically a SELECT
//func (x Query) QueryRowContext(ctx context.Context, qr Querier) *sql.Row {
//	return qr.QueryRowContext(ctx, x.wr.String(), x.args...)
//}
//
//// QueryRowContext executes a query that returns one row. typically a SELECT
//func (x Query) QueryRow(qr Querier) *sql.Row {
//	return x.QueryRowContext(context.Background(), qr)
//}
