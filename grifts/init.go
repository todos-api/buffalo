package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/todos-api/buffalo/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
