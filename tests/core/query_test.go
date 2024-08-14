package core

import (
	"testing"
	"toasty/core"
	"toasty/tests"

	"github.com/stretchr/testify/assert"
)

func TestExecuteQuery(t *testing.T) {
	containerMetas := tests.Setup(t, 5)
	//TODO: This part will be replaced with the actual tests and assertions later
	for _, containerMeta := range containerMetas {
		result, err := core.ExecuteQuery(containerMeta.Config, "SELECT * FROM users")
		assert.Equal(t, nil, err, "error should be nil")
		assert.Equal(t, nil, result, "result should be nil")
	}

	tests.Teardown(containerMetas)
}
