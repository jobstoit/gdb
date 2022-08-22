package query

import (
	"io"

	"github.com/myceliums/gdb/dialect"
)

type Condition interface {
	SQL(wr io.Writer, dialect dialect.Dialect)
	Args() []interface{}
}

type primeCondition struct {
	a         Field
	b         Field
	opperator string
	args      []interface{}
}

func newPrimeCondition(a, b Field, opperator string) *primeCondition {
	x := &primeCondition{}
	x.a = a
	x.b = b
	x.opperator = opperator

	return x
}

// SQL is an implementation of Condition
func (x primeCondition) SQL(wr io.Writer, dialect dialect.Dialect) {
	//
}

// Args is an implementation of Condition
func (x primeCondition) Args() []interface{} {
	return x.args
}

type chainCondition struct {
	cds       []Condition
	opperator string
}

// And is an and statement
func (x chainCondition) SQL(wr io.Writer, dialect dialect.Dialect) {
	//
}

// Args is an implementation
func (x chainCondition) Args() []interface{} {
	var args []interface{}

	for _, c := range x.cds {
		args = append(args, c.Args()...)
	}

	return args
}

// And returns an AND statement
func And(condition ...Condition) Condition {
	var x chainCondition
	x.opperator = `AND`
	x.cds = condition

	return x
}

func Or(condition ...Condition) Condition {
	var x chainCondition
	x.opperator = `OR`
	x.cds = condition

	return x
}

// Gt is a greater than statement
func Gt(a, b Field) Condition {
	return newPrimeCondition(a, b, `>`)
}

// Gte is a greater than or eaqual to statement
func Gte(a, b Field) Condition {
	return newPrimeCondition(a, b, `>=`)
}

// Lt is a lesser than statement
func Lt(a, b Field) Condition {
	return newPrimeCondition(a, b, `<`)
}

// Lte is a lesser than or equal to statment
func Lte(a, b Field) Condition {
	return newPrimeCondition(a, b, `<=`)
}

// Eq is a equal to statement
func Eq(a, b Field) Condition {
	return newPrimeCondition(a, b, `=`)
}

// Neq is a not equal to statement
func Neq(a, b Field) Condition {
	return newPrimeCondition(a, b, `!=`)
}
