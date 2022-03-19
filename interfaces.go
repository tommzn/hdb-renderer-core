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

	// All returns a list of available events for passed datasource.
	All(core.DataSource) ([]proto.Message, error)

	// Observe returns a channels to listen for new events.
	// You can pass a filter to listen for events from specific datasource, only.
	// Datasource will not block if channel capacity is reached. Events will be skipped!
	Observe(*[]core.DataSource) <-chan proto.Message
}

// Renderer generates content based on templates and event data.
type Renderer interface {

	// Content returns rendered elements.
	Content() (string, error)
}

// TimestampManager checks timestamp of events.
type TimestampManager interface {

	// IsLatest will returns true if there's no similar event in local storage
	// or if local timestamp is older than in passed event.
	IsLatest(proto.Message) bool

	// IsLatestWithSuffix acts in the same way as IsLatest with an optional
	// identifier suffix.
	IsLatestWithSuffix(proto.Message, string) bool

	// Add timestamp of passed event to local storage for later checks.
	// Identifier suffix is optional.
	Add(proto.Message)

	// AddWithSuffix works like Add with an additional type name suffix.
	AddWithSuffix(proto.Message, string)
}
