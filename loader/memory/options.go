package memory

import (
	"github.com/argent-smith/go-config/loader"
	"github.com/argent-smith/go-config/reader"
	"github.com/argent-smith/go-config/source"
)

// WithSource appends a source to list of sources
func WithSource(s source.Source) loader.Option {
	return func(o *loader.Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}
