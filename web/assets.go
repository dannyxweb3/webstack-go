package assets

import "embed"

var (
	//go:embed static
	Static embed.FS

	//go:embed templates
	Templates embed.FS

	//# go:embed wp-content
	// WpContent embed.FS
)
