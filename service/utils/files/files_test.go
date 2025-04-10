//go:build unit
// +build unit

package files_test

import (
	"testing"

	"github.com/TeddyCr/priceitt/service/utils/files"
	"github.com/stretchr/testify/assert"
)

func TestGetRoot(t *testing.T) {
	path := files.GetRoot(rootFileTest)
	assert.Equal(t, "/my", path)

	path = files.GetRootDefault()
	assert.NotEmpty(t, path)
}

func rootFileTest() (string, error) {
	return "/my/fake/path", nil
}
