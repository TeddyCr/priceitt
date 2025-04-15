//go:build unit
// +build unit

package repository

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryFilter(t *testing.T) {
	filter := NewQueryFilter(map[string]string{
		"id": "1",
	})
	assert.Equal(t, filter.String(), "id=$2")

	filter = NewQueryFilter(map[string]string{
		"city": "Montreal",
		"name": "test",
	})
	assert.True(t, filter.String() == "city=$2 AND name=$3" || filter.String() == "name=$2 AND city=$3")

	formattedQuery := fmt.Sprintf(GetById, "test", filter.String())
	assert.True(t, formattedQuery == "SELECT json FROM test WHERE id = $1 AND city=$2 AND name=$3")

	filter = NewQueryFilter(map[string]string{})
	assert.Equal(t, filter.String(), "True")

	formattedQuery = fmt.Sprintf(GetById, "test", filter.String())
	assert.Equal(t, formattedQuery, "SELECT json FROM test WHERE id = $1 AND True")
}
