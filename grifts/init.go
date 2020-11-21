package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/todos-api/coke/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
