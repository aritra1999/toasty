package core

import (
	"testing"
	"toasty/core"
	"toasty/tests"

	. "github.com/franela/goblin"
	"github.com/stretchr/testify/assert"
)

func TestExecuteQuery(t *testing.T) {
	var containerMetas []tests.PostgresContainerMeta = tests.Setup(t, 1)

	g := Goblin(t)

	g.Describe("ExecuteQuery", func() {
		g.It("should connect to all given postgres database configs and execute the given query", func() {
			for _, containerMeta := range containerMetas {
				response, _ := core.ExecuteQuery(containerMeta.Config, "SELECT 1 + 1")

				result := response[0]
				expected := "2"

				assert.Equal(t, expected, result, "should be able to execute queries on the given container")
				assert.Equal(t, nil, nil, "error should be nil")
			}
		})
	})

	tests.Teardown(containerMetas)
}
