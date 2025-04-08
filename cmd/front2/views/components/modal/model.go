package modal

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type Model struct {
	id     string
	header templ.Component
	body   templ.Component
	footer templ.Component
}

func Make(id string, header, body, footer templ.Component) *Model {
	return &Model{
		id:     id,
		header: header,
		body:   body,
		footer: footer,
	}
}

func (m *Model) Render(ctx context.Context, w io.Writer) error {
	return component(m).Render(ctx, w)
}
