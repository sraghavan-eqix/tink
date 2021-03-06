package mock

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
)

type Template struct {
	ID   uuid.UUID
	Data string
}

// CreateTemplate creates a new workflow template
func (d DB) CreateTemplate(ctx context.Context, name string, data string, id uuid.UUID) error {
	if d.TemplateDB == nil {
		d.TemplateDB = make(map[string]interface{})
	}

	if _, ok := d.TemplateDB[name]; ok {
		return errors.New("Template name already exist in the database")
	}

	d.TemplateDB[name] = Template{
		ID:   id,
		Data: data,
	}

	return nil
}

// GetTemplate returns a workflow template
func (d DB) GetTemplate(ctx context.Context, id string) (string, string, error) {
	return d.GetTemplateFunc(ctx, id)
}

// DeleteTemplate deletes a workflow template
func (d DB) DeleteTemplate(ctx context.Context, name string) error {
	if d.TemplateDB != nil {
		delete(d.TemplateDB, name)
	}

	return nil
}

// ListTemplates returns all saved templates
func (d DB) ListTemplates(fn func(id, n string, in, del *timestamp.Timestamp) error) error {
	return nil
}

// UpdateTemplate update a given template
func (d DB) UpdateTemplate(ctx context.Context, name string, data string, id uuid.UUID) error {
	return nil
}

// ClearTemplateDB clear all the templates
func (d DB) ClearTemplateDB() {
	d.TemplateDB = make(map[string]interface{})
}
