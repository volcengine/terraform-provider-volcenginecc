package ccschema

// FilterProperties lists the properties a List request can filter on and the
// filters the request requires.
type FilterProperties struct {
	FilterableProperties PropertyJsonPointers `json:"filterableProperties,omitempty"`
	Required             PropertyJsonPointers `json:"required,omitempty"`
}
