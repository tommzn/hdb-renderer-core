package core

// Template renders content.
type Template interface {

	// Render generated content based on a template with given values.
	RenderWith(interface{}) (string, error)
}

// UITemplate is used to render content for eink displays.
type UITemplate interface {

	// SetAnchor for the top left corner of a canvas.
	SetAnchor(Point)

	// GetSize returns height and width a rendered canvas uses.
	GetSize() Size

	// Render generated content based assigned position and event values.
	Render() (string, error)
}

// DataSource is used to get data for templates.
type DataSource interface {

	// Get should return data for template rendering.
	Get() (interface{}, error)
}
