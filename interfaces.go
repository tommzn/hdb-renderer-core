package core

import (
	"github.com/golang/protobuf/proto"
	core "github.com/tommzn/hdb-core"
)

// Template renders content.
type Template interface {

	// Render generated content based on a template with given values.
	RenderWith(interface{}) (string, error)
}

// DataSource is used to get data for templates.
type DataSource interface {

	// Latest willreturn latest element for given datasource.
	Latest(core.DataSource) (proto.Message, error)
}

// Renderer generates content based on templates and event data.
type Renderer interface {

	// Size returns height and width a rendered element uses.
	Size() Size

	// Content returns rendered elements.
	Content() (string, error)
}
