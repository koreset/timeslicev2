package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/koreset/timeslice/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
