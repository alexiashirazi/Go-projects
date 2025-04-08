package button

import (
	"context"
	"io"
)

type Model struct {
	id                     string
	title                  string
	disabled               bool
	toggleModalBehaviourId string
	dismissModalBehaviour  bool
}

func Standard(id, title string) *Model {
	return &Model{
		id:    id,
		title: title,
	}
}

func (m *Model) WithDisabled() *Model {
	m.disabled = true

	return m
}

func (m *Model) WithDismissModalBehaviour() *Model {
	m.dismissModalBehaviour = true

	return m
}

func (m *Model) WithToggleModalBehaviour(id string) *Model {
	m.toggleModalBehaviourId = id

	return m
}

func (m *Model) Render(ctx context.Context, w io.Writer) error {
	return component(m).Render(ctx, w)
}
