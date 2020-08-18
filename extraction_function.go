package godruid

// See https://druid.apache.org/docs/latest/querying/dimensionspecs.html#extraction-functions
// for details.
type ExtractionFn struct {
	Type string `json:"type"`
}

