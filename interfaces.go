package core

import (
	"context"

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

	// All returns a list of available events for passed datasource.
	All(core.DataSource) ([]proto.Message, error)

	// Observe returns a channels to listen for new events.
	// You can pass a filter to listen for events from specific datasource, only.
	// Datasource will not block if channel capacity is reached. Events will be skipped!
	Observe(*[]core.DataSource) <-chan proto.Message
}

// Renderer generates content based on templates and event data.
type Renderer interface {

	// Size returns height and width a rendered element uses.
	Size() Size

	// Content returns rendered elements.
	Content() (string, error)

	// ObserveDataSource will start listen for new events provided by used datasource to get updated data for rendering.
	// Uses passed context to wait for cancelations.
	ObserveDataSource(context.Context)
}
