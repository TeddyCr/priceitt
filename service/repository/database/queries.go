package repository

import (
	"bytes"
	"fmt"
)

const (
	InsertQuery = `INSERT INTO %s (json) VALUES ($1::jsonb)`
)

const (
	DeleteById = `DELETE FROM %s WHERE id = $1 AND %s`
)

const (
	DeleteByName = `DELETE FROM %s WHERE name = $1 AND %s`
)

const (
	GetByName = `SELECT json FROM %s WHERE name = $1 AND %s`
)

const (
	GetById = `SELECT json FROM %s WHERE id = $1 AND %s`
)

func NewQueryFilter(m map[string]string) *QueryFilter {
	if m == nil {
		m = make(map[string]string)
	}
	return &QueryFilter{
		Filter: m,
	}
}

type QueryFilter struct {
	Filter map[string]string
}

func (q *QueryFilter) String() string {
	if len(q.Filter) == 0 {
		return `True`
	}
	b := new(bytes.Buffer)
	i := 0
	for k := range q.Filter {
		if i > 0 {
			b.WriteString(` AND `)
		}
		fmt.Fprintf(b, `%s=$%d`, k, i+2)
		i++
	}
	return b.String()
}

func (q *QueryFilter) Args() []any {
	args := make([]any, len(q.Filter))
	i := 0
	for _, v := range q.Filter {
		args[i] = v
		i++
	}
	return args
}
