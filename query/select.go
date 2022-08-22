package query

import (
	"bytes"
	"context"
	"database/sql"
)

// NewSelect returns a new SelectQuery
func NewSelect(source string, columns ...Column) *SelectQuery {
	x := &SelectQuery{}
	x.context = newContext(source, columns)
	return x
}

// SelectQuery builds a select query with the given arguments
type SelectQuery struct {
	context *context
	limit   int
	offset  int
}

// SQL is an implementation of Field
func (x SelectQuery) SQL() string {
	wr := &bytes.Buffer{}

	// TODO
	return wr.String()
}

// Args is an implementation of Field
func (x SelectQuery) Args() []interface{} {
	return x.builder.values
}

// Limit adds a LIMIT statement to the query
func (x *SelectQuery) Limit(i int) *SelectQuery {
	x.limit = i
	return x
}

// Offset adds an OFFSET statement to the query
func (x *SelectQuery) Offset(i int) *SelectQuery {
	x.offset = i
	return x
}

// Where adds WHERE conditions to the query
func (x *SelectQuery) Where(condition Condition) *SelectQuery {
	x.context.where = condition
	return x
}

// InnerJoin adds a join statement to the query
func (x *SelectQuery) InnerJoin(col, ref Column) *SelectQuery {
	x.context.InnerJoin(col, ref)
	return x
}

// OuterJoin adds a join statement to the query
func (x *SelectQuery) OuterJoin(col, ref Column) *SelectQuery {
	x.context.OuterJoin(col, ref)
	return x
}

// LeftJoin adds a join statement to the query
func (x *SelectQuery) LeftJoin(col, ref Column) *SelectQuery {
	x.context.LeftJoin(col, ref)
	return x
}

// RightJoin adds a join statement to the query
func (x *SelectQuery) RightJoin(col, ref Column) *SelectQuery {
	x.context.RightJoin(col, ref)
	return x
}

// Query runs the query and returns the rows
func (x SelectQuery) Query(qr Querier) (*sql.Rows, error) {
	return x.QueryContext(context.Background(), qr)
}

// QueryContext runs the query and returns the rows
func (x SelectQuery) QueryContext(ctx context.Context, qr Querier) (*sql.Rows, error) {
	return qr.QueryContext(ctx, x.SQL(), x.args...)
}

// QueryRow runs the query and returns the first row
func (x SelectQuery) QueryRow(qr Querier) *sql.Row {
	return x.QueryRowContext(context.Background(), qr)
}

// QueryRowContext runs the query and returns the first row
func (x SelectQuery) QueryRowContext(ctx context.Context, qr Querier) *sql.Row {
	return qr.QueryRowContext(ctx, x.SQL(), x.args...)
}

// Cursor creates a DECLARE CURSOR statment and excetutes the statement and returns the Cursor
func (x SelectQuery) Cursor(qr Querier) (*Cursor, error) {
	return newCursor(x.dia, qr, x)
}

// CursorContext creates a DECLARE CURSOR statment and excetutes the statement and returns the Cursor
func (x SelectQuery) CursorContext(ctx context.Context, qr Querier) (*Cursor, error) {
	return newCursorContext(ctx, x.dia, qr, x)
}
